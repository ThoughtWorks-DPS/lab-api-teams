package repository

// type NamespaceRepository interface {
// 	GetNamespaces() ([]Namespace, error)
// 	GetNamespacesByAttribute(attrKey string, attrValue string) ([]Namespace, error)
// 	GetNamespaceByID(namespaceID string) (Namespace, error)
// 	AddNamespace(namespace Namespace) error
// 	UpdateNamespace(namespace Namespace) error
// 	RemoveNamespace(namespace Namespace) (Namespace, error)
// }

// TODO - redis repository tests

import (
	"context"
	"encoding/json"
	"log"

	"github.com/RBMarketplace/di-api-teams/pkg/domain"
	"github.com/go-redis/redis/v8"
)

type RedisNamespaceRepository struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisNamespaceRepository() *RedisNamespaceRepository {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	return &RedisNamespaceRepository{
		client: rdb,
		ctx:    context.Background(),
	}
}

func (store *RedisNamespaceRepository) GetNamespaces() ([]domain.Namespace, error) {
	namespaceIDs, err := store.client.SMembers(store.ctx, "namespaces").Result()
	if err != nil {
		return nil, err
	}

	var namespaces []domain.Namespace
	for _, namespaceID := range namespaceIDs {
		namespaceHash, err := store.client.HGet(store.ctx, "namespace:"+namespaceID, "data").Result()
		if err != nil {
			return nil, err
		}

		var namespace domain.Namespace
		err = json.Unmarshal([]byte(namespaceHash), &namespace)
		if err != nil {
			return nil, err
		}

		namespaces = append(namespaces, namespace)
	}

	return namespaces, nil
}

func (store *RedisNamespaceRepository) AddNamespace(namespace domain.Namespace) error {
	serializedNamespace, err := json.Marshal(namespace)
	if err != nil {
		return err
	}

	// Save the namespace in a hash
	if err := store.client.HSet(store.ctx, "namespace:"+namespace.NamespaceID, "data", serializedNamespace).Err(); err != nil {
		return err
	}

	// Add namespace's ID to the set of all namespaces
	if err := store.client.SAdd(store.ctx, "namespaces", namespace.NamespaceID).Err(); err != nil {
		return err
	}

	// Update the genre index
	if err := store.client.SAdd(store.ctx, "index:namespace:namespaceType:"+namespace.NamespaceType, namespace.NamespaceID).Err(); err != nil {
		return err
	}

	return nil
}

func (store *RedisNamespaceRepository) GetNamespacesByType(nsType string) ([]domain.Namespace, error) {
	namespaceIDs, err := store.client.SMembers(store.ctx, "index:namespace:namespaceType:"+nsType).Result()
	log.Printf("%+v", namespaceIDs)
	if err != nil {
		return nil, err
	}

	var namespaces []domain.Namespace
	for _, namespaceID := range namespaceIDs {
		nsHash, err := store.client.HGet(store.ctx, "namespace:"+namespaceID, "data").Result()
		if err != nil {
			return nil, err
		}

		var namespace domain.Namespace
		err = json.Unmarshal([]byte(nsHash), &namespace)
		if err != nil {
			return nil, err // TODO transient/err
		}

		namespaces = append(namespaces, namespace)
	}

	return namespaces, nil
}
