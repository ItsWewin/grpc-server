package services

import (
	"context"
	"fmt"
	"io"
	"log"
)

type UserServer struct {

}

func (u *UserServer) GetUsersInfo(ctx context.Context, userReq *UserRequest) (*UserResponse, error) {
	userIDList := userReq.UserIds
	fmt.Println(userIDList)
	var users []*User
	for _, id := range userIDList {
		user := &User{
			Id:            id,
			Score:         1000 + id,
		}
		users = append(users, user)
	}

	return &UserResponse{Users:users}, nil
}

func (u *UserServer) GetUserInfoByServerStream(userReq *UserRequest, stream UserService_GetUserInfoByServerStreamServer) error {

	var score int32=101
	users := make([]*User, 0)
	for index, id := range userReq.UserIds {
		user := &User{
			Id:            id,
			Score:         score,
		}
		score++
		users = append(users, user)

		if (index+1) % 2 == 0 && index > 0 {
			err := stream.Send(&UserResponse{Users: users})
			if err != nil {
				return err
			}

			users =(users)[0:0]
		}
	}

	if len(users) > 0 {
		err := stream.Send(&UserResponse{Users:users})
		if err != nil {
			return err
		}
	}

	return nil
}


func (u *UserServer) GetUserInfoByClientStream(stream UserService_GetUserInfoByClientStreamServer) error {
	users := make([]*User, 0)
	for {
		req, err := stream.Recv()
		switch {
		case err == io.EOF:
			return stream.SendAndClose(&UserResponse{Users: users})
		case err != nil:
			log.Fatal(err)
		}

		for index, id := range req.UserIds {
			user := &User{
				Id:            id,
				Score:         int32(index) + id,
			}
			users = append(users, user)
		}
	}
}

func (u *UserServer) GetUserInfoByBothSideStream(stream UserService_GetUserInfoByBothSideStreamServer) error {
	users := make([]*User, 0)
	for {
		req, err := stream.Recv()
		switch  {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}
		for _, id := range req.UserIds {
			u := &User{
				Id:id,
				Score: 100000000 + id,
			}
			users = append(users, u)
		}
		err = stream.Send(&UserResponse{Users: users})
		if err != nil {
			log.Fatal(err)
		}
		users = (users)[0:0]
	}
}