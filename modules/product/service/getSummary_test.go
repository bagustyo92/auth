package service

import (
	"errors"
	"net/http"
	"testing"

	reqModel "github.com/bagustyo92/auth/modules/request/models"
	"github.com/golang/mock/gomock"
)

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

func (s *productServiceSuite) Test_productService_GetSummaryPriceList() {
	var (
		forceErr = errors.New("force err")
	)

	s.Run("#Case1: Failed get GetPriceListIncludingPriceInUSD", func() {
		var floatPtr *float64
		flt := 10.0
		floatPtr = &flt

		s.cacher.EXPECT().Get(gomock.Any()).Return(floatPtr, nil).Times(1)
		s.requestEfishery.EXPECT().GetPriceLists(gomock.Any(), gomock.Any()).
			Times(1).Return(forceErr)

		data, err := s.productService.GetSummaryPriceList()
		s.NotNil(err)
		s.Equal(err, forceErr)
		s.Nil(data)
	})

	s.Run("#Case2: Failed get GetPriceListIncludingPriceInUSD", func() {
		var floatPtr *float64
		flt := 10.0
		floatPtr = &flt

		s.cacher.EXPECT().Get(gomock.Any()).Return(floatPtr, nil).Times(1)
		s.requestEfishery.EXPECT().GetPriceLists(gomock.Any(), gomock.Any()).
			Times(1).Do(func(args1 *http.Header, args2 **[]reqModel.Price) error {
			out := &[]reqModel.Price{
				{
					Province: "bandung",
					Price:    "1000",
				},
				{
					Province: "bandung",
					Price:    "asdas",
				},
			}
			*args2 = out
			return forceErr
		})

		data, err := s.productService.GetSummaryPriceList()
		s.Nil(err)
		s.NotNil(data)
	})
}
