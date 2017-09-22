// Copyright (c) 2015-2017 The grab developers. All rights reserved.
// Project site: https://github.com/cumulusware/grab
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package helpers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func Mkdir(perm os.FileMode, x ...string) error {
	path := path.Join(x...)

	err := os.Mkdir(path, perm)
	return err
}

func MkdirAll(perm os.FileMode, x ...string) error {
	path := path.Join(x...)

	err := os.MkdirAll(path, perm)
	return err
}

func DetermineImportPath(givenPath string) (goSrc, importPath string, err error) {
	var basePath string

	parentPath, err := filepath.Abs(filepath.Clean(givenPath))
	if err != nil {
		return "", "", err
	}

	if parentPath == "." {
		err := fmt.Errorf("Given path %s appears to be empty string", givenPath)
		return "", "", err
	}

	if parentPath == "/" {
		err := fmt.Errorf("Given path %s appears to be root directory", givenPath)
		return "", "", err
	}

	for parentPath != "." {
		basePath = filepath.Base(parentPath)
		parentPath = filepath.Dir(parentPath)

		if basePath == "src" {
			goSrc = filepath.Join(parentPath, basePath)
			return goSrc, importPath, nil
		} else {
			importPath = filepath.Join(basePath, importPath)
		}
	}

	err = fmt.Errorf("Unable to find Go src path for %s", givenPath)
	return "", "", err
}
