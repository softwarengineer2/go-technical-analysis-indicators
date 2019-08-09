package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Aroon
//Source : https://www.tradingview.com/wiki/Aroon
type aroon struct {
	period    uint
	prev      float64
	priceType string
	buf       *utils.OHLCVBuffer
}

//NewAROON 	: To return aroon struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewAROON(period uint, priceType string) *aroon {
	return &aroon{
		period:    period,
		prev:      math.NaN(),
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(period + 1),
	}
}

//Calculate : method to Calculate aroon and return results as float array
//Return	: Aroon High and Low (2 values in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *aroon) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	if ins.buf.Pushes < ins.buf.Size {
		return []float64{math.NaN()}
	}

	aroonHigh := ((float64(ins.period) - (float64(ins.period) - ins.buf.MaxIndex().High)) / float64(ins.period)) * 100.0
	aroonLow := ((float64(ins.period) - (float64(ins.period) - ins.buf.MinIndex().Low)) / float64(ins.period)) * 100.0

	ins.prev = aroonHigh - aroonLow

	return []float64{ins.prev}
}
