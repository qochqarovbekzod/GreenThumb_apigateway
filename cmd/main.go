package main

import (
	"api-gateway-service/api"
	"api-gateway-service/api/handler"
	"api-gateway-service/config"
	"api-gateway-service/generated/community"
	"api-gateway-service/generated/gardenManagement"
	"api-gateway-service/generated/sustainability"
	"api-gateway-service/generated/users"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	cfg := config.Load()
	hand := NewConnect()

	router := api.Router(hand)

	log.Fatal(router.Run(cfg.HTTP_PORT))
}

func NewConnect() *handler.Handler {
	usersConn, err := grpc.Dial("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error: ", err.Error())
		return nil
	}
	usersService := users.NewUserManagementClient(usersConn)

	gardenManagementConn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return nil
	}
	gardenManagementService := gardenManagement.NewGardenManagementClient(gardenManagementConn)

	communityConn, err := grpc.Dial("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return nil
	}
	communityService := community.NewComunityServiceClient(communityConn)

	sustainabilityConn, err := grpc.Dial("localhost:50053", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("ERROR: ", err.Error())
		return nil
	}
	sustainabilityService := sustainability.NewSustainabilityimpactServiceClient(sustainabilityConn)

	return &handler.Handler{
		User:           usersService,
		Garden:         gardenManagementService,
		Community:      communityService,
		Sustainability: sustainabilityService,
	}
}
