package service

import (
	"errors"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
)

type NamespaceQuery struct {
	Filters   map[string]string
	Page      int
	MaxResult int
}

type InvalidPageError struct {
	Err error
}

func (e *InvalidPageError) Error() string {
	return e.Err.Error()
}


type NamespaceService interface {
	GetNamespaces() ([]domain.Namespace, error)
	AddNamespace(ns domain.Namespace) error
	GetNamespacesMaster() ([]domain.Namespace, error)
	GetNamespacesStandard() ([]domain.Namespace, error)
	GetNamespacesCustom() ([]domain.Namespace, error)
	GetNamespacesByFilterWithPagination(query NamespaceQuery) ([]domain.Namespace, error)
}

type namespaceServiceImpl struct {
	repo domain.NamespaceRepository
}

func NewNamespaceService(repo domain.NamespaceRepository) NamespaceService {
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

func (s *namespaceServiceImpl) GetNamespacesByFilterWithPagination(query NamespaceQuery) ([]domain.Namespace, error){
	// [] should return error if filter key is not valid
	// [] should return error if filter value is not valid
	if query.Page < 0 {
		return nil, &InvalidPageError{Err: errors.New("page value is invalid")}
	}

	if query.MaxResult < -1 || query.MaxResult == 0 {
		return nil, &InvalidPageError{Err: errors.New("maxResult value is invalid")}
	}

	filter := &domain.Namespace{}

	namespaces, err := s.repo.GetNamespacesByFilterWithPagination(filter, query.Page, query.MaxResult)

	// namespaces, err := s.repo.GetNamespaces()

	if err != nil {
		return nil, err
	}

	return namespaces, nil

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
