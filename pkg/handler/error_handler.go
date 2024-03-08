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

		if len(c.Errors) < 1 {
			return
		}

		err := c.Errors[0].Err

		errorResponseBody := NewErrorResponseBody(getErrorCode(err), err.Error())

		switch err.(type) {
		case *service.InvalidPageError, *service.InvalidMaxResultsError:
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

func getErrorCode(err error) string {
	errorType := reflect.TypeOf(err)
	if errorType.Kind() == reflect.Ptr {
		return errorType.Elem().Name()
	} else {
		return errorType.Name()
	}
}

type ErrorResponseBody struct {
	Code        string
	Description string
}

func NewErrorResponseBody(code string, desc string) ErrorResponseBody {
	return ErrorResponseBody{
		Code:        code,
		Description: desc,
	}
}
