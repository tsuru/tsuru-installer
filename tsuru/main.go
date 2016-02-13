// Copyright 2015 yati authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"os"

	"github.com/tsuru/tsuru/cmd"
)

const version = "0.1.0"

func buildManager(name string) *cmd.Manager {
	m := cmd.BuildBaseManager(name, version, "", nil)
	return m
}

func main() {
	name := cmd.ExtractProgramName(os.Args[0])
	manager := buildManager(name)
	manager.Register(&install{})
	manager.Register(&uninstall{})
	manager.Run(os.Args[1:])
}
