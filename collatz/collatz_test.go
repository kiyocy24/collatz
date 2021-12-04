package collatz

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestCollatz(t *testing.T) {
	type args struct {
		n uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "n = 1",
			args: args{n: 1},
			want: 4,
		},
		{
			name: "n = 2",
			args: args{n: 2},
			want: 1,
		},
		{
			name: "n = 7",
			args: args{n: 7},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := collatz(tt.args.n)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("collatz() is mismatch (-want +got)\n%s", diff)
			}

		})
	}
}

func TestCollatzs(t *testing.T) {
	type args struct {
		start uint64
		end   uint64
	}
	tests := []struct {
		name string
		args args
		want [][]uint64
	}{
		{
			name: "start = 1, end = 1",
			args: args{start: 1, end: 1},
			want: [][]uint64{
				{1},
			},
		},
		{
			name: "start = 1, end = 2",
			args: args{start: 1, end: 2},
			want: [][]uint64{
				{1},
				{2, 1},
			},
		},
		{
			name: "start = 1, end = 3",
			args: args{start: 1, end: 3},
			want: [][]uint64{
				{1},
				{2, 1},
				{3, 10, 5, 16, 8, 4, 2, 1},
			},
		},
		{
			name: "start = 2, end = 3",
			args: args{start: 2, end: 3},
			want: [][]uint64{
				{2, 1},
				{3, 10, 5, 16, 8, 4, 2, 1},
			},
		},
		{
			name: "start = 3, end = 3",
			args: args{start: 3, end: 3},
			want: [][]uint64{
				{3, 10, 5, 16, 8, 4, 2, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Collatzs(tt.args.start, tt.args.end)
			if diff := cmp.Diff(tt.want, got); diff != "" {
				t.Errorf("Collatzs() is mismatch (-want +got)\n%s", diff)
			}

		})
	}
}
