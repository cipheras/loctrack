package handler

import (
	"errors"
	"fmt"
	"net/http"

	c "github.com/cipheras/gohelper"
)

// Hndl ...
type Hndl int

// Demo ...
func (h Hndl) Demo() {
	fs := http.FileServer(http.Dir("template/demo-template"))
	http.Handle("/static/", fs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case "GET": //GET request handler
			// Your code here

			/* Example:
			tpl, err := template.ParseFiles("template/demo/index.html")
			c.Try(err, false, "parsing demo/index.html")
			err = tpl.Execute(w, data)
			c.Try(err, false, "executing demo template")
			*/

		default:
			fmt.Fprintln(w, "Request not supported")
			fmt.Println("# Unsupported request")
			c.Try(errors.New("unsupported request"), false)
			return
		}
	})
}
