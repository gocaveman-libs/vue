package vue

import "time"

import "github.com/gocaveman/webresource" // FIXME: prototype import path for now

func ModulePROTO() webresource.Module {
	fs := webresource.NewFileSet("github.com/gocaveman-libs/vue", requires()...)
	addFiles(fs)
	return fs
}

func requires() []webresource.Module { return nil }

func addFiles(fs *webresource.FileSet) {
}
