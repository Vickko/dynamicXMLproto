package protofactory

import (
	"unicode"

	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

var optional = descriptorpb.FieldDescriptorProto_Label(protoreflect.Optional).Enum()

type ProtoFieldBuilder struct {
	descriptorpb.FieldDescriptorProto
}

func NewFieldBuilder() *ProtoFieldBuilder {
	p := ProtoFieldBuilder{}
	p.Options = new(descriptorpb.FieldOptions)
	return &p
}

// Invalid Kinds below:
// protoreflect.BoolKind,
// protoreflect.BytesKind,
// protoreflect.DoubleKind,
// protoreflect.EnumKind,
// protoreflect.Fixed32Kind,
// protoreflect.Fixed64Kind,
// protoreflect.FloatKind,
// protoreflect.GroupKind,
// protoreflect.Int32Kind,
// protoreflect.Int64Kind,
// protoreflect.MessageKind,
// protoreflect.Sfixed32Kind,
// protoreflect.Sfixed64Kind,
// protoreflect.Sint32Kind,
// protoreflect.Sint64Kind,
// protoreflect.StringKind,
// protoreflect.Uint32Kind,
// protoreflect.Uint64Kind.
// only these kinds' exact prefix are permitted,
// otherwise we assume that it is a user-defined msg type.
func ParseType(s string) *descriptorpb.FieldDescriptorProto_Type {
	var k protoreflect.Kind
	switch s {
	case "Bool":
		k = protoreflect.BoolKind
	case "Bytes":
		k = protoreflect.BytesKind
	case "Double":
		k = protoreflect.DoubleKind
	case "Enum":
		k = protoreflect.EnumKind
	case "Fixed32":
		k = protoreflect.Fixed32Kind
	case "Fixed64":
		k = protoreflect.Fixed64Kind
	case "Float":
		k = protoreflect.FloatKind
	case "Group":
		k = protoreflect.GroupKind
	case "Int32":
		k = protoreflect.Int32Kind
	case "Int64":
		k = protoreflect.Int64Kind
	case "Message":
		k = protoreflect.MessageKind
	case "Sfixed32Kind":
		k = protoreflect.Sfixed32Kind
	case "Sfixed64Kind":
		k = protoreflect.Sfixed64Kind
	case "Sint32Kind":
		k = protoreflect.Sint32Kind
	case "Sint64Kind":
		k = protoreflect.Sint64Kind
	case "String":
		k = protoreflect.StringKind
	case "Uint32Kind":
		k = protoreflect.Uint32Kind
	case "Uint64Kind":
		k = protoreflect.Uint64Kind
	// If an unknown type name is provided,
	// We assume that it is a user-defined msg
	default:
		k = protoreflect.MessageKind
	}
	return descriptorpb.FieldDescriptorProto_Type(k).Enum()
}

func ParseLabel(label string) *descriptorpb.FieldDescriptorProto_Label {
	var res *descriptorpb.FieldDescriptorProto_Label
	switch label {
	case "":
		res = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
	case "optional":
		res = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
	case "Optional":
		res = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
	case "OPTIONAL":
		res = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
	case "required":
		res = descriptorpb.FieldDescriptorProto_LABEL_REQUIRED.Enum()
	case "Required":
		res = descriptorpb.FieldDescriptorProto_LABEL_REQUIRED.Enum()
	case "REQUIRED":
		res = descriptorpb.FieldDescriptorProto_LABEL_REQUIRED.Enum()
	case "repeated":
		res = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
	case "Repeated":
		res = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
	case "REPEATED":
		res = descriptorpb.FieldDescriptorProto_LABEL_REPEATED.Enum()
	default:
		res = descriptorpb.FieldDescriptorProto_LABEL_OPTIONAL.Enum()
	}
	return res
}

// field content differs according to type is msg or not.
func NewField(label, typeName, name string, num int32) *descriptorpb.FieldDescriptorProto {
	t := ParseType(typeName)
	res := &descriptorpb.FieldDescriptorProto{}
	if *t == descriptorpb.FieldDescriptorProto_TYPE_MESSAGE {
		res.Name = proto.String(name)
		res.Number = proto.Int32(num)
		res.Type = t
		res.TypeName = proto.String(string(unicode.ToUpper(rune(name[0]))) + name[1:]) // Titled
	} else {
		res.Name = proto.String(name)
		res.Number = proto.Int32(num)
		res.Label = ParseLabel(label)
		res.Type = t
		res.JsonName = proto.String(name)
	}
	return res
}
