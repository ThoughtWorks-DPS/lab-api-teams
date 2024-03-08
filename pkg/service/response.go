package service

import "github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"

type ListResponse[K domain.Namespace | domain.Team] struct {
	Items      []K `json:"items"`
	Page       int `json:"page"`
	MaxResults int `json:"maxResults"`
}

type ListNamespaceResponse ListResponse[domain.Namespace]

type ListTeamResponse ListResponse[domain.Team]
