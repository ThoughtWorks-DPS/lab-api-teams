## Contributing Guidelines

Welcome to the Platform Starter Kit. We are excited about the prospect of you joining our community. 

Links:
- Backlog and Board: https://github.com/orgs/ThoughtWorks-DPS/projects/5/views/10

### Installing and developing the Teams API

#### Prerequisites
##### Running Postgres
###### Option 1: Run Postgres directly on host
(The not so great but totally works way)
- `brew install postgres`
- `initdb .postgres`
- `pg_ctl -D ./.postgres -o "-F -p 5433" start`
- `psql postgres -p 5433`
- ```postgresql
  CREATE DATABASE gorm
  CREATE USER postgres
  ```
###### Option 2: Run Postgres Docker container (using docker client with colima on MacOS)
- Install colima and docker cli: `brew install colima docker`
- Run postgres container within lima VM (unauthenticated, only use for local development):
  - `colima start`
  - `docker context use colima`
  - ```bash
    PORT=5433
    docker run --rm -d --name teams-api-db \
    -e POSTGRES_HOST_AUTH_METHOD=trust 
    -e POSTGRES_USER=postgres \
    -e POSTGRES_DB=gorm \
    -p $PORT:5432 \
    postgres
    ```
    or with docker compose:
    ```bash
    docker compose up -d
    ```
- Optional: Confirm db connectivity
   - Install psql client (if you do not have one yet) and add to PATH:
     - `brew install libpq`
     - `echo 'export PATH="/opt/homebrew/opt/libpq/bin:$PATH"' >> ~/.zshrc`
     - `source ~/.zshrc`
   - Test connectivity host with psql
     - `psql --host=localhost --port=$PORT -U postgres -d gorm`
- Continue to "Bootstrap database" below

##### Bootstrap database
- `export DATABASE_URL="localhost"`
- Migration script to create initial schemas: `go run ./scripts/migrate.go`

##### Starting teams API
Running the teams api (TODO: update makefile)
```
go get ./...
go run cmd/lab-api-teams/main.go
```

##### Linting
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


### Workflows and pipelines

This repository is trunk based, so you may contribute directly to main if you're comfortable doing so. On commit to main, our pipeline in `.circleci/config.yaml` will do a static check on the changes (fmt, lint, unit test) and do a dry run image build, after which it will trigger a release. At this point, the release pipeline picks up and attempts to deploy the new verison to dev, test, and production.

You can see examples of this in action:
- [Trunk Based Build 1.3.5](https://app.circleci.com/pipelines/github/ThoughtWorks-DPS/lab-api-teams/239/workflows/dde9f913-1266-44ae-8c01-a85394b91b04)
- [Triggered release 1.3.5](https://app.circleci.com/pipelines/github/ThoughtWorks-DPS/lab-api-teams/240/workflows/29b55a2a-6182-4171-9b4e-91beae1746bf)

You may notice that the release pipeline has an `e2e` step after each environment. These tests are maintained in a folder seperate from the `pkg` folder in `./test-e2e`. Please notice that these tests have 
go build flags at the top:

```
//go:build e2e
// +build e2e
```

This allows you to run the `make e2e` command, which translates to `go test ./... -tags=e2e -v`, which targets the e2e tests and ignores the other types of tests (that were run in the static stages). 

The release trigger is based on tags in the Github repo. To facilitate this, in the trunk/build pipeline we use semantic release: https://github.com/ThoughtWorks-DPS/lab-api-teams/blob/main/.circleci/config.yml#L187-L200

This does mean that your commits need to follow the conventions of conventional commits. You can find more details on how to do this here: https://github.com/semantic-release/semantic-release#commit-message-format

The table below shows which commit message gets you which release type when `semantic-release` runs (using the default configuration):

| Commit message                                                                                                                                                                                   | Release type                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------- |
| `fix(pencil): stop graphite breaking when too much pressure applied`                                                                                                                             | ~~Patch~~ Fix Release                                                                                           |
| `feat(pencil): add 'graphiteWidth' option`                                                                                                                                                       | ~~Minor~~ Feature Release                                                                                       |
| `perf(pencil): remove graphiteWidth option`<br><br>`BREAKING CHANGE: The graphiteWidth option has been removed.`<br>`The default graphite width of 10mm is always used for performance reasons.` | ~~Major~~ Breaking Release <br /> (Note that the `BREAKING CHANGE: ` token must be in the footer of the commit) |


