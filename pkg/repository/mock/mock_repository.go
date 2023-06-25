package mock

import (
	"fmt"

	"twdps.io/lab-api-teams/pkg/domain"
)

type MockRepository struct {
	Teams []domain.Team
}

func (m *MockRepository) GetTeams() ([]domain.Team, error) {
	return m.Teams, nil
}

func (m *MockRepository) AddTeam(newTeam domain.Team) error {
	m.Teams = append(m.Teams, newTeam)
	return nil
}

func (m *MockRepository) GetTeam(id string) (domain.Team, error) {
	for _, team := range m.Teams {
		if id == team.TeamID {
			return team, nil
		}
	}

	return domain.Team{}, fmt.Errorf("Team not found")
}
