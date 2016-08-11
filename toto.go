package main

import (
    "fmt"
        "net/http"
	)

type BasicAuthFunc func(username, password string) bool

func (f BasicAuthFunc) RequireAuth(w http.ResponseWriter) {
    w.Header().Set("WWW-Authenticate", `Basic realm="Authorization Required"`)
        w.WriteHeader(401)
	}

func (f BasicAuthFunc) Authenticate(r *http.Request) bool {
    username, password, ok := r.BasicAuth()
        return ok && f(username, password)
	}

func home(w http.ResponseWriter, r *http.Request) {
    // Identifiant et mot de passe attendu
        username, password := "foo", "bar"

    // Création d'une fonction de test de type BasicAuthFunc
        f := BasicAuthFunc(func(user, pass string) bool {
	        return username == user && password == pass
		    })

    if !f.Authenticate(r) {
            f.RequireAuth(w)
	        } else {
		        fmt.Fprintf(w, "<h1>%s</h1>", "Welcome to the private zone")
			    }
			    }

func main() {
    fmt.Println("Server is running...")

    http.Handle("/", http.HandlerFunc(home))
        http.ListenAndServe(":3000", nil)
	}