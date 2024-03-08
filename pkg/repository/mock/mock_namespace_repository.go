package mock

import (
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/stretchr/testify/mock"
)

type MockNamespaceRepository struct {
	mock.Mock
}

func (m *MockNamespaceRepository) GetNamespaces() ([]domain.Namespace, error) {
	args := m.Called()
	return args.Get(0).([]domain.Namespace), nil
}

func (m *MockNamespaceRepository) GetNamespacesByFilterWithPagination(filters map[string]interface{}, page int, maxResults int) ([]domain.Namespace, error) {
	args := m.Called(filters, page, maxResults)
	return args.Get(0).([]domain.Namespace), nil
}

func (m *MockNamespaceRepository) GetNamespacesByType(nsType string) ([]domain.Namespace, error) {
	args := m.Called(nsType)
	return args.Get(0).([]domain.Namespace), nil
}

func (m *MockNamespaceRepository) AddNamespace(namespace domain.Namespace) error {
	args := m.Called(namespace)
	result := args.Get(0)
	if result == nil {
		return nil
	} else {
		return result.(error)
	}
}

func (m *MockNamespaceRepository) GetNamespace(namespaceID string) (domain.Namespace, error) {
	args := m.Called(namespaceID)
	if args.Get(1) == nil {
		return args.Get(0).(domain.Namespace), nil
	} else {
		return args.Get(0).(domain.Namespace), args.Get(1).(error)
	}
}
