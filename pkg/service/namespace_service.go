package service

import (
	"errors"
	"slices"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/repository"
)

type NamespaceService interface {
	GetNamespaces() ([]domain.Namespace, error)
	AddNamespace(ns domain.Namespace) error
	GetNamespacesMaster() ([]domain.Namespace, error)
	GetNamespacesStandard() ([]domain.Namespace, error)
	GetNamespacesCustom() ([]domain.Namespace, error)
	GetNamespacesByFilterWithPagination(query Query) (*ListNamespaceResponse, error)
}

type namespaceServiceImpl struct {
	repo repository.NamespaceRepository
}

func NewNamespaceService(repo repository.NamespaceRepository) NamespaceService {
	return &namespaceServiceImpl{
		repo: repo,
	}
}

func (s *namespaceServiceImpl) GetNamespaces() ([]domain.Namespace, error) {
	namespaces, err := s.repo.GetNamespaces()

	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (s *namespaceServiceImpl) GetNamespacesByFilterWithPagination(query Query) (*ListNamespaceResponse, error) {

	if query.Page < 0 {
		return nil, &InvalidPageError{Err: errors.New("page value is invalid")}
	}

	if query.MaxResults < -1 || query.MaxResults == 0 {
		return nil, &InvalidPageError{Err: errors.New("maxResults value is invalid")}
	}

	legalFilters := []string{"team", "type"}
	illegalFilters := []string{}

	for k := range query.Filters {
		contains := slices.Contains(legalFilters, k)
		if contains {
			illegalFilters = append(illegalFilters, k)
		}
	}
	if len(illegalFilters) > 0 {
		return nil, &InvalidFilterError{Err: errors.New("invalid filter. only allow filter by team or type")}
	}

	namespaces, err := s.repo.GetNamespacesByFilterWithPagination(query.Filters, query.Page, query.MaxResults)

	if err != nil {
		return nil, err
	}

	return &ListNamespaceResponse{
		Items:      namespaces,
		Page:       query.Page,
		MaxResults: query.MaxResults,
	}, nil

}

func (s *namespaceServiceImpl) GetNamespacesStandard() ([]domain.Namespace, error) {
	namespaces, err := s.repo.GetNamespacesByType("standard")
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (s *namespaceServiceImpl) GetNamespacesCustom() ([]domain.Namespace, error) {
	namespaces, err := s.repo.GetNamespacesByType("custom")
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (s *namespaceServiceImpl) GetNamespacesMaster() ([]domain.Namespace, error) {
	namespaces, err := s.repo.GetNamespacesByType("master")
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (s *namespaceServiceImpl) AddNamespace(namespace domain.Namespace) error {
	if err := s.repo.AddNamespace(namespace); err != nil {
		return err
	}
	return nil
}

// func (s *namespaceServiceImpl) UpdateNamespace(namespace domain.Namespace) error {
// 	if err := s.repo.UpdateNamespace(namespace); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (s *namespaceServiceImpl) GetTeamNamespacesByTeamID(teamID string) ([]domain.Namespace, error) {
// 	ns, err := s.repo.GetNamespacesByAttribute("namespaceTeamID", teamID)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return ns, nil
// }

// func (s *namespaceServiceImpl) GetTeamNamespaceByID(nsID string) (domain.Namespace, error) {
// 	ns, err := s.repo.GetNamespaceByID(nsID)
// 	if err != nil {
// 		return domain.Namespace{}, err
// 	}

// 	return ns, nil
// }
