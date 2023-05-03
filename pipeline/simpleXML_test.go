package pipeline

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"

	"github.com/Vickko/dynamicXMLproto/protofactory"
	"github.com/Vickko/dynamicXMLproto/simplexmlwrapper"
	"github.com/VictorLowther/simplexml/dom"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
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

var data1 string = `
<?xml version="1.0" encoding="UTF-8"?>
<address>
<street>123 Main St</street>
<city>Anytown</city>
<state>CA</state>
<zip>12345</zip>
</address>`

func Test_Parse(t *testing.T) {
	doc, _ := dom.Parse(strings.NewReader(data))
	fmt.Println(doc)
	r := doc.Root()
	fmt.Println("children: \n\t", r.Children()[0].Children())

}

func judgeValueType(str string) string {
	n, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		if n < math.MaxInt32 {
			return "Int32"
		} else {
			return "Int64"
		}
	}
	_, err = strconv.ParseFloat(str, 64)
	if err == nil {
		return "Double"
	}
	_, err = strconv.ParseBool(str)
	if err == nil {
		return "Bool"
	}
	return "String"
}

func Test_SetField(t *testing.T) {
	str := "12345"
	var value any

	if n1, err := strconv.ParseInt(str, 10, 64); err == nil {
		if n1 < math.MaxInt32 {
			value = int32(n1)
			fmt.Println(1)
		} else {
			value = int64(n1)
		}
	} else if n2, err := strconv.ParseFloat(str, 64); err == nil {
		value = n2
	} else if n3, err := strconv.ParseBool(str); err == nil {
		value = n3
	} else {
		value = str
	}

	fmt.Println(value)
	fmt.Printf("%T\n", value)
}

func Test_engine(t *testing.T) {

	// XML document
	doc, _ := dom.Parse(strings.NewReader(data1))
	// root element of the DOM tree
	r := doc.Root()
	fmt.Println("original DOM data: ")
	fmt.Println(r)

	// protobuf reflect file created by protofactory
	f := protofactory.NewFileBuilder().SetConfig(
		"ai.momenta/hdmap/data/mdf/test.proto",
		"proto3",
		"hdmap.data.mdf",
		"ai.momenta/hdmap/data/mdf",
	)

	// DOM element wrapper aim to apply new functions
	wr := simplexmlwrapper.Element{Element: r}
	//fmt.Println(string(wr.Children()[0].Content))

	fmt.Println("building msg fields:")
	// create protobuf reflect msg using protofactory
	msg := protofactory.NewMsgBuilder().SetName(wr.Name.Space + wr.Name.Local)
	for i, child := range wr.Children() {
		msg.AppendField(
			protofactory.NewField(
				"",
				judgeValueType(string(child.Content)),
				child.Name.Space+child.Name.Local,
				int32(i+1),
			))
		fmt.Println(judgeValueType(string(child.Content)),
			child.Name.Space+child.Name.Local,
			int32(i+1))
	}
	// push msg into file
	f.AppendMsg(msg.ExportMsg())

	fmt.Println()
	fmt.Println("assign value for msg instance: ")
	// build new instance of the msg just created
	adr := f.MsgInstance("address")
	// assign value for each property of the instance
	for _, child := range wr.Children() {
		fmt.Println(child.Name.Space+child.Name.Local, string(child.Content))
		adr.SetFieldByType(child.Name.Space+child.Name.Local, string(child.Content))
	}
	fmt.Println()
	fmt.Println("Final msg: ")
	fmt.Println(adr)
	// f.AppendMsg(
	// 	protofactory.NewMsgBuilder().Set(wr.Name.Space+wr.Name.Local,
	// 		protofactory.NewField("String", "firstName", 1),
	// 		protofactory.NewField("String", "lastName", 2),
	// 		protofactory.NewField("Bool", "isAlive", 3),
	// 		protofactory.NewField("Int64", "age", 4),
	// 	).ExportMsg())
}

func Test_GetMinParentMsgNodes(t *testing.T) {
	doc, _ := dom.Parse(strings.NewReader(data))
	tree := simplexmlwrapper.NewDomTree(doc)
	_ = tree
	fmt.Println("NodeCount & MaxDepth:\n\t", tree.NodeCount, tree.MaxDepth)
	minMsg := GetMinParentMsgNodes(tree.DpNodes)
	_ = minMsg
	fmt.Println("minMsg:\n\t", minMsg)
}

func Test_RepeatedChildren(t *testing.T) {
	var data1 string = `
	<job>
	  <title>Software Engineer Intern</title>
	  <title>Software Engineer Intern</title>
	  <title>Software Engineer Intern</title>
	  <title>Software Engineer Intern</title>
	  <company>Google</company>
	  <start-date>June 2019</start-date>
	  <end-date>September 2019</end-date>
  	</job>
	`
	_ = data1
	doc, _ := dom.Parse(strings.NewReader(data1))
	tree := simplexmlwrapper.NewDomTree(doc)
	res := RepeatedChildren(&simplexmlwrapper.Element{tree.Root()})
	fmt.Println(res)
	fmt.Println()
	desc := BuildMsgProtoDescFromNode(&simplexmlwrapper.Element{tree.Root()})
	fmt.Println(desc)
	fmt.Printf("desc.GetField()[0].Label: %v\n", desc.GetField()[0].Label)
	fmt.Println(*desc.GetField()[0].Label == *descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum())
	fmt.Println(*desc.GetField()[0].Type)
	fmt.Println(*desc.GetField()[0].Type == *descriptorpb.FieldDescriptorProto_Type(protoreflect.StringKind).Enum())

}

func Test_relatives(t *testing.T) {
	var data1 string = `
	<jobs>
	  <job>
	    <title>Software Engineer Intern</title>
	    <title>Hardware Engineer Intern</title>
	    <company>Google</company>
	    <start-date>June 2019</start-date>
	    <end-date>September 2019</end-date>
  	  </job>
	  <job>
	    <title>Game Engineer Intern</title>
	    <title>Web Engineer Intern</title>
	    <end-date>September 2019</end-date>
		<a>123</a>
		<b>abc</b>
  	  </job>
	</jobs>
	`
	doc, _ := dom.Parse(strings.NewReader(data1))
	tree := simplexmlwrapper.NewDomTree(doc)
	job0 := &simplexmlwrapper.Element{tree.Root().Children()[0]}
	res := MergeMsgFromRelatives(tree, job0)
	fmt.Println("res: \n\t", res)

}
