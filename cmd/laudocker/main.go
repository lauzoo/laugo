package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "laudocker",
	Short: "Laudocker is a tool for docker",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Docker tools for liuliqiang.")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func showOpenSourceProtocol() {
	//https://i.loli.net/2018/12/31/5c29e84b2746c.png
}
