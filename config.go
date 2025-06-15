package config

import (
	"log/slog"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func MustLoadConfig() {
	viper.AutomaticEnv()

	pflag.String("config", "", "Path to the config file")
	pflag.Int("port", 8081, "api port")

	pflag.Parse()

	// Bind all flags to viper
	pflag.VisitAll(func(flag *pflag.Flag) {
		viper.BindPFlag(flag.Name, flag)
	})

	viper.BindEnv("port", "PORT")

	if configPath := viper.GetString("config"); configPath != "" {
		viper.SetConfigFile(configPath)

		if err := viper.ReadInConfig(); err != nil {
			panic("error reading config file")
		}
	}

	//Watch for changes in the configuration
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		slog.Info("config file changed:", slog.String("file:", e.Name))
	})
}
