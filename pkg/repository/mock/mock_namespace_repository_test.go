package mock

import (
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/service"
	"testing"
)

func TestMockNamespaceREpository(t *testing.T) {
	mockRepository := &MockNamespaceRepository{
		Namespaces: []domain.Namespace{
			{NamespaceType: "master", NamespaceTeamID: "team1", NamespaceID: "dev", NamespaceRam: 1, NamespaceCpu: 1, NamespaceInMesh: true, NamespaceFromDefault: true},
		},
	}

	// If Mock Repo didn't implement expected functions this will throw an InvalidIFaceAssign error
	// - This is itself a test, albeit it is basically a compiler error test
	namespaceService := service.NewNamespaceService(mockRepository)

	// Sanity check that mock implements actual functionality. Rest of service tests are in service_test.go
	namespaces, _ := namespaceService.GetNamespaces()
	if len(namespaces) != 1 {
		t.Errorf("expected %d teams got %d", 1, len(namespaces))
	}
}
