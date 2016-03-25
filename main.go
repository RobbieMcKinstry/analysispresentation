package main

import (
	"fmt"
	"log"
	"path/filepath"

	. "github.com/CMU-15819O/rmckinst/fib"
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	fmt.Println(Fib(0))
	fmt.Println(Fib(1))
	fmt.Println(Fib(2))
	fmt.Println(Fib(3))
	fmt.Println(Fib(4))
	fmt.Println(Fib(5))
	fmt.Println(Fib(6))
	fmt.Println(Fib(7))
	RunSSA()
}

func RunSSA() {
	var conf loader.Config
	filename := filepath.Join("./fib", "fib.go")
	conf.CreateFromFilenames("fib", filename)
	prog, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	pg := ssautil.CreateProgram(prog, ssa.PrintPackages|ssa.PrintFunctions|ssa.LogSource)
	pg.Build()
	for _, pack := range pg.AllPackages() {
		AnalyzePackage(pack)
	}
}

func AnalyzePackage(pack *ssa.Package) {
	for _, member := range pack.Members {
		switch member.(type) {
		case *ssa.NamedConst:
		case *ssa.Global:
		case *ssa.Function:
		case *ssa.Type:
		}
		fmt.Printf("Name: %v\n", member.Name())
		fmt.Printf("Type: %v\n", member.Type())
		fmt.Printf("Token: %v\n", member.Token())
		fmt.Println()
	}
}
