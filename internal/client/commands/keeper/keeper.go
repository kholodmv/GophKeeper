package keeper

import (
	"github.com/kholodmv/GophKeeper/internal/client/services/keeper"
	"github.com/spf13/cobra"
)

var (
	keeperService keeper.KeeperService
	// Cmd represents the auth command.
	Cmd = &cobra.Command{
		Use:   "keeper",
		Short: "keeper commands",
		Long:  "A parent command for secrets.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			baseURL := cmd.Flag("server").Value.String()
			keeperService = keeper.NewKeeperService(baseURL)
		},
	}
)

func init() {
	Cmd.PersistentFlags().StringP("title", "t", "", "title")
	Cmd.PersistentFlags().StringP("comment", "k", "", "comment")

	Cmd.MarkPersistentFlagRequired("title")
}
