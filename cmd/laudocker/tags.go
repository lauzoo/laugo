package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"gopkg.in/resty.v1"
)

var (
	imageName string
)

func init() {
	tagCmd.PersistentFlags().StringVarP(&imageName, "image.name", "n", "laudocker", "image for search")
	rootCmd.AddCommand(tagCmd)
}

var tagCmd = &cobra.Command{
	Use:   "tags",
	Short: "show docker images tags",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			logrus.Error("invalid args count, user -help for more info.")
			return
		}

		logrus.Debug("List tags for image: ", args[0])
		var dockerUrl = fmt.Sprintf("https://registry.hub.docker.com/v1/repositories/%s/tags", args[0])
		logrus.Debugf("request url: %s", dockerUrl)
		var respStruct []dockerTag
		resp, err := resty.R().SetResult(&respStruct).Get(dockerUrl)
		if err != nil {
			logrus.Errorf("Failed to fetch tags: %v", err)
			return
		}
		if resp.StatusCode() != http.StatusOK {
			logrus.Errorf("Failed to fetch tags with http status: %s", resp.Status())
			return
		}

		printTags(respStruct)
	},
}

type dockerTag struct {
	Layer string `json:"layer,omitempty"`
	Name  string `json:"name,omitempty"`
}

func printTags(tags []dockerTag) {
	for _, tag := range tags {
		fmt.Printf("%s\n", tag.Name)
	}
}
