//go:build !e2e
// +build !e2e

package service

import (
	"testing"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func Test_should_create_namespace(t *testing.T) {
	//given
	//when
	//then
}

func Test_should_get_a_namespace(t *testing.T) {
	//given
	//when
	//then
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
	mockNamespaceRepository := new(mock.MockNamespaceRepositoryV2)
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

func Test_should_return_error_when_page_number_less_than_1(t *testing.T) {
	//given
	query := Query{
		Filters:    make(map[string]interface{}),
		Page:       0,
		MaxResults: 10,
	}

	mockNamespaceRepository := new(mock.MockNamespaceRepositoryV2)
	mockNamespaceRepository.On("GetNamespacesByFilterWithPagination", query.Filters, query.Page, query.MaxResults).Return([]domain.Namespace{}, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespacesByFilterWithPagination(query)

	//then
	mockNamespaceRepository.AssertNumberOfCalls(t, "GetNamespacesByFilterWithPagination", 0)
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.EqualError(t, err, NewInvalidPageError().Error())
}

func Test_should_return_error_when_maxResult_less_than_minus_1(t *testing.T) {
	//given
	query := Query{
		Filters:    make(map[string]interface{}),
		Page:       1,
		MaxResults: 0,
	}

	mockNamespaceRepository := new(mock.MockNamespaceRepositoryV2)
	mockNamespaceRepository.On("GetNamespacesByFilterWithPagination", query.Filters, query.Page, query.MaxResults).Return([]domain.Namespace{}, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespacesByFilterWithPagination(query)

	//then
	mockNamespaceRepository.AssertNumberOfCalls(t, "GetNamespacesByFilterWithPagination", 0)
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.EqualError(t, err, NewInvalidMaxResultsError().Error())
}

func Test_should_return_error_when_maxResult_greater_than_25(t *testing.T) {
	//given
	query := Query{
		Filters:    make(map[string]interface{}),
		Page:       1,
		MaxResults: 26,
	}

	mockNamespaceRepository := new(mock.MockNamespaceRepositoryV2)
	mockNamespaceRepository.On("GetNamespacesByFilterWithPagination", query.Filters, query.Page, query.MaxResults).Return([]domain.Namespace{}, nil)
	namespaceService := NewNamespaceService(mockNamespaceRepository)

	//when
	resp, err := namespaceService.GetNamespacesByFilterWithPagination(query)

	//then
	mockNamespaceRepository.AssertNumberOfCalls(t, "GetNamespacesByFilterWithPagination", 0)
	assert.Nil(t, resp)
	assert.Error(t, err)
	assert.EqualError(t, err, NewInvalidMaxResultsError().Error())
}
