package main

import (
	"os"

	"github.com/RBMarketplace/di-api-teams/pkg/handler"
	"github.com/RBMarketplace/di-api-teams/pkg/repository"
	"github.com/RBMarketplace/di-api-teams/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisUrl := os.Getenv("REDIS_URL")
	if len(redisUrl) == 0 {
		redisUrl = "redis-master.twdps-core-labs-team-dev.svc.cluster.local"
	}

	router := gin.Default()

	// TODO - consider how to dynamically handle different data stores here
	teamRepo := repository.NewRedisTeamRepository(redisPassword, redisUrl)
	teamService := service.NewTeamService(teamRepo)
	teamHandler := handler.NewTeamHandler(teamService)

	namespaceRepo := repository.NewRedisNamespaceRepository()
	namespaceService := service.NewNamespaceService(namespaceRepo)
	namespaceHandler := handler.NewNamespaceHandler(namespaceService)

	router.GET("/teams/healthz/readiness", teamHandler.Readiness)
	router.GET("/teams/healthz/liveness", teamHandler.Liveness)
	router.GET("/teams/:teamID", teamHandler.GetTeam)
	router.GET("/teams", teamHandler.GetTeams)
	router.POST("/teams", teamHandler.AddTeam)
	router.DELETE("/teams/:teamID", teamHandler.RemoveTeam)
	router.DELETE("/teams/:teamID/confirm", teamHandler.ConfirmRemoveTeam)

	router.GET("/namespaces", namespaceHandler.GetNamespaces)
	router.GET("/namespaces/master", namespaceHandler.GetNamespacesMaster)
	router.GET("/namespaces/standard", namespaceHandler.GetNamespacesStandard)
	router.GET("/namespaces/custom", namespaceHandler.GetNamespacesCustom)
	router.POST("/namespaces", namespaceHandler.AddNamespace)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
