package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/nats-io/stan.go"
	"github.com/siuyin/dflt"
)

func main() {
	stage := dflt.EnvString("STAGE", "Test")
	fmt.Printf("HCEN %s: Recommendations and price check", stage)
	clientID := "hcen-rec"
	clusterID := "test-cluster"
	natsURL := dflt.EnvString("NATS_URL", "nats://hcen-dev:4222")
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Could not connect to NATS at %s", natsURL)
	}
	defer sc.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Rec World.")
		sc.Publish("hello", []byte(fmt.Sprintf("Hello: %s", time.Now().Format("15:04:05.000"))))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
