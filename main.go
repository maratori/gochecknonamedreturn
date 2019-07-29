package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/maratori/gochecknonamedreturn/gochecknonamedreturn"
)

func main() {
	singlechecker.Main(gochecknonamedreturn.Analyzer)
}
