package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Momentum
//Source : https://www.thebalance.com/how-to-trade-with-the-momentum-indicator-1031195
type mom struct {
	period    uint
	prev      float64
	priceType string
	buf       *utils.OHLCVBuffer
}

//NewMOM 	: To return mom struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewMOM(period uint, priceType string) *mom {
	return &mom{
		period:    period + 1,
		prev:      math.NaN(),
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(period + 1),
	}
}

//Calculate : method to Calculate mom and return results as float array
//Return	: Momentum result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *mom) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	if ins.buf.Pushes < ins.buf.Size {
		return []float64{math.NaN()}
	}

	ins.prev = (newPrice - ins.buf.Vals[0].GetByType(ins.priceType))

	return []float64{ins.prev}
}
