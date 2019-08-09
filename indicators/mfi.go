package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Money Flow Index
//Source : https://www.tradingview.com/wiki/Money_Flow_(MFI)
type mfi struct {
	period     uint
	prev       float64
	priceType  string
	buf        *utils.OHLCVBuffer
	typicalBuf *utils.OHLCVBuffer
}

//NewMFI 	: To return mfi struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewMFI(period uint, priceType string) *mfi {
	return &mfi{
		period:     period,
		prev:       math.NaN(),
		priceType:  priceType,
		buf:        utils.NewOHLCVBuffer(period + 1),
		typicalBuf: utils.NewOHLCVBuffer(period + 1),
	}
}

//Calculate : method to Calculate mfi and return results as float array
//Return	: mf index result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *mfi) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	typicalPriceVal := (newData.High + newData.Low + newData.Close) / 3.0
	typicalData := utils.NewOHLCV(ins.priceType, typicalPriceVal)

	ins.typicalBuf.Add(typicalData)

	if ins.typicalBuf.Pushes < ins.typicalBuf.Size-1 {
		return []float64{math.NaN()}
	}

	pos_period := 0.0
	neg_period := 0.0

	for i := int(ins.typicalBuf.Capacity) - 1; i > 0; i-- {
		typicalPrice := ins.typicalBuf.Vals[i].GetByType(ins.priceType)
		preTypicalPrice := ins.typicalBuf.Vals[i-1].GetByType(ins.priceType)

		raw := typicalPrice * ins.buf.Vals[i].Volume

		pos_day := 0.0
		neg_day := 0.0
		if typicalPrice >= preTypicalPrice {
			pos_day = raw
		}
		if typicalPrice <= preTypicalPrice {
			neg_day = raw
		}

		pos_period += pos_day
		neg_period += neg_day

	}
	mf_ratio := (pos_period / neg_period)
	ins.prev = 100 - (100 / (1 + mf_ratio)) //Mf_index

	return []float64{ins.prev}
}
