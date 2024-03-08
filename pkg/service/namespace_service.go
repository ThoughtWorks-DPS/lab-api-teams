package service

import (
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/repository"
)

const (
	NAMESACE_TYPE_MASTER    = "master"
	NAMESPACE_TYPE_STANDARD = "standard"
	NAMESACE_TYPE_CUSTOM    = "custom"
)

type NamespaceService interface {
	GetNamespace(namespaceID string) (*domain.Namespace, error)
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

func (s *namespaceServiceImpl) GetNamespace(namespaceID string) (*domain.Namespace, error) {
	namespace, err := s.repo.GetNamespace(namespaceID)
	if err != nil {
		return nil, NewResourceNotExistError()
	}
	return &namespace, nil
}

func (s *namespaceServiceImpl) GetNamespaces() ([]domain.Namespace, error) {
	namespaces, err := s.repo.GetNamespaces()
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (s *namespaceServiceImpl) GetNamespacesByFilterWithPagination(query Query) (*ListNamespaceResponse, error) {
	if query.Page < 1 {
		return nil, NewInvalidPageError()
	}

	if query.MaxResults < -1 || query.MaxResults == 0 || query.MaxResults > MAX_RESULTS {
		return nil, InvalidMaxResultsError{"maxResults value is invalid"}
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
	if ns, _ := s.repo.GetNamespace(namespace.NamespaceID); ns.NamespaceID != "" {
		return NewResourceAlreadyExistError()
	}
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
