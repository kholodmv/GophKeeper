package keeper

import (
	"fmt"
	"github.com/kholodmv/GophKeeper/internal/client/models"
	"github.com/kholodmv/GophKeeper/internal/utils/binary"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "add secret command",
		Long:  `add secret command.`,
		RunE: func(cmd *cobra.Command, args []string) error {

			secret := &models.Secret{}

			secret.Title = cmd.Flag("title").Value.String()
			secret.Username = cmd.Flag("username").Value.String()
			secret.Password = cmd.Flag("password").Value.String()
			secret.Comment = cmd.Flag("comment").Value.String()
			secret.Number = cmd.Flag("num").Value.String()
			secret.Date = cmd.Flag("date").Value.String()
			secret.Text = cmd.Flag("string").Value.String()
			secret.Cvv = cmd.Flag("cvv").Value.String()
			secret.FilePath = cmd.Flag("bin").Value.String()
			if secret.FilePath != "" {
				bin, err := binary.ReadFile(secret.FilePath)
				if err != nil {
					fmt.Println(err)
					return err
				}

				secret.Binary = bin
			}
			err := keeperService.Create(secret)
			if err != nil {
				fmt.Println(err)
				return err
			}
			fmt.Println("the secret has been added to server")
			return nil
		},
	}
)

func init() {
	Cmd.PersistentFlags().StringP("username", "u", "", "username")
	Cmd.PersistentFlags().StringP("password", "p", "", "password")
	Cmd.PersistentFlags().StringP("num", "n", "", "card number")
	Cmd.PersistentFlags().StringP("date", "d", "", "date")
	Cmd.PersistentFlags().StringP("string", "s", "", "string")
	Cmd.PersistentFlags().StringP("bin", "b", "", "binary")
	Cmd.PersistentFlags().StringP("cvv", "c", "", "cvv code")
	Cmd.AddCommand(createCmd)
}
