package grpc

import (
	"crypto/tls"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/cmd"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/infra/grpc"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/terror"
	"github.com/spf13/cobra"
)

type Option struct {
	*cmd.RootOption

	host     string
	port     string
	insecure bool
	data     string
	dataFile string
}

// RequestData リクエストデータの取得
func (opt *Option) RequestData() ([]byte, error) {
	if opt.data == "" && opt.dataFile == "" {
		return []byte{}, nil
	}

	data := []byte(opt.data)
	var err error
	if opt.dataFile != "" {
		data, err = readFile(opt.dataFile)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

// NewGRPCClient
func (opt *Option) NewGRPCClient() (*grpc.GRPCClient, error) {
	var tlsConfig *tls.Config
	if !opt.insecure {
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}

	opt.Logf("[INFO] Connection gRPC %s:%s...", opt.host, opt.port)
	grpcClient, err := grpc.NewGRPCClient(opt.host, opt.port, tlsConfig)
	if err != nil {
		return nil, err
	}

	return grpcClient, nil
}

// NewGRPCCmd
func NewGRPCCmd(rootOpt *cmd.RootOption) (*cobra.Command, *Option) {
	opts := &Option{
		RootOption: rootOpt,
		// RequestTimeout: 10 * time.Second,
	}

	cmd := &cobra.Command{
		Use:   "grpc",
		Short: "Call gRPC",
		Long:  "Call gRPC",
		Example: strings.Join([]string{
			fmt.Sprintf("%s grpc <service_name> <rpc_name> [-h <host>] [-p <port>] [-d <message_data> | --data-file <message_json>]", opts.RootCmdName),
			fmt.Sprintf(`%s grpc <service_name> <rpc_name> -d '{id: 1, name: "sample"}'`, opts.RootCmdName),
			fmt.Sprintf("%s grpc <service_name> <rpc_name> --data-file message.json", opts.RootCmdName),
		}, "\n"),
		RunE: func(cmd *cobra.Command, args []string) error {
			return handleGRPCCmd(cmd, opts, args)
		},
	}

	// flags
	cmd.PersistentFlags().BoolP("help", "", false, "help for grpc command")
	cmd.PersistentFlags().BoolVarP(&opts.insecure, "insecure", "", false, "insecure gRPC connection")
	cmd.PersistentFlags().StringVarP(&opts.host, "host", "h", "localhost", "gRPC server host")
	cmd.PersistentFlags().StringVarP(&opts.port, "port", "p", "50051", "gRPC server port")
	cmd.PersistentFlags().StringVarP(&opts.data, "data", "d", "", "json structure request message of call gRPC")
	cmd.PersistentFlags().StringVar(&opts.dataFile, "data-file", "", "json file path for gRPC request message")

	return cmd, opts
}

// handleGRPCCmd
func handleGRPCCmd(cmd *cobra.Command, _ *Option, _ []string) error {
	if err := cmd.Help(); err != nil {
		return terror.Wrapf(terror.CodeInternal, err, "failed to cmd help")
	}
	return nil
}

func readFile(filePath string) ([]byte, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to open file")
	}

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, terror.Wrapf(terror.CodeInternal, err, "failed to read file")
	}

	return data, nil
}
