package spec

import (
	"strings"

	"github.com/iancoleman/strcase"
	"google.golang.org/protobuf/types/descriptorpb"
)

type GRPCSubCmdSpec struct {
	// Services 定義されているサービス
	Services []*ServiceSpec
}

// ServiceSpec
type ServiceSpec struct {
	// IsRequestAllEmpty リクエストがすべてEmpty
	IsRequestAllEmpty bool
	// PackageName goのパッケージ名
	PackageName string
	// SubCmdName サブコマンド名
	SubCmdName string
	// ServiceName protobuf定義のサービス名
	ServiceName string
	// Comment サービスのコメント
	Comment string
	// Methods 定義されているメソッド
	Methods []*MethodSpec
}

// MethodSpec
type MethodSpec struct {
	// MethodName メソッド名
	MethodName string
	// InputType リクエストの型名
	InputType string
	// OutputType レスポンスの型名
	OutputType string
	// Comment コメント
	Comment string
}

// GenerateGRPCSubCmdSpec
func GenerateGRPCSubCmdSpec(file *descriptorpb.FileDescriptorProto) *GRPCSubCmdSpec {
	// game_serviceのprotoを解析する
	serviceSpec := make([]*ServiceSpec, 0, len(file.GetService()))
	for _, svc := range file.GetService() {
		methodSpecs, isAllRequestEmpty := generateMethodSpec(svc)

		serviceSpec = append(serviceSpec, &ServiceSpec{
			IsRequestAllEmpty: isAllRequestEmpty,

			PackageName: strings.ToLower(strings.TrimSuffix(svc.GetName(), "Service")),
			SubCmdName:  strcase.ToKebab(strings.TrimSuffix(svc.GetName(), "Service")),
			ServiceName: svc.GetName(),
			// Comment:     strings.Join(commentBuffer, "\n"),
			Methods: methodSpecs,
		})
	}

	return &GRPCSubCmdSpec{
		Services: serviceSpec,
	}
}

// generateMethodSpec methodをspecに変換する
// methodspec, existEmpty, isAllRequestEmpty が返却される
func generateMethodSpec(svc *descriptorpb.ServiceDescriptorProto) ([]*MethodSpec, bool) {
	methodSpecs := make([]*MethodSpec, 0, len(svc.GetMethod()))
	isAllRequestEmpty := true
	for _, method := range svc.GetMethod() {
		methodSpecs = append(methodSpecs, &MethodSpec{
			MethodName: method.GetName(),
			InputType:  extractTypeName(method.GetInputType()),
			OutputType: extractTypeName(method.GetOutputType()),
			// Comment: method.GetOptions(),
		})

		if method.GetInputType() != ".service.Empty" {
			isAllRequestEmpty = false
		}
	}

	return methodSpecs, isAllRequestEmpty
}

// extractTypeName
func extractTypeName(protoType string) string {
	packagePath := strings.Split(protoType, ".")
	return packagePath[len(packagePath)-1]
}
