package types

type FileType string

var (
	Ebook   FileType = "ebook"
	Url     FileType = "url"
	UrlFile FileType = "urlfile"
)

type Request struct {
	Path    string
	Type    FileType
	Options map[string]string
}

func NewRequest(path string, fileType FileType, opts map[string]string) Request {
	return Request{path, fileType, opts}
}
