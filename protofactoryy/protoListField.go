package protofactory

import "google.golang.org/protobuf/reflect/protoreflect"

type ProtoListField struct {
	protoreflect.List
	Kind protoreflect.Kind
}

func (p *ProtoListField) ListAppend(value any) *ProtoListField {
	p.Append(ParseValue(p.Kind, value))
	return p
}

func (p *ProtoListField) ListSet(index int, value any) *ProtoListField {
	p.Set(index, ParseValue(p.Kind, value))
	return p
}

// TODO: add pop, remove and so on