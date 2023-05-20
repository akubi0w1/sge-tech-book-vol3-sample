package spec

import (
	"fmt"
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Spec struct {
	Methods        []*MethodSpec
	MasterMessages []*MessageSpec
}

type MethodSpec struct {
	Name          string
	OutputMessage *MessageSpec
}

type MessageSpec struct {
	Name   string
	Fields []*MessageFieldSpec
}

type MessageFieldSpec struct {
	Name   string
	Type   string
	IsEnum bool
}

// GenerateMasterHandlerSpec
func GenerateMasterHandlerSpec(pbSet *descriptorpb.FileDescriptorSet) *Spec {
	spec := new(Spec)
	msgSpecMap := map[string]*MessageSpec{}
	for _, file := range pbSet.GetFile() {
		// マスター配布用のサービスを解析
		if file.GetName() == "service/master.proto" {
			for _, msg := range file.GetMessageType() {
				fields := make([]*MessageFieldSpec, 0, len(msg.GetField()))
				for _, msgField := range msg.GetField() {
					fields = append(fields, &MessageFieldSpec{
						Name: strcase.ToCamel(msgField.GetName()),
						Type: messageFieldType(msgField.GetTypeName()),
					})
				}
				msgSpec := &MessageSpec{
					Name:   msg.GetName(),
					Fields: fields,
				}
				msgSpecMap[fmt.Sprintf("*pb.%s", msg.GetName())] = msgSpec
			}

			for _, svc := range file.GetService() {
				methodSpecs := make([]*MethodSpec, 0, len(svc.GetMethod()))
				for _, method := range svc.GetMethod() {
					methodSpecs = append(methodSpecs, &MethodSpec{
						Name:          method.GetName(),
						OutputMessage: msgSpecMap[messageType(method.GetOutputType())],
					})
				}
				spec.Methods = methodSpecs
			}
		}

		// masterデータのメッセージを解析
		if strings.HasPrefix(file.GetName(), "master/") {
			for _, msg := range file.GetMessageType() {
				fields := make([]*MessageFieldSpec, 0, len(msg.GetField()))
				for _, msgField := range msg.GetField() {
					fields = append(fields, &MessageFieldSpec{
						Name:   strcase.ToCamel(msgField.GetName()),
						Type:   messageFieldType(msgField.GetTypeName()),
						IsEnum: strings.HasPrefix(msgField.GetTypeName(), ".enums"),
					})
				}
				spec.MasterMessages = append(spec.MasterMessages, &MessageSpec{
					Name:   msg.GetName(),
					Fields: fields,
				})
			}
		}
	}

	return spec
}

func messageFieldType(base string) string {
	sep := strings.Split(base, ".")
	return sep[len(sep)-1]
}

func messageType(base string) string {
	return fmt.Sprintf("*pb.%s", strings.TrimPrefix(base, ".service."))
}
