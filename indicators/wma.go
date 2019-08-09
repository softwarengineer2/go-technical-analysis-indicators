package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Weighted Moving Average
//Source : https://www.tradingview.com/wiki/Moving_Average
type wma struct {
	period    uint
	prev      float64
	priceType string
	buf       *utils.OHLCVBuffer
}

//NewWMA 	: To return wma struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewWMA(period uint, priceType string) *wma {
	return &wma{
		period:    period,
		prev:      math.NaN(),
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(period),
	}
}

//Calculate : method to Calculate wma and return results as float array
//Return	: wma result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *wma) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	if ins.buf.Pushes < ins.buf.Size {
		return []float64{math.NaN()}
	}

	var result float64
	for j := 0; j < int(ins.period); j++ {
		result += ins.buf.Vals[j].GetByType(ins.priceType) * float64(j+1)
	}
	ins.prev = result / (float64(ins.period) * (float64(ins.period) + 1.0) / 2.0)

	return []float64{ins.prev}
}
