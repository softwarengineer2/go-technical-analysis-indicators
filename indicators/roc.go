package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Rate of Change
//Source : https://www.tradingview.com/wiki/Rate_of_Change_(ROC)
type roc struct {
	prev      float64
	period    uint
	priceType string
	buf       *utils.OHLCVBuffer
}

//NewROC 	: To return roc struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewROC(period uint, priceType string) *roc {
	return &roc{
		prev:      math.NaN(),
		period:    period,
		priceType: priceType,
		buf:       utils.NewOHLCVBuffer(period + 1),
	}
}

//Calculate : method to Calculate roc and return results as float array
//Return	: roc result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *roc) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	if ins.buf.Pushes < ins.buf.Size {
		return []float64{math.NaN()}
	}

	ins.prev = ((newPrice - ins.buf.Vals[0].GetByType(ins.priceType)) / ins.buf.Vals[0].GetByType(ins.priceType)) * 100.0

	return []float64{ins.prev}
}
