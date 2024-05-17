package main

import (
	"fmt"
	"log"
	"os"

	"grpc_sample/app"
	"grpc_sample/app/database"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/joho/godotenv"
)

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
