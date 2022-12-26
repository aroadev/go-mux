package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=aroa password=aroadev dbname=gorm port=5432" /* String de conexión */
var DB *gorm.DB

func DBConnection() {
	var error error                                           /* Variable para el manejo de errores */
	DB, error = gorm.Open(postgres.Open(dsn), &gorm.Config{}) /* Conexión a base de datos */
	if error != nil {                                         /* Comparamos si hay errores para terminar la ejecución */
		log.Fatal(error)
	} else {
		log.Println("DB Connected")
	}
}
