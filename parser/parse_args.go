package parser

import (
	"flag"
	"strconv"
)

const (
	ModuleName = "mod"
	Home       = "home"
)

type parser struct {
	args map[string]string
}

var (
	Parser *parser
)

func (p *parser) parseArgs() {
	mod := flag.String(ModuleName, "", "module name to generate")
	homeDir := flag.Bool(Home, false, "if set, the project will be at home dir, else it will be at the pwd")
	flag.Parse()

	if *mod == "" {
		panic("Please, provide the module name to generate the go project")
	}

	p.args[ModuleName] = *mod
	p.args[Home] = strconv.FormatBool(*homeDir)
}

func (p *parser) GetModName() string {
	return p.args[ModuleName]
}

func (p *parser) GetIsHome() bool {
	return p.args[Home] == "true"
}

func init() {
	Parser = &parser{args: make(map[string]string)}
	Parser.parseArgs()
}
