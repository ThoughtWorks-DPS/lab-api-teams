package service

import (
	"errors"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/repository"
)

const NAMESACE_TYPE_MASTER = "master"
const NAMESPACE_TYPE_STANDARD = "standard"
const NAMESACE_TYPE_CUSTOM = "custom"

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
	namespaces, err := s.repo.GetNamespacesByType(NAMESPACE_TYPE_STANDARD)
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (s *namespaceServiceImpl) GetNamespacesCustom() ([]domain.Namespace, error) {
	namespaces, err := s.repo.GetNamespacesByType(NAMESACE_TYPE_CUSTOM)
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (s *namespaceServiceImpl) GetNamespacesMaster() ([]domain.Namespace, error) {
	namespaces, err := s.repo.GetNamespacesByType(NAMESACE_TYPE_MASTER)
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
