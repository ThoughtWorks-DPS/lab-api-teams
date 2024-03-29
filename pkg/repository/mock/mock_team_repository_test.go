//go:build !e2e
// +build !e2e

package mock

import (
	"testing"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/service"
)

func TestMock(t *testing.T) {
	mockRepository := &MockRepository{
		Teams: []domain.Team{
			{TeamID: "team-sapphire", TeamType: "normal", TeamDescription: "Sapphire frontend team", TeamRAM: 32, TeamCPU: 12, TeamRamLimit: 64, TeamCpuLimit: 24},
		},
	}

	// If Mock Repo didn't implement expected functions this will throw an InvalidIFaceAssign error
	// - This is itself a test, albeit it is basically a compiler error test
	teamService := service.NewTeamService(mockRepository)

	// Sanity check that mock implements actual functionality. Rest of service tests are in service_test.go
	teams, _ := teamService.GetTeams()
	if len(teams) != 1 {
		t.Errorf("expected %d teams got %d", 1, len(teams))
	}
}
