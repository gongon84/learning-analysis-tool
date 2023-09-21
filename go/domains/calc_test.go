package domains

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case1",
			args: args{1, 2},
			want: 3,
		},
		{
			name: "case2",
			args: args{0, 0},
			want: 0,
		},
		{
			name: "case3",
			args: args{-1, -2},
			want: -3,
		},
		{
			name: "case4",
			args: args{5, -5},
			want: 0,
		},
	}

	for _, tt := range tests {
		result := Add(tt.args.x, tt.args.y)
		if result != tt.want {
			t.Errorf("%s: For %d + %d, want %d but got %d", tt.name, tt.args.x, tt.args.y, tt.want, result)
		}
	}
}
