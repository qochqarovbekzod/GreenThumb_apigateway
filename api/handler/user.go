package handler

import (
	pb "api-gateway-service/generated/users"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUserByIdHandler(c *gin.Context) {
	userId := c.Param("user-id")

	resp, err := h.User.GetUserById(c, &pb.GetUserByIdRequest{UserId: userId})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, resp)
}

func (h *Handler) CreateUserHandler(c *gin.Context) {
	var user pb.CreateUsersRequest
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.User.CreateUser(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteUserHandler(c *gin.Context) {
	userId:=c.Param("user-id")

	resp,err:=h.User.DeleteUser(c,&pb.DeleteUserRequest{UserId: userId})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK,resp)

}

func (h *Handler)  GetUserByIdProfileHandler(c *gin.Context) {
	userId:=c.Param("user-id")

	resp,err:=h.User.GetUserByIdProfile(c,&pb.GetUserByIdProfileRequest{UserId: userId})

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK,resp)
}

func (h *Handler) UpdateUserProfileHandler(c *gin.Context)  {
	id := c.Param("user-id")
	var userProfil pb.UpdateUserProfileRequest
	err:=c.BindJSON(&userProfil)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userProfil.UserId = id
	resp,err:=h.User.UpdateUserProfile(c,&userProfil)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK,resp)

}
