package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/ibmdb/go_ibm_db"
	"github.com/nats-io/stan.go"

	"github.com/siuyin/dflt"
)

func main() {
	stage := dflt.EnvString("STAGE", "Test")
	fmt.Printf("HCEN %s: Point of Sale\n", stage)

	passwd := dflt.EnvString("DB2_PASSWORD", "aPassword")
	con := fmt.Sprintf("HOSTNAME=hcen-dev-db2;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=%s", passwd)
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {

		fmt.Println(err)
	}
	defer db.Close()

	clientID := "hcen-pos"
	clusterID := "test-cluster"
	natsURL := dflt.EnvString("NATS_URL", "nats://hcen-dev:4222")
	sc, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		log.Fatalf("Could not connect to NATS at %s: %v", natsURL, err)
	}
	defer sc.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello POS World.")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
