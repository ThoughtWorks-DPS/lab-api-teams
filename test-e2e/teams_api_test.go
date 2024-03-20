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
	"github.com/google/uuid"
)

func TestE2ETeamsAPi(t *testing.T) {

	apiURL := os.Getenv("TEAMS_API_URL")

	if strings.Contains(apiURL, "prod") {
		t.Log(fmt.Sprintf("Using prod api url: %s", apiURL))
		apiURL = "https://twdps.io"
	} else if apiURL == "" {
		t.Log("Using local api url")
		apiURL = "http://localhost:8080" // default value
	}

	e := httpexpect.Default(t, apiURL)

	teamId := uuid.New()

	newTeam := map[string]any{
		"teamID":                teamId,
		"teamType":              "normal",
		"teamDescription":       "frontend team",
		"teamRAM":               32,
		"teamCPU":               12,
		"teamRamLimit":          64,
		"teamCPULimit":          24,
		"teamMarkedForDeletion": "",
	}

	e.POST("/teams").
		WithJSON(newTeam).
		Expect().
		Status(http.StatusCreated).JSON().Object().IsEqual(newTeam)

	teams := e.GET("/teams").
		Expect().
		Status(http.StatusOK).JSON().Array()

	teams.ContainsAll(newTeam)

	e.DELETE(fmt.Sprintf("/teams/%s", teamId)).
		Expect().
		Status(http.StatusOK).
		JSON().Object().HasValue("message", "Team delete requested")

	addedTeam := e.GET(fmt.Sprintf("/teams/%s", teamId)).
		Expect().
		Status(http.StatusOK).JSON().Object()

	addedTeam.HasValue("teamMarkedForDeletion", "Requested")

	e.DELETE(fmt.Sprintf("/teams/%s/confirm", teamId)).
		Expect().
		Status(http.StatusOK).
		JSON().Object().HasValue("message", "Team deleted")

	updatedTeams := e.GET("/teams").
		Expect().
		Status(http.StatusOK).JSON().Array()

	updatedTeams.NotContainsAll(newTeam)
}
