package cli

import (
	"github.com/junjl1/tagee-dto/fetcher"
	"github.com/spf13/cobra"
)

var printCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen gen gen",
	Args:  cobra.ExactArgs(1), // 确保传递一个参数
	Run: func(cmd *cobra.Command, args []string) {
		str := args[0]
		fetcher.Fetch(str)
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
