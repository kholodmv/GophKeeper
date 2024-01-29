package auth

import (
	"github.com/kholodmv/GophKeeper/internal/client/services/auth"
	"github.com/spf13/cobra"
)

var (
	// authService is a service used for a command implementation.
	authService auth.AuthService
	// Cmd represents the auth command.
	Cmd = &cobra.Command{
		Use:   "auth",
		Short: "authorization and registration commands",
		Long:  "A parent command for login and register.",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			baseURL := cmd.Flag("server").Value.String()
			authService = auth.NewAuthService(baseURL)
		},
	}
)

func init() {
	Cmd.PersistentFlags().StringP("username", "u", "", "username to authorize")
	Cmd.PersistentFlags().StringP("password", "p", "", "password to authorize")

	for _, flag := range []string{"username", "password"} {
		Cmd.MarkPersistentFlagRequired(flag)
	}
}
