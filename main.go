package main

import (
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type config struct {
	// extension to filter out
	ext string
	// min file size
	size int64
	// list files
	list bool
	// delete files
	del bool
	// log destination writer
	wLog io.Writer
	// archive directory
	archive string
	// modification time
	modSince string
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "%s tool. Developed for educational purposes.\n", os.Args[0])
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	root := flag.String("root", ".", "Root directory to start.")
	ext := flag.String("ext", "", "File extensions to filter out.")
	list := flag.Bool("list", false, "List files only.")
	modSince := flag.String("since", "", "Time of file modification.")
	size := flag.Int64("size", 0, "Minimum file size in bytes.")
	del := flag.Bool("del", false, "Delete files.")
	logFile := flag.String("log", "", "Log deletes to this file.")
	archive := flag.String("archive", "", "Archive directory.")
	flag.Parse()

	var (
		f   = os.Stdout
		err error
	)

	if *logFile != "" {
		f, err = os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer f.Close()
	}

	c := config{
		ext:      *ext,
		size:     *size,
		list:     *list,
		del:      *del,
		wLog:     f,
		archive:  *archive,
		modSince: *modSince,
	}

	if err = run(*root, os.Stdout, c); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	delLogger := log.New(cfg.wLog, "DELETED FILE: ", log.LstdFlags)
	return filepath.Walk(root,
		func(path string, info fs.FileInfo, err error) error {
			if err != nil {
				return err
			}

			extList := strings.Fields(cfg.ext)

			var afterRFC822Z time.Time
			if cfg.modSince != "" {
				afterRFC822Z, err = time.Parse(time.RFC822Z, cfg.modSince)
				if err != nil {
					return err
				}

			}

			if filterOut(path, extList, cfg.size, afterRFC822Z, info) {
				return nil
			}

			// If list was explicitly set, don't do anything else
			if cfg.list {
				return listFile(path, out)
			}

			// Archive files and continue if successful
			if cfg.archive != "" {
				if err = archiveFile(cfg.archive, root, path); err != nil {
					return err
				}
			}

			// Delete file
			if cfg.del {
				return delFile(path, delLogger)
			}

			// List is default function is nothing else was set
			return listFile(path, out)
		})
}
