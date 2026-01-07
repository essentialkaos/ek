package interactive

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"github.com/essentialkaos/ek/v13/events"
	"github.com/essentialkaos/ek/v13/fmtc"
	"github.com/essentialkaos/ek/v13/fmtutil"
	"github.com/essentialkaos/ek/v13/selfupdate"
	"github.com/essentialkaos/ek/v13/spinner"
	"github.com/essentialkaos/ek/v13/terminal"
)

// ////////////////////////////////////////////////////////////////////////////////// //

var dispatcher *events.Dispatcher

// ////////////////////////////////////////////////////////////////////////////////// //

// Dispatcher returns event dispatcher for interactive update
func Dispatcher() *events.Dispatcher {
	if dispatcher != nil {
		return dispatcher
	}

	dispatcher = events.NewDispatcher()

	dispatcher.AddHandler(selfupdate.EV_UPDATE_START, updateStart)
	dispatcher.AddHandler(selfupdate.EV_UPDATE_ERROR, updateError)
	dispatcher.AddHandler(selfupdate.EV_UPDATE_COMPLETE, updateComplete)

	dispatcher.AddHandler(selfupdate.EV_SIGNATURE_DOWNLOAD_START, signatureDownloadStart)
	dispatcher.AddHandler(selfupdate.EV_SIGNATURE_DOWNLOAD_ERROR, signatureDownloadError)
	dispatcher.AddHandler(selfupdate.EV_SIGNATURE_DOWNLOAD_COMPLETE, signatureDownloadComplete)

	dispatcher.AddHandler(selfupdate.EV_SIGNATURE_PARSE_START, signatureParseStart)
	dispatcher.AddHandler(selfupdate.EV_SIGNATURE_PARSE_ERROR, signatureParseError)
	dispatcher.AddHandler(selfupdate.EV_SIGNATURE_PARSE_COMPLETE, signatureParseComplete)

	dispatcher.AddHandler(selfupdate.EV_BINARY_DOWNLOAD_START, binaryDownloadStart)
	dispatcher.AddHandler(selfupdate.EV_BINARY_DOWNLOAD_SIZE, binaryDownloadSize)
	dispatcher.AddHandler(selfupdate.EV_BINARY_DOWNLOAD_ERROR, binaryDownloadError)
	dispatcher.AddHandler(selfupdate.EV_BINARY_DOWNLOAD_COMPLETE, binaryDownloadComplete)

	dispatcher.AddHandler(selfupdate.EV_BINARY_VERIFY_START, binaryVerifyStart)
	dispatcher.AddHandler(selfupdate.EV_BINARY_VERIFY_ERROR, binaryVerifyError)
	dispatcher.AddHandler(selfupdate.EV_BINARY_VERIFY_OK, binaryVerifyComplete)

	dispatcher.AddHandler(selfupdate.EV_BINARY_REPLACE_START, binaryReplaceStart)
	dispatcher.AddHandler(selfupdate.EV_BINARY_REPLACE_ERROR, binaryReplaceError)
	dispatcher.AddHandler(selfupdate.EV_BINARY_REPLACE_COMPLETE, binaryReplaceComplete)

	return dispatcher
}

// ////////////////////////////////////////////////////////////////////////////////// //

// updateStart prints message about start update process
func updateStart(payload any) {
	info, ok := payload.(selfupdate.Update)

	if !ok {
		fmtc.Println("\n{s}Updating binary…{!}\n")
	} else {
		fmtc.Printfn("\n{s}Updating binary to %s…{!}\n", info.Version)
	}
}

// updateError prints message about update error
func updateError(payload any) {
	message, ok := payload.(string)

	if !ok {
		terminal.Error("\nCan't update binary\n")
	} else {
		terminal.Error("\nCan't update binary: %s\n", message)
	}
}

// updateComplete prints message about successful update
func updateComplete(payload any) {
	info, ok := payload.(selfupdate.Update)

	fmtc.NewLine()

	if !ok {
		fmtc.Println("{g}Binary successfully updated!{!}\n")
	} else {
		fmtc.Printfn("{g}Binary successfully updated to {*}%s{!}\n", info.Version)
	}
}

// signatureDownloadStart shows spinner with message about start ECDSA signature downloading
func signatureDownloadStart(_ any) {
	spinner.Show("Download ECDSA signature")
}

// signatureDownloadError stops spinner with message about ECDSA signature downloading error
func signatureDownloadError(_ any) {
	spinner.Done(false)
}

// signatureDownloadComplete stops spinner with message about successfully downloaded 
// ECDSA signature
func signatureDownloadComplete(_ any) {
	spinner.Done(true)
}

// signatureParseStart shows spinner with message about start ECDSA signature parsing
func signatureParseStart(_ any) {
	spinner.Show("Parse ECDSA signature data")
}

// signatureParseError stops spinner with message about ECDSA signature parsing error
func signatureParseError(_ any) {
	spinner.Done(false)
}

// signatureParseComplete stops spinner with message about successfully parsed
// ECDSA signature
func signatureParseComplete(_ any) {
	spinner.Done(true)
}

// binaryDownloadStart shows spinner with message about start binary downloading
func binaryDownloadStart(_ any) {
	spinner.Show("Download binary")
}

// binaryDownloadSize updates spinner with message about binary downloading size
func binaryDownloadSize(payload any) {
	binarySize, ok := payload.(int64)

	if ok {
		spinner.Update("Download binary {s}(%s){!}", fmtutil.PrettySize(binarySize))
	}
}

// binaryDownloadError stops spinner with message about binary downloading error
func binaryDownloadError(_ any) {
	spinner.Done(false)
}

// binaryDownloadComplete stops spinner with message about successfully downloaded
// binary
func binaryDownloadComplete(_ any) {
	spinner.Done(true)
}

// binaryVerifyStart shows spinner with message about start binary ECDSA signature
// verification
func binaryVerifyStart(_ any) {
	spinner.Show("Verify binary ECDSA signature")
}

// binaryVerifyError stops spinner with message about binary ECDSA signature verification error
func binaryVerifyError(_ any) {
	spinner.Done(false)
}

// binaryVerifyComplete stops spinner with message about successfully verified
// binary ECDSA signature
func binaryVerifyComplete(_ any) {
	spinner.Done(true)
}

// binaryReplaceStart shows spinner with message about start binary replacement
func binaryReplaceStart(_ any) {
	spinner.Show("Replace binary")
}

// binaryReplaceError stops spinner with message about binary replacement error
func binaryReplaceError(_ any) {
	spinner.Done(false)
}

// binaryReplaceComplete stops spinner with message about successfully replaced
// binary
func binaryReplaceComplete(_ any) {
	spinner.Done(true)
}
