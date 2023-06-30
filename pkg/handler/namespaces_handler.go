package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"twdps.io/lab-api-teams/pkg/domain"
	"twdps.io/lab-api-teams/pkg/service"
)

type NamespaceHandler struct {
	namespaceService service.NamespaceService
}

func NewNamespaceHandler(namespaceService service.NamespaceService) *NamespaceHandler {
	return &NamespaceHandler{namespaceService: namespaceService}
}

func (handler *NamespaceHandler) GetNamespaces(c *gin.Context) {
	namespaces, err := handler.namespaceService.GetNamespaces()
	if err != nil {
		log.Fatalf("Failed to call GetNamespaces %v", err)
	}

	c.IndentedJSON(http.StatusOK, namespaces)
}

func (handler *NamespaceHandler) AddNamespace(c *gin.Context) {
	var newNamespace domain.Namespace

	if err := c.BindJSON(&newNamespace); err != nil {
		log.Printf("error %+v", err)
		return
	}

	err := handler.namespaceService.AddNamespace(newNamespace)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, err)
	}

	c.IndentedJSON(http.StatusCreated, newNamespace)
}

func (handler *NamespaceHandler) GetNamespacesMaster(c *gin.Context) {
	namespaces, err := handler.namespaceService.GetNamespacesMaster()

	if err != nil {
		log.Fatalf("Failed to call get namespaces master: %v", err)
	}

	c.IndentedJSON(http.StatusOK, namespaces)
}

func (handler *NamespaceHandler) GetNamespacesStandard(c *gin.Context) {
	namespaces, err := handler.namespaceService.GetNamespacesStandard()

	if err != nil {
		log.Fatalf("Failed to call get namespaces master: %v", err)
	}

	c.IndentedJSON(http.StatusOK, namespaces)
}

func (handler *NamespaceHandler) GetNamespacesCustom(c *gin.Context) {
	namespaces, err := handler.namespaceService.GetNamespacesCustom()

	if err != nil {
		log.Fatalf("Failed to call get namespaces master: %v", err)
	}

	c.IndentedJSON(http.StatusOK, namespaces)
}
