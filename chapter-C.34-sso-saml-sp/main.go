package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/crewjam/saml/samlsp"
)

var (
	samlCertificatePath = "./myservice.cert"
	samlPrivateKeyPath  = "./myservice.key"
	samlIDPMetadata     = "https://samltest.id/saml/idp"

	webserverPort    = 9000
	webserverRootURL = fmt.Sprintf("http://localhost:%d", webserverPort)
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := samlsp.Token(r.Context()).Attributes.Get("displayName")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	sp, err := newSamlMiddleware()
	if err != nil {
		log.Fatal(err.Error())
	}

	http.Handle("/", sp.RequireAccount(
		http.HandlerFunc(hello),
	))

	http.Handle("/saml/", sp)

	portString := fmt.Sprintf(":%d", webserverPort)
	fmt.Println("server started at", portString)
	http.ListenAndServe(portString, nil)
}
