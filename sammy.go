package sammy

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

// SupportedFileExtensions is a list over file extensions that are considered audio samples.
var SupportedFileExtensions []string = []string{".wav", ".wave", ".flac", ".mp3", ".mp4", ".aiff", ".ogg", ".ogv", ".oga", ".ogx", ".ogm", ".spx", ".opus"}

// GenerateChangeSet returns a map of filenames and their transformed counterparts
// after applying the transformers in tfs recursively to all files in dir.
func GenerateChangeSet(l *log.Logger, dir string, tfs ...Transformer) (map[string]string, error) {
	l.Printf("Indexing %s\n", dir)
	cs := make(map[string]string)

	scanned := 0
	registered := 0
	start := time.Now()

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
			l.Printf("%s -> %s\n", path, trans)
			cs[path] = trans
			registered++
		}

		scanned++

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed traversing dir: %v", err)
	}

	duration := time.Now().Sub(start)

	l.Println()
	l.Printf("Scanned %d samples in %s and found %d candidates for renaming.\n", scanned, duration, registered)

	return cs, nil
}

func Rename(cs map[string]string) error {
	return nil
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
