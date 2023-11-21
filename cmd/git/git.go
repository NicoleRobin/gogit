package git

import (
	"os"

	"github.com/nicolerobin/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gogit",
	Short: "gogit is a git implementation in pure go",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

var Verbose bool

func init() {
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "V", false, "verbose output")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error("rootCmd.Execute() failed, err:%s", err)
		os.Exit(1)
	}
}
