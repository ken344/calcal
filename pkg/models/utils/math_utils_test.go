package utils

import "testing"

func TestRounding(t *testing.T) {
	type args struct {
		f    float64
		prec int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{
			name: "Test case prec0-down",
			args: args{
				f:    1.23456789,
				prec: 0},
			want: 1,
		},
		{
			name: "Test case prec0-up",
			args: args{
				f:    1.56789999,
				prec: 0},
			want: 2,
		},
		{
			name: "Test case prec1-down",
			args: args{
				f:    1.23456789,
				prec: 1},
			want: 1.2,
		},
		{
			name: "Test case prec1-up",
			args: args{
				f:    1.56789999,
				prec: 1},
			want: 1.6,
		},
		{
			name: "Test case prec2-down",
			args: args{
				f:    1.23456789,
				prec: 2},
			want: 1.23,
		},
		{
			name: "Test case prec2-up",
			args: args{
				f:    1.56789999,
				prec: 2},
			want: 1.57,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Rounding(tt.args.f, tt.args.prec); got != tt.want {
				t.Errorf("Rounding() = %v, want %v", got, tt.want)
			}
		})
	}
}
