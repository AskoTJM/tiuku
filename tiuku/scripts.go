package swagger

import (
	"fmt"
	"net/http"
)

func testResponse(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
