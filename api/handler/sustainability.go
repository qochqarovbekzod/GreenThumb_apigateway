package handler

import (
	"api-gateway-service/api/token"
	"api-gateway-service/generated/community"
	pb "api-gateway-service/generated/sustainability"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) LogImpactHandle(ctx *gin.Context) {
	var logImpact pb.LogImpactRequest

	claims, _ := token.ExtractClaims(ctx.GetHeader("Authorization"))
	userId, _ := (*claims)["user_id"].(string)
	logImpact.UserId = userId

	err := ctx.ShouldBindJSON(&logImpact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Sustainability.LogImpact(ctx, &logImpact)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Success": resp.Success,
		"Message": "Log impact created successfully",
	})
}

func (h *Handler) GetUserImpactHandle(ctx *gin.Context) {
	userId := ctx.Param("user-id")

	lomImpact, err := h.Sustainability.GetUserImpact(ctx, &pb.GetUserImpactRequest{UserId: userId})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, lomImpact)
}

func (h *Handler) GetCommunityImpactHandle(ctx *gin.Context) {
	id := ctx.Param("community-id")

	members, err := h.Community.CommunityMembers(ctx, &community.CommunityMembersRequest{ComunityId: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	logImpact, err := h.Sustainability.GetCommunityImpact(ctx, &pb.GetCommunityImpactRequest{
		CommunityId: id,
		Members: members.Members,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, logImpact)

}

func (h *Handler) GetChallengesHandle(ctx *gin.Context) {
	challenges, err := h.Sustainability.GetChallenges(ctx, &pb.GetChallengesRequest{})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, challenges)
}

func (h *Handler) JoinChallengeHendler1(ctx *gin.Context) {
	id := ctx.Param("challenge-id")
	var req pb.JoinChallengeRequest

	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	req.ChallengeId = id

	resp, err := h.Sustainability.JoinChallenge(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) UpdateChallengeProgressHendler1(ctx *gin.Context) {
	id := ctx.Param("challenge-id")

	var req pb.UpdateChallengeProgressRequest

	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	req.ChallengeId = id
	resp, err := h.Sustainability.UpdateChallengeProgress(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUserChallengesHandler(ctx *gin.Context) {
	id := ctx.Param("user-id")

	resp, err := h.Sustainability.GetUserChallenges(ctx, &pb.GetUserChallengesRequest{UserId: id})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetUsersLeaderboardHandler(ctx *gin.Context) {
	resp, err := h.Sustainability.GetCommunitiesLeaderboard(ctx, &pb.GetCommunitiesLeaderboardRequest{})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) GetCommunitiesLeaderboardHandler(ctx *gin.Context) {
	var req pb.GetCommunitiesLeaderboardRequest
	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Sustainability.GetCommunitiesLeaderboard(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
