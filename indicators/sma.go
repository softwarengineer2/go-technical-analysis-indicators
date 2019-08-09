package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Simple Moving Average
//Source : https://www.tradingview.com/wiki/Moving_Average
type sma struct {
	period    uint
	prev      float64
	priceType string
	buf       *utils.OHLCVBuffer
}

//NewSMA 	: To return sma struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewSMA(period uint, priceType string) *sma {
	return &sma{
		period:    period,
		prev:      math.NaN(),
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(period),
	}
}

//Calculate : method to Calculate sma and return results as float array
//Return	: sma result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *sma) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	ins.prev = ins.buf.Sum.GetByType(ins.priceType) * (1.0 / float64(ins.period))

	if ins.buf.Pushes < ins.buf.Size {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
