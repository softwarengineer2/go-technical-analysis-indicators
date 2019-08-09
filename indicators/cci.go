package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Commodity Channel Index
//Source : https://www.tradingview.com/wiki/Commodity_Channel_Index_(CCI)
type cci struct {
	period    uint
	prev      float64
	priceType string
	buf       *utils.OHLCVBuffer
}

//NewCCI 	: To return cci struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewCCI(period uint, priceType string) *cci {
	return &cci{
		period:    period,
		priceType: priceType,
		prev:      math.NaN(),
		buf:       utils.NewOHLCVBuffer(period),
	}
}

//Calculate : method to Calculate aroon and return results as float array
//Return	: CCI Result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *cci) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	if ins.buf.Pushes < ins.buf.Size {
		return []float64{math.NaN()}
	}

	avg := ins.buf.Total().GetByType(ins.priceType) * (1. / float64(ins.period))

	acc := 0.
	for _, oneValue := range ins.buf.Vals {
		acc += math.Abs(avg - oneValue.GetByType(ins.priceType))
	}
	acc *= (1. / float64(ins.period))

	ins.prev = (newPrice - avg) / acc / 0.015

	return []float64{ins.prev}
}
