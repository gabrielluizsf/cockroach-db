package main

import (
	"github.com/GabrielLuizSF/cockroach-db/database"
	"github.com/GabrielLuizSF/cockroach-db/server"
)

func main() {
	database.Conect();
	s := server.NewServer();

	s.Run();
}