package mock

import (
	"fmt"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
)

type MockNamespaceRepository struct {
	Namespaces []domain.Namespace
}

func (m *MockNamespaceRepository) GetNamespaces() ([]domain.Namespace, error) {
	return m.Namespaces, nil
}

func (m *MockNamespaceRepository) GetNamespacesByFilterWithPagination(filters map[string]interface{}, page int, maxResult int) ([]domain.Namespace, error) {
	return nil, fmt.Errorf("not implemented!")
}

func (m *MockNamespaceRepository) GetNamespace(id string) (domain.Namespace, error) {
	for _, namespace := range m.Namespaces {
		if id == namespace.NamespaceID {
			return namespace, nil
		}
	}

	return domain.Namespace{}, fmt.Errorf("namespace not found")
}

func (m *MockNamespaceRepository) GetNamespacesByType(nsType string) ([]domain.Namespace, error) {
	var namespaces []domain.Namespace
	for _, namespace := range m.Namespaces {
		if nsType == namespace.NamespaceType {
			namespaces = append(namespaces, namespace)
			return namespaces, nil
		}
	}

	return nil, fmt.Errorf("namespace not found")
}

func (m *MockNamespaceRepository) AddNamespace(newNamespace domain.Namespace) error {
	t, err := m.GetNamespace(newNamespace.NamespaceID)
	if err == nil {
		return fmt.Errorf("team already exists: %v", t)
	}

	m.Namespaces = append(m.Namespaces, newNamespace)
	return nil
}
