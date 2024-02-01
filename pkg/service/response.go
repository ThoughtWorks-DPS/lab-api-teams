package service

import "github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"

type ListNamespaceResponse struct {
	Items      []domain.Namespace `json:"items"`
	Page       int                `json:"page"`
	MaxResults int                `json:"maxResults"`
}
