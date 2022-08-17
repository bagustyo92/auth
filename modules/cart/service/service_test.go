package service

func (c *CartSuite) TestNewCartService() {
	res := NewCartService(c.cartRepo)
	c.NotNil(res)
	c.Equal(c.cartService, res)
}
