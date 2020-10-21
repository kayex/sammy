package main

import (
	"errors"
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

// Debug indicates whether debug log levels should be used and should be set at compile time.
var Debug string

func main() {
	debug := Debug == "true"
	debug = true

	var l *log.Logger
	l = log2.Discard()
	if debug {
		l = log2.StdErr()
	}

	defer func(l *log.Logger) {
		if err := recover(); err != nil {
			if e, ok := err.(error); ok {
				handleError(l, e, debug)
			} else if s, ok := err.(string); ok {
				e := errors.New(s)
				handleError(l, e, debug)
			}

			panic(err)
		}
	}(l)

	title := "Select sample directory. Samples in this directory and all sub-directories will be affected."
	dir, err := dialog.Directory().Title(title).Browse()
	if err != nil {
		if err == dialog.ErrCancelled {
			os.Exit(0)
		}

		handleError(l, err, debug)
	}

	cs, err := sammy.GenerateChangeSet(l, dir, sammy.ExtendMajor, sammy.ExtendMinor, sammy.NormalizeAccidentals)
	if err != nil {
		handleError(l, err, debug)
	}

	if len(cs) == 0 {
		dialog.Message("Could not find any samples to rename.").Title("Rename complete").Info()
		os.Exit(0)
	}

	ok := dialog.Message("%d %s will be renamed in %s.\n\nContinue?", len(cs), sammy.StrSamples(len(cs)), dir).Title("Confirm rename").YesNo()
	if !ok {
		os.Exit(0)
	}

	l.Printf("Renaming %d samples. This may take a while...", len(cs))
	errs := sammy.Rename(l, cs)
	if !errs.Check() {
		handleWarning(l, errs, debug)
	}

	err = printChangeSet(l, dir, cs, errs)
	if err != nil {
		handleError(l, err, debug)
	}

	fmt.Printf("Successfully renamed %d samples.\n", len(cs)-len(errs))
	if debug {
		fmt.Println("Press any key to continue.")
		fmt.Scanf("h")
	}
}

func handleWarning(l *log.Logger, err error, debug bool) {
	dialog.Message("Error: %v", err).Error()
	l.Printf("Error: %v\n", err)
}

func handleError(l *log.Logger, err error, debug bool) {
	dialog.Message("Error: %v", err).Error()

	if debug {
		l.Printf("Error: %v\n", err)
		l.Println("Press any key to dump error and exit program.")
		fmt.Scanf("h")
		fmt.Println(err)
	} else {
		fmt.Println(err)
	}

	os.Exit(1)
}

func printChangeSet(l *log.Logger, dir string, cs map[string]string, errs sammy.RenameErrors) error {
	logFile := filepath.Join(dir, fmt.Sprintf("sammy-log-%d.txt", time.Now().Unix()))
	err := writeLog(logFile, dir, cs, errs)
	if err != nil {
		return err
	}

	l.Printf("Wrote log file to %v", logFile)
	c := len(cs) - len(errs)
	dialog.Message("Successfully renamed %d %s.\n\nLog file can be found in:\n\n%s", c, sammy.StrSamples(c), logFile).Title("Rename complete").Info()

	return nil
}

func writeLog(logFile, dir string, cs map[string]string, errs sammy.RenameErrors) error {
	f, err := os.Create(logFile)
	if err != nil {
		return fmt.Errorf("could not create log file: %v", err)
	}
	defer f.Close()

	w := tabwriter.NewWriter(f, 0, 0, 3, ' ', 0)
	fmt.Fprintf(w, "Status\tOriginal filename\tNew filename\tError\n")
	fmt.Fprintf(w, "\t\t\t\n")
	for o, n := range cs {
		status := "SUCCESS"
		errMsg := ""
		if e, ok := errs[o]; ok {
			status = "FAILED"
			errMsg = fmt.Sprintf("%v", e.Err)
		}

		ot := strings.TrimPrefix(o, dir+string(os.PathSeparator))
		nt := strings.TrimPrefix(n, dir+string(os.PathSeparator))

		fmt.Fprintf(w, "%s\t%v\t%v\t%v\n", status, ot, nt, errMsg)
	}
	w.Flush()

	return nil
}
