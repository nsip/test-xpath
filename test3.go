package main

// https://github.com/subchen/go-xmldom

import (
	"github.com/subchen/go-xmldom"
  "fmt"
)

func main() {
	doc := xmldom.Must(xmldom.ParseFile("sif.xml"))
  fmt.Printf("Loaded ok\n")
	fmt.Printf("DOC\n%+v\n", doc)

	root := doc.Root
	nodeList := root.Query("//SchoolInfo")
	for _, c := range nodeList {
		fmt.Printf("%v: RefId = %v\n", c.Name, c.GetAttributeValue("RefId"))
	}

	nodeList2 := root.Query("/NAPStudentResponseSet[PathTakenForDomain and not(PathTakenForDomain/@xsi:nil = ‘true’)]/TestletList/Testlet/ItemResponseList/ItemResponse/Response")
	for _, c := range nodeList2 {
		fmt.Printf("%v: RefId = %v\n", c.Name, c.GetAttributeValue("RefId"))
	}

  // channel := xmlquery.FindOne(doc, "//SchoolInfo")
  // channel2 := xmlquery.FindOne(
  //   doc,
  //   "/NAPStudentResponseSet[PathTakenForDomain and not(PathTakenForDomain/@xsi:nil = ‘true’)]/TestletList/Testlet/ItemResponseList/ItemResponse/Response")

}
