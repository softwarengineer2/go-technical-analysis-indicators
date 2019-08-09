package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Exponential Moving Average
//Source : https://www.tradingview.com/wiki/Moving_Average
type ema struct {
	count     uint
	prev      float64
	period    uint
	priceType string
	sma1      *sma
}

//NewEMA 	: To return ema struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewEMA(period uint, priceType string) *ema {
	return &ema{
		prev:      math.NaN(),
		period:    period,
		priceType: priceType,
		sma1:      NewSMA(period, priceType),
	}
}

//Calculate : method to Calculate ema and return results as float array
//Return	: ema result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *ema) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	if ins.count <= ins.period {
		ins.prev = ins.sma1.Calculate(newData)[0]
	}
	if ins.count > ins.period {
		ins.prev = (newPrice-ins.prev)*(2.0/(float64(ins.period)+1.0)) + ins.prev
	}

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
