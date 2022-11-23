package main

import (
        "os"
	"fmt"
        "log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"

        "github.com/gorilla/handlers"
)

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "hello_processed_total",
		Help: "The total number of processed calls",
	})
)

func main() {
	log.Println("Starting hello-world server...")
	mux := http.NewServeMux()
	mux.Handle("/metrics", handlers.LoggingHandler(os.Stdout, promhttp.Handler()))
	mux.HandleFunc("/", helloServer)

	log.Fatal(http.ListenAndServe(ListenPort(), mux))
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")

	log.Println("increasing hello_processed_total counter")
	opsProcessed.Inc()
}

func ListenPort() string {
        return fmt.Sprintf(":%s", getEnvVar("PORT", "8080"))
}

func getEnvVar(key, def string) string {
        if value, ok := os.LookupEnv(key); ok {
                return value
        }
        return def
}
