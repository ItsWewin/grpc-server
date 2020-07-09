 cd profiles
 protoc --go_out=plugins=grpc:../services Prod.proto
 protoc --go_out=plugins=grpc:../services Models.proto
 protoc --go_out=plugins=grpc:../services Order.proto
 cd ..
