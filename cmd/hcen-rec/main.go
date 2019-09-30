package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/ibmdb/go_ibm_db"
	"github.com/siuyin/dflt"
)

func main() {
	stage := dflt.EnvString("STAGE", "Test")
	fmt.Printf("HCEN %s: Recommendations and price check", stage)

	passwd := dflt.EnvString("DB2_PASSWORD", "aPassword")
	con := fmt.Sprintf("HOSTNAME=hcen-dev-db2;DATABASE=testdb;PORT=50000;UID=db2inst1;PWD=%s", passwd)
	db, err := sql.Open("go_ibm_db", con)
	if err != nil {

		fmt.Println(err)
	}
	db.Close()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Rec World.")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
