package service

import (
	"fmt"

	domain "twdps.io/lab-api-teams/pkg/domain"
)

type TeamService interface {
	GetTeam(teamID string) (domain.Team, error)
	GetTeams() ([]domain.Team, error)
	AddTeam(team domain.Team) error
	RequestRemoveTeam(teamID string) error
	ConfirmRemoveTeam(teamID string) error
	DatabaseAvailable() (bool, error)
}

type teamServiceImpl struct {
	repo domain.TeamRepository
}

func NewTeamService(repo domain.TeamRepository) TeamService {
	return &teamServiceImpl{
		repo: repo,
	}
}

func (s *teamServiceImpl) DatabaseAvailable() (bool, error) {
	dbAvailable, err := s.repo.DatabaseAvailable()
	if err != nil {
		return false, err
	}

	if dbAvailable == false {
		return false, fmt.Errorf("DB is not available yet")
	}

	return true, nil
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

// TODO - SYNC ALL? Teams, or just THAT team
func (s *teamServiceImpl) UpdateTeam(team domain.Team) error {
	err := s.repo.UpdateTeam(team)
	if err != nil {
		return err // TODO transient/status
	}

	return nil
}

func (s *teamServiceImpl) RequestRemoveTeam(teamID string) error {
	team, err := s.repo.GetTeam(teamID)
	if err != nil {
		return err // TODO transient/status errors
	}

	team.TeamMarkedForDeletion = "Requested"

	err = s.repo.UpdateTeam(team)
	if err != nil {
		return err // TODO transient
	}

	return nil
}

func (s *teamServiceImpl) ConfirmRemoveTeam(teamID string) error {
	team, err := s.repo.GetTeam(teamID)
	if err != nil {
		return err // TODO transient/status errors
	}

	if team.TeamMarkedForDeletion != "Requested" {
		return fmt.Errorf("Team %s is not requested for deletion", teamID)
	}

	err = s.repo.RemoveTeam(teamID)
	if err != nil {
		return err // TODO transient
	}

	return nil
}
