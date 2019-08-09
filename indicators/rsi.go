package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Relative Strength Index
//Source : https://www.tradingview.com/wiki/Relative_Strength_Index_(RSI)
type rsi struct {
	count     uint
	prev      float64
	prevPrice float64
	period    uint
	priceType string
	up        *wilders
	down      *wilders
}

//NewRSI 	: To return rsi struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewRSI(period uint, priceType string) *rsi {
	return &rsi{
		prev:      math.NaN(),
		prevPrice: math.NaN(),
		period:    period,
		priceType: priceType,
		up:        NewWILDERS(period, priceType),
		down:      NewWILDERS(period, priceType),
	}
}

//Calculate : method to Calculate rsi and return results as float array
//Return	: rsi result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *rsi) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	if math.IsNaN(ins.prevPrice) {
		ins.prevPrice = newData.GetByType(ins.priceType)
		return []float64{math.NaN()}
	}

	upResult := math.Max(newData.GetByType(ins.priceType)-ins.prevPrice, 0)
	downResult := math.Max(ins.prevPrice-newData.GetByType(ins.priceType), 0)

	up := utils.NewOHLCV(ins.priceType, upResult)
	down := utils.NewOHLCV(ins.priceType, downResult)

	ins.prevPrice = newData.GetByType(ins.priceType)
	var sup, sdown float64 = ins.up.Calculate(up)[0], ins.down.Calculate(down)[0]
	ins.prev = 100.0 - (100.0 / (1.0 + (sup / sdown)))
	return []float64{ins.prev}
}
