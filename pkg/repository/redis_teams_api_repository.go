package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
	"twdps.io/lab-api-teams/pkg/domain"
)

type RedisTeamRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisTeamRepository() *RedisTeamRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	return &RedisTeamRepository{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (store *RedisTeamRepository) GetTeam(id string) (domain.Team, error) {
	teamJson, err := store.client.Get(store.ctx, fmt.Sprintf("team:%s", id)).Result()
	if err != nil {
		return domain.Team{}, err
	}

	var team domain.Team
	if err := json.Unmarshal([]byte(teamJson), &team); err != nil {
		return domain.Team{}, err
	}

	return team, nil
}

func (store *RedisTeamRepository) GetTeams() ([]domain.Team, error) {
	teamKeys, err := store.client.Keys(store.ctx, "team:*").Result()
	if err != nil {
		return nil, err // TODO Redis Err handling
	}

	fmt.Printf("%+v", teamKeys)

	var teams []domain.Team
	for _, key := range teamKeys {
		teamJson, err := store.client.Get(store.ctx, key).Result()
		if err != nil {
			return nil, err
		}

		var team domain.Team
		if err := json.Unmarshal([]byte(teamJson), &team); err != nil {
			return nil, err
		}

		teams = append(teams, team)

	}

	return teams, nil
}

func (store *RedisTeamRepository) AddTeam(newTeam domain.Team) error {
	teamJson, err := json.Marshal(newTeam)
	if err != nil {
		log.Fatalf("failed to marshal")
		return err // Transient error - TODO
	}

	return store.client.Set(store.ctx, fmt.Sprintf("team:%s", newTeam.TeamID), teamJson, 0).Err()
}
