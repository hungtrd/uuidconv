package cmd

import (
	"fmt"
	"os"

	"github.com/hungtrd/uuidconv/pkg/uuid"
	"github.com/spf13/cobra"
)

var (
	from string
	to   string
	cp   string
)

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&from, "from", "f", "", "From format: string, base64, base62")
	rootCmd.PersistentFlags().StringVarP(&to, "to", "t", "", "To format: string, base64, base62")
	rootCmd.PersistentFlags().StringVarP(&cp, "copy", "c", "string", "Copy format to clipboard: string, base64, base62")

	rootCmd.AddCommand(convertCmd)
}

func initConfig() {
	// Do stuff here
}

var rootCmd = &cobra.Command{
	Use:   "uuidconv",
	Short: "generate or decode UUID string to base64 encoded and vice versa",
	Long:  "generate or decode UUID string to base64 encoded and vice versa",
	Run: func(cmd *cobra.Command, _ []string) {
		newUUID(cmd)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func newUUID(_ *cobra.Command) {
	u := uuid.New()

	printUUID(u)
	copyToClipboard(u, uuid.Format(cp))
}
