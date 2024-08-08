package handler

import (
	"api-gateway-service/generated/gardenManagement"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateGardenHandler(ctx *gin.Context) {
	var garden gardenManagement.CreateGardenRequest

	err := ctx.ShouldBindJSON(&garden)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})

		return 
	}
	resp, err := h.Garden.CreateGarden(ctx, &garden)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"succuss": resp.Success,
		"message": "Garden created succussfully",
	})
}

func (h *Handler) ViewGardeHandler(ctx *gin.Context) {
	id := ctx.Param("garden-id")
	req := gardenManagement.ViewGardenRequest{Id: id}

	garden, err := h.Garden.ViewGarden(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, garden)
}

func (h *Handler) UpdateGardenHandler(ctx *gin.Context) {
	id := ctx.Param("garden-id")

	var updateGarden gardenManagement.UpdateGardenRequest
	
	err := ctx.ShouldBindJSON(&updateGarden)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})

		return 
	}

	updateGarden.Id = id

	resp, err := h.Garden.UpdateGarden(ctx, &updateGarden)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if !resp.Success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": resp.Success,
			"message": "Garden not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"message": "Garden updeted successfully",
	})
}

func (h *Handler) DeleteGardenHandler(ctx *gin.Context) {
	id := ctx.Param("garden-id")

	req  := gardenManagement.DeleteGardenRequest{Id: id}

	resp, err := h.Garden.DeleteGarden(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if !resp.Success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": resp.Success,
			"message": "Garden not found or already deleted",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"message": "Garden deleted succussfully",
	})
}

func (h *Handler) ViewUserGardensHandler(ctx *gin.Context) {
	id := ctx.Param("user-id")
	req := gardenManagement.ViewUserGardensRequest{UserId: id}

	gardens, err := h.Garden.ViewUserGardens(ctx, &req)
	
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gardens)
}

func (h *Handler) AddPlanttoGarden(ctx *gin.Context) {
	id := ctx.Param("garden-id")
	var plant gardenManagement.AddPlanttoGardenRequest

	err := ctx.ShouldBindJSON(&plant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	plant.GardenId = id

	resp, err := h.Garden.AddPlanttoGarden(ctx, &plant)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"succuss": resp.Success,
		"Message": "Plant added to Garden succussfulle",
	})
}

func (h *Handler) ViewGardenPlantsHandler(ctx *gin.Context) {
	id := ctx.Param("garden-id")
	req := gardenManagement.ViewGardenPlantsRequest{GardenId: id}

	plants, err := h.Garden.ViewGardenPlants(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, plants)
}

func (h *Handler) UpdatePlantHandler(ctx *gin.Context) {
	id := ctx.Param("plant-id")
	var plant gardenManagement.UpdatePlantRequest
	
	err := ctx.ShouldBindJSON(&plant)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	plant.Id = id

	resp, err := h.Garden.UpdatePlant(ctx, &plant)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if !resp.Success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succuss": resp.Success,
			"Message": "Plant not found",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"Message": "PLant updated succussfully",
	})
}

func (h *Handler) DeletePlantHandler(ctx *gin.Context) {
	id := ctx.Param("plant-id")
	req := gardenManagement.DeletePlantRequest{Id: id}

	resp, err := h.Garden.DeletePlant(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if !resp.Success {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": resp.Success,
			"Message": "Plant not found or already deleted",
		})

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": resp.Success,
		"Message": "Pland deleted succsussfully",
	})
}

func (h *Handler) AddPlantCareLogHandler(ctx *gin.Context) {
	id := ctx.Param("plant-id")

	var careLog gardenManagement.CareLog

	err := ctx.ShouldBindJSON(&careLog)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	careLog.PlantId = id

	resp, err := h.Garden.AddPlantCareLog(ctx, (*gardenManagement.AddPlantCareLogResquest)(&careLog))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"succuss": resp.Success,
		"Message": "Care log added successfully to Plant",
	})
}

func (h *Handler) ViewPlantCareLogsHandler(ctx *gin.Context) {
	id := ctx.Param("plant-id")
	fmt.Println(id)
	req := gardenManagement.ViewPlantCareLogsRequest{PlantId: id}

	carelogs, err := h.Garden.ViewPlantCareLogs(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, carelogs)
}