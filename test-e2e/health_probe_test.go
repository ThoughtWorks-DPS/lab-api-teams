//go:build e2e
// +build e2e

package e2e

import (
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestE2EYourAPI(t *testing.T) {

	// Get API URL from environment variable or use default
	apiURL := os.Getenv("TEAMS_API_URL")
	if apiURL == "" {
		apiURL = "http://localhost:8080" // default value
	}

	e := httpexpect.Default(t, apiURL)

	// Simple test: check liveness
	e.GET("/teams/healthz/liveness").
		Expect().
		Status(http.StatusOK).
		JSON().Object().HasValue("message", "Teams api is live")

	// Simple test: check readiness
	e.GET("/teams/healthz/readiness").
		Expect().
		Status(http.StatusOK).
		JSON().Object().HasValue("message", "DB is available")

	// TODO
	// Get teams
	// Post team
	// Delete team
	// PUT team

}
