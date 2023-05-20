// {{ headerComment }}

package {{ .PackageName }}

import (
	"context"
	{{ if not .IsRequestAllEmpty -}}
	"encoding/json"
	{{ end -}}
	"fmt"
	"strings"
	"time"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/cmd/grpc"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	pb "github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/service"
	"github.com/spf13/cobra"
)

type Option struct {
	*grpc.Option
}

func NewCmd(parentOpt *grpc.Option) (*cobra.Command, *Option) {
	opt := &Option{
		Option: parentOpt,
	}

	cmd := &cobra.Command{
		Use:   "{{ .SubCmdName }}",
		Short: "Call {{ .ServiceName }} rpc.",
		Long:  "Call {{ .ServiceName }} rpc.",
		Example: strings.Join([]string{
			fmt.Sprintf("%s grpc {{ .SubCmdName }} <rpc_name> [-d '{}']", opt.RootCmdName),
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return terror.Newf(terror.CodeInvalidArgument, "requires <rpc_name>")
			}
			rpcName := args[0]

			// get data body
			data, err := opt.RequestData()
			if err != nil {
				return err
			}

			grpcClient, err := opt.NewGRPCClient()
			if err != nil {
				return err
			}
			defer grpcClient.Close()

			client := pb.New{{ .ServiceName }}Client(grpcClient)

			switch rpcName {
			{{ range .Methods -}}
			case "{{ .MethodName }}":
				return handle{{ .MethodName }}(cmd.Context(), opt, client, data)
			{{ end -}}
			}

			return terror.Newf(terror.CodeInvalidArgument, "invalid rpc name.")
		},
	}

	return cmd, opt
}
{{ range .Methods }}
func handle{{ .MethodName }}(
	ctx context.Context,
	opt *Option,
	cli pb.{{ $.ServiceName }}Client,
	data []byte,
) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	{{ if eq .InputType "Empty" -}}
	req := &pb.Empty{}
	{{- else -}}
	req := new(pb.{{ .InputType }})
	if err := json.Unmarshal(data, &req); err != nil {
		return terror.Wrapf(terror.CodeInvalidArgument, err, "failed to unmarshal {{ .InputType }} request.")
	}
	{{- end }}
	
	res, err := cli.{{ .MethodName }}(ctx, req)
	if err != nil {
		opt.Logf("returns error from server.")
		opt.Logf("%+s", err)
		return nil
	}

	opt.Logf("returns success from server.")
	opt.Format(res)
	return nil
}
{{ end -}}
