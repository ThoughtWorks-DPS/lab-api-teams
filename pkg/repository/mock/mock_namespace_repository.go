package mock

import (
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/stretchr/testify/mock"
)

type MockNamespaceRepositoryV2 struct {
	mock.Mock
}

func (m *MockNamespaceRepositoryV2) GetNamespaces() ([]domain.Namespace, error) {
	args := m.Called()
	return args.Get(0).([]domain.Namespace), nil
}

func (m *MockNamespaceRepositoryV2) GetNamespacesByFilterWithPagination(filters map[string]interface{}, page int, maxResults int) ([]domain.Namespace, error) {
	args := m.Called(filters, page, maxResults)
	return args.Get(0).([]domain.Namespace), nil
}

func (m *MockNamespaceRepositoryV2) GetNamespacesByType(nsType string) ([]domain.Namespace, error) {
	args := m.Called(nsType)
	return args.Get(0).([]domain.Namespace), nil
}

func (m *MockNamespaceRepositoryV2) AddNamespace(namespace domain.Namespace) error {
	m.Called(namespace)
	return nil
}
