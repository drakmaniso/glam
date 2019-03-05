package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

////////////////////////////////////////////////////////////////////////////////

func main() {
	err := run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() (err error) {
	if len(os.Args) != 4 {
		return errors.New("Usage: cozypack <directory name> <package name.variable name> <output.go>")
	}

	names := strings.Split(os.Args[2], ".")
	if len(names) != 2 {
		return errors.New("Usage: cozypack <directory name> <package name.variable name> <output.go>")
	}

	//TODO: check if directory

	// Zip input
	buf := strings.Builder{}
	zbuf := zip.NewWriter(&buf)
	root := os.Args[1]
	err = filepath.Walk(root, func(p string, fi os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if fi.IsDir() || strings.HasPrefix(fi.Name(), ".") {
			return nil
		}
		h, err := zip.FileInfoHeader(fi)
		if err != nil {
			return err
		}
		h.Name, err = filepath.Rel(root, p)
		if err != nil {
			return err
		}
		h.Name = filepath.ToSlash(h.Name)
		h.Method = zip.Deflate
		f, err := zbuf.CreateHeader(h)
		if err != nil {
			return err
		}
		b, err := ioutil.ReadFile(p)
		if err != nil {
			return err
		}
		_, err = f.Write(b)
		return err
	})
	if err != nil {
		return err
	}
	err = zbuf.Close()
	if err != nil {
		return err
	}

	// Write output
	//TODO: check if overwrite?
	dest, err := os.Create(os.Args[3])
	if err != nil {
		return err
	}

	fmt.Fprintf(
		dest, format,
		os.Args[1], names[1], //os.Args[2], os.Args[3],
		names[0],
		names[1], os.Args[1],
		names[1], buf.String(),
	)

	err = dest.Close()
	if err != nil {
		return err
	}

	return nil
}

const format = `// Code generated by cozypack. DO NOT EDIT.

//go:generate cozypack "%s" $GOPACKAGE.%s $GOFILE

package %s

// %s contains zipped resources from directory "%s".
const %s = %#v

`

//// Copyright (c) 2018-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).