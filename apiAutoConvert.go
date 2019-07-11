package main

import (
	"bytes"
	"io/ioutil"
	"net/http"

	"code.sajari.com/docconv"
	"github.com/gin-gonic/gin"
	"github.com/vimeo/go-magic/magic"
)

func apiAutoConvert(c *gin.Context) {
	var resp struct {
		FileName string `json:"filename"`
		MimeType string `json:"mime"`
		Text     string `json:"text"`
		Error    error  `json:"error,omitempty"`
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
		text, _, err := docconv.ConvertDoc(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text
	// DOCX
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		text, _, err := docconv.ConvertDocx(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text
	// ODT
	case "application/vnd.oasis.opendocument.text":
		text, _, err := docconv.ConvertODT(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text
	// PDF
	case "application/pdf":
		text, _, err := docconv.ConvertPDF(bytes.NewReader(buf))
		if err != nil {
			resp.Error = err
			break
		}
		resp.Text = text
	}

	c.JSON(http.StatusOK, resp)
}
