package git

import (
	"fmt"
	"github.com/nicolerobin/gogit/constant"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of gogit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("gogit %s\n", constant.Version)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
