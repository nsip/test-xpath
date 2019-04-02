package main

// https://github.com/beevik/etree

import (
	"github.com/beevik/etree"
  "fmt"
)

func main() {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile("sif.xml"); err != nil {
		panic(err)
	}
	fmt.Printf("Loaded\n")

  // channel := xmlquery.FindOne(doc, "//SchoolInfo")
  // fmt.Printf("CHANNEL\n%+v\n", channel)

  // Query for Nick
  // channel2 := xmlquery.FindOne(
  //  doc,
  //  "/NAPStudentResponseSet[PathTakenForDomain and not(PathTakenForDomain/@xsi:nil = ‘true’)]/TestletList/Testlet/ItemResponseList/ItemResponse/Response")
  // fmt.Printf("CHANNEL2\n%+v\n", channel2)

}
