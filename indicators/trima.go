package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Triangular Moving Average
//Source : https://www.thebalance.com/triangular-moving-average-tma-description-and-uses-1031203
type trima struct {
	count     uint
	limit     uint
	period    uint
	prev      float64
	priceType string
	smaArr    []*sma
}

//NewTRIMA 	: To return trima struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewTRIMA(period uint, priceType string) *trima {
	return &trima{
		limit:     (period + ((period) - 1)),
		period:    period,
		prev:      math.NaN(),
		priceType: priceType,
		smaArr:    []*sma{NewSMA(period, priceType), NewSMA(period, priceType)},
	}
}

//Calculate : method to Calculate trima and return results as float array
//Return	: trima result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *trima) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	SMAResult := ins.smaArr[0].Calculate(newData)[0]
	sma1 := utils.NewOHLCV(ins.priceType, SMAResult)
	if !math.IsNaN(SMAResult) {
		ins.prev = ins.smaArr[1].Calculate(sma1)[0]
	} else {
		return []float64{math.NaN()}
	}

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
