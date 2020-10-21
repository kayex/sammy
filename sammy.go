package sammy

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

		l.Printf("Processing %s\n", path)

		dir := filepath.Dir(path)
		original := filepath.Base(path)
		filename := original
		for _, t := range tfs {
			filename = t(filename)
		}

		if filename != original {
			o := strings.TrimPrefix(filepath.Join(dir, original), dir+string(os.PathSeparator))
			n := strings.TrimPrefix(filepath.Join(dir, filename), dir+string(os.PathSeparator))
			l.Printf("%s -> %s\n", o, n)
			cs[path] = filepath.Join(dir, filename)
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

type RenameError struct {
	NewPath string
	Err     error
}

func (re RenameError) Error() string {
	return re.Err.Error()
}

type RenameErrors map[string]RenameError

func (re RenameErrors) Check() bool {
	return len(re) == 0
}

func (re RenameErrors) Error() string {
	c := len(re)
	return fmt.Sprintf("failed renaming %d %s", c, StrSamples(c))
}

func Rename(l *log.Logger, cs map[string]string) RenameErrors {
	var errs = make(RenameErrors)

	for o, n := range cs {
		err := os.Rename(o, n)
		if err != nil {
			if e, ok := err.(*os.LinkError); ok {
				err = e.Err
			}

			re := RenameError{
				NewPath: n,
				Err:     err,
			}
			errs[o] = re

			l.Printf("Failed renaming %s: %v\n", o, re.Err.Error())
		}
	}

	return errs
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

// StrSamples returns the word "sample" pluralized according to count.
func StrSamples(count int) string {
	if count == 1 {
		return "sample"
	}

	return "samples"
}
