// Copyright 2023 Vickko. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package protofactory

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/dynamicpb"
)

type ProtoMessage struct {
	protoreflect.MessageDescriptor
}

func (p *ProtoMessage) NewInstance() *ProtoInstance {
	return &ProtoInstance{
		*dynamicpb.NewMessage(p.MessageDescriptor),
	}
}

// TODO: need a stringer to format output?
// func (p *ProtoMessage)String() string {
// 	return "MessageDescriptor{\n\t" + fmt.Sprint(p.FullName())
// }
