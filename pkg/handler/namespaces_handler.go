package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/service"
	"github.com/gin-gonic/gin"
)

type NamespaceHandler struct {
	namespaceService service.NamespaceService
}

type NamespaceQueryResult struct {
	Items     []domain.Namespace
	Page      int
	MaxResult int
}

func NewNamespaceHandler(namespaceService service.NamespaceService) *NamespaceHandler {
	return &NamespaceHandler{namespaceService: namespaceService}
}

// /platform/namespaces?filters[team]=marketplace-demo&filters[type]=standard&page=0&maxResults=25
func (handler *NamespaceHandler) GetNamespaces(c *gin.Context) {

	namespaceQuery := service.Query{
		Filters:   make(map[string]string),
		Page:      0,  // should set page to 0 if page is not provided
		MaxResult: 25, // should set maxResult to 25 if maxResult is not provided
	}

	// should return namespaces based on filters
	filters, exist := c.GetQueryMap("filters")
	if exist {
		namespaceQuery.Filters = filters
	}

	page, exist := c.GetQuery("page")
	if exist {
		pageInt, err := strconv.Atoi(page)
		if err != nil {
			// should return 400 if page value is not a integer
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid page value"})
			return
		}
		namespaceQuery.Page = pageInt
	}

	mapResult, exist := c.GetQuery("maxResults")
	if exist {
		mapResultInt, err := strconv.Atoi(mapResult)

		if err != nil {
			// should return 400 if maxResult value is not a integer
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid maxResult value"})
			return
		}
		// should set maxResult to 25 if maxResult is greatedr than 25
		if mapResultInt < namespaceQuery.MaxResult {
			namespaceQuery.MaxResult = mapResultInt
		}
	}

	namespaces, err := handler.namespaceService.GetNamespacesByFilterWithPagination(namespaceQuery)

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
