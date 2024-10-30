package gh

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/thorstenhans/spin-gh-plugin/internal/detective"
	gh "github.com/thorstenhans/spin-gh-plugin/internal/github"
)

type CreateActionOptions struct {
	Output       string
	Name         string
	ToolVersions gh.ToolVersions
	TemplatePath string
	Plugins      []string
	Overwrite    bool
	BranchName   string
}

var options CreateActionOptions = CreateActionOptions{
	ToolVersions: gh.DefaultToolVersions(),
}
var createActionCmd = &cobra.Command{
	Use:   "create-action",
	Short: "Examines your Spin App and creates a working GitHub Action for CI",
	Run: func(cmd *cobra.Command, args []string) {
		apps := detective.FindAllSpinApps()
		if len(apps) == 0 {
			fmt.Println("Could not find Spin App(s) under the current directory")
			os.Exit(1)
		}

		renderOptions := gh.RenderActionOptions{
			SpinApps:           apps,
			Output:             options.Output,
			Overwrite:          options.Overwrite,
			Plugins:            options.Plugins,
			TargetBranch:       options.BranchName,
			CustomTemplatePath: options.TemplatePath,
			Name:               options.Name,
			ToolVersions:       options.ToolVersions,
		}
		err := gh.RenderAction(renderOptions)
		if err != nil {
			fmt.Printf("Error while rendering template %v", err)
			os.Exit(1)
		}
	},
}

func init() {
	defaultPlugins := []string{"js2wasm", "py2wasm", "kube"}
	createActionCmd.Flags().StringVarP(&options.Output, "output", "o", ".github/workflows/ci.yaml", "Path where the GitHub Action will be created")
	createActionCmd.Flags().StringVarP(&options.Name, "name", "n", "CI", "Name for the GitHub Action")

	createActionCmd.Flags().StringVarP(&options.ToolVersions.Rust, "rust-version", "", options.ToolVersions.Rust, "Set Rust version for GitHub Actions")
	createActionCmd.Flags().StringVarP(&options.ToolVersions.Go, "go-version", "", options.ToolVersions.Go, "Set Go version for GitHub Actions")
	createActionCmd.Flags().StringVarP(&options.ToolVersions.TinyGo, "tinygo-version", "", options.ToolVersions.TinyGo, "Set TinyGo version for GitHub Actions")
	createActionCmd.Flags().StringVarP(&options.ToolVersions.Node, "node-version", "", options.ToolVersions.Node, "Set Node.js version for GitHub Actions")
	createActionCmd.Flags().StringVarP(&options.ToolVersions.Python, "python-version", "", options.ToolVersions.Python, "Set Python version for GitHub Actions")
	createActionCmd.Flags().StringVarP(&options.ToolVersions.Spin, "spin-version", "", "", "Set Spin version for GitHub Actions (default: current stable Spin release)")

	createActionCmd.Flags().BoolVarP(&options.Overwrite, "overwrite", "", false, "Overwrite existing output file")
	createActionCmd.Flags().StringVarP(&options.TemplatePath, "template", "t", "", "Specify the path to a custom template for creating the GitHub Action")
	createActionCmd.Flags().StringSliceVarP(&options.Plugins, "plugins", "p", defaultPlugins, "Specify required Spin plugins")
	createActionCmd.Flags().StringVarP(&options.BranchName, "branch", "b", "main", "Specify the desired branch")
	rootCmd.AddCommand(createActionCmd)
}
