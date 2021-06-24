package zaper

import (
	"net/http"
	"reflect"
	"testing"

	"go.uber.org/zap"
)

func TestLogMiddleware(t *testing.T) {
	type args struct {
		h http.Handler
	}
	tests := []struct {
		name string
		args args
		want http.Handler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogMiddleware(tt.args.h); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_wrapReqFields(t *testing.T) {
	type args struct {
		logger *zap.Logger
		r      *http.Request
	}
	tests := []struct {
		name string
		args args
		want *zap.Logger
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := wrapReqFields(tt.args.logger, tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrapReqFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
