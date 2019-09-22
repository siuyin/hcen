package main

import (
	"fmt"

	"github.com/siuyin/dflt"
)

func main() {
	stage := dflt.EnvString("STAGE", "Test")
	fmt.Printf("HCEN %s: Recommendations and price check", stage)
}
