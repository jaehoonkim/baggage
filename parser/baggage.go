package parser

import (
	"encoding/xml"
)

type Baggage struct {
	XMLName  xml.Name  `xml:"baggage"`
	Version  string    `xml:"version,attr"`
	Projects []Project `xml:"project"`
}

type Project struct {
	XMLName xml.Name `xml:"project"`
	Name    string   `xml:"name,attr"`
	Members Members  `xml:"members"`
	Items   Items    `xml:"items"`
}

type Members struct {
	XMLName xml.Name `xml:"members"`
	Member  []Member `xml:"member"`
}

type Member struct {
	XMLName xml.Name `xml:"member"`
	Name    string   `xml:"name,attr"`
}

type Items struct {
	XMLName xml.Name `xml:"items"`
	Item    []Item   `xml:"item"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Name    string   `xml:"name,attr"`
	Owner   string   `xml:"owner,attr"`
	Checked string   `xml:"checked,attr"`
}

func (b *Baggage) AddProject(name string) {
	prj := Project{Name: name}
	b.Projects = append(b.Projects, prj)
}
