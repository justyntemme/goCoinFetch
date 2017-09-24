package goCoinFetch
import (
	"github.com/justyntemme/goCoinFetch/web"
)

func GrabTicker(ticker string) string {
	return web.GrabTicker(ticker)
	
}
