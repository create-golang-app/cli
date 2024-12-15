package filesystem

import (
	"os"
	"path/filepath"
)

// RemoveFolders function for massively remove folders.
func RemoveFolders(rootFolder string, foldersToRemove []string) {
	for _, folder := range foldersToRemove {
		_ = os.RemoveAll(filepath.Join(rootFolder, folder))
	}
}
