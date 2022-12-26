package routes

import "net/http"

func HandleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
