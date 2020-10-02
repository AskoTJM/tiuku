package scripts

import (
	"fmt"
	"net/http"
	"strconv"
)

// desc: Turn uint64 to string
func uint64ToString(newUint uint64) string {
	//var u uint32 = newUint
	//var s = strconv.FormatUint(uint64(u), 10)
	var s = strconv.FormatUint(uint64(newUint), 10)
	return s
}
func TestResponse(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}
