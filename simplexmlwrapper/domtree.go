package simplexmlwrapper

import (
	"github.com/VictorLowther/simplexml/dom"
)

type DomTree struct {
	*dom.Document
	MaxDepth   int
	NodeCount  int
	DpNodes    []*Element
	NodesIndex map[string][]*Element
}

func NewDomTree(dom *dom.Document) *DomTree {
	d := &DomTree{dom, 0, 1, []*Element{}, make(map[string][]*Element)}
	d.PreOrderTraverse()
	return d
}

// outer wrapper for traverse
func (t *DomTree) PreOrderTraverse() {
	ps := t.preorderTraverse(&Element{t.Root()})

	// fmt.Println()
	// fmt.Println("ps: ")
	// for _, p := range ps {
	// 	fmt.Println(p.FullName(), p.GetDepth())
	// }
	// fmt.Println("====ps ends==== ")
	// fmt.Println()
	t.DpNodes = ps
	// fmt.Println("DpNodes count:\n\t", len(t.DpNodes))
}

// update nodeCount and maxDepth when traversing,
// return every node pointer at maxDepth
func (t *DomTree) preorderTraverse(p *Element) []*Element {

	res := []*Element{}

	// do sth, prepare return list, and terminate logic
	t.NodeCount++
	fname := p.FullName()
	t.NodesIndex[fname] = append(t.NodesIndex[fname], p)
	// fmt.Println(p.FullName(), p.GetDepth())
	if p.IsAtBottom() {
		check := t.UpdateDepth(p)
		if check == 1 {
			res = []*Element{}
			res = append(res, p)
		}
		if check == 0 {
			res = append(res, p)
		}
		return res
	}

	// Traverse each child
	for _, child := range p.Children() {
		cres := t.preorderTraverse(child)
		if len(res) > 0 &&
			len(cres) > 0 &&
			res[len(res)-1].GetDepth() < cres[0].GetDepth() {
			res = []*Element{}
		}
		res = append(res, cres...)
	}
	return res
}

// if update, return 1
// if equals, return 0
// if less, return -1
func (t *DomTree) UpdateDepth(p *Element) int {
	cur := p.GetDepth()
	if cur == t.MaxDepth {
		return 0
	}
	if cur > t.MaxDepth {
		t.MaxDepth = cur
		return 1
	}

	return -1
}
