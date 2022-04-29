package pkg

import (
	"reflect"
	"testing"
)

func TestNewSvc(t *testing.T) {
	tests := []struct {
		name string
		want TrackerSvc
	}{
		{
			name: "success",
			want: NewSvc(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSvc(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSvc() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrackerSvc_GetImg(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "fail",
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := TrackerSvc{}
			if got := svc.GetImg(); got != tt.want {
				t.Errorf("GetImg() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTrackerSvc_GetPing(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{
			name: "fail",
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := TrackerSvc{}
			if got := svc.GetPing(); got != tt.want {
				t.Errorf("GetPing() = %v, want %v", got, tt.want)
			}
		})
	}
}
