// Copyright 2023 Vickko. All rights reserved.
// Use of this source code is governed by MIT
// license that can be found in the LICENSE file.

package dynamicxmlproto

import (
	log "github.com/Vickko/dynamicXMLproto/logplus"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
)

type ProtoFileBuilder struct {
	descriptorpb.FileDescriptorProto
}

func NewFileBuilder() *ProtoFileBuilder {
	p := ProtoFileBuilder{}
	p.Options = new(descriptorpb.FileOptions)
	return &p
}

func (p *ProtoFileBuilder) IsInitialized() bool {
	return p.Options != nil
}

func (p *ProtoFileBuilder) InitializationGuard(s string) {
	if !p.IsInitialized() {
		log.Errorf("Attemp to %s on uninitialized FileBuilder", s)
	}
}

func (p *ProtoFileBuilder) SetName(s string) *ProtoFileBuilder {
	p.InitializationGuard("SetName")
	p.Name = proto.String(s)
	return p
}

func (p *ProtoFileBuilder) GetName() string {
	p.InitializationGuard("GetName")
	if p.Name == nil {
		return ""
	}
	return *p.Name
}

func (p *ProtoFileBuilder) SetSyntax(s string) *ProtoFileBuilder {
	p.InitializationGuard("SetSyntax")
	if s != "proto2" && s != "proto3" {
		log.Logln("Invalid syntax string")
		// TODO: now only print a warning but do nothing, considering do a panic or sth
	}
	p.Syntax = proto.String(s)
	return p
}

func (p *ProtoFileBuilder) GetSyntax() string {
	p.InitializationGuard("GetSyntax")
	if p.Syntax == nil {
		return ""
	}
	return *p.Syntax
}

func (p *ProtoFileBuilder) SetPackage(s string) *ProtoFileBuilder {
	p.InitializationGuard("SetPackage")
	p.Package = proto.String(s)
	return p
}

func (p *ProtoFileBuilder) GetPackage() string {
	p.InitializationGuard("GetPackage")
	if p.Package == nil {
		return ""
	}
	return *p.Package
}

// TODO: more language package options function needs implement
func (p *ProtoFileBuilder) SetGoPackage(pkg string) *ProtoFileBuilder {
	p.InitializationGuard("SetGoPackage")
	p.Options.GoPackage = proto.String(pkg)
	return p
}

func (p *ProtoFileBuilder) GetGoPackage() string {
	p.InitializationGuard("GetGoPackage")
	return p.Options.GetGoPackage()
}

// TODO: Msg series methods needs implement, like findby, delete or sth
func (p *ProtoFileBuilder) AppendMsg(msgs ...*descriptorpb.DescriptorProto) *ProtoFileBuilder {
	p.InitializationGuard("AppendMsg")
	p.MessageType = append(p.MessageType, msgs...)
	return p
}

func (p *ProtoFileBuilder) SetConfig(Name, Syntax, Package, GoPackage string) *ProtoFileBuilder {
	// return p.SetName(Name).SetSyntax(Syntax).SetPackage(Package).SetGoPackage(GoPackage)
	p.SetName(Name)
	p.SetSyntax(Syntax)
	p.SetPackage(Package)
	p.SetGoPackage(GoPackage)
	return p
}

func (p *ProtoFileBuilder) EnoughInfo() bool {
	p.InitializationGuard("CheckEnoughInfo")
	return p.Name != nil && p.Syntax != nil && p.Package != nil && p.Options.GoPackage != nil && len(p.MessageType) != 0
}

func (p *ProtoFileBuilder) registerFile() *protoregistry.Files {
	set := descriptorpb.FileDescriptorSet{}
	set.File = append(set.File, &p.FileDescriptorProto)
	f, err := protodesc.NewFiles(&set)
	if err != nil {
		log.Errorln(err)
	}
	return f
}

func (p *ProtoFileBuilder) Build() *ProtoFile {
	if !p.EnoughInfo() {
		log.Errorln("Not enough Info to build a FileDescriptor")
	}

	fd, err := protodesc.NewFile(&p.FileDescriptorProto, p.registerFile())
	if err != nil {
		log.Errorln(err)
	}
	return &ProtoFile{fd}
}

func (p *ProtoFileBuilder) MsgInstance(name string) *ProtoInstance {
	return p.Build().GetMessage(name).NewInstance()
}

// TODO: maybe implement a stringer? for debug use
