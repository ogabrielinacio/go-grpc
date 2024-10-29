package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/ogabrielinacio/go-grpc/internal/database"
	"github.com/ogabrielinacio/go-grpc/internal/pb"
	"github.com/ogabrielinacio/go-grpc/internal/service"
	"github.com/ogabrielinacio/go-grpc/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	connStr := utils.GetConnectionString()
	fmt.Println(connStr)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	defer db.Close()

	categoryDb := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDb)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":5051")
	if err != nil {
		panic(err)
	}
	if err := grpcServer.Serve(listener); err != nil {
		panic(err)
	}
}
