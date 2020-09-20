package handler

import (
	"fmt"
	"net/http"
)

// Gdrive ...
// func (h Hndl) Gdrive() {
// 	http.HandleFunc("/", gdrive)
// }

func gdrive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "test")
	// t, err := template.ParseFiles("in.htm")
	// Try("", err, true)
	// err:= t.Execute(w,var)
	// fmt.Println("-" + r.URL.Path)
	// fmt.Fprintln(w, "test")
}
