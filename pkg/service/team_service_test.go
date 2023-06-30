package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"twdps.io/lab-api-teams/pkg/domain"
	"twdps.io/lab-api-teams/pkg/repository/mock"
)

func TestGetTeams(t *testing.T) {
	mockRepository := &mock.MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}

	teamService := NewTeamService(mockRepository)
	teams, _ := teamService.GetTeams()
	if len(teams) != 1 {
		t.Errorf("expected %d teams got %d", 1, len(teams))
	}
}

func TestAddTeam(t *testing.T) {
	mockRepository := &mock.MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}

	teamService := NewTeamService(mockRepository)
	var newTeam = domain.Team{
		TeamID: "team-jade", TeamType: "normal", TeamDescription: "jade frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24,
	}

	err := teamService.AddTeam(newTeam)
	if err != nil {
		t.Errorf("failed to add team")
	}

	teams, _ := teamService.GetTeams()
	if len(teams) != 2 {
		t.Errorf("expected %d teams got %d", 2, len(teams))
	}

}

func TestUpdateTeam(t *testing.T) {
	mockRepository := &mock.MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}

	teamService := NewTeamService(mockRepository)
	err := teamService.RequestRemoveTeam("team-sapphire")
	if err != nil {
		t.Errorf("failed to add team")
	}

	ts, _ := teamService.GetTeam("team-sapphire")
	if ts.TeamMarkedForDeletion != "Requested" {
		t.Errorf("expected %s team to be marked for delete", "team-sapphire")
	}
}

func TestConfirmFailedDeleteTeam(t *testing.T) {
	mockRepository := &mock.MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}
	teamID := "team-sapphire"

	teamService := NewTeamService(mockRepository)

	err := teamService.ConfirmRemoveTeam(teamID)
	expectedErrorMsg := fmt.Sprintf("Team %s is not requested for deletion", teamID)
	assert.EqualErrorf(t, err, expectedErrorMsg, "Error should be: %v, got %v", expectedErrorMsg, err)
}

func TestConfirmDeleteTeam(t *testing.T) {
	mockRepository := &mock.MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}
	teamID := "team-sapphire"

	teamService := NewTeamService(mockRepository)
	err := teamService.RequestRemoveTeam(teamID)
	if err != nil {
		t.Errorf("failed to add team")
	}

	teamRequestedForDelete, _ := teamService.GetTeam(teamID)
	if teamRequestedForDelete.TeamMarkedForDeletion != "Requested" {
		t.Errorf("expected %s team to be marked for delete", teamID)
	}

	err = teamService.ConfirmRemoveTeam(teamID)
	if err != nil {
		t.Errorf("failed to delete team")
	}

	teamRequestedForDelete, _ = teamService.GetTeam(teamID)
	if teamRequestedForDelete.TeamID == "team-sapphire" {
		t.Errorf("Failed to remove team from database")
	}
}
