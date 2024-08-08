package handler

import (
	"api-gateway-service/api/token"
	"api-gateway-service/generated/users"
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func (h *Handler) Register(ctx *gin.Context) {
	auth := users.CreateUsersRequest{}
	

	if err := ctx.ShouldBindJSON(&auth); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	_, err := h.User.CreateUser(ctx, &auth)
	if err != nil {
		if pgErr, ok := err.(*pq.Error); ok && pgErr.Code == "23505" {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "The email is already used",
			})
			return
		}
		ctx.AbortWithStatusJSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Successfully created",
	})
}

func (h *Handler) Login(ctx *gin.Context) {
	var req users.LoginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	exists, err := h.User.EmailExists(ctx, &users.EmailExistsRequest{Email: req.Email})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if !exists.Exists {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "Email is not found",
		})
		return
	}

	res, err := h.User.Login(ctx, &users.LoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "Password is incorrect",
			})
		}
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	err = token.GeneratedJWTToken(res)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"error": err.Error(),
		})
	}

	ctx.IndentedJSON(http.StatusOK, res)
}
