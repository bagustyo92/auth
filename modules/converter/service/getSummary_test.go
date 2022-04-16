package service

import "testing"

func Test_findMedian(t *testing.T) {
	type args struct {
		data []float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			"#Case1",
			args{[]float64{1, 2, 3, 4, 5, 6, 7}},
			4,
		},
		{
			"#Case2",
			args{[]float64{1, 2, 3, 4, 5, 6}},
			float64(7) / float64(2),
		},
		{
			"#Case3",
			args{[]float64{4, 5, 1, 2, 7, 3, 6}},
			4,
		},
		{
			"#Case4",
			args{[]float64{5, 3, 6, 4, 2, 1}},
			float64(7) / float64(2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedian(tt.args.data); got != tt.want {
				t.Errorf("findMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
