package pipeline

import (
	"sort"

	"github.com/Vickko/dynamicXMLproto/logplus"
	"github.com/Vickko/dynamicXMLproto/protofactory"
	"github.com/Vickko/dynamicXMLproto/simplexmlwrapper"
	"google.golang.org/protobuf/types/descriptorpb"
)

func NodeIsDuplicated(nlist []*simplexmlwrapper.Element, n *simplexmlwrapper.Element) bool {
	for _, e := range nlist {
		if e.FullName() == n.FullName() {
			return true
		}
	}
	return false
}

func GetMinParentMsgNode(p *simplexmlwrapper.Element) *simplexmlwrapper.Element {
	res := p.Parent()
	if res.IsMinimumMsg() {
		return res
	}
	return nil
}

func GetMinParentMsgNodes(ps []*simplexmlwrapper.Element) []*simplexmlwrapper.Element {
	res := []*simplexmlwrapper.Element{}
	for _, p := range ps {
		pr := p.Parent()
		// fmt.Println(&pr)
		if pr.IsMinimumMsg() && !NodeIsDuplicated(res, pr) {
			res = append(res, pr)
		}
	}
	return res
}

func RepeatedChildren(p *simplexmlwrapper.Element) map[string]int {
	appearance := make(map[string]int)
	for _, child := range p.Children() {
		_, exist := appearance[child.FullName()]
		if exist {
			appearance[child.FullName()]++
		} else {
			appearance[child.FullName()] = 1
		}
	}
	return appearance
}

func BuildFieldProtoDescFromNode(p *simplexmlwrapper.Element, isRepeated bool, index int) *descriptorpb.FieldDescriptorProto {
	t := protofactory.JudgeValueType(string(p.Content))
	if len(p.Children()) != 0 || len(p.Content) == 0 {
		// TODO: Msg logic
		logplus.Warningln("BuildFieldProtoDescFromNode, ", p, "is a msg")
		t = p.FullName()
	}
	// TODO: optional and required?
	r := ""
	if isRepeated {
		r = "repeated"
	}
	return protofactory.NewField(r, t, p.FullName(), int32(index))

}

func BuildMsgProtoDescFromNode(p *simplexmlwrapper.Element) *descriptorpb.DescriptorProto {
	fields := []*descriptorpb.FieldDescriptorProto{}
	repeatedMap := RepeatedChildren(p)
	done := map[string]bool{}
	index := 0
	for _, field := range p.Children() {
		if done[field.FullName()] {
			continue
		}
		index++
		isRepeated := false
		if repeatedMap[field.FullName()] > 1 {
			isRepeated = true
		}
		fields = append(fields, BuildFieldProtoDescFromNode(field, isRepeated, index))
		done[field.FullName()] = true
	}
	return protofactory.NewMsgBuilder().Set(p.FullName(), fields...).ExportMsg()
}

// TODO: now sort fields by dict order, need a more realistic sorting scheme
// TODO: add a map to mark only message in the whole tree
func MergeMsgFromRelatives(d *simplexmlwrapper.DomTree, p *simplexmlwrapper.Element) *descriptorpb.DescriptorProto {
	relatives, exist := d.NodesIndex[p.FullName()]
	if !exist {
		// TODO: Error handling logic
		logplus.Errorln("BuildMsgProtoDescFromNode, ", p.FullName(), " not found")
	}
	fieldsMap := make(map[string][]*descriptorpb.FieldDescriptorProto)
	for _, relative := range relatives {
		for _, field := range BuildMsgProtoDescFromNode(relative).Field {
			fieldsMap[field.GetJsonName()] = append(fieldsMap[field.GetJsonName()], field)
		}
	}

	// TODO: add type Consistency check
	res := protofactory.NewMsgBuilder().Set(p.FullName())
	index := 0

	fieldNameList := make([]string, 0, len(fieldsMap))
	for k := range fieldsMap {
		fieldNameList = append(fieldNameList, k)
	}
	sort.Strings(fieldNameList)
	fieldListList := make([][]*descriptorpb.FieldDescriptorProto, 0, len(fieldsMap))
	for _, k := range fieldNameList {
		fieldListList = append(fieldListList, fieldsMap[k])
	}

	for _, fieldList := range fieldListList {
		index++
		label := ""
		typeName := ""
		_ = typeName
		for _, fieldInstance := range fieldList {
			if *fieldInstance.Label == *descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum() {
				label = "repeated"
			}
			// if fieldInstance.GetTypeName()
			// fmt.Printf("fieldInstance.Type: %v\n", fieldInstance.Type)
			// fmt.Printf("protofactory.ParseTypeString(fieldInstance.Type): %v\n", protofactory.ParseTypeString(fieldInstance.Type))
			typeName = protofactory.ParseTypeString(fieldInstance.Type)
		}
		// fmt.Println("l: ", label)
		currentField := protofactory.NewField(label, typeName, *fieldList[0].JsonName, int32(index))

		// if label == "repeated" {
		// 	fmt.Println("cf: ", currentField)
		// }

		res.AppendField(currentField)
	}
	return res.ExportMsg()
}
