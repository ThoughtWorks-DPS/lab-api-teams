package service

import (
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
)

type NamespaceService interface {
	GetNamespaces() ([]domain.Namespace, error)
	AddNamespace(ns domain.Namespace) error
	GetNamespacesMaster() ([]domain.Namespace, error)
	GetNamespacesStandard() ([]domain.Namespace, error)
	GetNamespacesCustom() ([]domain.Namespace, error)
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
