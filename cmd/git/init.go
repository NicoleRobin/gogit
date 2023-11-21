package git

import (
	"fmt"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init a local git repo",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("init a local repo")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
}
