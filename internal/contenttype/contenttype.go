package contenttype

const (
	// Header is HTTP Content-Type header name.
	Header = "Content-Type"
	// CharsetUTF8 is Content-Type header's character set portition with a value of UTF-8.
	CharsetUTF8 = "; charset=utf-8"
)

// Type defines all supported Content-Type header values.
type Type string

// String returns Content-Type header value. For example `application/json` or `text/html`.
func (t Type) String() string {
	return string(t)
}

const (
	// JSON content type.
	JSON Type = "application/json"
	// HTML content type.
	HTML Type = "text/html"
	// Text content type.
	Text Type = "text/plain"
)
