package services

import (
	"context"
	"fmt"
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
