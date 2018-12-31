package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	. "github.com/lauzoo/laugo/internal/pkgs/log"

	"github.com/spf13/cobra"
)

var (
	templateUrl = "https://raw.githubusercontent.com/lauzoo/laugo/master/statics/ides/gitignores/%s.gitignore"
)

func init() {
	rootCmd.AddCommand(gitIgnoreCmd)
}

var gitIgnoreCmd = &cobra.Command{
	Use:   "gitignore",
	Short: "Get ignore file for local project",
	Long:  `Get ignore file for local project`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			Log.Error("invalid args count, user -help for more info.")
			return
		}

		var url string
		switch strings.ToUpper(args[0])[:2] {
		case "GO":
			url = fmt.Sprintf(templateUrl, "go")
		case "PY":
			url = fmt.Sprintf(templateUrl, "py")
		default:
			Log.Error("unsupport language: " + args[0])
			return
		}

		downloadFileToPwd(url, ".gitignore")
	},
}

func downloadFileToPwd(url string, fileName string) (err error) {
	var pwd string
	if pwd, err = filepath.Abs(filepath.Dir(os.Args[0])); err != nil {
		return err
	}

	var filePath = path.Join(pwd, fileName)
	return downloadFile(url, filePath)
}

func downloadFile(url string, filePath string) (err error) {
	var file *os.File
	if file, err = os.Create(filePath); err != nil {
		return err
	}
	defer file.Close()

	// Get the data
	var resp *http.Response
	if resp, err = http.Get(url); err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
