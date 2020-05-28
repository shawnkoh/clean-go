package frameworks

import (
	"fmt"
	"net/http"

	"example/interfaceadapters"
)

// Router is a router...
func Router(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hi there, I love %s!", request.URL.Path[1:])
	interfaceadapters.UserController{}.CreateUser("abc")
}
