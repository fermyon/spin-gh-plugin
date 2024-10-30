package gh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	gh "github.com/thorstenhans/spin-gh-plugin/internal/github"
)

var output string
var overwrite bool
var ejectCmd = &cobra.Command{
	Use:   "eject",
	Short: "Render the default GitHub Actions template",
	Run: func(cmd *cobra.Command, args []string) {
		template := gh.GetDefaultTemplate()
		if output != "" {
			if !overwrite {
				if _, err := os.Stat(output); err == nil {
					fmt.Printf("File %s already exists. Use --overwrite to overwrite it.\n", output)
					return
				}
			}
			err := os.WriteFile(output, []byte(template), 0644)
			if err != nil {
				fmt.Printf("Failed to write to file %s: %v\n", output, err)
				return
			}
			fmt.Printf("Template written to %s\n", output)
		}
		fmt.Println(template)
	},
}

func init() {
	ejectCmd.Flags().StringVarP(&output, "output", "o", "", "Path where the default template should be rendered to")
	ejectCmd.Flags().BoolVarP(&overwrite, "overwrite", "", false, "Overwrite existing output file")
	rootCmd.AddCommand(ejectCmd)
}
