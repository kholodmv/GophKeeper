package keeper

import (
	"fmt"
	"github.com/kholodmv/GophKeeper/internal/client/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get command",
	Long:  `The get command.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		title := cmd.Flag("title").Value.String()
		secret, err := keeperService.Get(title)
		if err != nil {
			fmt.Println(err)
			return err
		}
		if secret.Binary != nil {
			err = os.WriteFile(secret.FilePath, secret.Binary, 0644)
			if err != nil {
				log.Printf("Failed writing to file: %s\n", err)
			}
			fmt.Println("secret: " + secret.Title + " the file has been saved in " + secret.FilePath)
		} else {
			utils.ShowResult(secret)
		}
		return nil
	},
}

func init() {
	Cmd.AddCommand(getCmd)
}
