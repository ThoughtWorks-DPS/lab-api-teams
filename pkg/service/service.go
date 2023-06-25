package service

import (
	"fmt"

	domain "twdps.io/lab-api-teams/pkg/domain"
)

type TeamService interface {
	GetTeam(teamID string) (domain.Team, error)
	GetTeams() ([]domain.Team, error)
	AddTeam(team domain.Team) error
}

type teamServiceImpl struct {
	repo domain.TeamRepository
}

func NewTeamService(repo domain.TeamRepository) TeamService {
	return &teamServiceImpl{
		repo: repo,
	}
}

func (s *teamServiceImpl) GetTeams() ([]domain.Team, error) {
	teams, err := s.repo.GetTeams()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch teams: %v", err)
	}
	return teams, err
}

func (s *teamServiceImpl) AddTeam(newTeam domain.Team) error {
	if err := s.repo.AddTeam(newTeam); err != nil {
		return fmt.Errorf("could not add team: %v", err)
	}
	return nil
}

func (s *teamServiceImpl) GetTeam(teamID string) (domain.Team, error) {
	team, err := s.repo.GetTeam(teamID)
	if err != nil {
		return domain.Team{}, err // TODO transient/status errors
	}

	return team, nil
}
