package main

import (
	c "asciicolor/colors"
	utils "asciicolor/utils"
	"testing"
)

func TestGetRGBValue(t *testing.T) {
	tests := []struct {
		name    string
		color   string
		want    []int
		wantErr bool
	}{
		{
			name:    "Test aquamarine color",
			color:   "aquamarine",
			want:    []int{127, 255, 212},
			wantErr: false,
		},
		{
			name:    "Test alice Blue color",
			color:   "alice Blue",
			want:    []int{240, 248, 255},
			wantErr: false,
		},
		{
			name:    "Test invalid color",
			color:   "invalid color",
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := c.GetRGBValue(tt.color)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRGBValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestRGB_CONVERTER(t *testing.T) {
	tests := []struct {
		name  string
		color []int
		s     string
		want  string
	}{
		{
			name:  "Red color",
			color: []int{255, 0, 0},
			s:     "Hello, World!",
			want:  "\x1b[38;2;255;0;0mHello, World!\x1b[0m",
		},
		{
			name:  "Green color",
			color: []int{0, 255, 0},
			s:     "Hello, World!",
			want:  "\x1b[38;2;0;255;0mHello, World!\x1b[0m",
		},
		{
			name:  "Blue color",
			color: []int{0, 0, 255},
			s:     "Hello, World!",
			want:  "\x1b[38;2;0;0;255mHello, World!\x1b[0m",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := c.RGB_CONVERTER(tt.color, tt.s); got != tt.want {
				t.Errorf("RGB_CONVERTER() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHandleRGB(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		wantErr bool
	}{
		{
			name:    "Argument 1",
			arg:     `rgb(123, 45, 67 , 90)`,
			wantErr: true,
		},
		{
			name:    "Argument 2",
			arg:     `rgb(255 , 45, 67)`,
			wantErr: false,
		},
		{
			name:    "Argument 3",
			arg:     `rgb(123, 45, 89)`,
			wantErr: false,
		},
		{
			name:    "Argument 4",
			arg:     `rgb(123, 45, 67)`,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := utils.HandleRGB(tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("HandleRGB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
