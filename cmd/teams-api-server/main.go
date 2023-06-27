package main

import (
	"github.com/gin-gonic/gin"
	"twdps.io/lab-api-teams/pkg/handler"
	"twdps.io/lab-api-teams/pkg/repository"
	"twdps.io/lab-api-teams/pkg/service"
)

func main() {
	router := gin.Default()

	teamRepo := repository.NewRedisTeamRepository()
	teamService := service.NewTeamService(teamRepo)
	teamHandler := handler.NewTeamHandler(teamService)

	router.GET("/teams/healthz/readiness", teamHandler.Readiness)
	router.GET("/teams/healthz/liveness", teamHandler.Liveness)
	router.GET("/teams/:teamID", teamHandler.GetTeam)
	router.GET("/teams", teamHandler.GetTeams)
	router.POST("/teams", teamHandler.AddTeam)
	router.DELETE("/teams/:teamID", teamHandler.RemoveTeam)
	router.DELETE("/teams/:teamID/confirm", teamHandler.ConfirmRemoveTeam)

	router.Run(":8080")
}
