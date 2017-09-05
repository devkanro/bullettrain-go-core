package main

import (
	"github.com/bullettrain-sh/bullettrain-go-core/car_context"
	"github.com/bullettrain-sh/bullettrain-go-core/car_date"
	"github.com/bullettrain-sh/bullettrain-go-core/car_directory"
	"github.com/bullettrain-sh/bullettrain-go-core/car_os"
	"github.com/bullettrain-sh/bullettrain-go-core/car_status"
	"github.com/bullettrain-sh/bullettrain-go-core/car_time"
	"github.com/bullettrain-sh/bullettrain-go-git"
	"github.com/bullettrain-sh/bullettrain-go-golang"
	"github.com/bullettrain-sh/bullettrain-go-nodejs"
	"github.com/bullettrain-sh/bullettrain-go-php"
	"github.com/bullettrain-sh/bullettrain-go-python"
	"github.com/bullettrain-sh/bullettrain-go-ruby"
)

const defaultCarOrder = "os time date context dir python go ruby nodejs php git status"

// trailers results in the list of cars to be available for use.
func trailers(currentWorkingDir string) map[string]carRenderer {
	return map[string]carRenderer{
		"context": &carContext.Car{},
		"date":    &carDate.Car{},
		"dir":     &carDirectory.Car{Pwd: currentWorkingDir},
		"git":     &carGit.Car{Pwd: currentWorkingDir},
		"go":      &carGo.Car{Pwd: currentWorkingDir},
		"nodejs":  &carNodejs.Car{Pwd: currentWorkingDir},
		"os":      &carOs.Car{},
		"php":     &carPhp.Car{Pwd: currentWorkingDir},
		"python":  &carPython.Car{Pwd: currentWorkingDir},
		"ruby":    &carRuby.Car{Pwd: currentWorkingDir},
		"status":  &carStatus.Car{},
		"time":    &carTime.Car{},
	}
}
