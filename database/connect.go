package database

import (
	"fmt"
	"os"
	"time"
	"github.com/GabrielLuizSF/cockroach-db/database/migrations"
	"github.com/GabrielLuizSF/cockroach-db/database/utils/errors/fatal"
	"github.com/GabrielLuizSF/cockroach-db/database/utils/errors/connection"	
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	
)

var database *gorm.DB;

func Conect(){
    
	err := godotenv.Load();
		fatal.PrintFatalError(err,"Error loading .env file");
	
	conectDB := os.Getenv("DATABASE_URI");
	db, err := gorm.Open(postgres.Open(conectDB), &gorm.Config{});
		fatal.PrintFatalError(err,"failed to connect database");
		fmt.Println("Database Connected");

	database = db;
		config, _ := database.DB();
			config.SetMaxIdleConns(10);
			config.SetMaxOpenConns(100);
			config.SetConnMaxLifetime(time.Hour);
		
		migrations.Run(database);
			fmt.Println("Migrations Finished");
	
}
func CloseConn(){
	config, err := database.DB();
		connection.CloseConnError(err);
	err = config.Close();
		connection.CloseConnError(err);
}

func GetDatabase() *gorm.DB {
	return database
}

