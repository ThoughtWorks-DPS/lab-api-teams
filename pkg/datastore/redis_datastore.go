package datastore

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
	"github.com/go-redis/redis/v8"
)

type RedisDatastore interface {
	Datastore
}

type redisDatastoreImpl struct {
	client *redis.Client
	ctx    context.Context
}

func NewRedisDatastore() RedisDatastore {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   0,
	})

	return &redisDatastoreImpl{
		client: rdb,
		ctx:    context.Background(),
	}
}

func getKey(data interface{}) string {
	switch data.(type) {
	case domain.Namespace:
		return "namespace:"
	case domain.Team:
		return "team:"
	// Add more cases for other entities
	default:
		return ""
	}
}

func (r *redisDatastoreImpl) Create(data interface{}) error {
	serializedData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := r.client.HSet(r.ctx, "data", serializedData).Err(); err != nil {
		return err
	}

	// Create indices
	attributes := make(map[string]interface{})
	v := reflect.ValueOf(data)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i).Interface()
		attributes[field.Name] = value
	}
	return r.CreateIndicesForObject("namesapce:", attributes)
}

func (r *redisDatastoreImpl) CreateIndicesForObject(id string, attributes map[string]interface{}) error {
	for attr, value := range attributes {
		key := fmt.Sprintf("index:%s:%s:%v", reflect.TypeOf(attributes).Name(), attr, value)
		if err := r.client.SAdd(r.ctx, key, id).Err(); err != nil {
			return err
		}
	}
	return nil
}

func (ds *redisDatastoreImpl) ReadByID(id string, out interface{}) error {
	key := getKey(out) + id // Define getKey similarly
	dataHash, err := ds.client.HGet(ds.ctx, key, "data").Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(dataHash), out)
}

func (ds *redisDatastoreImpl) ReadAll(data interface{}) error {
	return fmt.Errorf("not implemented")
}

func (r *redisDatastoreImpl) Delete(data interface{}) error {

	// Construct the key for the team
	teamKey := fmt.Sprintf("team:%s", "asdf")

	// Delete the team from Redis
	err := r.client.Del(r.ctx, teamKey).Err()
	if err != nil {
		return fmt.Errorf("could not delete team: %v", err)
	}

	return nil
}

func (r *redisDatastoreImpl) IsDatabaseAvailable() (bool, error) {
	res, err := r.client.Ping(r.ctx).Result()
	if err != nil {
		return false, fmt.Errorf("redis was not reachable: %v", err)
	}
	if res != "PONG" {
		return false, fmt.Errorf("unexpected response from redis PING: %s", res)
	}

	return true, nil
}

func (store *redisDatastoreImpl) Update(data interface{}) error {
	return fmt.Errorf("not implemented")
}

func (r *redisDatastoreImpl) ReadByAttributes(filter Filter, out interface{}) error {
	// Create a slice of all index keys to intersect
	keysToIntersect := make([]string, 0, len(filter))
	for attr, value := range filter {
		key := fmt.Sprintf("index:%s:%s:%v", reflect.TypeOf(filter).Name(), attr, value)
		keysToIntersect = append(keysToIntersect, key)
	}

	// Intersect the sets to find common IDs
	matchingIDs, err := r.client.SInter(r.ctx, keysToIntersect...).Result()
	if err != nil {
		return err
	}

	// Fetch each entity by ID and append to the out slice
	resultVal := reflect.ValueOf(out).Elem()
	for _, id := range matchingIDs {
		hashData, err := r.client.HGetAll(r.ctx, id).Result()
		if err != nil {
			return err
		}

		serializedData, ok := hashData["data"]
		if !ok {
			continue
		}

		objType := reflect.TypeOf(out).Elem().Elem() // Assuming out is a pointer to a slice
		obj := reflect.New(objType).Interface()

		if err := json.Unmarshal([]byte(serializedData), &obj); err != nil {
			return err
		}

		resultVal.Set(reflect.Append(resultVal, reflect.ValueOf(obj).Elem()))
	}
	return nil
}

func (g *redisDatastoreImpl) ReadByAttributesWithPagination(filter map[string]interface{}, out interface{}, page int, maxResult int) error {
	return fmt.Errorf("not implemented")
}
