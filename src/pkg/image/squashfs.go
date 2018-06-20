// Copyright (c) 2018, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package image

import (
	"os"
)

// SQUASHFS defines constants for squashfs format
const SQUASHFS = 1

type squashfsFormat struct {
	file *os.File
}

func (f *squashfsFormat) Validate(file *os.File) bool {
	f.file = file
	return false
}

func (f *squashfsFormat) Init(img *Image) error {
	return nil
}

func init() {
	registerFormat(&squashfsFormat{})
}
