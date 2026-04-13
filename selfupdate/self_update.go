// Package selfupdate provides methods and structs for application self-update
package selfupdate

// ////////////////////////////////////////////////////////////////////////////////// //
//                                                                                    //
//                         Copyright (c) 2026 ESSENTIAL KAOS                          //
//      Apache License, Version 2.0 <https://www.apache.org/licenses/LICENSE-2.0>     //
//                                                                                    //
// ////////////////////////////////////////////////////////////////////////////////// //

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/asn1"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/essentialkaos/ek/v14/events"
	"github.com/essentialkaos/ek/v14/fsutil"
	"github.com/essentialkaos/ek/v14/hashutil"
	"github.com/essentialkaos/ek/v14/req"
	"github.com/essentialkaos/ek/v14/strutil"
)

// ////////////////////////////////////////////////////////////////////////////////// //

const (
	// EV_UPDATE_START is the event name for the start of the update process
	EV_UPDATE_START = "update.start"

	// EV_UPDATE_COMPLETE is the event name for successful completion of the update
	EV_UPDATE_COMPLETE = "update.complete"

	// EV_UPDATE_ERROR is the event name for any error during the update
	EV_UPDATE_ERROR = "update.error"

	// EV_SIGNATURE_DOWNLOAD_START is the event name for starting signature download
	EV_SIGNATURE_DOWNLOAD_START = "signature.download.start"

	// EV_SIGNATURE_DOWNLOAD_ERROR is the event name for signature download errors
	EV_SIGNATURE_DOWNLOAD_ERROR = "signature.download.error"

	// EV_SIGNATURE_DOWNLOAD_COMPLETE is the event name for completed signature download
	EV_SIGNATURE_DOWNLOAD_COMPLETE = "signature.download.complete"

	// EV_SIGNATURE_PARSE_START is the event name for starting signature parsing
	EV_SIGNATURE_PARSE_START = "signature.parse.start"

	// EV_SIGNATURE_PARSE_ERROR is the event name for signature parsing errors
	EV_SIGNATURE_PARSE_ERROR = "signature.parse.error"

	// EV_SIGNATURE_PARSE_COMPLETE is the event name for completed signature parsing
	EV_SIGNATURE_PARSE_COMPLETE = "signature.parse.complete"

	// EV_BINARY_DOWNLOAD_START is the event name for starting binary download
	EV_BINARY_DOWNLOAD_START = "binary.download.start"

	// EV_BINARY_DOWNLOAD_SIZE is the event name for reporting binary download size
	EV_BINARY_DOWNLOAD_SIZE = "binary.download.size"

	// EV_BINARY_DOWNLOAD_ERROR is the event name for binary download errors
	EV_BINARY_DOWNLOAD_ERROR = "binary.download.error"

	// EV_BINARY_DOWNLOAD_COMPLETE is the event name for completed binary download
	EV_BINARY_DOWNLOAD_COMPLETE = "binary.download.complete"

	// EV_BINARY_VERIFY_START is the event name for starting binary signature verification
	EV_BINARY_VERIFY_START = "binary.verify.start"

	// EV_BINARY_VERIFY_ERROR is the event name for binary signature verification errors
	EV_BINARY_VERIFY_ERROR = "binary.verify.error"

	// EV_BINARY_VERIFY_OK is the event name for successful binary signature verification
	EV_BINARY_VERIFY_OK = "binary.verify.ok"

	// EV_BINARY_REPLACE_START is the event name for starting binary replacement
	EV_BINARY_REPLACE_START = "binary.replace.start"

	// EV_BINARY_REPLACE_ERROR is the event name for binary replacement errors
	EV_BINARY_REPLACE_ERROR = "binary.replace.error"

	// EV_BINARY_REPLACE_COMPLETE is the event name for completed binary replacement
	EV_BINARY_REPLACE_COMPLETE = "binary.replace.complete"
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Update holds information about a pending update including URLs and signature
type Update struct {
	BinaryURL    string // BinaryURL is the URL of the new binary
	SignatureURL string // SignatureURL is the URL of the signature file
	Version      string // Version is the version of the new binary
	Signature    []byte // Signature contains the raw signature data
}

// ////////////////////////////////////////////////////////////////////////////////// //

// ecdsaSignature contain ECDSA signature data
type ecdsaSignature struct {
	R *big.Int
	S *big.Int
}

// ////////////////////////////////////////////////////////////////////////////////// //

var (
	// ErrEmptyPubKey indicates that the provided public key data is empty
	ErrEmptyPubKey = fmt.Errorf("public key is empty")

	// ErrNoSignature indicates that no signature data or URL was provided in update info
	ErrNoSignature = fmt.Errorf("update info has no signature")

	// ErrNoBinaryURL indicates that the update info does not contain a binary URL
	ErrNoBinaryURL = fmt.Errorf("update info has no binary URL")

	// ErrNoVersion indicates that the update info does not contain a version
	ErrNoVersion = fmt.Errorf("update info has no version")
)

// ////////////////////////////////////////////////////////////////////////////////// //

// Run executes the self-update process using the given update info and public key
func Run(info Update, pubKeyData string, dispatcher *events.Dispatcher) error {
	if pubKeyData == "" {
		return dispatchError(dispatcher, ErrEmptyPubKey)
	}

	err := info.Validate()

	if err != nil {
		return dispatchError(dispatcher, err)
	}

	curBinary, updBinary, err := getBinaryPaths()

	if err != nil {
		return dispatchError(dispatcher, err)
	}

	pubKey, err := parsePublicKey(formatPubKey(pubKeyData))

	if err != nil {
		return dispatchError(dispatcher, fmt.Errorf("can't parse public key: %w", err))
	}

	dispatcher.DispatchAndWait(EV_UPDATE_START, info)

	if info.SignatureURL != "" {
		signatureData, err := downloadSignature(info.SignatureURL, dispatcher)

		if err != nil {
			return dispatchError(dispatcher, fmt.Errorf("can't download ECDSA signature: %w", err))
		}

		info.Signature = signatureData
	}

	signature, err := parseSignature(info.Signature, dispatcher)

	if err != nil {
		return dispatchError(dispatcher, fmt.Errorf("can't parse ECDSA signature: %w", err))
	}

	hash, err := downloadBinary(info.BinaryURL, updBinary, dispatcher)

	if err != nil {
		return dispatchError(dispatcher, fmt.Errorf("can't download new binary: %w", err))
	}

	isSignatureValid := validateSignature(pubKey, signature, hash, dispatcher)

	if !isSignatureValid {
		return dispatchError(dispatcher, fmt.Errorf("binary signature is invalid"))
	}

	err = replaceBinary(curBinary, updBinary, dispatcher)

	if err != nil {
		return dispatchError(dispatcher, err)
	}

	dispatcher.DispatchAndWait(EV_UPDATE_COMPLETE, info)

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// Validate checks that the update struct contains all required fields
func (i Update) Validate() error {
	switch {
	case len(i.Signature) == 0 && i.SignatureURL == "":
		return ErrNoSignature
	case i.BinaryURL == "":
		return ErrNoBinaryURL
	case i.Version == "":
		return ErrNoVersion
	}

	return nil
}

// ////////////////////////////////////////////////////////////////////////////////// //

// getBinaryPaths retrieves paths of current binary and update binary
func getBinaryPaths() (string, string, error) {
	curBinary, err := os.Executable()

	if err != nil {
		return "", "", fmt.Errorf("can't get current binary name: %w", err)
	}

	curBinary, err = filepath.EvalSymlinks(curBinary)

	if err != nil {
		return "", "", fmt.Errorf("can't get current binary path: %w", err)
	}

	curBinaryDir := filepath.Dir(curBinary)
	curBinaryFile := filepath.Base(curBinary)
	updBinary := filepath.Join(curBinaryDir, "_"+curBinaryFile)

	if !fsutil.IsWritable(curBinaryDir) {
		return "", "", fmt.Errorf("binary directory %q is not writable", curBinaryDir)
	}

	if !fsutil.IsWritable(curBinary) {
		return "", "", fmt.Errorf("binary %q is not writable", curBinary)
	}

	return curBinary, updBinary, nil
}

// parsePublicKey parses ECDSA public key from PEM-encoded data
func parsePublicKey(data string) (*ecdsa.PublicKey, error) {
	pubKeyBlock, _ := pem.Decode([]byte(data))

	if pubKeyBlock == nil {
		return nil, fmt.Errorf("public key is invalid")
	}

	key, err := x509.ParsePKIXPublicKey(pubKeyBlock.Bytes)

	if err != nil {
		return nil, err
	}

	pubKey, ok := key.(*ecdsa.PublicKey)

	if !ok {
		return nil, fmt.Errorf("key is not ECDSA public key (%T)", key)
	}

	return pubKey, nil
}

// downloadSignature downloads ECDSA signature from the specified URL
func downloadSignature(signatureURL string, dispatcher *events.Dispatcher) ([]byte, error) {
	dispatcher.DispatchAndWait(EV_SIGNATURE_DOWNLOAD_START, nil)

	resp, err := req.Request{
		URL:         signatureURL,
		Timeout:     10 * time.Second,
		AutoDiscard: true,
	}.Get()

	if err != nil {
		dispatcher.DispatchAndWait(EV_SIGNATURE_DOWNLOAD_ERROR, nil)
		return nil, fmt.Errorf("can't download signature: %w", err)
	}

	if resp.StatusCode != req.STATUS_OK {
		dispatcher.DispatchAndWait(EV_SIGNATURE_DOWNLOAD_ERROR, nil)
		return nil, fmt.Errorf(
			"server returned non-ok status code (%d)",
			resp.StatusCode,
		)
	}

	data, err := resp.Bytes()

	if err != nil {
		dispatcher.DispatchAndWait(EV_SIGNATURE_DOWNLOAD_ERROR, nil)
		return nil, fmt.Errorf("can't read signature data: %w", err)
	}

	dispatcher.DispatchAndWait(EV_SIGNATURE_DOWNLOAD_COMPLETE, nil)
	return data, nil
}

// parseSignature parses ECDSA signature from ASN.1-encoded data
func parseSignature(data []byte, dispatcher *events.Dispatcher) (*ecdsaSignature, error) {
	dispatcher.DispatchAndWait(EV_SIGNATURE_PARSE_START, nil)

	signature := &ecdsaSignature{}
	_, err := asn1.Unmarshal(data, signature)

	if err != nil {
		dispatcher.DispatchAndWait(EV_SIGNATURE_PARSE_ERROR, nil)
		return nil, err
	}

	dispatcher.DispatchAndWait(EV_SIGNATURE_PARSE_COMPLETE, nil)
	return signature, nil
}

// downloadBinary downloads binary from the specified URL and saves it to the output file
func downloadBinary(binaryURL, outputFile string, dispatcher *events.Dispatcher) (hashutil.Hash, error) {
	dispatcher.DispatchAndWait(EV_BINARY_DOWNLOAD_START, nil)

	resp, err := req.Request{
		URL:         binaryURL,
		Timeout:     10 * time.Second,
		AutoDiscard: true,
	}.Get()

	if err != nil {
		dispatcher.DispatchAndWait(EV_BINARY_DOWNLOAD_ERROR, nil)
		return nil, err
	}

	if resp.StatusCode != req.STATUS_OK {
		dispatcher.DispatchAndWait(EV_BINARY_DOWNLOAD_ERROR, nil)
		return nil, fmt.Errorf(
			"server returned non-ok status code (%d)",
			resp.StatusCode,
		)
	}

	dispatcher.DispatchAndWait(EV_BINARY_DOWNLOAD_SIZE, resp.ContentLength)

	hash, err := resp.SaveWithHash(outputFile, 0640, sha256.New())

	if err != nil {
		dispatcher.DispatchAndWait(EV_BINARY_DOWNLOAD_ERROR, nil)
		return nil, err
	}

	dispatcher.DispatchAndWait(EV_BINARY_DOWNLOAD_COMPLETE, nil)

	return hash, nil
}

// validateSignature validates ECDSA signature for the given hash using the provided public key
func validateSignature(pubKey *ecdsa.PublicKey, signature *ecdsaSignature, hash hashutil.Hash, dispatcher *events.Dispatcher) bool {
	dispatcher.DispatchAndWait(EV_BINARY_VERIFY_START, nil)

	if ecdsa.Verify(pubKey, hash.Bytes(), signature.R, signature.S) {
		dispatcher.DispatchAndWait(EV_BINARY_VERIFY_OK, nil)
		return true
	}

	dispatcher.DispatchAndWait(EV_BINARY_VERIFY_ERROR, nil)

	return false
}

// replaceBinary replaces the current binary with the new binary
func replaceBinary(curBinary, newBinary string, dispatcher *events.Dispatcher) error {
	dispatcher.DispatchAndWait(EV_BINARY_REPLACE_START, nil)

	err := fsutil.CopyAttr(curBinary, newBinary)

	if err != nil {
		dispatcher.DispatchAndWait(EV_BINARY_REPLACE_ERROR, nil)
		return fmt.Errorf("can't copy attributes to new binary: %w", err)
	}

	tmpBinary := curBinary + "_old"
	err = os.Rename(curBinary, tmpBinary)

	if err != nil {
		dispatcher.DispatchAndWait(EV_BINARY_REPLACE_ERROR, nil)
		return fmt.Errorf("can't rename current binary: %w", err)
	}

	err = os.Rename(newBinary, curBinary)

	if err != nil {
		dispatcher.DispatchAndWait(EV_BINARY_REPLACE_ERROR, nil)
		return fmt.Errorf("can't rename new binary: %w", err)
	}

	os.Remove(tmpBinary)

	dispatcher.DispatchAndWait(EV_BINARY_REPLACE_COMPLETE, nil)

	return nil
}

// formatPubKey formats the public key string to ensure it is in the correct PEM format
func formatPubKey(key string) string {
	if !strings.ContainsRune(key, '\n') {
		key = strutil.Wrap(key, 64)
	}

	if !strings.Contains(key, "-----BEGIN") {
		key = "-----BEGIN PUBLIC KEY-----\n" + key + "\n-----END PUBLIC KEY-----"
	}

	return key
}

// dispatchError dispatches an error event with the provided error message
func dispatchError(dispatcher *events.Dispatcher, err error) error {
	dispatcher.DispatchAndWait(EV_UPDATE_ERROR, err.Error())
	return err
}
