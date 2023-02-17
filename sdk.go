package sdk_html_2_image_pdf

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/WebXense/ginger/ginger"
	"github.com/WebXense/http"
)

type html2ImagePdfSdk struct {
	host string
}

func NewHtml2ImagePdfSdk(host string) *html2ImagePdfSdk {
	return &html2ImagePdfSdk{
		host: host,
	}
}

func (s *html2ImagePdfSdk) Html2Image(html string, width, height, delaySecond int) (*StaticResponseDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](s.host+url_html_2_image, nil, nil, &html2ImageRequest{
		Html:        html,
		Width:       width,
		Height:      height,
		DelaySecond: delaySecond,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("html 2 image failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &StaticResponseDTO{}), resp, nil
}

func (s *html2ImagePdfSdk) Html2Pdf(html string, delaySecond int, pdfOpt *PdfOpt) (*StaticResponseDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](s.host+url_html_2_pdf, nil, nil, &html2PdfRequest{
		Html:        html,
		DelaySecond: delaySecond,
		PdfOpt:      pdfOpt,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("html 2 pdf failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &StaticResponseDTO{}), resp, nil
}

func mapByJson[T any](from interface{}, to *T) *T {
	j, err := json.Marshal(from)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(j, to)
	if err != nil {
		return nil
	}
	return to
}
