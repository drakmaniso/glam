package resource

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/cozely/cozely/internal"
)

////////////////////////////////////////////////////////////////////////////////

type source interface {
	exist(n string) bool
	open(n string) (io.ReadCloser, error)
}

var sources []source

////////////////////////////////////////////////////////////////////////////////

type path string

// Path adds p to the stack of sources for resource look-up.
func Path(p string) error {
	if internal.Running {
		return errors.New("unable to add resource path while the framework is running")
	}
	if p[len(p)-1] != '/' {
		return errors.New(`invalid resource path "` + p + `"`)
	}
	_, err := os.Stat(p)
	if err != nil {
		return internal.Wrap(`invalid resource path "`+p+`"`, err)
	}
	sources = append(sources, path(p))
	return nil
}

func (p path) exist(n string) bool {
	_, err := os.Stat(string(p) + n)
	return !os.IsNotExist(err)
}

func (p path) open(n string) (io.ReadCloser, error) {
	f, err := os.Open(string(p) + n)
	if !os.IsNotExist(err) {
		return f, err
	}
	return nil, err
}

////////////////////////////////////////////////////////////////////////////////

type pack struct {
	*zip.Reader
	files map[string]*zip.File
}

// Pack adds a zipped string to the stack of sources for resource look-up.
func Pack(content string) error {
	var err error
	if internal.Running {
		return errors.New("unable to add resource path while the framework is running")
	}
	p := pack{}
	p.Reader, err = zip.NewReader(strings.NewReader(content), int64(len(content)))
	if err != nil {
		return fmt.Errorf("resource.Pack: %s", err)
	}
	p.files = map[string]*zip.File{}
	for _, f := range p.Reader.File {
		p.files[f.Name] = f
	}
	sources = append(sources, p)
	return nil
}

func (p pack) exist(n string) bool {
	_, ok := p.files[n]
	return ok
}

func (p pack) open(n string) (io.ReadCloser, error) {
	zf, ok := p.files[n]
	if !ok {
		return nil, os.ErrNotExist
	}
	f, err := zf.Open()
	if err != nil {
		return nil, fmt.Errorf("open resource pack: %s", err)
	}
	return f, nil
}

//// Copyright (c) 2013-2018 Laurent Moussault. All rights reserved.
//// Licensed under a simplified BSD license (see LICENSE file).
