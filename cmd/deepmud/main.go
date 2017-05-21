package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/Sirupsen/logrus"
	"github.com/kratenko/deepmud-go/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cmd := &cobra.Command{
		Use:   "deepmud",
		Short: "Going deeper into MUDs",
		Run:   server.RunServer,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			SetupConfig(cmd, false)
		},
	}

	cmd.PersistentFlags().StringP("config", "c", "config", "Config file name, default extension is JSON if not given")
	cmd.PersistentFlags().StringP("addr", "", ":4242", "Address and port to bind")

	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}

func SetupConfig(cmd *cobra.Command, requireFile bool) *viper.Viper {
	cfg := viper.GetViper()
	configName, err := cmd.Flags().GetString("config")
	if err != nil {
		panic(err)
	}
	cfg.SetConfigName(configName)
	cfg.SetConfigType("json")
	// Add working dir
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	cfg.AddConfigPath(dir)
	// Add binary dir
	dir, err = filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	cfg.AddConfigPath(dir)
	cfg.BindPFlags(cmd.Flags())

	err = cfg.ReadInConfig()
	if err != nil && requireFile {
		logrus.Fatal(errors.New(err.Error() + ": Failed to read config file"))
	}
	return cfg
}
