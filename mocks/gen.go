//go:build ignore

package main

//go:generate go run github.com/golang/mock/mockgen -source=./template-service/internal/ports/repository.go -destination=./mocks/mockrepo.go -package=main
//go:generate go run github.com/golang/mock/mockgen -source=./template-service/internal/usecase/interface.go -destination=./mocks/mockusecase.go -package=main
