package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/log"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/tools/clisubcmdgen/spec"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/descriptorpb"
)

const (
	descriptorPath = "tools/clisubcmdgen/tmp/protodesc.json"

	grpcSubCmdTemplatePath = "tools/clisubcmdgen/template/grpc_subcmd.tpl"
	grpcSubCmdOutDir       = "internal/cli/cmd/grpc"
	headerComment          = "Code generated by clisubcmd generator. DO NOT EDIT."
)

// 1サービスから1ファイルを生成する
func main() {
	desc, err := parseProtoFileDescriptor(descriptorPath)
	if err != nil {
		log.Errorf("failed to parse proto file: %v", err)
		os.Exit(1)
	}

	for _, file := range desc.GetFile() {
		if !strings.HasPrefix(file.GetName(), "service/") {
			continue
		}

		cmdSpec := spec.GenerateGRPCSubCmdSpec(file)
		for _, svcSpec := range cmdSpec.Services {
			outPath := fmt.Sprintf("%s/%s/root.gen.go", grpcSubCmdOutDir, svcSpec.PackageName)
			if err := executeTemplate(grpcSubCmdTemplatePath, outPath, svcSpec, map[string]interface{}{
				"headerComment": func() string {
					return headerComment
				},
			}); err != nil {
				log.Errorf("failed to execute template: %v", err)
				os.Exit(1)
			}

			log.Infof("generate success: %s => %s", file.GetName(), outPath)
		}
	}

	log.Infof("finish all")
}

// executeTemplate テンプレートを使ってファイルに書き込みを行う
func executeTemplate(
	templatePath, outFile string,
	spec interface{},
	funcMap map[string]interface{},
) error {
	if err := os.MkdirAll(filepath.Dir(outFile), 0700); err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to mkdir. dir=%s", filepath.Dir(outFile))
	}

	f, err := os.Create(outFile)
	if err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to open file. filename=%s", outFile)
	}
	defer f.Close()

	tpl, err := template.New(filepath.Base(templatePath)).Funcs(funcMap).ParseFiles(templatePath)
	if err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to new template. templatePath=%s", templatePath)
	}

	err = tpl.Execute(f, spec)
	if err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to execute template. templateName=%s", tpl.Name())
	}

	return nil
}

// parseProtoFileDescriptor
func parseProtoFileDescriptor(descriptorFilePath string) (*descriptorpb.FileDescriptorSet, error) {
	f, err := os.ReadFile(descriptorFilePath)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to open descriptor file. descriptorFilePath=%s", descriptorFilePath)
	}

	pbSet := new(descriptorpb.FileDescriptorSet)
	err = protojson.UnmarshalOptions{DiscardUnknown: true}.Unmarshal(f, pbSet)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to unmarshal descriptor file. descriptorFilePath=%s", descriptorFilePath)
	}

	return pbSet, nil
}
