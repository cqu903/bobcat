package utils

import (
	"strings"
)

//根据对应的文件后缀名返回合适的content-type
func DecideRightContentType(suffix string) string {
	suffix = strings.ToLower(suffix)
	switch suffix {
	case ".tif":
		return "image/tiff"
	case "tiff":
		return "image/tiff"
	case "gif":
		return "image/gif"
	case "jpg":
		return "image/jpeg"
	case "jpeg":
		return "image/jpeg"
	case "png":
		return "image/png"
	case "css":
		return "text/css"
	case "htm":
		return "text/html"
	case "html":
		return "text/html"
	case "svg":
		return "text/xml"
	case "xhtml":
		return "text/html"
	case "txt":
		return "text/plain"
	case "js":
		return "application/x-javascript"
	case "xls":
		return "application/vnd.ms-excel"
	case "mp4":
		return "video/mpeg4"
	case "pdf":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}
