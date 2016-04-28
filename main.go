package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	. "github.com/CMU-15819O/rmckinst/fib"
	"golang.org/x/tools/go/loader"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

func main() {
	t1, t2 := Fib(7)
	fmt.Println(t1)
	fmt.Println(t2)
	//	fmt.Println(Fib(0))
	//	fmt.Println(Fib(1))
	//	fmt.Println(Fib(2))
	//	fmt.Println(Fib(3))
	//	fmt.Println(Fib(4))
	//	fmt.Println(Fib(5))
	//	fmt.Println(Fib(6))
	//	fmt.Println(Fib(7))
	RunSSA()
}

func main() {
	if len(os.Args) > 2 {
		log.Fatal("Please input the path to the file you want to analyze.")
	}

	filename := os.Args[1]

	var conf loader.Config
	filename := filepath.Join("./", filename) // ./fib/ + fib.go
	conf.CreateFromFilenames("fib", filename)
	prog, err := conf.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Receive the package you want to test
	// Pass in a package name relative to the current package -- like "fib"
	// Take the package, and fetch the main funciton out of it.

	// Then, run the analysis over the main function.
	//
}

// Analysis:
//
// An analysis consists of:
// 		Recursive call.
//		AnalyzeFunction(Takes a function and returns a function with the same number of args that transforms the args into a mapping)
//

type (

// Mapping of names to aliases?
// Oh, it's a set of ordered pairs!

)

func RunSSA() {

	pg := ssautil.CreateProgram(prog, ssa.PrintPackages|ssa.PrintFunctions|ssa.LogSource)
	pg.Build()
	for _, pack := range pg.AllPackages() {
		AnalyzePackage(pack)
	}
}

func GetMain(prog *ssa.Program) *ssa.Package {
	pkgs := prog.AllPackages()
	var main *ssa.Package = nil
	for _, pkg := range pkgs {
		if pkg.Pkg.Name() == "main" {
			main = pkg
			if main.Func("main") == nil {
				log.Fatal("no func main() in main package")
			}
			break
		}
	}

	if main == nil {
		log.Fatal("No main function in this package")
	}
	return main
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

// First thing we need to do is collect all of the global variables.
// Next, we need to find the main function.
// After that, we can collect the basicblocks, and iterate across them.

// Data Structures:
// ---------------------------------------------
// Need to specify the AbstractDomain
// Need to specify the flow functions
// Might need to specify the Join operator
// ---------------------------------------------
// Ok, then we need a framework for processing
// the ssa.Instructions
// We should have a function that takes an instruction and type switches over it, calling the right flow function based on the instruction type.
type Framework struct {
	analyses []Analysis
}

func (f *Framework) AddAnalysis(a Analysis) {
	f.analyses = append(f.analyses, a)
}

func (f *Framework) AnalyzeMain(*ssa.Function) {
	// Iterate over each of the blocks, and each of the instructions in each block, flowing through each instruction for all Analyses in the framework -- something something DomPreorder
}

type Analysis interface {
	Flow(*ssa.Instruction)
}
