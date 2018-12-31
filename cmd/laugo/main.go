package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "laugo",
	Short: "Laugo is a tool build by liqiang lau",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello world.")
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
