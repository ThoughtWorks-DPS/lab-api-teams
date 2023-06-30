package main

import (
	"github.com/gin-gonic/gin"
	"twdps.io/lab-api-teams/pkg/handler"
	"twdps.io/lab-api-teams/pkg/repository"
	"twdps.io/lab-api-teams/pkg/service"
)

// TODO - PSK
// github.com/thoughtworks-dps/lab-api-teams/pkg/domain
// becomes
// github.com/thoughtworks-dps/psk-api-teams/pkg/domain

func main() {
	router := gin.Default()

	// TODO - consider how to dynamically handle different data stores here
	teamRepo := repository.NewRedisTeamRepository()
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

	router.Run(":8080")
}
