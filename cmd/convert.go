package cmd

import (
	"errors"
	"fmt"
	"slices"

	"github.com/atotto/clipboard"
	"github.com/hungtrd/uuidconv/pkg/uuid"
	"github.com/spf13/cobra"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Convert back and forth between different formats",
	Long: `Convert back and forth between different formats: string, base64, base62.
            The default is to print all formats.`,
	Args: func(_ *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("requires one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		convertUUID(cmd, args)
	},
}

func convertUUID(cmd *cobra.Command, args []string) {
	input := args[0]

	var err error
	from, err := cmd.Flags().GetString("from")
	if err != nil {
		fmt.Println(err)
	}

	var u uuid.UUID
	switch uuid.Format(from) {
	case uuid.FormatString:
		u, err = uuid.NewFromString(input)
		if err != nil {
			fmt.Println(err)
		}
	case uuid.FormatBase64:
		u, err = uuid.NewFromBase64(input)
		if err != nil {
			fmt.Println(err)
		}
	case uuid.FormatBase62:
		u, err = uuid.NewFromBase62(input)
		if err != nil {
			fmt.Println(err)
		}
	default:
		u, err = uuid.NewFromUnknow(input)
		if err != nil {
			fmt.Println(err)
		}
	}

	to, err := cmd.Flags().GetString("to")
	if err != nil {
		fmt.Println(err)
	}
	switch uuid.Format(to) {
	case uuid.FormatString:
		printUUID(u, uuid.Format(to))
		copyToClipboard(u, uuid.Format(to))
	case uuid.FormatBase64:
		printUUID(u, uuid.Format(to))
		copyToClipboard(u, uuid.Format(to))
	case uuid.FormatBase62:
		printUUID(u, uuid.Format(to))
		copyToClipboard(u, uuid.Format(to))
	default:
		printUUID(u)
		cp, err := cmd.Flags().GetString("copy")
		if err != nil {
			fmt.Println(err)
		}
		copyToClipboard(u, uuid.Format(cp))
	}
}

func printUUID(u uuid.UUID, formats ...uuid.Format) {
	fmt.Println("-------------------------")
	if slices.Contains(formats, uuid.FormatString) || len(formats) == 0 {
		fmt.Println("UUID string:    ", u.NormalString)
	}
	if slices.Contains(formats, uuid.FormatBase64) || len(formats) == 0 {
		fmt.Println("Base64 encoded: ", u.Base64Encoded)
	}
	if slices.Contains(formats, uuid.FormatBase62) || len(formats) == 0 {
		fmt.Println("Base62 encoded: ", u.Base62Encoded)
	}
	fmt.Println("-------------------------")
}

func copyToClipboard(u uuid.UUID, format uuid.Format) {
	var s string
	switch format {
	case uuid.FormatString:
		s = u.NormalString
	case uuid.FormatBase64:
		s = u.Base64Encoded
	case uuid.FormatBase62:
		s = u.Base62Encoded
	default:
		s = u.NormalString
	}
	if err := clipboard.WriteAll(s); err != nil {
		fmt.Printf("write to clipboard failed: %v", err)
	}
	fmt.Printf("copied %s(%s) to clipboard!\n", format, s)
}
