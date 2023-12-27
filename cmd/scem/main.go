package main

import (
	"log"
	"os"

	"github.com/hkm12345123/transport_system/api/middleware"
	"github.com/hkm12345123/transport_system/api/server"
	"github.com/hkm12345123/transport_system/internal/handler"
	CommonService "github.com/hkm12345123/transport_system/internal/service/common"
	CommonMessage "github.com/hkm12345123/transport_system/internal/service/common_message"
	ZBMessage "github.com/hkm12345123/transport_system/internal/service/zeebe/message"
	ZBWorker "github.com/hkm12345123/transport_system/internal/service/zeebe/worker"
	ZBWorkflow "github.com/hkm12345123/transport_system/internal/service/zeebe/workflow"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("RUNENV") != "docker" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}

	// Initial web auth middleware
	if os.Getenv("RUN_WEB_AUTH") == "yes" {
		runWebAuth()
	}

	// Select app auth middleware
	if os.Getenv("RUN_APP_AUTH") == "redis" {
		runAppAuthRedis()
	} else if os.Getenv("RUN_APP_AUTH") == "buntdb" {
		log.Println("Selected BuntDB to run app auth")
	}

	// Select database
	if os.Getenv("SELECT_DATABASE") == "1" {
		connectPostgress()
	} else if os.Getenv("SELECT_DATABASE") == "2" {
		connectMySQL()
	} else if os.Getenv("SELECT_DATABASE") == "3" {
		connectSQLite()
	} else {
		log.Println("No database selected!")
		os.Exit(1)
	}

	gormDB := handler.GetGormInstance()
	CommonService.MappingGormDBConnection(gormDB)
	CommonMessage.MappingGormDBConnection(gormDB)

	if err := handler.RefreshDatabase(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	if err := handler.MigrationDatabase(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Print("Database refreshed!")

	if os.Getenv("STATE_SERVICE") == "1" {
		connectZeebeClient()
	}

	// Our servers will live in the routes package
	server.RunServer()

}

// Source code: https://www.devdungeon.com/content/working-files-go#read_all
func runWebAuth() {
	sessionKey := []byte(os.Getenv("SESSION_KEY"))
	if err := middleware.RunWebAuth(sessionKey); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("Web authenticate activated!")
}

func runAppAuthRedis() {
	if err := middleware.RunAppAuthRedis(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("Selected Redis to run app auth!")
}

func connectPostgress() {
	if err := handler.ConnectPostgres(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	log.Println("Connected with posgres database!")
}

func connectMySQL() {
	if err := handler.ConnectMySQL(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Println("Connected with posgres database!")
}

func connectSQLite() {
	if err := handler.ConnectSQLite(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Println("Connected with sqlite database!")
}

func connectZeebeClient() {
	if err := ZBWorkflow.ConnectZeebeEngine(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Println("Zeebe workflow package connected with zeebe!")
	if err := ZBMessage.ConnectZeebeEngine(); err != nil {
		log.Print(err)
		os.Exit(1)
	}
	log.Println("Zeebe message package connected with zeebe!")
	// Run Zebee service
	ZBWorker.RunOrderLongShip()
	ZBWorker.RunOrderShortShip()
	ZBWorker.RunLongShipFinish()
}
