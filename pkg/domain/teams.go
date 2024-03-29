package domain

// team-type: [ normal | master | admin ]
// team-id: (team github name)
// team-description: freeform
// team-ram: integer
//      In the master record this
//      represents the default pool
//      size to assign to teams.
// team-cpu: integer
// team-integrations: json
//      list of the supported integrations the
//      team wants the platform to maintain on their
//      behalf.
// team-ram-limit: only admin edit.
//      in master this is the max
//      self-managed resource limit
// team-cpu-limit: only admin edit.
//      in master this is the max
//      self-managed resource limit
// team-marked-for-deletion:
//      [ requested | pending | done ]

type Team struct {
	TeamType              string `json:"teamType"`
	TeamID                string `json:"teamID" gorm:"primaryKey"`
	TeamDescription       string `json:"teamDescription"`
	TeamRAM               int    `json:"teamRAM"`
	TeamCPU               int    `json:"teamCPU"`
	TeamRamLimit          int    `json:"teamRamLimit"`
	TeamCpuLimit          int    `json:"teamCPULimit"`
	TeamMarkedForDeletion string `json:"teamMarkedForDeletion"`
}

type TeamIntegration struct {
	IntegrationName string `json:"integrationName"`
}
