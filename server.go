//go:generate protoc ./student.proto --go_out=plugins=grpc:./pb

package student

import (
	"context"
	"fmt"
	"github.com/jinzhu/copier"
	"github.com/tyagip966/common-repo/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"student/constants"
	"student/models/postgres"
	"student/pb"
)

type GrpcServer struct {
	Service *postgres.StudentService
}

func (g GrpcServer) AddStudent(ctx context.Context, request *pb.AddStudentRequest) (*pb.AddStudentResponse, error) {
	var input models.Student
	_ = copier.Copy(&input, &request.Input)
	result, err := g.Service.AddStudent(input)
	if err != nil {
		return nil, err
	}
	respone := new(pb.Student)
	copier.Copy(respone, result)
	return &pb.AddStudentResponse{
		Student: respone,
	}, nil
}

func (g GrpcServer) GetStudent(ctx context.Context, request *pb.GetStudentRequest) (*pb.AddStudentResponse, error) {
	result, err := g.Service.GetStudent(int(request.ID))
	if err != nil {
		return nil, err
	}
	respone := new(pb.Student)
	copier.Copy(respone, result)
	return &pb.AddStudentResponse{
		Student: respone,
	}, nil
}

func (g GrpcServer) UpdateStudent(ctx context.Context, request *pb.UpdateStudentRequest) (*pb.AddStudentResponse, error) {
	var input models.Student
	_ = copier.Copy(&input, &request.Input)
	result, err := g.Service.UpdateStudent(int(request.ID), input)
	if err != nil {
		return nil, err
	}
	respone := new(pb.Student)
	copier.Copy(respone, result)
	return &pb.AddStudentResponse{
		Student: respone,
	}, nil
}

func (g GrpcServer) DeleteStudent(ctx context.Context, request *pb.DeleteStudentRequest) (*pb.AddStudentResponse, error) {
	result, err := g.Service.DeleteStudent(int(request.ID))
	if err != nil {
		return nil, err
	}
	respone := new(pb.Student)
	copier.Copy(respone, result)
	return &pb.AddStudentResponse{
		Student: respone,
	}, nil
}

func (g GrpcServer) GetStudents(ctx context.Context, request *pb.GetStudentsRequest) (*pb.GetStudentsResponse, error) {
	result, err := g.Service.GetStudents(int(request.SchoolCode))
	if err != nil {
		return nil, err
	}
	var respone []*pb.Student
	for _, i := range result {
		res := new(pb.Student)
		_ = copier.Copy(res, &i)
		respone = append(respone, res)
	}
	copier.Copy(respone, result)
	return &pb.GetStudentsResponse{
		Student: respone,
	}, nil
}

func ListenGRPC(s postgres.StudentService, port int) (*postgres.StudentService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))

	if err != nil {
		return nil, err
	}

	serv := grpc.NewServer()
	pb.RegisterStudentServiceServer(serv, &GrpcServer{&s})
	reflection.Register(serv)
	log.Println("Server Started at ", constants.ServerPort)
	return &s, serv.Serve(lis)
}
