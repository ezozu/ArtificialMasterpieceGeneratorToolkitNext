// cmd/artificialmasterpiecegeneratortoolkitnext/main.go
package main

import (
"flag"
"log"
"os"

"artificialmasterpiecegeneratortoolkitnext/internal/artificialmasterpiecegeneratortoolkitnext"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := artificialmasterpiecegeneratortoolkitnext.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
