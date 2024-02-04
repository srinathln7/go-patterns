package main

import "fmt"

// Formatter: Interface that formats a given text
type Formatter interface {
	format(string) string
}

type PlainText struct{}

func (pt *PlainText) format(str string) string {
	return str
}

type BoldText struct{}

func (bt *BoldText) format(str string) string {
	return "/b( " + str + " /b)"
}

type ItalicText struct{}

func (it *ItalicText) format(str string) string {
	return "/i( " + str + " /i)"
}

type UnderlineText struct{}

func (it *UnderlineText) format(str string) string {
	return "__________" + str + "_________"
}

func main() {

	str := "Hello, Srinath!"
	pt := &PlainText{}
	bt := &BoldText{}
	it := &ItalicText{}
	ut := &UnderlineText{}

	fmt.Println(pt.format(str))
	fmt.Println(bt.format(str))
	fmt.Println(it.format(str))
	fmt.Println(ut.format(str))

	// But what if we want bold-italic combo, italic-underline or bold-italic-underline combo ?
	// How to accomplish that without writing extra code?

	// We will have to do chaining and this will result in spagethi like code
	// You will have to pass in the req'd string in every function
	fmt.Println(bt.format(fmt.Sprint(it.format(str))))
}
