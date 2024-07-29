package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

// printCmd represents the print command
var printCmd = &cobra.Command{
	Use:   "gen",
	Short: "gen gen gen",
	Args:  cobra.ExactArgs(1), // 确保传递一个参数
	Run: func(cmd *cobra.Command, args []string) {
		// 获取字符串参数
		str := args[0]

		// 打印字符串
		fmt.Println("Printing:", str)
	},
}

func init() {
	rootCmd.AddCommand(printCmd)
}
