package config_test

import (
	"os"
	"testing"

	"github.com/chiefagu/config"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func resetFlagsAndViper() {
	viper.Reset()
	os.Args = []string{"cmd"}
	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError)
}

func TestMustLoadConfig(t *testing.T) {
	t.Run("default values apply", func(t *testing.T) {
		resetFlagsAndViper()
		os.Unsetenv("PORT")

		assert.NotPanics(t, func() {
			config.MustLoadConfig()
		})

		assert.Equal(t, 8081, viper.Get("port"))
	})

	t.Run("environment variables apply", func(t *testing.T) {
		resetFlagsAndViper()
		os.Setenv("PORT", "8000")
		defer os.Unsetenv("PORT")

		assert.NotPanics(t, func() {
			config.MustLoadConfig()
		})

		assert.Equal(t, "8000", viper.Get("port"))
	})

	t.Run("flags apply", func(t *testing.T) {
		resetFlagsAndViper()
		os.Args = []string{"cmd", "--port=7070"}
		pflag.CommandLine = pflag.NewFlagSet(os.Args[1], pflag.ExitOnError)

		assert.NotPanics(t, func() {
			config.MustLoadConfig()
		})

		assert.Equal(t, 7070, viper.Get("port"))

	})

	t.Run("using a config file", func(t *testing.T) {
		resetFlagsAndViper()
		os.Args = []string{"cmd", "--config=./invalid/path/config.env"}
		pflag.CommandLine = pflag.NewFlagSet(os.Args[1], pflag.ExitOnError)

		assert.Panics(t, func() {
			config.MustLoadConfig()
		})
	})

}
