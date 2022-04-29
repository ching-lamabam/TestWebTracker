package pkg

import "os"

type TrackerSvc struct {
}

func NewSvc() TrackerSvc {
	return TrackerSvc{}
}

func (svc TrackerSvc) GetPing() int {
	_, err := os.Stat("/tmp/ok")
	if err == nil {
		return 1
	}
	return -1
}

func (svc TrackerSvc) GetImg() string {
	_, err := os.Stat("/tmp/img.png")
	if err == nil {
		return "/tmp/img.png"
	}
	return ""
}
