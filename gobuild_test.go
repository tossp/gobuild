// Copyright 2015 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"testing"

	"github.com/issue9/assert"
)

func TestGetExts(t *testing.T) {
	a := assert.New(t)

	a.Equal(getExts(""), []string{})
	a.Equal(getExts(",, ,"), []string{})
	a.Equal(getExts(",.go, ,.php,"), []string{".go", ".php"})
	a.Equal(getExts(",go,.php,"), []string{".go", ".php"})
	a.Equal(getExts(",go , .php,"), []string{".go", ".php"})
}

func TestRecursivePath(t *testing.T) {
	a := assert.New(t)

	a.Equal(recursivePath(false, []string{"./testdir"}), []string{
		"./testdir",
	})

	a.Equal(recursivePath(true, []string{"./testdir"}), []string{
		"./testdir",
		"testdir/testdir1",
		"testdir/testdir2",
		"testdir/testdir2/testdir3",
	})

	a.Equal(recursivePath(true, []string{"./testdir/testdir1", "./testdir/testdir2"}), []string{
		"./testdir/testdir1",
		"./testdir/testdir2",
		"testdir/testdir2/testdir3",
	})

	a.Equal(recursivePath(true, []string{"./testdir/testdir2"}), []string{
		"./testdir/testdir2",
		"testdir/testdir2/testdir3",
	})
}
