package issue

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

type config struct {
	ProjectName string `toml:"project_name"`
}

// Init initialises issue tracking in the current dir
func Init() error {
	// TODO: handle re-init

	if err := os.MkdirAll(issuesDir, defaultDirPerms); err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	var cfg config
	_, cfg.ProjectName = filepath.Split(cwd)

	b, err := toml.Marshal(cfg)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(configFile, b, defaultFilePerms)
}

func loadConfig() (*config, error) {
	var cfg config
	b, err := ioutil.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	if err := toml.Unmarshal(b, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
