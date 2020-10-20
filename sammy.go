package sammy

import (
	"fmt"
	"os"
	"path/filepath"
)

// SupportedFileExtensions is a list over file extensions that are considered audio samples.
var SupportedFileExtensions []string = []string{".wav", ".wave", ".flac", ".mp3", ".mp4", ".aiff", ".ogg", ".ogv", ".oga", ".ogx", ".ogm", ".spx", ".opus"}

// GenerateChangeSet returns a map of filenames and their transformed counterparts
// after applying the transformers in tfs recursively to all files in dir.
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

		trans := path
		for _, t := range tfs {
			trans = t(trans)
		}

		if trans != path {
			cs[path] = trans
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed traversing dir: %v", err)
	}

	return cs, nil
}

func Rename(cs map[string]string) error {
	for o, n := range cs {
		err := os.Rename(o, n)
		if err != nil {
			return fmt.Errorf("failed renaming %s: %v", o, err)
		}
	}

	return nil
}

// sample returns a bool indicating if path is an audio sample file.
func sample(path string) bool {
	e := filepath.Ext(path)
	for _, ext := range SupportedFileExtensions {
		if e == ext {
			return true
		}
	}

	return false
}
