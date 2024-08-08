package handler

import (
	"api-gateway-service/api/token"
	"api-gateway-service/generated/community"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateCommunityHandler(ctx *gin.Context) {
	var req community.CreateCommunityRequest

	err := ctx.ShouldBind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	resp, err := h.Community.CreateCommunity(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, &resp)
}

func (h *Handler) GetCommunityHandler(ctx *gin.Context) {
	id := ctx.Param("community-id")

	req := community.GetCommunityRequest{Id: id}

	resp, err := h.Community.GetCommunity(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) UpdateCommunityHendler(ctx *gin.Context) {
	id := ctx.Param("community-id")
	var req *community.UpdateCommunityRequest

	err := ctx.ShouldBind(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	req.Id = id

	resp, err := h.Community.UpdateCommunity(ctx, req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) DeleteCommunityHandler(ctx *gin.Context) {
	id := ctx.Param("community-id")

	req := community.DeleteCommunityRequest{Id: id}

	resp, err := h.Community.DeleteCommunity(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ListCommunitiesHandler(ctx *gin.Context) {
	var fCommunity community.ListCommunitiesRequest

	if err := ctx.ShouldBindQuery(&fCommunity); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid query parameters",
		})
		return
	}

	defaultLimit := 10
	defaultOffset := 0

	if limitStr := ctx.Query("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid limit parameter",
			})
			return
		}
		fCommunity.Limit = int64(limit)
	} else {
		fCommunity.Limit = int64(defaultLimit)
	}
	if offsetStr := ctx.Query("offset"); offsetStr != "" {
		offset, err := strconv.Atoi(offsetStr)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid offset parameter",
			})
			return
		}
		fCommunity.Offset = int64(offset)
	} else {
		fCommunity.Offset = int64(defaultOffset)
	}
	resp, err := h.Community.ListCommunities(ctx, &fCommunity)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) JoinCommunityHendler(ctx *gin.Context) {
	id := ctx.Param("community-id")

	claims, _ := token.ExtractClaims(ctx.GetHeader("Authorization"))
	userId, _ := (*claims)["user_id"].(string)

	req := community.JoinCommunityRequest{
		UserId:      userId,
		CommunityId: id,
	}

	resp, err := h.Community.JoinCommunity(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) LeaveCommunityHandler(ctx *gin.Context) {
	id := ctx.Param("community-id")

	claims, _ := token.ExtractClaims(ctx.GetHeader("Authorization"))
	userId, _ := (*claims)["user_id"].(string)

	req := community.LeaveCommunityRequest{
		CommunityId: id,
		UserId: userId,
	}

	resp, err := h.Community.LeaveCommunity(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) CreateCommunityEventHendler(ctx *gin.Context) {
	id := ctx.Param("community-id")
	var req community.CreateCommunityEventRequest

	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	req.ComunityId = id

	resp, err := h.Community.CreateCommunityEvent(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ListCommunityEventsHandler(ctx *gin.Context) {
	id := ctx.Param("community-id")

	req := community.ListCommunityEventsRequest{CommunityId: id}

	resp, err := h.Community.ListCommunityEvents(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)

}

func (h *Handler) CreateCommunityForumPostHendler(ctx *gin.Context) {
	id := ctx.Param("community-id")

	var req community.CreateCommunityForumPostRequest

	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	req.CommunityId = id

	resp, err := h.Community.CreateCommunityForumPost(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ListCommunityForumPostsHandler(ctx *gin.Context) {
	id := ctx.Param("community-id")

	req := community.ListCommunityForumPostsRequest{ComunityId: id}

	resp, err := h.Community.ListCommunityForumPosts(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) AddForumPostCommentHendler(ctx *gin.Context) {

	id := ctx.Param("post-id")

	var req community.AddForumPostCommentRequest

	err := ctx.BindJSON(&req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	req.PostId = id

	resp, err := h.Community.AddForumPostComment(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}

func (h *Handler) ListForumPostCommentsHandler(ctx *gin.Context) {
	id := ctx.Param("post-id")

	req := community.ListForumPostCommentsRequest{PostId: id}

	resp, err := h.Community.ListForumPostComments(ctx, &req)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, resp)
}
