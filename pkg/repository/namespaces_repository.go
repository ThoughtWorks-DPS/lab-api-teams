package repository

import (
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/datastore"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
)

type NamespaceRepository struct {
	datastore datastore.Datastore
}

func NewNamespaceRepository(store datastore.Datastore) *NamespaceRepository {
	return &NamespaceRepository{datastore: store}
}

func (repo *NamespaceRepository) GetNamespaces() ([]domain.Namespace, error) {
	var namespaces []domain.Namespace
	err := repo.datastore.ReadAll(&namespaces)
	if err != nil {
		return nil, err
	}

	return namespaces, nil
}

func (repo *NamespaceRepository) AddNamespace(namespace domain.Namespace) error {
	return repo.datastore.Create(&namespace)
}

func (repo *NamespaceRepository) GetNamespacesByQuery() ([]domain.Namespace, error) {
	var namespaces []domain.Namespace

	// filter := 
	err := repo.datastore.ReadByAttributesWithPagination(nil, &namespaces)

	if err != nil {
		return nil, err
	}
	return namespaces, nil
}

func (repo *NamespaceRepository) GetNamespacesByType(nsType string) ([]domain.Namespace, error) {
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

func (repo *NamespaceRepository) GetNamespaceByID(namespaceID string) (domain.Namespace, error) {
	var namespace domain.Namespace
	err := repo.datastore.ReadByID(namespaceID, &namespace)
	if err != nil {
		return domain.Namespace{}, err
	}

	return namespace, nil
}

func (repo *NamespaceRepository) UpdateNamespace(namespace domain.Namespace) error {
	return repo.datastore.Update(&namespace)
}

func (repo *NamespaceRepository) RemoveNamespace(namespace domain.Namespace) error {
	return repo.datastore.Delete(&namespace)
}
