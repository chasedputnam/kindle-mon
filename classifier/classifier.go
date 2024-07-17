package classifier

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/chasedputnam/kindle-mon/types"
)

func isUrl(u string) bool {
	for _, proto := range []string{"http://", "https://"} {
		if strings.HasPrefix(u, proto) {
			return true
		}
	}
	return false
}

func isUrlFile(u string) bool {
	file, err := os.Open(u)
	if err != nil {
		return false
	}
	defer file.Close()
	buf := make([]byte, 1024)
	n, _ := file.Read(buf)
	content := string(buf[:n])
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}
		if !strings.HasPrefix(line, "http") {
			return false
		}
	}
	return true
}

func isBook(u string) bool {
	extension := filepath.Ext(u)
	// does file exist
	_, err := os.Stat(u)
	if err != nil {
		return false
	}
	for _, ext := range []string{".mobi", ".pdf", ".epub", ".azw3", ".txt"} {
		if extension == ext {
			return true
		}
	}
	return false
}

func Classify(arg string) []types.Request {
	var requests []types.Request
	if isBook(arg) {
		requests = append(requests, types.NewRequest(arg, types.Ebook, nil))
	}
	return requests
}
