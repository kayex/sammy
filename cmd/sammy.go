package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/kayex/sammy"
	log2 "github.com/kayex/sammy/log"
	"github.com/sqweek/dialog"

	"text/tabwriter"

	"path/filepath"
)

var Debug string

func main() {
	debug := Debug == "true"

	var l *log.Logger
	l = log2.Discard()
	if debug {
		l = log2.StdErr()
	}

	title := "Select sample directory. Samples in this directory and all sub-directories will be affected."
	dir, err := dialog.Directory().Title(title).Browse()
	if err != nil {
		if err == dialog.ErrCancelled {
			os.Exit(0)
		}

		showError(l, err, debug)
	}

	cs, err := sammy.GenerateChangeSet(l, dir, sammy.ExtendMajor, sammy.ExtendMinor, sammy.NormalizeAccidentals)
	if err != nil {
		showError(l, err, debug)
	}

	if len(cs) == 0 {
		dialog.Message("Could not find any samples to rename.").Title("Rename complete").Info()
		os.Exit(0)
	}

	ok := dialog.Message("%d %s will be renamed in %s.\n\nContinue?", len(cs), strSamples(len(cs)), dir).Title("Confirm rename").YesNo()
	if !ok {
		os.Exit(0)
	}

	/*
		err = sammy.Rename(cs)
		if err != nil {
			showError(err)
			os.Exit(1)
		}
	*/

	err = printChangeSet(l, dir, cs)
	if err != nil {
		showError(l, err, debug)
	}

	if debug {
		fmt.Scanf("h")
	}
}

func showError(l *log.Logger, err error, debug bool) {
	dialog.Message("Error: %v", err).Error()

	if debug {
		l.Println(err)
		fmt.Scanf("h")
	} else {
		l.Panicln(err)
	}
}

func printChangeSet(l *log.Logger, dir string, cs map[string]string) error {
	logFile := filepath.Join(dir, fmt.Sprintf("sammy-log-%d.txt", time.Now().Unix()))
	err := writeLog(logFile, dir, cs)
	if err != nil {
		return err
	}

	l.Printf("Wrote log file to %v", logFile)
	dialog.Message("Successfully renamed %d %s.\n\nCheck %s for details.", len(cs), strSamples(len(cs)), logFile).Title("Rename complete").Info()

	return nil
}

func writeLog(logFile, dir string, cs map[string]string) error {
	f, err := os.Create(logFile)
	if err != nil {
		return fmt.Errorf("could not create log file: %v", err)
	}
	defer f.Close()

	w := tabwriter.NewWriter(f, 0, 0, 3, '.', 0)
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

func strSamples(count int) string {
	if count == 1 {
		return "sample"
	}

	return "samples"
}
