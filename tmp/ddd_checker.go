package main

import (
	"fmt"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Define the layers
const (
	Interface      = "interface"
	Application    = "application"
	Domain         = "domain"
	Infrastructure = "infrastructure"
	Unknown        = "unknown"
)

// Allowed dependencies
var allowedDependencies = map[string][]string{
	Interface:      {Application, Domain},
	Application:    {Domain},
	Infrastructure: {Application, Domain},
	Domain:         {},
}

func getLayer(path string) string {
	switch {
	case strings.Contains(path, "/internal/interface/"):
		return Interface
	case strings.Contains(path, "/internal/application/"):
		return Application
	case strings.Contains(path, "/internal/domain/"):
		return Domain
	case strings.Contains(path, "/internal/infrastructure/"):
		return Infrastructure
	default:
		return Unknown
	}
}

func main() {
	violations := false
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !strings.HasSuffix(path, ".go") {
			return nil
		}

		fromLayer := getLayer(path)
		if fromLayer == Unknown {
			return nil
		}

		fset := token.NewFileSet()
		node, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			return err
		}

		for _, imp := range node.Imports {
			importPath := strings.Trim(imp.Path.Value, "`")
			if !strings.HasPrefix(importPath, "github.com/FeisalDy/go-ddd") {
				continue
			}

			toLayer := getLayer(importPath)
			if toLayer == Unknown {
				continue
			}

			if fromLayer == toLayer {
				continue
			}

			allowed := false
			for _, allowedDep := range allowedDependencies[fromLayer] {
				if toLayer == allowedDep {
					allowed = true
					break
				}
			}

			if !allowed {
				fmt.Printf("DDD Violation: %s layer (%s) cannot depend on %s layer (%s)\n", fromLayer, path, toLayer, importPath)
				violations = true
			}
		}
		return nil
	})

	if err != nil {
		fmt.Printf("Error walking the path: %v\n", err)
		os.Exit(1)
	}

	if violations {
		os.Exit(1)
	} else {
		fmt.Println("No DDD violations found.")
	}
}
