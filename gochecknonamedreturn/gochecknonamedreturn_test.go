package gochecknonamedreturn_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/maratori/gochecknonamedreturn/gochecknonamedreturn"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()
	testdata := analysistest.TestData()
	analysistest.Run(t, testdata, gochecknonamedreturn.Analyzer)
}
