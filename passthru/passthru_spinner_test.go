package passthru_test

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2025 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io"
	"os"
	"time"

	"github.com/essentialkaos/ek/v13/fmtutil"
	"github.com/essentialkaos/ek/v13/passthru"
	"github.com/essentialkaos/ek/v13/path"
	"github.com/essentialkaos/ek/v13/req"
	"github.com/essentialkaos/ek/v13/spinner"
	"github.com/essentialkaos/ek/v13/terminal"
)

type DLSpinner struct {
	file       string
	lastUpdate time.Time
	reader     *passthru.Reader
}

func Example() {
	file := "https://mirror.yandex.ru/fedora/linux/releases/39/Server/x86_64/iso/Fedora-Server-netinst-x86_64-39-1.5.iso"
	filename := path.Base(file)

	spinner.Show("Download file {c}%s{!}", filename)

	resp, err := req.Request{URL: file, AutoDiscard: true}.Get()

	if err != nil {
		spinner.Done(false)
		terminal.Error(err)
		return
	}

	fd, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		spinner.Done(false)
		terminal.Error(err)
		return
	}

	defer resp.Body.Close()

	r := passthru.NewReader(resp.Body, resp.ContentLength)
	s := &DLSpinner{file: filename, reader: r}
	s.reader.Update = s.Update

	_, err = io.Copy(fd, r)

	spinner.Update(
		"Download file {c}%s{!} {s}(%s){!}",
		filename, fmtutil.PrettySize(resp.ContentLength),
	)

	spinner.Done(err == nil)

	if err != nil {
		terminal.Error(err)
		return
	}
}

func (s *DLSpinner) Update(_ int) {
	now := time.Now()

	if now.Sub(s.lastUpdate) < time.Second/10 {
		return
	}

	speed, _ := s.reader.Speed()
	spinner.Update(
		"{s}[ %4.1f%% | %6s/s ]{!} Download file {c}%s{!}",
		s.reader.Progress(), fmtutil.PrettySize(speed), s.file,
	)

	s.lastUpdate = now
}
