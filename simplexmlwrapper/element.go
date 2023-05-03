package simplexmlwrapper

import "github.com/VictorLowther/simplexml/dom"

// import "github.com/Vickko/dynamicXMLproto/domcore"
type Element struct {
	*dom.Element
}

func (node *Element) Children() (res []*Element) {
	for _, r := range node.Element.Children() {
		res = append(res, &Element{r})
	}
	return res
}

func (node *Element) Parent() *Element {

	return &Element{node.Element.Parent()}
}

// func (node *Element) Children() (res []domcore.Noder) {
// 	for _,r := range node.Element.Children() {
// 		res = append(res, domcore.Noder(&Element{*r}))
// 	}
// 	return res
// }

func (node *Element) IsAtBottom() bool {
	return len(node.Element.Children()) == 0
}

func (node *Element) IsMinimumMsg() bool {
	if len(node.Element.Children()) == 0 {
		return false
	}
	for _, r := range node.Children() {
		if !r.IsAtBottom() {
			return false
		}
	}
	return true
}

// TODO: replace string add with buffer
func (node *Element) FullName() string {
	return node.Name.Space + node.Name.Local
}

func (node *Element) GetDepth() int {
	res := 0
	n := node
	for n.Parent().Element!= nil {
		res ++
		n = n.Parent()
	}
	return res
}
