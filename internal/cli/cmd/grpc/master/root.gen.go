// Code generated by clisubcmd generator. DO NOT EDIT.

package master

import (
	"context"
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
		Use:   "master",
		Short: "Call MasterService rpc.",
		Long:  "Call MasterService rpc.",
		Example: strings.Join([]string{
			fmt.Sprintf("%s grpc master <rpc_name> [-d '{}']", opt.RootCmdName),
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

			client := pb.NewMasterServiceClient(grpcClient)

			switch rpcName {
			case "GetAll":
				return handleGetAll(cmd.Context(), opt, client, data)
			case "GetCard":
				return handleGetCard(cmd.Context(), opt, client, data)
			case "GetCharacter":
				return handleGetCharacter(cmd.Context(), opt, client, data)
			}

			return terror.Newf(terror.CodeInvalidArgument, "invalid rpc name.")
		},
	}

	return cmd, opt
}

func handleGetAll(
	ctx context.Context,
	opt *Option,
	cli pb.MasterServiceClient,
	data []byte,
) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req := &pb.Empty{}
	
	res, err := cli.GetAll(ctx, req)
	if err != nil {
		opt.Logf("returns error from server.")
		opt.Logf("%+s", err)
		return nil
	}

	opt.Logf("returns success from server.")
	opt.Format(res)
	return nil
}

func handleGetCard(
	ctx context.Context,
	opt *Option,
	cli pb.MasterServiceClient,
	data []byte,
) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req := &pb.Empty{}
	
	res, err := cli.GetCard(ctx, req)
	if err != nil {
		opt.Logf("returns error from server.")
		opt.Logf("%+s", err)
		return nil
	}

	opt.Logf("returns success from server.")
	opt.Format(res)
	return nil
}

func handleGetCharacter(
	ctx context.Context,
	opt *Option,
	cli pb.MasterServiceClient,
	data []byte,
) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	req := &pb.Empty{}
	
	res, err := cli.GetCharacter(ctx, req)
	if err != nil {
		opt.Logf("returns error from server.")
		opt.Logf("%+s", err)
		return nil
	}

	opt.Logf("returns success from server.")
	opt.Format(res)
	return nil
}