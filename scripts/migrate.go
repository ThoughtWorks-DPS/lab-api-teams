package main

import (
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/datastore"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/domain"
)

// manual migration steps
// set DATABASE_URL env variable to localhost
// port-forward the yugabyte db
// run this go program
func main() {
	ds_tm := datastore.NewGormDatastore("team")
	ds_ns := datastore.NewGormDatastore("namespaces")
	// ds_gw := datastore.NewGormDatastore("gateways")

	if migrator, ok := ds_tm.(datastore.Migratable); ok {
		err := migrator.Migrate(&domain.Team{})
		if err != nil {
			panic(err)
		}
		err = migrator.Migrate(&domain.Namespace{})
		if err != nil {
			panic(err)
		}
		err = migrator.Migrate(&domain.Gateway{})
		if err != nil {
			panic(err)
		}
	}

	if migrator, ok := ds_ns.(datastore.Migratable); ok {
		err := migrator.Migrate(&domain.Team{})
		if err != nil {
			panic(err)
		}
		err = migrator.Migrate(&domain.Namespace{})
		if err != nil {
			panic(err)
		}
		err = migrator.Migrate(&domain.Gateway{})
		if err != nil {
			panic(err)
		}
	}
}
