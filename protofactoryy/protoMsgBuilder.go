package protofactory
import (
	log "github.com/Vickko/dynamicXMLproto/logplus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type ProtoMsgBuilder struct {
	descriptorpb.DescriptorProto
}

func NewMsgBuilder() *ProtoMsgBuilder {
	p := ProtoMsgBuilder{}
	p.Options = new(descriptorpb.MessageOptions)
	return &p
}

func (p *ProtoMsgBuilder) IsInitialized() bool {
	return p.Options != nil
}

func (p *ProtoMsgBuilder) InitializationGuard(s string) {
	if !p.IsInitialized() {
		log.Errorf("Attemp to %s on uninitialized MsgBuilder", s)
	}
}

func (p *ProtoMsgBuilder) SetName(s string) *ProtoMsgBuilder {
	p.InitializationGuard("SetName")
	p.Name = proto.String(s)
	return p
}

func (p *ProtoMsgBuilder) GetName() string {
	p.InitializationGuard("GetName")
	if p.Name == nil {
		return ""
	}
	return *p.Name
}

func (p *ProtoMsgBuilder) AppendField(fields ...*descriptorpb.FieldDescriptorProto) *ProtoMsgBuilder {
	p.InitializationGuard("AppendField")
	p.Field = append(p.Field, fields...)
	return p
}

func (p *ProtoMsgBuilder) Set(Name string, Fields ...*descriptorpb.FieldDescriptorProto) *ProtoMsgBuilder {
	p.SetName(Name)
	p.AppendField(Fields...)
	return p
}

func (p *ProtoMsgBuilder)RegisterNestedMsgs(msgs ...*descriptorpb.DescriptorProto) *ProtoMsgBuilder {
	p.InitializationGuard("RegisterNestedMsgs")
	p.NestedType = append(p.NestedType, msgs...)
	return p
}


func (p *ProtoMsgBuilder) ExportMsg() *descriptorpb.DescriptorProto {
	return &p.DescriptorProto
}
