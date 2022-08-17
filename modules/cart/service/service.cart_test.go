package service

import (
	"errors"
	"reflect"
	"testing"

	mock_repo "github.com/bagustyo92/auth/mocks/modules/cart/repository"
	"github.com/bagustyo92/auth/modules/cart/models"
	"github.com/golang/mock/gomock"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

func Test_gatheringMoneyOnRupiah(t *testing.T) {
	type args struct {
		moneyInput int
	}
	tests := []struct {
		name string
		args args
		want map[int]int
	}{
		{
			"100.000",
			args{100000},
			map[int]int{100000: 1},
		},
		{
			"102.000",
			args{102000},
			map[int]int{100000: 1, 2000: 1},
		},
		{
			"40.000",
			args{40000},
			map[int]int{20000: 2},
		},
		{
			"40.250",
			args{40250},
			map[int]int{100: 1, 200: 1, 20000: 2},
		},
		{
			"1.052.250",
			args{1052250},
			map[int]int{100000: 10, 50000: 1, 2000: 1, 100: 1, 200: 1},
		},
		{
			"0",
			args{0},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := gatheringMoneyOnRupiah(tt.args.moneyInput); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gatheringMoneyOnRupiah() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_stringEditAssessment(t *testing.T) {
	type args struct {
		text1 string
		text2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"TEST 1",
			args{
				"telkom",
				"telecom",
			},
			3,
		},
		{
			"TEST 2",
			args{
				"telkom",
				"telcom",
			},
			2,
		},
		{
			"TEST 3",
			args{
				"telkom",
				"telkoom",
			},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := stringEditAssessment(tt.args.text1, tt.args.text2); got != tt.want {
				t.Errorf("stringEditAssessment() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeMoneyResponse(t *testing.T) {
	type args struct {
		moneyGathered map[int]int
	}
	tests := []struct {
		name string
		args args
		want map[string]int
	}{
		{
			"Test1",
			args{
				map[int]int{100000: 1, 50000: 1},
			},
			map[string]int{"Rp 100.000,-": 1, "Rp 50.000,-": 1},
		},
		{
			"Test2",
			args{
				map[int]int{20000: 1},
			},
			map[string]int{"Rp 20.000,-": 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeMoneyResponse(tt.args.moneyGathered); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeMoneyResponse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateQty(t *testing.T) {
	type args struct {
		qty int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{0},
			1,
		},
		{
			"Test 2",
			args{15},
			15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateQty(tt.args.qty); got != tt.want {
				t.Errorf("validateQty() = %v, want %v", got, tt.want)
			}
		})
	}
}

type CartSuite struct {
	suite.Suite
	*require.Assertions

	ctrl        *gomock.Controller
	cartRepo    *mock_repo.MockCartRepoInterface
	cartService CartInterface
}

func TestCartSuite(t *testing.T) {
	suite.Run(t, new(CartSuite))
}

func (c *CartSuite) SetupTest() {
	c.Assertions = require.New(c.T())
	c.ctrl = gomock.NewController(c.T())

	c.cartRepo = mock_repo.NewMockCartRepoInterface(c.ctrl)
	c.cartService = NewCartService(c.cartRepo)
}

func (c *CartSuite) Test_cartService_GetCart() {
	c.Run("Success", func() {
		mockRes := models.Cart{}
		c.cartRepo.EXPECT().GetCart(gomock.Any(), gomock.Any()).
			Times(1).Return(&mockRes, nil)
		actual, err := c.cartService.GetCart(uuid.UUID{}, models.ProductCartFilter{})
		c.Nil(err)
		c.NotNil(actual)
		c.Equal(&mockRes, actual)
	})

	c.Run("Failed", func() {
		mockErr := errors.New("forceErr")
		c.cartRepo.EXPECT().GetCart(gomock.Any(), gomock.Any()).
			Times(1).Return(nil, mockErr)
		actual, err := c.cartService.GetCart(uuid.UUID{}, models.ProductCartFilter{})
		c.Nil(actual)
		c.NotNil(err)
		c.Equal(mockErr, err)
	})
}

func (c *CartSuite) Test_cartService_DeleteProductFromCart() {
	c.Run("Success", func() {
		c.cartRepo.EXPECT().DeleteProductCart(gomock.Any(), gomock.Any()).
			Times(1).Return(nil)
		err := c.cartService.DeleteProductFromCart(uuid.UUID{}, "")
		c.Nil(err)
	})

	c.Run("Fail", func() {
		mockErr := errors.New("forceErr")
		c.cartRepo.EXPECT().DeleteProductCart(gomock.Any(), gomock.Any()).
			Times(1).Return(mockErr)
		err := c.cartService.DeleteProductFromCart(uuid.UUID{}, "")
		c.NotNil(err)
		c.Equal(mockErr, err)
	})
}

func (c *CartSuite) Test_cartService_AddProductToCart() {
	c.Run("#Case1: Fail when insert cart", func() {
		mockErr := errors.New("forceErr")
		c.cartRepo.EXPECT().InsertCart(gomock.Any()).
			Times(1).Return(nil, mockErr)
		res, err := c.cartService.AddProductToCart(models.Cart{})
		c.Nil(res)
		c.NotNil(err)
		c.Equal(mockErr, err)
	})

	c.Run("#Case2: Fail when get cart", func() {
		mockid := uuid.NewV4()
		mockErr := errors.New("forceErr")
		c.cartRepo.EXPECT().GetCart(gomock.Any(), gomock.Any()).
			Times(1).Return(nil, mockErr)
		res, err := c.cartService.AddProductToCart(models.Cart{
			Base: models.Base{
				ID: mockid,
			},
		})
		c.Nil(res)
		c.NotNil(err)
		c.Equal(mockErr, err)
	})

	c.Run("#Case3: Fail when insert new product", func() {
		mockid := uuid.NewV4()
		mockErr := errors.New("forceErr")
		c.cartRepo.EXPECT().GetCart(gomock.Any(), gomock.Any()).
			Times(1).Return(&models.Cart{Products: []models.ProductCart{}}, nil)
		c.cartRepo.EXPECT().InsertProductCart(gomock.Any()).Times(1).Return(mockErr)

		res, err := c.cartService.AddProductToCart(models.Cart{
			Base: models.Base{
				ID: mockid,
			},
			Products: []models.ProductCart{
				{ProductCode: "111"},
			},
		})
		c.Nil(res)
		c.NotNil(err)
		c.Equal(mockErr, err)
	})

	c.Run("#Case4: Fail when insert new product", func() {
		mockid := uuid.NewV4()
		mockErr := errors.New("forceErr")
		c.cartRepo.EXPECT().GetCart(gomock.Any(), gomock.Any()).
			Times(1).Return(&models.Cart{Products: []models.ProductCart{
			{ProductCode: "111"},
		}}, nil)
		c.cartRepo.EXPECT().UpdateProductCart(gomock.Any(), gomock.Any()).Times(1).Return(mockErr)

		res, err := c.cartService.AddProductToCart(models.Cart{
			Base: models.Base{
				ID: mockid,
			},
			Products: []models.ProductCart{
				{ProductCode: "111"},
			},
		})
		c.Nil(res)
		c.NotNil(err)
		c.Equal(mockErr, err)
	})

	c.Run("#Case5: Success", func() {
		mockid := uuid.NewV4()
		c.cartRepo.EXPECT().GetCart(gomock.Any(), gomock.Any()).
			Times(1).Return(&models.Cart{Products: []models.ProductCart{
			{ProductCode: "111"},
		}}, nil)
		c.cartRepo.EXPECT().UpdateProductCart(gomock.Any(), gomock.Any()).Times(1).Return(nil)

		res, err := c.cartService.AddProductToCart(models.Cart{
			Base: models.Base{
				ID: mockid,
			},
			Products: []models.ProductCart{
				{ProductCode: "111"},
			},
		})
		c.Nil(res)
		c.Nil(err)
	})
}

func (c *CartSuite) Test_cartService_MoneyTest() {
	res := c.cartService.MoneyTest(1000)
	c.Equal(map[string]int{
		"Rp 1.000,-": 1,
	}, res)
}

func (c *CartSuite) Test_StringTest() {
	res := c.cartService.StringTest("telkom", "tlkom")
	c.True(res)
}
