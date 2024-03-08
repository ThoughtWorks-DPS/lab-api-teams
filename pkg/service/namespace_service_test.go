//go:build !e2e
// +build !e2e

package service

import (
	"errors"
	"testing"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func Test_should_add_namespace(t *testing.T) {
	//given
	namespace := domain.Namespace{
		NamespaceType: "normal", NamespaceTeamID: "foo", NamespaceID: "bar-dev", NamespaceRam: 1, NamespaceCpu: 1, NamespaceInMesh: true, NamespaceFromDefault: true,
	}
	mockNamespaceRepository := new(mock.MockNamespaceRepository)
	mockNamespaceRepository.On("GetNamespace", namespace.NamespaceID).Return(domain.Namespace{}, nil)
	mockNamespaceRepository.On("AddNamespace", namespace).Return(nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	err := namespaceService.AddNamespace(namespace)

	//then
	mockNamespaceRepository.AssertExpectations(t)
	assert.NoError(t, err)
}

func Test_should_return_error_given_namespace_exists(t *testing.T) {
	//given
	namespace := domain.Namespace{
		NamespaceType: "normal", NamespaceTeamID: "foo", NamespaceID: "bar-dev", NamespaceRam: 1, NamespaceCpu: 1, NamespaceInMesh: true, NamespaceFromDefault: true,
	}
	mockNamespaceRepository := new(mock.MockNamespaceRepository)
	mockNamespaceRepository.On("GetNamespace", namespace.NamespaceID).Return(namespace, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	err := namespaceService.AddNamespace(namespace)

	//then
	mockNamespaceRepository.AssertExpectations(t)
	assert.Error(t, err)
	assert.Equal(t, NewResourceAlreadyExistError(), err)
}

func Test_should_return_error_given_team_not_exists(t *testing.T) {

}

func Test_should_return_error_given_team_is_marked_for_deletion(t *testing.T) {

}

func Test_should_get_a_namespace(t *testing.T) {
	//given
	namespace := domain.Namespace{
		NamespaceType: "standard", NamespaceTeamID: "foo", NamespaceID: "bar-dev", NamespaceRam: 1, NamespaceCpu: 1, NamespaceInMesh: true, NamespaceFromDefault: true,
	}
	mockNamespaceRepository := new(mock.MockNamespaceRepository)
	mockNamespaceRepository.On("GetNamespace", namespace.NamespaceID).Return(namespace, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespace(namespace.NamespaceID)

	//then
	mockNamespaceRepository.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, &namespace, resp)
}

func Test_should_return_error_given_namespace_not_exists(t *testing.T) {
	//given
	namespaceID := "bar-dev"
	mockNamespaceRepository := new(mock.MockNamespaceRepository)
	mockNamespaceRepository.On("GetNamespace", namespaceID).Return(domain.Namespace{}, errors.New(""))
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespace(namespaceID)

	//then
	mockNamespaceRepository.AssertExpectations(t)
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, NewResourceNotExistError(), err)
}

func Test_should_list_namespaces_without_pagination(t *testing.T) {
	//given
	namespaces := []domain.Namespace{
		{NamespaceType: "master", NamespaceTeamID: "foo-1", NamespaceID: "bar", NamespaceRam: 1, NamespaceCpu: 1, NamespaceInMesh: true, NamespaceFromDefault: true},
		{NamespaceType: "standard", NamespaceTeamID: "foo-2", NamespaceID: "bar-dev", NamespaceRam: 1, NamespaceCpu: 1, NamespaceInMesh: true, NamespaceFromDefault: true},
	}
	mockNamespaceRepository := new(mock.MockNamespaceRepository)
	mockNamespaceRepository.On("GetNamespaces").Return(namespaces, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespaces()

	//then
	mockNamespaceRepository.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, len(namespaces), len(resp))
}

func Test_should_list_namespaces(t *testing.T) {
	//given
	namespaces := []domain.Namespace{
		{NamespaceType: "master", NamespaceTeamID: "foo-1", NamespaceID: "bar", NamespaceRam: 1, NamespaceCpu: 1, NamespaceInMesh: true, NamespaceFromDefault: true},
		{NamespaceType: "standard", NamespaceTeamID: "foo-2", NamespaceID: "bar-dev", NamespaceRam: 1, NamespaceCpu: 1, NamespaceInMesh: true, NamespaceFromDefault: true},
	}
	query := Query{
		Filters:    make(map[string]interface{}),
		Page:       1,
		MaxResults: 10,
	}
	mockNamespaceRepository := new(mock.MockNamespaceRepository)
	mockNamespaceRepository.On("GetNamespacesByFilterWithPagination", query.Filters, query.Page, query.MaxResults).Return(namespaces, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespacesByFilterWithPagination(query)

	//then
	mockNamespaceRepository.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, query.Page, resp.Page)
	assert.Equal(t, query.MaxResults, resp.MaxResults)
	assert.Equal(t, len(namespaces), len(resp.Items))
}

func Test_should_return_error_given_page_number_less_than_1(t *testing.T) {
	//given
	query := Query{
		Filters:    make(map[string]interface{}),
		Page:       0,
		MaxResults: 10,
	}

	mockNamespaceRepository := new(mock.MockNamespaceRepository)
	mockNamespaceRepository.On("GetNamespacesByFilterWithPagination", query.Filters, query.Page, query.MaxResults).Return([]domain.Namespace{}, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespacesByFilterWithPagination(query)

	//then
	mockNamespaceRepository.AssertNumberOfCalls(t, "GetNamespacesByFilterWithPagination", 0)
	assert.Nil(t, resp)
	assert.Error(t, err)
}

func Test_should_return_error_given_maxResult_less_than_minus_1(t *testing.T) {
	//given
	query := Query{
		Filters:    make(map[string]interface{}),
		Page:       1,
		MaxResults: -2,
	}

	mockNamespaceRepository := new(mock.MockNamespaceRepository)
	mockNamespaceRepository.On("GetNamespacesByFilterWithPagination", query.Filters, query.Page, query.MaxResults).Return([]domain.Namespace{}, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespacesByFilterWithPagination(query)

	//then
	mockNamespaceRepository.AssertNumberOfCalls(t, "GetNamespacesByFilterWithPagination", 0)
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, NewInvalidMaxResultsError(), err)
}

func Test_should_return_error_given_maxResult_greater_than_25(t *testing.T) {
	//given
	query := Query{
		Filters:    make(map[string]interface{}),
		Page:       1,
		MaxResults: 26,
	}

	mockNamespaceRepository := new(mock.MockNamespaceRepository)
	mockNamespaceRepository.On("GetNamespacesByFilterWithPagination", query.Filters, query.Page, query.MaxResults).Return([]domain.Namespace{}, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespacesByFilterWithPagination(query)

	//then
	mockNamespaceRepository.AssertNumberOfCalls(t, "GetNamespacesByFilterWithPagination", 0)
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.Equal(t, NewInvalidMaxResultsError(), err)
}
