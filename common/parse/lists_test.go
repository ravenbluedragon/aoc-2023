package parse

import (
	"slices"
	"testing"
)

func TestStringToIntList(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		sep     string
		want    []int
		wantErr bool
	}{
		{
			name:    "Normal case",
			s:       "1,2,3,4,5",
			sep:     ",",
			want:    []int{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "Extra spaces",
			s:       "1, 2, 3, 4, 5",
			sep:     ",",
			want:    []int{1, 2, 3, 4, 5},
			wantErr: false,
		},
		{
			name:    "Non-numeric characters",
			s:       "1,a,3,4,5",
			sep:     ",",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := StringToIntList(tt.s, tt.sep)
			if (err != nil) != tt.wantErr {
				t.Errorf("StringToIntList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("StringToIntList() = %v, want %v", got, tt.want)
			}
		})
	}
}
