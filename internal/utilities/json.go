package utilities

import (
	"encoding/json"
	"io"
)

// DecodeJson декодирует тело реквеста
func DecodeJson(i interface{}, request io.ReadCloser) error {
	decode := json.NewDecoder(request)
	decode.DisallowUnknownFields()
	return decode.Decode(i)
}
