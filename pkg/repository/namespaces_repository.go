package repository

import (
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/datastore"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
)

type NamespaceRepository interface {
	GetNamespaces() ([]domain.Namespace, error)
	GetNamespacesByFilterWithPagination(filters map[string]string, page int, maxResult int) ([]domain.Namespace, error)
	GetNamespacesByType(nsType string) ([]domain.Namespace, error)
	// GetNamespaceByID(namespaceID string) (Namespace, error)
	AddNamespace(namespace domain.Namespace) error
	// UpdateNamespace(namespace Namespace) error
	// RemoveNamespace(namespace Namespace) (Namespace, error)
}

type NamespaceRepositoryImpl struct {
	datastore datastore.Datastore
}

func NewNamespaceRepository(store datastore.Datastore) NamespaceRepository {
	return &NamespaceRepositoryImpl{datastore: store}
}

func (repo *NamespaceRepositoryImpl) GetNamespaces() ([]domain.Namespace, error) {
	var namespaces []domain.Namespace
	err := repo.datastore.ReadAll(&namespaces)
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (repo *NamespaceRepositoryImpl) AddNamespace(namespace domain.Namespace) error {
	return repo.datastore.Create(&namespace)
}

func (repo *NamespaceRepositoryImpl) GetNamespacesByFilterWithPagination(filters map[string]string, page int, maxResult int) ([]domain.Namespace, error) {
	var namespaces []domain.Namespace

	// filter :=
	// filter by 
	err := repo.datastore.ReadByAttributesWithPagination(nil, &namespaces, page, maxResult)

	if err != nil {
		return nil, err
	}
	return namespaces, nil
}

func (repo *NamespaceRepositoryImpl) GetNamespacesByType(nsType string) ([]domain.Namespace, error) {
	// This is a bit tricky. Since querying by type is specific to Namespace and how you'd handle this
	// differs between Redis and SQL databases, this might require custom logic within the underlying Datadatastore.
	// For now, let's mock it up and you can implement the specifics.
	var namespaces []domain.Namespace
	filter := map[string]interface{}{
		"NamespaceType": nsType,
	}

	err := repo.datastore.ReadByAttributes(filter, &namespaces)
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (repo *NamespaceRepositoryImpl) GetNamespaceByID(namespaceID string) (domain.Namespace, error) {
	var namespace domain.Namespace
	err := repo.datastore.ReadByID(namespaceID, &namespace)
	if err != nil {
		return domain.Namespace{}, err
	}

	return namespace, nil
}

func (repo *NamespaceRepositoryImpl) UpdateNamespace(namespace domain.Namespace) error {
	return repo.datastore.Update(&namespace)
}

func (repo *NamespaceRepositoryImpl) RemoveNamespace(namespace domain.Namespace) error {
	return repo.datastore.Delete(&namespace)
}
