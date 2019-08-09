package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Generalized Double Exponential Moving Average
//Source : https://www.tradingview.com/wiki/Moving_Average
type gdema struct {
	count        uint
	period       uint
	prev         float64
	volumeFactor float64
	priceType    string
	emaArr       [2]*ema
}

//NewGDEMA : To return gdema struct instance
//Params
//period	 : calculation period
//vFactor	 : volume factor
//priceType  : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewGDEMA(period uint, vFactor float64, priceType string) *gdema {
	return &gdema{
		period:       period,
		prev:         math.NaN(),
		volumeFactor: vFactor,
		priceType:    priceType,
		emaArr:       [2]*ema{NewEMA(period, priceType), NewEMA(period, priceType)},
	}
}

//Calculate : method to Calculate gdema and return results as float array
//Return	: Kaufman result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *gdema) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	EMAResult := ins.emaArr[0].Calculate(newData)[0]

	ema1 := utils.NewOHLCV(ins.priceType, EMAResult)

	ins.prev = (((1.0 + ins.volumeFactor) * EMAResult) - (ins.emaArr[1].Calculate(ema1)[0])*ins.volumeFactor)

	return []float64{ins.prev}
}
