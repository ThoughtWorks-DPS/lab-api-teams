package handler

import (
	"net/http"
	"strconv"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/service"
	"github.com/gin-gonic/gin"
)

type NamespaceHandler struct {
	namespaceService service.NamespaceService
}

func NewNamespaceHandler(namespaceService service.NamespaceService) *NamespaceHandler {
	return &NamespaceHandler{namespaceService: namespaceService}
}

func (handler *NamespaceHandler) GetNamespace(c *gin.Context) {
	namespaceID := c.Param("namespaceID")
	namespace, err := handler.namespaceService.GetNamespace(namespaceID)
	if err != nil {
    _ = c.AbortWithError(-1, err)
		return
	}
	c.IndentedJSON(http.StatusOK, namespace)
}

func (handler *NamespaceHandler) GetNamespaces(c *gin.Context) {
	namespaceQuery := service.Query{
		Page:       1,  // should set page to 1 if page is not provided
		MaxResults: 25, // should set maxResults to 25 if maxResults is not provided
	}

	// should return namespaces based on filters
	filters, exist := c.GetQueryMap("filters")
	if exist {
		f := make(map[string]interface{})
		for k, v := range filters {
			f[k] = v
		}
		namespaceQuery.Filters = f
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
			// should return 400 if maxResults value is not a integer
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid maxResults value"})
			return
		}
		// should set maxResults to 25 if maxResults is greatedr than 25
		if mapResultInt < namespaceQuery.MaxResults {
			namespaceQuery.MaxResults = mapResultInt
		}
	}

	resp, err := handler.namespaceService.GetNamespacesByFilterWithPagination(namespaceQuery)
	if err != nil {
		_ = c.AbortWithError(-1, err)
		return
	}

	c.IndentedJSON(http.StatusOK, resp)
}

func (handler *NamespaceHandler) AddNamespace(c *gin.Context) {
	var newNamespace domain.Namespace

	if err := c.BindJSON(&newNamespace); err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := handler.namespaceService.AddNamespace(newNamespace)
	if err != nil {
		_ = c.AbortWithError(-1, err)
		return
	}

	c.IndentedJSON(http.StatusCreated, newNamespace)
}

func (handler *NamespaceHandler) GetNamespacesMaster(c *gin.Context) {
	namespaces, err := handler.namespaceService.GetNamespacesMaster()
	if err != nil {
		_ = c.AbortWithError(-1, err)
		return
	}

	c.IndentedJSON(http.StatusOK, namespaces)
}

func (handler *NamespaceHandler) GetNamespacesStandard(c *gin.Context) {
	namespaces, err := handler.namespaceService.GetNamespacesStandard()
	if err != nil {
		_ = c.AbortWithError(-1, err)
		return
	}

	c.IndentedJSON(http.StatusOK, namespaces)
}

func (handler *NamespaceHandler) GetNamespacesCustom(c *gin.Context) {
	namespaces, err := handler.namespaceService.GetNamespacesCustom()
	if err != nil {
		_ = c.AbortWithError(-1, err)
		return
	}

	c.IndentedJSON(http.StatusOK, namespaces)
}
