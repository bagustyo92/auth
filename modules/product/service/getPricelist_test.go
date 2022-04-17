package service

import (
	"errors"
	"net/http"

	"github.com/bagustyo92/auth/modules/request/models"
	"github.com/golang/mock/gomock"
)

func (s *productServiceSuite) Test_productService_GetPriceListIncludingPriceInUSD() {
	var (
		forceErr = errors.New("force err")
	)

	s.Run("#Case1: Failed get cache and get curr usd", func() {
		s.cacher.EXPECT().Get(gomock.Any()).Return(nil, forceErr).Times(1)
		s.requestConverter.EXPECT().GetCurrUSD(gomock.Any(), gomock.Any()).Return(forceErr).Times(1)

		data, err := s.productService.GetPriceListIncludingPriceInUSD()
		s.NotNil(err)
		s.Nil(data)
		s.Equal(err, forceErr)
	})

	s.Run("#Case2: Failed get priceList", func() {
		s.cacher.EXPECT().Get(gomock.Any()).Return(nil, forceErr).Times(1)
		s.requestConverter.EXPECT().GetCurrUSD(gomock.Any(), gomock.Any()).Return(nil).Times(1)
		s.cacher.EXPECT().Set(gomock.Any(), gomock.Any(), gomock.Any()).Times(1)
		s.requestEfishery.EXPECT().GetPriceLists(gomock.Any(), gomock.Any()).Return(forceErr).Times(1)

		data, err := s.productService.GetPriceListIncludingPriceInUSD()
		s.NotNil(err)
		s.Nil(data)
		s.Equal(err, forceErr)
	})

	s.Run("#Case3: Success get from cache and success", func() {
		var floatPtr *float64
		flt := 10.0
		floatPtr = &flt

		s.cacher.EXPECT().Get(gomock.Any()).Return(floatPtr, nil).Times(1)
		s.requestEfishery.EXPECT().
			GetPriceLists(gomock.Any(), gomock.Any()).Times(1).Do(func(args1 *http.Header, args2 **[]models.Price) error {
			out := &[]models.Price{{
				Province: "bandung",
				Price:    "1000",
			},
				{
					Province: "bandung",
					Price:    "asdas",
				}}
			*args2 = out
			return nil
		})

		data, err := s.productService.GetPriceListIncludingPriceInUSD()
		s.NotNil(data)
		s.Nil(err)
	})
}
