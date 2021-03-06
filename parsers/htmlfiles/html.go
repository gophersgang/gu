package htmlfiles

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

// ParseDir returns a new instance of all CSS files located within the provided directory.
func ParseDir(dir string) (map[string]string, error) {
	items := make(map[string]string)

	// Walk directory pulling contents into css items.
	if cerr := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if cerr := walkDir(items, dir, path, info, err); cerr != nil {
			return cerr
		}

		return nil
	}); cerr != nil {
		return nil, cerr
	}

	return items, nil
}

var allowedExtensions = []string{".html", ".xhtml", ".xml", ".ghtml", ".gml", ".tml"}

// validExension returns true/false if the extension provide is a valid acceptable one
// based on the allowedExtensions string slice.
func validExtension(ext string) bool {
	for _, es := range allowedExtensions {
		if es != ext {
			continue
		}

		return true
	}

	return false
}

// walkDir adds the giving path if it matches certain criterias into the items map.
func walkDir(items map[string]string, root string, path string, info os.FileInfo, err error) error {
	if err != nil {
		return err
	}

	if !info.Mode().IsRegular() {
		return nil
	}

	// Is file an exension we allow else skip.
	if !validExtension(filepath.Ext(path)) {
		return nil
	}

	rel, err := filepath.Rel(root, path)
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	items[rel] = string(data)
	return nil
}
