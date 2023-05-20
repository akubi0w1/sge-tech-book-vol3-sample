package cmd

import (
	"io"

	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/output"
	"github.com/spf13/cobra"
)

const (
	CmdName string = "api-cli"
)

var (
	CmdVersion string
)

// RootOption ルートコマンドのオプション値
type RootOption struct {
	RootCmdName string

	formatter string
	output.Output
}

// NewCmdRoot
func NewCmdRoot(outWriter, errWriter io.Writer) (*cobra.Command, *RootOption) {
	opts := &RootOption{
		RootCmdName: CmdName,
	}

	cmd := &cobra.Command{
		Use:           CmdName,
		Version:       CmdVersion,
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	cmd.SetOut(outWriter)
	cmd.SetErr(errWriter)

	// flags
	cmd.PersistentFlags().StringVarP(&opts.formatter, "formatter", "f", "json", "select output format. [ json | console ]")

	cmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		var err error
		opts.Output, err = output.NewOutputWithFormatter(opts.formatter, outWriter)
		if err != nil {
			return err
		}

		return nil
	}

	return cmd, opts
}
