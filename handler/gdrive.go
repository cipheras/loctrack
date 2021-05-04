package handler

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	c "github.com/cipheras/gohelper"
)

// Demo ...
func (h Hndl) Gdrive() {
	fs := http.FileServer(http.Dir("template/gdrive"))
	http.Handle("/static/", fs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case "GET": //GET request handler
			tpl, err := template.ParseFiles("template/gdrive/index.html")
			c.Try(err, false, "parsing gdrive/index.html")
			err = tpl.Execute(w, nil)
			c.Try(err, false, "executing template")

		default:
			fmt.Fprintln(w, "Request not supported")
			fmt.Println("# Unsupported request")
			c.Try(errors.New("unsupported request"), false)
			return
		}
	})
}
