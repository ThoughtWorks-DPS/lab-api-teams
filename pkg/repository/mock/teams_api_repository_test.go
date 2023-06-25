package mock

import (
	"testing"

	"twdps.io/lab-api-teams/pkg/domain"
	"twdps.io/lab-api-teams/pkg/service"
)

func TestGetTeams(t *testing.T) {
	mockRepository := &MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}

	teamService := service.NewTeamService(mockRepository)
	teams, _ := teamService.GetTeams()
	if len(teams) != 1 {
		t.Errorf("expected %d teams got %d", 1, len(teams))
	}
}

func TestAddTeam(t *testing.T) {
	mockRepository := &MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}

	teamService := service.NewTeamService(mockRepository)
	var newTeam = domain.Team{
		TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24,
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
