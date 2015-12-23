// Copyright 2015 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/tsuru/tsuru/cmd"
	"gopkg.in/check.v1"
)

type S struct{}

var _ = check.Suite(&S{})
var manager *cmd.Manager

func Test(t *testing.T) { check.TestingT(t) }

func (s *S) SetUpTest(c *check.C) {
	var stdout, stderr bytes.Buffer
	manager = cmd.NewManager("yati", version, "", &stdout, &stderr, os.Stdin, nil)
}
