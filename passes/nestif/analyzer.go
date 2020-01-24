package nestif

import (
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "nestif",
	Doc:  "detect deeply nested if statements",
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	var issues []Issue
	for _, file := range pass.Files {

		checker := &Checker{
			MinComplexity: 1,
		}
		i := checker.Check(file, pass.Fset)
		issues = append(issues, i...)
	}
	for _, issue := range issues {
		pass.Reportf(issue.Pos, "%s", issue.Message)
	}

	return nil, nil
}
