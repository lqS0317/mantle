package tokenprice

import (
	"net/http"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

func TestGetTokenPrice(t *testing.T) {
	tokenPricer := NewClient("https://api.bybit.com", 3)
	ethPrice, err := tokenPricer.Query("ETHUSDT")
	require.NoError(t, err)
	t.Logf("Btc price:%v", ethPrice)

	bitPrice, err := tokenPricer.Query("BITUSDT")
	require.NoError(t, err)
	t.Logf("Btc price:%v", bitPrice)

	t.Log(ethPrice.Quo(ethPrice, bitPrice))
	ratio, err := tokenPricer.PriceRatio()
	require.NoError(t, err)
	t.Logf("ratio:%v", ratio)

}

func TestMockTokenPricer(t *testing.T) {
	route := gin.New()
	tokenPricer := NewClient("https://api.bybit.com", 3)
	route.GET("/spot/quote/v1/ticker/price", func(context *gin.Context) {
		symbol := context.Query("symbol")
		t.Logf("symbol:%v", symbol)
		p, err := tokenPricer.Query(symbol)
		if err != nil {
			t.Log(err)
		}
		context.JSON(http.StatusOK, p)
	})
	route.Run(":8000")
}
