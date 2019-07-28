package gochecknonamedreturn

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer that reports usage of functions with named return value
var Analyzer = &analysis.Analyzer{ // nolint:gochecknoglobals
	Name:     "gochecknonamedreturn",
	Doc:      "report usage of functions with named return value",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	traverse := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector) // nolint:errcheck // let's panic
	filter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
	}
	traverse.Preorder(filter, func(node ast.Node) {
		switch fnNode := node.(type) {
		case *ast.FuncDecl:
			checkFunctionResults(pass, fnNode.Type.Results)
		case *ast.FuncLit:
			checkFunctionResults(pass, fnNode.Type.Results)
		}
	})
	return nil, nil
}

func checkFunctionResults(pass *analysis.Pass, results *ast.FieldList) {
	pos := namedReturnPos(results)
	if pos != token.NoPos {
		pass.Reportf(pos, "don't use named return values")
	}
}

// namedReturnPos finds position of the first named return value
func namedReturnPos(results *ast.FieldList) token.Pos {
	if results == nil {
		return token.NoPos
	}
	for _, types := range results.List {
		if types != nil {
			for _, name := range types.Names {
				if name != nil && name.Name != "" {
					return name.NamePos
				}
			}
		}
	}
	return token.NoPos
}
