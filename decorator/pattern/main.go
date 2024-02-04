package main

import "fmt"

// Formatter: Interface that formats a given text
type Formatter interface {
	format() string
}

type PlainText struct {
	Text string
}

func (pt *PlainText) format() string {
	return pt.Text
}

type BoldText struct {
	component Formatter
}

func (bt *BoldText) format() string {
	return "/b( " + bt.component.format() + " /b)"
}

type ItalicText struct {
	component Formatter
}

func (it *ItalicText) format() string {
	return "/i( " + it.component.format() + " /i)"
}

type UnderlineText struct {
	component Formatter
}

func (ut *UnderlineText) format() string {
	return "__________" + ut.component.format() + "_________"
}

func main() {

	str := "Hello, Srinath!"

	pt := &PlainText{Text: str}
	bt := &BoldText{component: pt}
	it := &ItalicText{component: pt}
	ut := &UnderlineText{component: pt}

	fmt.Println("********************* BASIC OPS ********************************")
	fmt.Println(pt.format())
	fmt.Println(bt.format())
	fmt.Println(it.format())
	fmt.Println(ut.format())

	// But what if we want bold-italic combo, italic-underline or bold-italic-underline combo ?
	// How to accomplish that without writing extra code?

	fmt.Println("********************* HYBRID OPS ********************************")

	// bold-italic combo
	biC := &ItalicText{component: bt}
	fmt.Println(biC.format())

	// italic-underline combo
	iuC := &UnderlineText{component: it}
	fmt.Println(iuC.format())

	// boild-italic-underline combo
	biuC := &UnderlineText{component: biC}
	fmt.Println(biuC.format())
}
