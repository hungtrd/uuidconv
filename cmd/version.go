package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of uuidconv",
	Long:  `All software has versions. This is uuidconv's`,
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Println("uuidconv v0.3.0")
	},
}
