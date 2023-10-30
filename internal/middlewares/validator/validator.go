package validator

import (
	"bufio"
	"io"
	"mime/multipart"
	"path"
	"path/filepath"
	"regexp"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) FileExtValidation(header *multipart.FileHeader) string {
	var message string

	ext := path.Ext(header.Filename)
	if ext != ".txt" && ext != ".pgp" && ext != ".gpg" {
		message = "File format is not allowed"
	}

	return message
}

func (v *Validator) ContentTypeValidation(header *multipart.FileHeader) string {
	var message string

	contentType := header.Header["Content-Type"][0]
	if contentType != "text/plain" && contentType != "application/pgp-encrypted" {
		message = "The provided file format is not allowed. Please upload TXT, PGP or GPG file"
	}

	return message
}

func (v *Validator) BinaryFileValidation(header *multipart.FileHeader) string {
	var message string

	file, _ := header.Open()
	r := bufio.NewReader(file)
	for {
		if c, _, err := r.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				message = "Error reading file"
			}
		} else {
			if c == '\000' {
				message = "Error! Malicious file format detected"
			}
		}
	}

	return message
}

func (v *Validator) SpecialCharacterFileNameValidation(header *multipart.FileHeader) string {
	var message string

	regex, _ := regexp.Compile("[!\"#$&'()*.,;<>?[]^`{|}]")
	filename := header.Filename[:len(header.Filename)-len(filepath.Ext(header.Filename))]

	if regex.MatchString(filename) {
		message = "Nama file yang diupload tidak boleh mengandung !\"#$&'()*.,;<>?[]^`{|}"
	}

	return message
}
