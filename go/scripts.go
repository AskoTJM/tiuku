package swagger

import (
	"fmt"
	"net/http"
)

type testData struct {
	UserID string `json:"id"`
}

func testResponse(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
