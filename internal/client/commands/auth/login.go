package auth

import (
	"fmt"
	"github.com/kholodmv/GophKeeper/internal/client/storage"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "login",
	Long: `The loginCmd command represents the login functionality, used for user authorization.
		   The command takes a username and password as arguments and returns a token,
		   which can be used for subsequent authenticated requests.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		username := cmd.Flag("username").Value.String()
		password := cmd.Flag("password").Value.String()
		tokenString, err := authService.Login(username, password)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Println("User has logged in")
		ts := storage.New()
		err = ts.SaveToken(tokenString)
		if err != nil {
			fmt.Println(err)
			return err
		}
		return nil
	},
}

func init() {
	Cmd.AddCommand(loginCmd)
}
