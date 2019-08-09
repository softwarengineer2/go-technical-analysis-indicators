package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Triple Exponential Moving Average
//Source : https://www.tradingview.com/wiki/Moving_Average
type tema struct {
	count     uint
	period    uint
	prev      float64
	priceType string
	emaArr    [3]*ema
}

//NewTEMA 	: To return tema struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewTEMA(period uint, priceType string) *tema {
	return &tema{
		period:    period,
		prev:      math.NaN(),
		priceType: priceType,
		emaArr:    [3]*ema{NewEMA(period, priceType), NewEMA(period, priceType), NewEMA(period, priceType)},
	}
}

//Calculate : method to Calculate tema and return results as float array
//Return	: tema result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *tema) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	EMAResult1 := ins.emaArr[0].Calculate(newData)[0]
	ema1 := utils.NewOHLCV(ins.priceType, EMAResult1)

	EMAResult2 := ins.emaArr[1].Calculate(ema1)[0]
	ema2 := utils.NewOHLCV(ins.priceType, EMAResult2)

	if !math.IsNaN(EMAResult1) && !math.IsNaN(EMAResult2) {
		ins.prev = ((3.0 * EMAResult1) - (3.0 * EMAResult2)) + ins.emaArr[2].Calculate(ema2)[0]
	} else {
		return []float64{math.NaN()}
	}

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
