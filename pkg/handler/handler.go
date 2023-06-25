package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"twdps.io/lab-api-teams/pkg/domain"
	"twdps.io/lab-api-teams/pkg/service"
)

type TeamHandler struct {
	teamService service.TeamService
}

func NewTeamHandler(teamService service.TeamService) *TeamHandler {
	return &TeamHandler{teamService: teamService}
}

func (handler *TeamHandler) GetTeams(c *gin.Context) {
	teams, err := handler.teamService.GetTeams()
	if err != nil {
		log.Fatalf("Failed to call GetTeams %v", err)
	}

	c.IndentedJSON(http.StatusOK, teams)
}

func (handler *TeamHandler) AddTeam(c *gin.Context) {
	var newTeam domain.Team

	if err := c.BindJSON(&newTeam); err != nil {
		log.Printf("error %+v", err)
		return
	}

	err := handler.teamService.AddTeam(newTeam)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusCreated, newTeam)
}
