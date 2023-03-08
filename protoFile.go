// Copyright 2023 Vickko. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package dynamicxmlproto

import "google.golang.org/protobuf/reflect/protoreflect"

type ProtoFile struct {
	protoreflect.FileDescriptor
}

func (p *ProtoFile) GetMessage(name string) *ProtoMessage {
	return &ProtoMessage{
		p.Messages().ByName(protoreflect.Name(name)),
	}
}
