//go:build e2e
// +build e2e

package e2e

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestE2EHealthProbe(t *testing.T) {

	// Get API URL from environment variable or use default
	apiURL := os.Getenv("TEAMS_API_URL")

	if strings.Contains(apiURL, "prod") {
		t.Log(fmt.Sprintf("Using prod api url: %s", apiURL))
		apiURL = "https://twdps.io"
	} else if apiURL == "" {
		t.Log("Using local api url")
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

}
