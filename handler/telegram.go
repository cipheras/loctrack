package handler

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"

	c "github.com/cipheras/gohelper"
)

// Telegram ...
func (h Hndl) Telegram() {
	fs := http.FileServer(http.Dir("template/telegram"))
	http.Handle("/static/", fs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET": //GET request handler
			tpl, err := template.ParseFiles("template/telegram/index.html")
			c.Try(err, false, "parsing whatsapp/index.html")
			err = tpl.Execute(w, nil)
			c.Try(err, false, "executing telegram template")

		default:
			fmt.Fprintln(w, "Request not supported")
			fmt.Println("# Unsupported request")
			c.Try(errors.New("unsupported request"), false)
			return
		}
	})
}
