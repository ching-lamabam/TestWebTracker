package rest

import (
	"github.com/ching-lamabam/web-tracker/internal/pkg"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want trackerRest
	}{
		{
			name: "success",
			want: New(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_trackerRest_GetImg(t *testing.T) {
	type fields struct {
		trackerSvc pkg.TrackerSvc
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name           string
		expectedStatus int
		args           args
	}{
		{
			name:           "fail",
			expectedStatus: 503,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rest := &trackerRest{
				trackerSvc: pkg.NewSvc(),
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "web-tracker/v1/img", nil)
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.RemoveExtraSlash = true
			r.GET("web-tracker/v1/img", rest.GetImg)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}

func Test_trackerRest_GetPing(t *testing.T) {
	type fields struct {
		trackerSvc pkg.TrackerSvc
	}
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name           string
		expectedStatus int
		args           args
	}{
		{
			name:           "fail",
			expectedStatus: 503,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rest := &trackerRest{
				trackerSvc: pkg.NewSvc(),
			}
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "web-tracker/v1/ping", nil)
			gin.SetMode(gin.TestMode)
			r := gin.Default()
			r.RemoveExtraSlash = true
			r.GET("web-tracker/v1/ping", rest.GetPing)
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}
