package core

import (
	"net/http"
	"net/url"
	"os"

	"git.atonline.com/tristantech/gophp/core/stream"
)

type Process struct {
	fHandler map[string]stream.Handler
}

func NewProcess() *Process {
	res := &Process{
		fHandler: make(map[string]stream.Handler),
	}
	res.fHandler["file"], _ = stream.NewFileHandler("/")
	res.fHandler["php"] = stream.PhpHandler()
	return res
}

func (p *Process) Open(u *url.URL) (*stream.Stream, error) {
	s := u.Scheme
	if s == "" {
		s = "file"
	}

	h, ok := p.fHandler[s]
	if !ok {
		return nil, os.ErrInvalid
	}

	return h.Open(u)
}

func (p *Process) Handler(docroot string) http.Handler {
	return nil // TODO
}
