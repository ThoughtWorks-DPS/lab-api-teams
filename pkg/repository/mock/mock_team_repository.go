package mock

import (
	"fmt"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
)

type MockRepository struct {
	Teams []domain.Team
}

func (m *MockRepository) GetTeams() ([]domain.Team, error) {
	return m.Teams, nil
}

func (m *MockRepository) AddTeam(newTeam domain.Team) error {
	t, err := m.GetTeam(newTeam.TeamID)
	if err == nil {
		return fmt.Errorf("team already exists: %v", t)
	}

	m.Teams = append(m.Teams, newTeam)
	return nil
}

func (m *MockRepository) GetTeam(id string) (domain.Team, error) {
	for _, team := range m.Teams {
		if id == team.TeamID {
			return team, nil
		}
	}

	return domain.Team{}, fmt.Errorf("team not found")
}

func update_slice_item(slice []domain.Team, index int, team domain.Team) []domain.Team {
	// Remove the index first
	teams := append(slice[:index], slice[index+1:]...)
	// Update with the new one
	return append(teams, team)
}

func delete_slice_item(slice []domain.Team, index int) []domain.Team {
	return append(slice[:index], slice[index+1:]...)
}

func (m *MockRepository) RemoveTeam(teamID string) error {
	for index, team := range m.Teams {
		if team.TeamID == teamID {
			m.Teams = delete_slice_item(m.Teams, index)
		}
	}

	return nil
}

func (m *MockRepository) UpdateTeam(teamToUpdate domain.Team) error {
	for index, team := range m.Teams {
		if team.TeamID == teamToUpdate.TeamID {
			m.Teams = update_slice_item(m.Teams, index, teamToUpdate)
		}
	}

	return nil
}

func (m *MockRepository) DatabaseAvailable() (bool, error) {
	return true, nil
}
