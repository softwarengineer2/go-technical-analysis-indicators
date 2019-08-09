package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Double Exponential Moving Average
//Source : https://www.tradingview.com/wiki/Moving_Average
type dema struct {
	count     uint
	period    uint
	prev      float64
	priceType string
	emaArr    [2]*ema
}

//NewDEMA 	: To return dema struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewDEMA(period uint, priceType string) *dema {
	return &dema{
		period:    period,
		prev:      math.NaN(),
		priceType: priceType,
		emaArr:    [2]*ema{NewEMA(period, priceType), NewEMA(period, priceType)},
	}
}

//Calculate : method to Calculate dema and return results as float array
//Return	: dema result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *dema) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	EMAResult := ins.emaArr[0].Calculate(newData)[0]
	ema1 := utils.NewOHLCV(ins.priceType, EMAResult)

	if !math.IsNaN(EMAResult) {
		ins.prev = ((2.0 * EMAResult) - (ins.emaArr[1].Calculate(ema1)[0]))
	} else {
		return []float64{math.NaN()}
	}

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
