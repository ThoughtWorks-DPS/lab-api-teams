## Contributing Guidelines

Welcome to the Platform Starter Kit. We are excited about the prospect of you joining our community. 

Links:
- Backlog and Board: https://github.com/orgs/ThoughtWorks-DPS/projects/5/views/10

### Installing and developing the Teams API

#### Prerequisites
- Running postgres database, connection string exported to env: `DATABASE_URL`
- Migration script to create initial schemas: `go run ./scripts/migrate.go`

Running the teams api (TODO: update makefile)
```
go get ./...
go run cmd/lab-api-teams/main.go
```

To run fmt, lint, test in one command
```
make static
```

### Getting Started, Purpose and Vision

First, you should checkout the Management APIs Mural to get a strong understanding of what we are trying to accomplish with this project:
[Mural, PSK Management APIs](https://app.mural.co/t/thoughtworksclientprojects1205/m/thoughtworksclientprojects1205/1687460544281/833975940ec11b5a3c7af94e4ec8cc4253a6187d?sender=ue017dd0a0ba865be72d75848)

If you are an outside collaborator (outside of thoughtworks) or an ex-thoughtworker that wants to continue contributing, please send us a request (file an issue on this repo) to get added to this Mural board.

To describe in text, the purpose of the teams api is to be the primary touchpoint for teams to self-service the management of their team within the context of the DI platform. To give an example, one can do things like: create a team, see other teams, create an integration for your team, create a new namespace for your team, create a gateway for your team, etc. This api sits within the control plane/mapi of our typical deployment.

To date, we have never "harvested" the teams-api implementation back to the DPS github org, we usually build some form of it at each engagement. This is partially because every client is very different, and they have different clouds and different preferred databases. 

So this teams api, is an attempt to create a multi-cloud multi-database teams api that is functional just about anywhere. To that end, it uses [GORM (Go ORM)](https://gorm.io/index.html) to support many database technologies. 

And it uses a Domain Driven Design approach to decouple the datastore implementation from the rest of the program. We accomplish this by using the architectural constructs of:

- [Datastore](https://github.com/ThoughtWorks-DPS/lab-api-teams/tree/main/pkg/datastore)
- [Repository](https://github.com/ThoughtWorks-DPS/lab-api-teams/tree/main/pkg/repository)
- [Handler](https://github.com/ThoughtWorks-DPS/lab-api-teams/tree/main/pkg/handler)
- [Service](https://github.com/ThoughtWorks-DPS/lab-api-teams/tree/main/pkg/service)
- [Domain](https://github.com/ThoughtWorks-DPS/lab-api-teams/tree/main/pkg/domain)

Because of this decoupling in our codebase, if we needed to later on add a databse that Gorm doesn't support, we would just implement a new datastore for that database. The rest of the code wouldn't be touched.

### Where to contribute

After going through the Mural, we recommend checking out the [teams api board](https://github.com/orgs/ThoughtWorks-DPS/projects/5/views/10)https://github.com/orgs/ThoughtWorks-DPS/projects/5/views/10)
Any stories that are in the `Ready` column are groomed and ready to be picked up and worked on. If you have clarifying questions please ask them in the thread of the story or in gchat.




