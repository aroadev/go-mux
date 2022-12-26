package main

import (
	"net/http"

	"github.com/artzdev/go-mux/db"
	"github.com/artzdev/go-mux/models"
	"github.com/artzdev/go-mux/routes"
	"github.com/gorilla/mux"
)

func main() {
	db.DBConnection() /* Objeto de conexión a la base de datos */
	db.DB.AutoMigrate(models.Task{})
	db.DB.AutoMigrate(models.User{})

	r := mux.NewRouter() /* Creamos un nuevo enrutador de MUX */

	r.HandleFunc("/", routes.HandleHome) /* Creamos la ruta principal */
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r) /* Asignamos el puerto para la ejecución sel servidor */
}
