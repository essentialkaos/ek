package progress_test

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2023 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"io"
	"os"
	"time"

	"github.com/essentialkaos/ek/v12/progress"
	"github.com/essentialkaos/ek/v12/req"
)

// ////////////////////////////////////////////////////////////////////////////////// //

func SimpleExample() {
	pb := progress.New(1000, "Counting…")

	// You can use default settings as a starting point
	pbs := progress.DefaultSettings
	pbs.RefreshRate = 50 * time.Millisecond
	pbs.IsSize = false

	// You can update all settings except RefreshRate before and after
	// calling Start method. RefreshRate must be set before calling
	// Start method.
	pb.UpdateSettings(pbs)

	pb.Start() // Start async progress handling

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Second / 100)
		pb.Add(1)
	}

	pb.Finish() // Stop async progress handling
}

func DownloadExample() {
	pb := progress.New(0, "file.zip")
	resp, err := req.Request{URL: "https://domain.com/file.zip"}.Get()

	if err != nil {
		panic(err.Error())
	}

	if resp.StatusCode != 200 {
		panic("Looks like something is wrong")
	}

	defer resp.Body.Close()

	fd, err := os.OpenFile("file.zip", os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err.Error())
	}

	defer fd.Close()

	// Set total size to content length
	pb.SetTotal(resp.ContentLength)
	pb.Start()
	io.Copy(fd, pb.Reader(resp.Body))
	pb.Finish()
}
