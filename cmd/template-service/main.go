package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"template-service/config"
	"template-service/internal/adapters/grpc/grpcservice"
	"template-service/internal/adapters/repository"
	"template-service/internal/usecase"
	"template-service/pkg/mongo"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Parse config
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("config.New(): %v", err)
	}

	// Mongo connect
	mongoDB, err := mongo.NewConnect(ctx, cfg.Mongo)
	if err != nil {
		log.Fatalf("mongo.NewConnect: %v", err)
	}

	// Net listener
	list, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("net.Listen: %v", err)
	}

	// gRPC server
	gRPCServer := grpcserver.New(ctx)

	userRepo := repository.NewUser(mongoDB)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userService := grpcservice.NewUserService(userUseCase)
	fbsvc.RegisterUserServiceServer(gRPCServer, userService)

	runErr := make(chan error, 1)
	go func() {
		log.Printf("gRPC server start listen on: %d", cfg.GRPC.Port)
		errSrv := gRPCServer.Serve(list)
		if errSrv != nil {
			runErr <- fmt.Errorf("can't start gRPC server: %w", err)
		}
	}()

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-runErr:
		log.Fatalf("Running error: %s", errRun)
	case s := <-shutdownCh:
		log.Printf("Run - signal: %v", s.String())

		gRPCServer.GracefulStop()

		log.Printf("%v", "Graceful shutdown completed!")
	}

}
