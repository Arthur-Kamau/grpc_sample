package main

import (
	"fmt"
	"log"
	"os"

	"govt-tracker/app"
	"govt-tracker/app/database"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

// @title Government Tracker API
// @version 1.0
// @description This is api des the following
//				<br/> 1. Auth process ( Create Users, login, validate token) <br/>
//				<br/> 2. Fetch and update profile <br/>
//				<br/> 3.  <br/>
//				<br/> 4. Create and view Gov projects <br/>
//				<br/> 5. Create and view Gov projects <br/>
//
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey
// @in header
// @name Authorization

// @host identity-service.dev.da-ride.com
// @BasePath /
// @schemes https
func main() {

	// load .env file
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load .env file. Err: %s \n Assuming you set env before go run main.go ", err.Error())

	}

	// setup database
	dbInstance := database.DbInstance(os.Getenv("DBNAME"))

	driver, err := postgres.WithInstance(dbInstance, &postgres.Config{})
	if err != nil {

		panic(err)
	}

	migrationPath := GetRootPath()
	if len(migrationPath) == 0 {
		panic(fmt.Sprintf("Migration file locations not found in env \n path %s", migrationPath))
	}

	migrationPathFile := fmt.Sprintf("file:///%s/app/migrations", migrationPath)
	fmt.Printf("migrationPathFile %s", migrationPathFile)
	m, err := migrate.NewWithDatabaseInstance(migrationPathFile, "postgres", driver)
	if err != nil {

		log.Printf("migration setup error %s ", err.Error())
	}

	err = m.Up() // or m.Step(2) if you want to explicitly set the number of migrations to run
	if err != nil {

		log.Printf("migration error %s ", err.Error())
	}

	router := &app.Server{}
	router.Initialize(dbInstance)
	select {}
}

func GetRootPath() string {

	// _, b, _, _ := runtime.Caller(0)

	// Root folder of this project
	// return filepath.Join(filepath.Dir(b), "./")
	return os.Getenv("MIGRATION_LOCATION")
}
