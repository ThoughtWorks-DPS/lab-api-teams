//go:build !e2e
// +build !e2e

package service

import (
	"testing"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/repository/mock"
	"github.com/stretchr/testify/assert"
)

func TestGetNamespaces(t *testing.T) {
	mockRepository := &mock.MockNamespaceRepository{
		Namespaces: []domain.Namespace{
			{NamespaceType: "master", NamespaceTeamID: "team1", NamespaceID: "dev", NamespaceRam: 1, NamespaceCpu: 1, NamespaceInMesh: true, NamespaceFromDefault: true},
		},
	}

	namespaceService := NewNamespaceService(mockRepository)
	namespaces, err := namespaceService.GetNamespaces()

	assert.NoError(t, err)
	if len(namespaces) != 1 {
		t.Errorf("expected %d teams got %d", 1, len(namespaces))
	}
}
