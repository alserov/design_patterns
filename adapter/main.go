package main

import "fmt"

func main() {

}

// some service
type Service interface {
	SendXML()
}

// service data
type XMLDoc struct{}

func (receiver XMLDoc) SendXML() {
	fmt.Println("sending xml data")
}

// internal data
type JSONDoc struct{}

func (j JSONDoc) ConvertToXML() string {
	return "<xml></xml>"
}

// adapter
type JSONDocAdapter struct {
	JSONDoc *JSONDoc
}

func (j JSONDocAdapter) SendXML() {
	j.JSONDoc.ConvertToXML()
	fmt.Println("sending converted xml")
}
