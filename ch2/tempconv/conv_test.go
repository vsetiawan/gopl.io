package tempconv

import (
	"testing"
)

const floatingPointPrecision = 1e-9

func TestCToK(t *testing.T) {
	type args struct {
		c Celsius
	}
	tests := []struct {
		name string
		args args
		want Kelvin
	}{
		{
			name: "melting",
			args: args{
				c: 0,
			},
			want: 273.15,
		},
		{
			name: "boiling",
			args: args{
				c: 100,
			},
			want: 373.15,
		},
		{
			name: "absolute zero",
			args: args{
				c: -273.15,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CToK(tt.args.c); got != tt.want {
				t.Errorf("CToK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFtoK(t *testing.T) {
	type args struct {
		f Fahrenheit
	}
	tests := []struct {
		name string
		args args
		want Kelvin
	}{
		{
			name: "melting",
			args: args{
				f: 32,
			},
			want: 273.15,
		},
		{
			name: "melting",
			args: args{
				f: 212,
			},
			want: 373.15,
		},
		{
			name: "absolute zero",
			args: args{
				f: -459.67,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FtoK(tt.args.f); (got - tt.want) > floatingPointPrecision {
				t.Errorf("FtoK() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKtoC(t *testing.T) {
	type args struct {
		k Kelvin
	}
	tests := []struct {
		name string
		args args
		want Celsius
	}{
		{
			name: "melting",
			args: args{
				k: 273.15,
			},
			want: 0,
		},
		{
			name: "boiling",
			args: args{
				k: 373.15,
			},
			want: 100,
		},
		{
			name: "absolute zero",
			args: args{
				k: 0,
			},
			want: -273.15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KtoC(tt.args.k); got != tt.want {
				t.Errorf("KtoC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKtoF(t *testing.T) {
	type args struct {
		k Kelvin
	}
	tests := []struct {
		name string
		args args
		want Fahrenheit
	}{
		{
			name: "melting",
			args: args{
				k: 273.15,
			},
			want: 32,
		},
		{
			name: "melting",
			args: args{
				k: 373.15,
			},
			want: 212,
		},
		{
			name: "absolute zero",
			args: args{
				k: 0,
			},
			want: -459.67,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KtoF(tt.args.k); (got - tt.want) > floatingPointPrecision {
				t.Errorf("KtoF() = %v, want %v", got, tt.want)
			}
		})
	}
}
