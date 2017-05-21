package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Printf("Hello Mud!\n")
}

func main() {
	cmd := &cobra.Command{
		Use:   "deepmud",
		Short: "Going deeper into MUDs",
		Run:   Run,
	}
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
