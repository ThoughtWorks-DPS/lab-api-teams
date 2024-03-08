package main

import (
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/datastore"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/handler"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/repository"
	"github.com/ThoughtWorks-DPS/lab-api-teams/pkg/service"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(handler.ErrorHandler())

	ds_tm := datastore.NewGormDatastore("team")
	ds_ns := datastore.NewGormDatastore("namespace")
	// ds_gw := datastore.NewGormDatastore("gateway")

	// You can choose to run migrations prior to server startup.
	//    Be warned, this can be very slow and the k8s probes are not tuned to wait,
	//		thus causing crashloops
	//
	// if migrator, ok := ds_tm.(datastore.Migratable); ok {
	// 	err := migrator.Migrate(&domain.Team{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	err = migrator.Migrate(&domain.Namespace{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	err = migrator.Migrate(&domain.Gateway{})
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }

	// Teams
	teamRepo := repository.NewTeamsRepo(ds_tm)
	teamService := service.NewTeamService(teamRepo)
	teamHandler := handler.NewTeamHandler(teamService)

	// Namespaces
	namespaceRepo := repository.NewNamespaceRepository(ds_ns)
	namespaceService := service.NewNamespaceService(namespaceRepo)
	namespaceHandler := handler.NewNamespaceHandler(namespaceService)

	router.GET("/teams/healthz/readiness", teamHandler.Readiness)
	router.GET("/teams/healthz/liveness", teamHandler.Liveness)
	router.GET("/teams/:teamID", teamHandler.GetTeam)
	router.GET("/teams", teamHandler.GetTeams)
	router.POST("/teams", teamHandler.AddTeam)
	router.DELETE("/teams/:teamID", teamHandler.RemoveTeam)
	router.DELETE("/teams/:teamID/confirm", teamHandler.ConfirmRemoveTeam)
	router.GET("/namespaces/:namespaceID", namespaceHandler.GetNamespace)
	router.GET("/namespaces", namespaceHandler.GetNamespaces)
	router.GET("/namespaces/master", namespaceHandler.GetNamespacesMaster)
	router.GET("/namespaces/standard", namespaceHandler.GetNamespacesStandard)
	router.GET("/namespaces/custom", namespaceHandler.GetNamespacesCustom)
	router.POST("/namespaces", namespaceHandler.AddNamespace)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
