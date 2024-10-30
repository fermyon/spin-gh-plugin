package detective

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/thorstenhans/spin-gh-plugin/internal/spinapp"
)

func FindAllSpinApps() []*spinapp.App {
	apps := make([]*spinapp.App, 0)
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.EqualFold(info.Name(), spinapp.ManifestName) {
			relativePath, err := filepath.Rel(".", filepath.Dir(path))
			if err != nil {
				return err
			}

			if relativePath != "." {
				relativePath = fmt.Sprintf("./%s", relativePath)
			}
			app, err := spinapp.NewApp(relativePath)
			if err != nil {
				fmt.Printf("%v", err)
				return err
			}
			fmt.Printf("Discovered Spin App %s\n", app.ToString())
			apps = append(apps, app)
		}
		return nil
	})
	if err != nil {
		return make([]*spinapp.App, 0)
	}
	return apps
}
