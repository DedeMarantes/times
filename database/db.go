package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DedeMarantes/times/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDb() {
	password := os.Getenv("DB_PASS")
	stringConexao := fmt.Sprintf("host=localhost user=root password=%s dbname=postgres port=5432 sslmode=disable", password)
	//stringConexao := "host=localhost user=root password=%s dbname=postgres port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		log.Fatal("Erro ao conectar banco de dados")
	}
	DB.AutoMigrate(&models.Time{}, &models.Jogador{})
}
