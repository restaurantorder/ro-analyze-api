package helper

import (
	"log"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"os"
)

type api struct {
	Host		string
	Port		int
}
type soundFiles struct {
	Path		string
	Type		string
}
type environment struct {
	OrderApi		api
	AuthApi			api
	SoundFiles		soundFiles
	NoAuthRoutes	[]string
}
type config struct {
	Dev 			environment
	Prod 			environment
	Local0 			environment
	Local1 			environment
	Local2 			environment
	NoAuthRoutes	[]string
}

func ReadConfig() environment {
	var config config
	absPath, _ := filepath.Abs("config.toml")
	if _, err := toml.DecodeFile(absPath, &config); err != nil {
		log.Fatal(err)
	}
	var configEnv environment

	if os.Getenv("DEV") != "" {
		configEnv = config.Dev
	} else if os.Getenv("PROD") != "" {
		configEnv = config.Prod
	} else if os.Getenv("LOCAL0") != "" {
		configEnv = config.Local0
	} else if os.Getenv("LOCAL1") != "" {
		configEnv = config.Local1
	} else if os.Getenv("LOCAL2") != "" {
		configEnv = config.Local2
	}
	configEnv.NoAuthRoutes = config.NoAuthRoutes
	return configEnv
}