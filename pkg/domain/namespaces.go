package domain

// ns-type: [ normal | master | standard | custom ]
//			normal:     the ns definition is a normal, team ns desired state definition. Normal
//                        namespaces are subject of the sync action.
//      master:     ns definition not associated with any team. The master record contains
//                       the default values used for ram, cpu, in-or-out of mesh.
//      standard: The standard ns records represent the default namespaces provisioned
//                       for a new team at creation.
//      custom:    Teams may define custom namespace. By convention this is the equivalent
//                        of creating an 'optional' standard entry. All teams can view the list of custom
//                        ns and choose to adopt it as well as define a new one.
// ns-team: team-id used by normal ns entries
// ns-id: name to append to team-id [dev | qa | etc]
// ns-ram: k8s resource
//      [pool resource] In the master record, define the default amount of ram/cpu to assign
//      to new ns (typically a number like 1/5 of the amount of the teams pool from the team
//      record
// ns-cpu: k8s resource [pool resource like ram]
// ns-in-mesh: istio managed? boolean, true by default
// ns-cluster-location: name of cluster where the ns resides TODO - discuss if this is needed
// ns-from-default: was this created at onboarding from defaults?

type Namespace struct {
	NamespaceType        string `json:"namespaceType"`   // normal, master, standard, custom
	NamespaceTeamID      string `json:"namespaceTeamID"` // ID from teams API
	NamespaceID          string `json:"namespaceID"`     // Dev, QA, etc
	NamespaceRam         int    `json:"namespaceRam"`
	NamespaceCpu         int    `json:"namespaceCpu"`
	NamespaceInMesh      bool   `json:"namespaceInMesh"`
	NamespaceFromDefault bool   `json:"namespaceFromDefault"`
}



// TODO the below functions (and other business logic specific api calls) will move to service layer
// GetNamespacesMaster() ([]Namespace, error)
// GetNamespacesStandard() ([]Namespace, error)
// GetNamespacesCustom() ([]Namespace, error)

// GetTeamNamespaces(teamID string) ([]Namespace, error)
// GetTeamNamespaceByNamespaceID(teamID string, namespaceID string) (Namespace, error)
// AddTeamNamespace(namespace Namespace) error
// UpdateTeamNamespace(namespace Namespace) error
// DeleteTeamNamespace(namespace Namespace) error
