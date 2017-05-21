package server

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func RunServer(cmd *cobra.Command, args []string) {
	ListenAndServe(viper.GetString("addr"))
}
