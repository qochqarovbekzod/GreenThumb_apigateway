package api

import (
	"api-gateway-service/api/handler"
	"api-gateway-service/api/middleware"

	"github.com/gin-gonic/gin"
)

func Router(hand *handler.Handler) *gin.Engine {
	router := gin.Default()

	auth := router.Group("/api/auth")
	auth.POST("/register", hand.Register)
	auth.POST("/login", hand.Login)

	protected := router.Group("/")
	protected.Use(middleware.Middleware)

	user := protected.Group("/api/users")
	{
		user.GET("/:user-id", hand.GetUserByIdHandler)
		user.POST("/:user-id", hand.CreateUserHandler)
		user.DELETE("/:user-id", hand.DeleteUserHandler)
		user.GET("/:user-id/profile", hand.GetUserByIdProfileHandler)
		user.PUT("/:user-id/profile", hand.UpdateUserProfileHandler)
	}

	garden := protected.Group("/api/gardens")
	{
		garden.POST("/", hand.CreateGardenHandler)
		garden.GET("/:garden-id", hand.ViewGardeHandler)
		garden.PUT("/:garden-id", hand.UpdateGardenHandler)
		garden.DELETE("/:garden-id", hand.DeleteGardenHandler)
		garden.POST("/:garden-id/plants", hand.AddPlanttoGarden)
		garden.GET("/:garden-id/plants", hand.ViewGardenPlantsHandler)
	}

	protected.GET("/api/users/:user-id/gardens", hand.ViewUserGardensHandler)

	plant := protected.Group("/api/plants")
	{
		plant.PUT("/:plant-id", hand.UpdatePlantHandler)
		plant.DELETE("/:plant-id", hand.DeletePlantHandler)
		plant.POST("/:plant-id/care-logs", hand.AddPlantCareLogHandler)
		plant.GET("/:plant-id/care-logs", hand.ViewPlantCareLogsHandler)
	}

	community := protected.Group("/api/communities")
	{
		community.POST("/", hand.CreateCommunityHandler)
		community.GET("/:community-id", hand.GetCommunityHandler)
		community.PUT("/:community-id", hand.UpdateCommunityHendler)
		community.DELETE("/:community-id", hand.DeleteCommunityHandler)
		community.GET("/", hand.ListCommunitiesHandler)
		community.POST("/:community-id/join", hand.JoinCommunityHendler)
		community.POST("/:community-id/leave", hand.LeaveCommunityHandler)
		community.POST("/:community-id/events", hand.CreateCommunityEventHendler)
		community.GET("/:community-id/events", hand.ListCommunityEventsHandler)
		community.POST("/:community-id/forum", hand.CreateCommunityForumPostHendler)
		community.GET("/:community-id/forum", hand.ListCommunityForumPostsHandler)
	}

	forum := protected.Group("/api/forum")
	{
		forum.POST("/:post-id", hand.AddForumPostCommentHendler)
		forum.GET("/:post-id", hand.ListForumPostCommentsHandler)
	}

	{
		protected.POST("/api/impact/log", hand.LogImpactHandle)
		protected.GET("/api/users/:user-id/impact", hand.GetUserImpactHandle)
		protected.GET("/api/communties/:community-id/impact", hand.GetCommunityImpactHandle)
		protected.GET("/api/challenges", hand.GetChallengesHandle)
		protected.POST("/api/challenges/:challenge-id", hand.JoinChallengeHendler1)
		protected.PUT("/api/challenges/:challenge-id/progress", hand.UpdateChallengeProgressHendler1)
		protected.GET("/api/users/:user-id/challenges", hand.GetUserChallengesHandler)
		protected.GET("/api/leaderboard/users", hand.GetUsersLeaderboardHandler)
		protected.GET("api/leaderboars/communities", hand.GetCommunitiesLeaderboardHandler)
	}

	return router
}
