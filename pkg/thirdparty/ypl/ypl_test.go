package ypl_test

import (
	"testing"

	"62tech.co/service/pkg/thirdparty/ypl"
)

func TestGetData(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := ypl.GetData()
			if err != nil {
				t.Errorf("GetData() error = %v", err)
			}
		})
	}
}
