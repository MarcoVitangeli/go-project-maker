package main

import (
	"github.com/MarcoVitangeli/go-project-maker/cli"
	"github.com/MarcoVitangeli/go-project-maker/parser"
)

func main() {
	modName := parser.Parser.GetModName()
	runner := cli.NewCommandRunner()
	var err error

	if parser.Parser.GetIsHome() {
		err = runner.GoHome()
	}

	if err != nil {
		panic(err)
	}

	err = runner.CreateDirAndGo(modName, 0777)
	if err != nil {
		panic(err)
	}

	err = runner.InitGoMod(modName)

	if err != nil {
		panic(err)
	}

	err = runner.CreateMainGoFile()

	if err != nil {
		panic(err)
	}
}
