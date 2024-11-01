package spinapp

import (
	"fmt"
	"os"
	"path/filepath"

	toml "github.com/pelletier/go-toml/v2"
)

const ManifestName = "spin.toml"

type spinManifest struct {
	App spinManifestApp `toml:"application"`
}
type spinManifestApp struct {
	Name string `toml:"name"`
}

func getAppNameFromManifest(appLocation string) (string, error) {
	filePath := filepath.Join(appLocation, ManifestName)
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("could not read spin.toml at %s", appLocation)
	}

	var manifest spinManifest
	err = toml.Unmarshal(data, &manifest)
	if err != nil {
		return "", fmt.Errorf("spin.toml is invalid! Please fix it first (%s)", filePath)
	}
	return manifest.App.Name, nil
}
