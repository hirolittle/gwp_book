package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/set_cookie", setCookie)
	http.HandleFunc("/get_cookie", getCookie)

	server.ListenAndServe()

}

func setCookie(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "first_cookie_value",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "second_cookie_value",
		HttpOnly: true,
	}

	//w.Header().Set("Set-Cookie", c1.String())
	//w.Header().Add("Set-cookie", c2.String())

	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	//c := r.Header["Cookie"]
	//fmt.Fprintf(w, "Cookie: %v\n", c)

	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintf(w, "Cookie not found")
	}
	cs := r.Cookies()
	fmt.Fprintf(w, "Cookie: %v\n", c1)
	fmt.Fprintf(w, "Cookies: %v\n", cs)
}
