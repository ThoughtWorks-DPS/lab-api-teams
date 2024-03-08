package handler

import (
	"net/http"
	"reflect"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/service"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		err := c.Errors[0].Err
		errorResponseBody := NewErrorResponseBody("Error happens. ", ErrorItem{Title: reflect.TypeOf(err).Name(), Message: err.Error()})
		switch err.(type) {
		case *service.InvalidPageError:
			c.IndentedJSON(http.StatusBadRequest, errorResponseBody)
		case *service.InvalidMaxResultsError:
			c.IndentedJSON(http.StatusBadRequest, errorResponseBody)
		case *service.ResourceAlreadyExistError:
			c.IndentedJSON(http.StatusConflict, errorResponseBody)
		case *service.ResourceNotExistError:
			c.IndentedJSON(http.StatusNotFound, errorResponseBody)
		default:
			c.IndentedJSON(http.StatusInternalServerError, errorResponseBody)
		}
	}
}

type ErrorItem struct {
	Title   string
	Message string
}

type ErrorResponseBody struct {
	Description string
	Errors      []ErrorItem
}

func NewErrorResponseBody(desc string, errorItem ...ErrorItem) ErrorResponseBody {
	result := ErrorResponseBody{Description: desc}
	result.Errors = append(result.Errors, errorItem...)
	return result
}
