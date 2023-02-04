package services

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/4ynyky/grpc_app/internal/domains"
	mock_services "github.com/4ynyky/grpc_app/internal/services/mock"
	"github.com/golang/mock/gomock"
)

func TestSet(t *testing.T) {
	type fields struct {
		mockStorer *mock_services.MockStorer
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    domains.Item
		wantErr bool
	}{
		{
			name: "Set empty domains.Item",
			prepare: func(f *fields) {
				f.mockStorer.EXPECT().Set(domains.Item{}).Return(fmt.Errorf("Err"))
			},
			args:    domains.Item{},
			wantErr: true,
		},
		{
			name: "Set ok domains.Item",
			prepare: func(f *fields) {
				f.mockStorer.EXPECT().Set(domains.Item{ID: "1", Value: "15"}).Return(nil)
			},
			args:    domains.Item{ID: "1", Value: "15"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				mockStorer: mock_services.NewMockStorer(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			ss := NewStorageService(f.mockStorer)

			if err := ss.Set(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type fields struct {
		mockStorer *mock_services.MockStorer
	}
	tests := []struct {
		name     string
		prepare  func(f *fields)
		args     string
		expected domains.Item
		wantErr  bool
	}{
		{
			name: "Positive test",
			prepare: func(f *fields) {
				f.mockStorer.EXPECT().Get("1").Return(domains.Item{ID: "1", Value: "ABC"}, nil)
			},
			args:     "1",
			expected: domains.Item{ID: "1", Value: "ABC"},
			wantErr:  false,
		},
		{
			name: "Negative test",
			prepare: func(f *fields) {
				f.mockStorer.EXPECT().Get("1").Return(domains.Item{ID: "1", Value: "ABC"}, fmt.Errorf("Error"))
			},
			args:     "1",
			expected: domains.Item{},
			wantErr:  true,
		},
		{
			name: "Get domains.Item with empty ID",
			prepare: func(f *fields) {
				f.mockStorer.EXPECT().Get("1").Return(domains.Item{ID: "", Value: "ABC"}, nil)
			},
			args:     "1",
			expected: domains.Item{},
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				mockStorer: mock_services.NewMockStorer(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			ss := NewStorageService(f.mockStorer)

			item, err := ss.Get(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(item, tt.expected) {
				t.Errorf("Expected: %v, got: %v", tt.expected, item)
			}
		})
	}
}

func TestDelete(t *testing.T) {
	type fields struct {
		mockStorer *mock_services.MockStorer
	}
	tests := []struct {
		name    string
		prepare func(f *fields)
		args    string
		wantErr bool
	}{
		{
			name: "Positive",
			prepare: func(f *fields) {
				f.mockStorer.EXPECT().Delete("1").Return(nil)
			},
			args:    "1",
			wantErr: false,
		},
		{
			name: "Negative",
			prepare: func(f *fields) {
				f.mockStorer.EXPECT().Delete("1").Return(fmt.Errorf("Error"))
			},
			args:    "1",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			f := fields{
				mockStorer: mock_services.NewMockStorer(ctrl),
			}
			if tt.prepare != nil {
				tt.prepare(&f)
			}

			ss := NewStorageService(f.mockStorer)

			if err := ss.Delete(tt.args); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
