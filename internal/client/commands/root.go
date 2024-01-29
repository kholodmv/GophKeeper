package commands

import (
	"github.com/kholodmv/GophKeeper/internal/client/commands/auth"
	"github.com/kholodmv/GophKeeper/internal/client/commands/keeper"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "welcome to client for a gophkeeper server",
}

func init() {
	rootCmd.AddCommand(auth.Cmd, keeper.Cmd)
	rootCmd.PersistentFlags().StringP("server", "a", "http://localhost:8080", "host address")
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
