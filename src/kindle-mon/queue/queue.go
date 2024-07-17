package queue

import (
	"github.com/chasedputnam/kindle-mon/mailer"
	"github.com/chasedputnam/kindle-mon/types"
)

func Queue(downloadRequests []types.Request) []types.Request {
	var processedRequests []types.Request
	for _, req := range downloadRequests {
		switch req.Type {
		case types.Ebook:
			processedRequests = append(processedRequests, req)
			continue
		}
	}
	return processedRequests
}

func SendMail(mailRequests []types.Request, timeout int) {
	var filePaths []string
	for _, req := range mailRequests {
		filePaths = append(filePaths, req.Path)
	}
	mailer.Send(filePaths, timeout)
}
