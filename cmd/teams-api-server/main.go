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

	router.GET("/teams", teamHandler.GetTeams)
	router.POST("/teams", teamHandler.AddTeam)

	router.Run(":8080")
}
