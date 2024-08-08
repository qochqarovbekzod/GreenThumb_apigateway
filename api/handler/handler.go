package handler

import (
	"api-gateway-service/generated/community"
	"api-gateway-service/generated/gardenManagement"
	"api-gateway-service/generated/sustainability"
	"api-gateway-service/generated/users"
)


type Handler struct{
	User users.UserManagementClient
	Garden gardenManagement.GardenManagementClient
	Community community.ComunityServiceClient
	Sustainability sustainability.SustainabilityimpactServiceClient
}

