package main

import (
	"strings"

	. "github.com/lauzoo/laugo/internal/pkgs/log"
	"github.com/lauzoo/laugo/internal/pkgs/ssh"

	"github.com/go-yaml/yaml"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
)

var (
	host     string
	port     int
	username string
	password string
	cmdLen   = 6
	sshTool  = ssh.Tools{}
)

func init() {
	sshCmd.PersistentFlags().StringVarP(&host, "host", "H", "127.0.0.1", "Host to login")
	sshCmd.PersistentFlags().IntVarP(&port, "port", "P", 22, "Port for host")
	sshCmd.PersistentFlags().StringVarP(&username, "username", "u", "root", "User for host")
	sshCmd.PersistentFlags().StringVarP(&password, "password", "p", "password", "Password for username")
	rootCmd.AddCommand(sshCmd)
}

var sshCmd = &cobra.Command{
	Use:   "ssh",
	Short: "ssh utils",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			Log.Error("invalid args count, user -help for more info.")
			return
		}

		switch strings.ToUpper(args[0])[:cmdLen] {
		case "COPY-ID"[:cmdLen]:
			sshTool.CopyKey(host, port, username, password)
		case "TUNNEL"[:cmdLen]:
			var config *ssh.YamlConfig
			config = readConfig()
			sshTool.CreateTunnel(config.Tunnels[0])
		default:
			Log.Error("unsupport operation: " + args[0])
			return
		}
	},
}

func readConfig() *ssh.YamlConfig {
	fs := afero.NewOsFs()
	file, err := afero.ReadFile(fs, "/root/.tunnel.yml")
	if err != nil {
		panic(err)
	}

	var config ssh.YamlConfig
	yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	return &config
}
