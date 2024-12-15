package git

import (
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"github.com/create-golang-app/cli/pkg/filesystem"
	"github.com/create-golang-app/cli/pkg/utils/message"
	"github.com/go-git/go-git/v5"
)

// GitClone function for `git clone` defined project template.
func GitClone(templateType, templateURL string) error {
	if templateType == "" || templateURL == "" {
		return message.Error("template type or template url is empty")
	}

	// Get current directory.
	currentDir, _ := os.Getwd()

	// Set project folder.
	folder := filepath.Join(currentDir, templateType)

	// Clone project template.
	_, errPlainClone := git.PlainClone(
		folder,
		false,
		&git.CloneOptions{
			URL: getAbsoluteURL(templateURL),
		},
	)

	if errPlainClone != nil {
		return message.Error(errPlainClone.Error())
	}

	// After clone project template, remove .git folder.
	filesystem.RemoveFolders(folder, []string{".git", ".github"})

	return nil
}

func getAbsoluteURL(templateURL string) string {
	templateURL = strings.TrimSpace(templateURL)
	u, _ := url.Parse(templateURL)

	if u.Scheme == "" {
		u.Scheme = "https"
	}

	return u.String()
}
