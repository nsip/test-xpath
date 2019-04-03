package main

// Experiment with:
//  - https://github.com/antchfx/xpath
//    . https://github.com/antchfx/xmlquery

// /NAPStudentResponseSet[PathTakenForDomain and not(PathTakenForDomain/@xsi:nil = ‘true’)]/TestletList/Testlet/ItemResponseList/ItemResponse/Response

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"os"
)

func main() {
	f, err := os.Open("sif.xml")
	doc, err := xmlquery.Parse(f)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loaded ok\n")
	fmt.Printf("DOC\n%+v\n", doc)

	channel := xmlquery.FindOne(doc, "//SchoolInfo")
	fmt.Printf("CHANNEL\n%+v\n", channel)

	// Query for Nick
	channel2 := xmlquery.FindOne(
		doc,
		"//NAPStudentResponseSet[PathTakenForDomain and not(PathTakenForDomain/@xsi:nil = 'true')]/TestletList/Testlet/ItemResponseList/ItemResponse/Response")
	fmt.Printf("CHANNEL2\n%+v\n", channel2.OutputXML(true))

}
