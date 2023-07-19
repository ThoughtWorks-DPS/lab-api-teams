package handler

import (
	"log"
	"net/http"

	"github.com/RBMarketplace/di-api-teams/pkg/domain"
	"github.com/RBMarketplace/di-api-teams/pkg/service"
	"github.com/gin-gonic/gin"
)

type TeamHandler struct {
	teamService service.TeamService
}

func NewTeamHandler(teamService service.TeamService) *TeamHandler {
	return &TeamHandler{teamService: teamService}
}

func (handler *TeamHandler) Liveness(c *gin.Context) {
	c.Status(http.StatusOK)
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Teams api is live"})
}

// TODO - given that we implement many different database types, the database available function
//        should return a string or object wsith the type in it, like { type: redis, conn_url: localhost:6379, available: true }
//				This is good info for error logs when debugging (Database %s is not available yet, db.Type)
func (handler *TeamHandler) Readiness(c *gin.Context) {
	dbReady, err := handler.teamService.DatabaseAvailable()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	if dbReady != true {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Database is not available"})
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "DB is available"})
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

func (handler *TeamHandler) GetTeam(c *gin.Context) {
	teamID := c.Param("teamID")

	team, err := handler.teamService.GetTeam(teamID)
	if err != nil { // TODO transient/status errors
		log.Printf("error %v", err)
		return
	}

	c.IndentedJSON(http.StatusOK, team)
}

func (handler *TeamHandler) RemoveTeam(c *gin.Context) {
	teamID := c.Param("teamID")

	err := handler.teamService.RequestRemoveTeam(teamID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err) // TODO
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Team delete requested"})
}

func (handler *TeamHandler) ConfirmRemoveTeam(c *gin.Context) {
	teamID := c.Param("teamID")

	err := handler.teamService.ConfirmRemoveTeam(teamID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, err.Error()) // TODO
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Team deleted"})
}
