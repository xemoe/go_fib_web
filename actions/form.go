package actions

import (
	"io"
	"os"
	"path/filepath"

	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

func FormIndex(c buffalo.Context) error {
	return c.Render(200, r.HTML("form/index.html"))
}

func FormUpload(c buffalo.Context) error {

	wf := &WithFile{}
	if err := c.Bind(wf); err != nil {
		return errors.WithStack(err)
	}

	//
	// Save upload file
	//
	dir := filepath.Join(".", "uploads")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, wf.MyFile.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, wf.MyFile)

	//
	// JSON response message
	//
	resp := &Message{
		Status: "OK",
		Result: wf.MyFile.Filename,
	}

	return c.Render(200, r.JSON(resp))
}
