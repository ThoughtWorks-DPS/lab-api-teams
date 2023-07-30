# lab-api-teams

### Local Execution

TODO: automate/script local setup 

```
$ go get
$ docker run -d --name redis-stack-server -p 6379:6379 redis/redis-stack-server:latest
$ go run cmd/teams-api-server/main.go
```

### Project structure 
``` bash
.
├── README.md
├── cmd
│   └── teams-api-server
│       └── main.go
├── go.mod
├── go.sum
├── pkg
│   ├── domain
│   │   ├── gateways.go
│   │   ├── namespaces.go
│   │   └── teams.go
│   ├── handler
│   │   └── handler.go
│   ├── repository
│   │   ├── mock
│   │   │   ├── mock_team_repository.go
│   │   │   └── mock_team_repository_test.go
│   │   └── redis_teams_api_repository.go
│   └── service
│       ├── service.go
│       └── service_test.go
└── scripts
    └── curl.sh
```
