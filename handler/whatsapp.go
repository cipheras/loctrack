package handler

import (
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	c "github.com/cipheras/gohelper"
)

// Hndl ...
// type Hndl int

// Whatsapp ...
func (h Hndl) Whatsapp() {
	var title, img string
	c.Cprint(c.S, "Enter group title: ")
	fmt.Scanf("%v\n", &title)
	c.Cprint(c.S, "Enter group image path (e.g. /home/$USER/...): ")
	fmt.Scanf("%v\n", &img)
	var newImgPath string
	if img != "" {
		iname := strings.Split(img, "/") //image name from path
		iloc, err := os.Open(img)        //open original img
		c.Try(err, false, "opening image")
		ilocNew, err := os.Create("template/whatsapp/static/images/" + iname[len(iname)-1]) //open copy img
		c.Try(err, false, "copying image to workspace")
		io.Copy(ilocNew, iloc) //copy img data
		iloc.Close()
		ilocNew.Close()
		newImgPath = "static/images/" + iname[len(iname)-1] //new copied img path
		defer os.Remove(newImgPath)
	} else {
		newImgPath = "/static/images/wJskj.jpg"
	}
	fs := http.FileServer(http.Dir("template/whatsapp"))
	http.Handle("/static/", fs)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {

		case "GET": //GET request handler
			type Data struct {
				Title string
				Img   string
			}
			data := Data{
				Title: title,
				Img:   newImgPath,
			}
			tpl, err := template.ParseFiles("template/whatsapp/index.html")
			c.Try(err, false, "parsing whatsapp/index.html")
			err = tpl.Execute(w, data)
			c.Try(err, false, "executing whatsapp template")

		default:
			fmt.Fprintln(w, "Request not supported")
			fmt.Println("# Unsupported request")
			c.Try(errors.New("unsupported request"), false)
			return
		}
	})
}
