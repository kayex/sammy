package sammy

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateChangeSet(dir string, tfs ...Transformer) (map[string]string, error) {
	cs := make(map[string]string)

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if path == dir {
			return nil
		}

		if !sample(path) {
			return nil
		}

		fmt.Println(path, info.Size())

		trans := path
		for _, t := range tfs {
			trans = t(trans)
		}

		cs[path] = trans

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed traversing dir: %v", err)
	}

	return cs, nil
}

func sample(path string) bool {
	switch filepath.Ext(path) {
	case ".wav", ".wave", ".flac", ".mp3", ".mp4", ".aiff", ".ogg", ".ogv", ".oga", ".ogx", ".ogm", ".spx", ".opus":
		return true
	}

	return false
}
