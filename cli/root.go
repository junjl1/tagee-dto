package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var (
	version = "v1.0.0"
)

var rootCmd = &cobra.Command{
	Use: "tagee-dto",
	Run: func(cmd *cobra.Command, args []string) {
		//if cmd.Flags().Changed("version") {
		//	fmt.Println("tagee-dto version", version)
		//	return
		//}
	},
}

func Execute() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	//rootCmd.Flags().BoolP("version", "v", false, "the version of tagee-dto")
	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {

}
