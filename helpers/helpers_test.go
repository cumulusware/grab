// Copyright (c) 2015-2017 The grab developers. All rights reserved.
// Project site: https://github.com/cumulusware/grab
// Use of this source code is governed by a MIT-style license that
// can be found in the LICENSE.txt file for the project.

package helpers

import (
	"testing"
)

const (
	failCheck = `✗` // UTF-8 u2717
	passCheck = `✓` // UTF-8 u2713
)

func TestGoodImportPaths(t *testing.T) {
	testCases := []struct {
		givenPath          string
		expectedGoSrc      string
		expectedImportPath string
	}{
		{
			"/Users/johndoe/development/go/src/github.com/johndoe/myproject",
			"/Users/johndoe/development/go/src",
			"github.com/johndoe/myproject",
		},
		{
			"/Users/johndoe/development/go/src/github.com/johndoe/myproject/",
			"/Users/johndoe/development/go/src",
			"github.com/johndoe/myproject",
		},
		{
			"foo/blah",
			"/Users/matthew/development/go/src",
			"github.com/cumulusware/grab/helpers/foo/blah",
		},
	}
	for _, testCase := range testCases {
		t.Log("Given the need to determine the Go src and import path")
		t.Logf("\tfor %s", testCase.givenPath)
		goSrc, importPath, err := DetermineImportPath(testCase.givenPath)
		if err != nil {
			t.Fatalf(
				"\t%v Should be able run DetermineImportPath on %s",
				failCheck,
				testCase.givenPath,
			)
		}
		if goSrc != testCase.expectedGoSrc {
			t.Fatalf(
				"\t%v Should get expected goSrc %s instead of %s",
				failCheck,
				testCase.expectedGoSrc,
				goSrc,
			)
		}
		t.Logf(
			"\t%v Should get goSrc %s",
			passCheck,
			goSrc,
		)
		if importPath != testCase.expectedImportPath {
			t.Fatalf(
				"\t%v Should get expected importPath %s instead of %s",
				failCheck,
				testCase.expectedImportPath,
				importPath,
			)
		}
		t.Logf(
			"\t%v Should get importPath %s",
			passCheck,
			importPath,
		)
	}
}
