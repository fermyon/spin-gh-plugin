package spinapp

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
)

const ManifestName = "spin.toml"

type Language struct {
	Name string
	hint string
}

type App struct {
	location  string
	languages []Language
	name      string
}

type spinManifest struct {
	App spinManifestApp `toml:"application"`
}
type spinManifestApp struct {
	Name string `toml:"name"`
}

func (app *App) GetName() string {
	return app.name
}

func (app *App) GetLanguages() []Language {
	return app.languages
}

func (app *App) GetLocation() string {
	return app.location
}

func (app *App) ToString() string {

	langs := ""
	for i, lang := range app.GetLanguages() {
		if i > 0 {
			langs += ", "
		}
		langs += lang.Name
	}

	return fmt.Sprintf("%s (%s) in %s", app.GetName(), langs, app.GetLocation())

}

func NewApp(location string) (*App, error) {
	app := App{
		location: location,
	}

	revealed := make([]Language, 0)
	for _, knownLang := range GetSupportedLanguages() {
		if isLanguageInUse(location, knownLang) {
			revealed = append(revealed, knownLang)
		}
	}
	app.languages = revealed
	filePath := filepath.Join(app.location, ManifestName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not read spin.toml at %s", app.location)
	}
	var manifest spinManifest
	err = toml.Unmarshal(data, &manifest)
	if err != nil {
		return nil, fmt.Errorf("spin.toml is invalid! Please fix it first (%s)", filePath)
	}
	app.name = manifest.App.Name
	return &app, nil
}

var Rust = Language{Name: "Rust", hint: "Cargo.toml"}
var GoLang = Language{Name: "Rust", hint: "go.mod"}
var JavaScript = Language{Name: "JavaScript", hint: "package.json"}
var Python = Language{Name: "Python", hint: "requirements.txt"}

func GetSupportedLanguages() []Language {
	return []Language{Rust, GoLang, JavaScript, Python}
}

func isLanguageInUse(appPath string, lang Language) bool {
	found := false
	err := filepath.Walk(appPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && info.Name() == lang.hint {
			found = true
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return false
	}
	return found

}
