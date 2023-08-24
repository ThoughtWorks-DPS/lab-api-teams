package repository

import (
	"github.com/RBMarketplace/di-api-teams/pkg/datastore"
	"github.com/RBMarketplace/di-api-teams/pkg/domain"
)

type TeamRepository interface {
	GetTeams() ([]domain.Team, error)
	GetTeam(id string) (domain.Team, error) // GetTeamByTeamID
	AddTeam(newTeam domain.Team) error
	UpdateTeam(team domain.Team) error
	RemoveTeam(id string) error
	DatabaseAvailable() (bool, error)
}

type TeamRepositoryImpl struct {
	store datastore.Datastore
}

func NewTeamsRepo(store datastore.Datastore) TeamRepository {
	return &TeamRepositoryImpl{store: store}
}

func (r *TeamRepositoryImpl) AddTeam(team domain.Team) error {
	return r.store.Create(team)
}

func (r *TeamRepositoryImpl) GetTeam(id string) (domain.Team, error) {
	var team domain.Team
	err := r.store.ReadByID(id, &team)
	if err != nil {
		return domain.Team{}, err
	}
	return team, nil
}

func (r *TeamRepositoryImpl) GetTeams() ([]domain.Team, error) {
	var teams []domain.Team
	err := r.store.ReadAll(&teams)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (r *TeamRepositoryImpl) UpdateTeam(team domain.Team) error {
	return r.store.Update(team)
}

func (r *TeamRepositoryImpl) RemoveTeam(teamID string) error {
	var team domain.Team
	err := r.store.ReadByID(teamID, &team)
	if err != nil {
		return err // TODO transient
	}

	return r.store.Delete(team)
}

func (r *TeamRepositoryImpl) DatabaseAvailable() (bool, error) {
	dbAvailable, err := r.store.IsDatabaseAvailable()
	if err != nil {
		return false, err
	}
	return dbAvailable, nil
}
