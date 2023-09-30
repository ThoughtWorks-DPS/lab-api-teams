//go:build !e2e
// +build !e2e

package service

import (
	"fmt"
	"testing"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetTeams(t *testing.T) {
	mockRepository := &mock.MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}

	teamService := NewTeamService(mockRepository)
	teams, err := teamService.GetTeams()

	assert.NoError(t, err)
	if len(teams) != 1 {
		t.Errorf("expected %d teams got %d", 1, len(teams))
	}
}

func TestGetTeam(t *testing.T) {
	mockRepository := &mock.MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}

	teamService := NewTeamService(mockRepository)
	team, err := teamService.GetTeam("team-sapphire")

	assert.NoError(t, err)
	assert.Equal(t, "team-sapphire", team.TeamID)

	team, err = teamService.GetTeam("team-jade")
	assert.Error(t, err)
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

	err = teamService.AddTeam(newTeam)
	assert.Error(t, err)

}

func TestUpdateTeam(t *testing.T) {
	mockRepository := &mock.MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}

	teamService := NewTeamService(mockRepository)
	teamOne := mockRepository.Teams[0]
	teamOne.TeamDescription = "Sapphire frontend team updated"
	err := teamService.UpdateTeam(teamOne)
	assert.NoError(t, err)

	teamOne.TeamID = "team-jade"
	err = teamService.UpdateTeam(teamOne)
	assert.Error(t, err)
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
	expectedErrorMsg := fmt.Sprintf("team %s is not requested for deletion", teamID)
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
