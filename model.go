package sdk_html_2_image_pdf

type StaticResponseDTO struct {
	UUID string `json:"uuid"`
}

type PdfOpt struct {
	Landscape           bool    `json:"landscape,omitempty"`           // Paper orientation. Defaults to false.
	DisplayHeaderFooter bool    `json:"displayHeaderFooter,omitempty"` // Display header and footer. Defaults to false.
	PrintBackground     bool    `json:"printBackground,omitempty"`     // Print background graphics. Defaults to false.
	Scale               float64 `json:"scale,omitempty"`               // Scale of the webpage rendering. Defaults to 1.
	PaperWidth          float64 `json:"paperWidth,omitempty"`          // Paper width in inches. Defaults to 8.5 inches.
	PaperHeight         float64 `json:"paperHeight,omitempty"`         // Paper height in inches. Defaults to 11 inches.
	MarginTop           float64 `json:"marginTop"`                     // Top margin in inches. Defaults to 1cm (~0.4 inches).
	MarginBottom        float64 `json:"marginBottom"`                  // Bottom margin in inches. Defaults to 1cm (~0.4 inches).
	MarginLeft          float64 `json:"marginLeft"`                    // Left margin in inches. Defaults to 1cm (~0.4 inches).
	MarginRight         float64 `json:"marginRight"`                   // Right margin in inches. Defaults to 1cm (~0.4 inches).
	PageRanges          string  `json:"pageRanges,omitempty"`          // Paper ranges to print, one based, e.g., '1-5, 8, 11-13'. Pages are printed in the document order, not in the order specified, and no more than once. Defaults to empty string, which implies the entire document is printed. The page numbers are quietly capped to actual page count of the document, and ranges beyond the end of the document are ignored. If this results in no pages to print, an error is reported. It is an error to specify a range with start greater than end.
	HeaderTemplate      string  `json:"headerTemplate,omitempty"`      // HTML template for the print header. Should be valid HTML markup with following classes used to inject printing values into them: - date: formatted print date - title: document title - url: document location - pageNumber: current page number - totalPages: total pages in the document  For example, <span class=title></span> would generate span containing the title.
	FooterTemplate      string  `json:"footerTemplate,omitempty"`      // HTML template for the print footer. Should use the same format as the headerTemplate.
	PreferCSSPageSize   bool    `json:"preferCSSPageSize,omitempty"`   // Whether or not to prefer page size as defined by css. Defaults to false, in which case the content will be scaled to fit the paper size.
}

type html2ImageRequest struct {
	Html        string `json:"html" binding:"required"`
	Width       int    `json:"width" binding:"required"`
	Height      int    `json:"height" binding:"required"`
	DelaySecond int    `json:"delay_second" binding:"required"`
}

type html2PdfRequest struct {
	Html        string  `json:"html" binding:"required"`
	DelaySecond int     `json:"delay_second" binding:"required"`
	PdfOpt      *PdfOpt `json:"pdf_opt,omitempty"`
}
