package simplexmlwrapper

import (
	"fmt"
	"strings"
	"testing"

	"github.com/VictorLowther/simplexml/dom"
)

var data string = `
<?xml version="1.0" encoding="UTF-8"?>
<resume>
<personal-info>
  <name>John Doe</name>
  <email>john.doe@email.com</email>
  <phone>555-1234</phone>
  <address>
	<street>123 Main St</street>
	<city>Anytown</city>
	<state>CA</state>
	<zip>12345</zip>
  </address>
</personal-info>
<education>
  <degree>Bachelor of Science</degree>
  <major>Computer Science</major>
  <school>University of California, Los Angeles</school>
  <grad-year>2020</grad-year>
</education>
<experience>
  <job>
	<title>Software Engineer Intern</title>
	<company>Google</company>
	<start-date>June 2019</start-date>
	<end-date>September 2019</end-date>
  </job>
  <job>
	<title>Software Engineer</title>
	<company>Apple</company>
	<start-date>January 2020</start-date>
	<end-date>Present</end-date>
  </job>
</experience>
</resume>`

func Test_PreOrderTraverse(t *testing.T) {
	doc, _ := dom.Parse(strings.NewReader(data))
	tree := NewDomTree(doc)
	_ = tree
	fmt.Println("NodeCount & MaxDepth:\n\t", tree.NodeCount, tree.MaxDepth)
	
}

func Test_ContentAndChild(t *testing.T) {
	data1 := `<abc>123</abc>`
	data2 := `<cba><abc>123</abc></cba>`
	doc, _ := dom.Parse(strings.NewReader(data2))
	tree := NewDomTree(doc)
	r := tree.Root()
	fmt.Println("content & children")
	fmt.Println(string(r.Content), ",", r.Children())
	_ = data1
	_ = data2
}
