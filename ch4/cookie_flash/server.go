package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func setMessage(w http.ResponseWriter, r *http.Request) {
	msg := []byte("Hello, hiro")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("flash")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			fmt.Fprintln(w, "No Message Found")
		}
	} else {
		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)
		val, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(val))
	}
}

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/set_message", setMessage)
	http.HandleFunc("/show_message", showMessage)

	server.ListenAndServe()

}
