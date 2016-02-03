package parser

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Parser struct {
	Path string
	Bag  Baggage
}

func NewParser(p string) *Parser {
	return &Parser{Path: p, Bag: Baggage{Version: "0.1"}}
}

func exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil && os.IsExist(err) == false {
		return false
	}
	return true
}

func (p *Parser) Parse() {
	var f *os.File
	var err error
	if exist(p.Path) == false {
		f, err = os.Create(p.Path)
		if err != nil {
			panic(err)
		}

	} else {
		f, err = os.Open(p.Path)
		if err != nil {
			panic(err)
		}

		data, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}

		err = xml.Unmarshal(data, &p.Bag)
		if err != nil {
			panic(err)
		}
	}
	defer f.Close()
}

func (p *Parser) Save(path string) {
	output, err := xml.MarshalIndent(p.Bag, "", "	")
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(path, output, 0644)
	if err != nil {
		panic(err)
	}
}
