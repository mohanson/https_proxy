package main

import (
	"flag"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/mohanson/doa"
)

var (
	flCertFile = flag.String("cert", "", "path to tls certificate file")
	flPrivFile = flag.String("priv", "", "path to tls private key file")
)

func main() {
	flag.Parse()
	u := doa.Try2(url.Parse("http://127.0.0.1:80")).(*url.URL)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy := httputil.NewSingleHostReverseProxy(u)
		proxy.ServeHTTP(w, r)
	})
	doa.Try1(http.ListenAndServeTLS(":443", *flCertFile, *flPrivFile, nil))
}
