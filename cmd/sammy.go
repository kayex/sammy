package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/kayex/sammy"
	"github.com/sqweek/dialog"

	"text/tabwriter"

	"path/filepath"
)

func main() {
	title := "Select sample directory. Samples in this directory and all sub-directories will be affected."

	dir, err := dialog.Directory().Title(title).Browse()
	if err != nil {
		showError(err)
		os.Exit(1)
	}

	cs, err := sammy.GenerateChangeSet(dir, sammy.ExtendMajor, sammy.ExtendMinor)
	if err != nil {
		showError(err)
		os.Exit(1)
	}

	if len(cs) == 0 {
		dialog.Message("Could not find any samples to rename.").Title("Rename complete").Info()
		os.Exit(0)
	}

	ok := dialog.Message("%d samples will be renamed in %s.\n\nContinue?", len(cs), dir).Title("Confirm rename").YesNo()
	if !ok {
		os.Exit(1)
	}

	err = sammy.Rename(cs)
	if err != nil {
		showError(err)
		os.Exit(1)
	}

	err = printChangeSet(dir, cs)
	if err != nil {
		showError(err)
		os.Exit(1)
	}
}

func showError(err error) {
	dialog.Message("Error: %v", err).Error()
}

func printChangeSet(dir string, cs map[string]string) error {
	logFile := filepath.Join(dir, "log.txt")
	err := writeLog(logFile, dir, cs)
	if err != nil {
		return err
	}

	dialog.Message("Successfully renamed %d samples.\n\nCheck %s for details.", len(cs), logFile).Title("Rename complete").Info()

	return nil
}

func writeLog(logFile, dir string, cs map[string]string) error {
	f, err := os.Create(logFile)
	if err != nil {
		return fmt.Errorf("could not create log file: %v", err)
	}
	defer f.Close()

	w := tabwriter.NewWriter(f, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "Original filename\tNew filename\t\n")
	fmt.Fprintf(w, "\t\t\n")
	for o, n := range cs {
		o := strings.TrimPrefix(o, dir+string(os.PathSeparator))
		n := strings.TrimPrefix(n, dir+string(os.PathSeparator))

		fmt.Fprintf(w, "%v\t%v\t\n", o, n)
	}
	w.Flush()

	return nil
}
