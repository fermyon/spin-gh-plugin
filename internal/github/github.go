package github

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/thorstenhans/spin-gh-plugin/internal/spinapp"
)

type RenderActionOptions struct {
	SpinApps           []*spinapp.App
	SpinVersion        string
	Output             string
	Overwrite          bool
	Name               string
	CustomTemplatePath string
	TargetBranch       string
	Plugins            []string
	ToolVersions       ToolVersions
}

func RenderAction(options RenderActionOptions) error {
	templateContent, err := getTemplateContents(options.CustomTemplatePath)
	if err != nil {
		return err
	}
	templ, err := template.New("ci.yaml").Parse(templateContent)
	if err != nil {
		return err
	}
	actionFile, err := getActionFile(options)
	if err != nil {
		return err
	}
	defer actionFile.Close()

	data := templateData{
		Rust:            isLanguageInUse(options.SpinApps, spinapp.Rust),
		Go:              isLanguageInUse(options.SpinApps, spinapp.GoLang),
		JavaScript:      isLanguageInUse(options.SpinApps, spinapp.JavaScript),
		Python:          isLanguageInUse(options.SpinApps, spinapp.Python),
		ActionName:      options.Name,
		TargetBranch:    options.TargetBranch,
		OperatingSystem: "ubuntu-latest",
		SpinPlugins:     strings.Join(options.Plugins, ","),
		ToolVersions:    options.ToolVersions,
		SpinApps:        newSpinAppMetadata(options.SpinApps),
	}

	err = templ.Execute(actionFile, data)
	if err != nil {
		return err
	}
	return nil
}

func isLanguageInUse(apps []*spinapp.App, lang spinapp.Language) bool {
	inUse := false
	for _, app := range apps {
		if slices.Contains(app.GetLanguages(), lang) {
			inUse = true
		}
	}
	return inUse
}

func getActionFile(options RenderActionOptions) (*os.File, error) {
	if !options.Overwrite {
		if _, err := os.Stat(options.Output); err == nil {
			return nil, fmt.Errorf("pass --force to overwrite an existing GitHub Action file")
		}
	}
	if _, err := os.Stat(options.Output); err == nil {
		if err := os.Remove(options.Output); err != nil {
			return nil, err
		}
	}

	if err := os.MkdirAll(filepath.Dir(options.Output), os.ModePerm); err != nil {
		return nil, err
	}

	file, err := os.Create(options.Output)
	if err != nil {
		return nil, err
	}
	return file, nil
}

//go:embed templates/_default.yaml.tmpl
var actionsTemplate string

func GetDefaultTemplate() string {
	return actionsTemplate
}

func getTemplateContents(customTemplateFilePath string) (string, error) {
	if len(customTemplateFilePath) == 0 {
		return actionsTemplate, nil
	}
	content, err := os.ReadFile(customTemplateFilePath)
	if err != nil {
		return "", fmt.Errorf("custom template not found: %w", err)
	}
	return string(content), nil
}

type templateData struct {
	Rust            bool
	Go              bool
	JavaScript      bool
	Python          bool
	SpinPlugins     string
	ActionName      string
	TargetBranch    string
	OperatingSystem string
	ToolVersions    ToolVersions
	SpinApps        []spinAppMetadata
}

type spinAppMetadata struct {
	Name string
	Path string
}

func newSpinAppMetadata(apps []*spinapp.App) []spinAppMetadata {
	res := make([]spinAppMetadata, len(apps))
	for idx, app := range apps {
		res[idx] = spinAppMetadata{
			Name: app.GetName(),
			Path: app.GetLocation(),
		}
	}
	return res
}

type ToolVersions struct {
	Rust   string
	Go     string
	TinyGo string
	Node   string
	Python string
	Spin   string
}

func DefaultToolVersions() ToolVersions {
	return ToolVersions{
		Rust:   "1.80.1",
		Go:     "1.23.2",
		TinyGo: "0.33.0",
		Python: "3.13.0",
		Node:   "22",
		Spin:   "",
	}
}
