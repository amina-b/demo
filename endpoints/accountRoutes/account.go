package accountRoutes

import (
	"io"
	"net/http"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "registerRuta")
}
