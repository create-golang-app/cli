package utils

import (
	"github.com/mattn/go-colorable"
)

var (
	Stdout = colorable.NewColorableStdout() // add a colorable std out
	Stderr = colorable.NewColorableStderr() // add a colorable std err
)
