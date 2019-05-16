package now_go_test

import (
	"fmt"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) { fmt.Fprintf(w, "Hello World") }
