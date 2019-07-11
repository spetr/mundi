package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spetr/docconv"
	"github.com/vimeo/go-magic/magic"
)

func apiAutoConvert(c *gin.Context) {
	var resp struct {
		FileName      string `json:"filename"`
		MimeType      string `json:"mime"`
		Text          string `json:"text"`
		SupportedType bool   `json:"supported_type"`
		Error         error  `json:"error,omitempty"`
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	resp.FileName = header.Filename

	buf, _ := ioutil.ReadAll(file)
	resp.MimeType = magic.MimeFromBytes(buf)

	switch resp.MimeType {

	// DOC
	case "application/msword":
		resp.SupportedType = true
		text, _, err := docconv.ConvertDoc(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text

	// DOCX
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		resp.SupportedType = true
		text, _, err := docconv.ConvertDocx(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text

	// ODT
	case "application/vnd.oasis.opendocument.text":
		resp.SupportedType = true
		text, _, err := docconv.ConvertODT(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text

	// PDF
	case "application/pdf":
		resp.SupportedType = true
		text, _, err := docconv.ConvertPDF(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text

	// RTF
	case "text/rtf":
		resp.SupportedType = true
		text, _, err := docconv.ConvertRTF(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text

	// HTML
	case "text/html":
		resp.SupportedType = true
		text, _, err := docconv.ConvertHTML(bytes.NewReader(buf), true)
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text

	// Apple Pages
	// TODO - sprava detekce mime typu!!!
	case "application/vnd.apple.pages":
		resp.SupportedType = true
		text, _, err := docconv.ConvertPages(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text

	// Image - OCR
	case "image/jpeg", "image/png", "image/tiff", "image/gif":
		resp.SupportedType = true
		text, _, err := docconv.ConvertImage(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text

	}

	c.JSON(http.StatusOK, resp)
}
