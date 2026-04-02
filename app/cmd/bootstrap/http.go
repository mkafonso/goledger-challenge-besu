package bootstrap

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mkafonso/goledger-challenge-besu/config"
	httpinfra "github.com/mkafonso/goledger-challenge-besu/infra/http"
)

func StartServer(h *httpinfra.Handlers) {
	router := httpinfra.NewRouter(h)

	addr := config.Env.APIHost + ":" + config.Env.APIPort
	fmt.Println("HTTP Server running on:", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
