package dto

import (
	"mime/multipart"
	"net/textproto"
)

type FileDetails struct {
	File     multipart.File       `json:"file"`
	Header   multipart.FileHeader `json:"header"`
	Name     string               `json:"name"`
	Size     int64                `json:"size"`
	MimeType textproto.MIMEHeader `json:"mime_type,omitempty"`
	Ext      string               `json:"ext"`
}
