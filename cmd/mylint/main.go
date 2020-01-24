package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/nakabonne/mylint/passes/nestif"
)

func main() {
	singlechecker.Main(nestif.Analyzer)
}
