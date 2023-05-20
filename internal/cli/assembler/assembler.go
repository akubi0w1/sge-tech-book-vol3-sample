package assembler

import (
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/cmd"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/cmd/grpc"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/cmd/grpc/card"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/cmd/grpc/master"
	"github.com/akubi0w1/sge-tech-book-vol3-sample/internal/cli/cmd/grpc/user"
	"github.com/spf13/cobra"
)

// AddSubCommands adds all the subcommands to the rootCmd.
func AddSubCommands(rootCmd *cobra.Command, rootOpts *cmd.RootOption) {
	// rootCmd
	grpcCmd, grpcOpt := grpc.NewGRPCCmd(rootOpts)
	rootCmd.AddCommand(grpcCmd)

	// rootCmd grpc
	grpcMasterCmd, _ := master.NewCmd(grpcOpt)
	grpcUserCmd, _ := user.NewCmd(grpcOpt)
	grpcCardCmd, _ := card.NewCmd(grpcOpt)

	grpcCmd.AddCommand(grpcMasterCmd)
	grpcCmd.AddCommand(grpcUserCmd)
	grpcCmd.AddCommand(grpcCardCmd)
}
