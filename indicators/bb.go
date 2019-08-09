package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Bollinger Bands
//Source : https://www.tradingview.com/wiki/Bollinger_Bands
type bb struct {
	period    uint
	prev      float64
	mult      float64
	priceType string
	long      *sma
	buf       *utils.OHLCVBuffer
}

//NewBB 	: To return bb struct instance
//Params
//period	: calculation period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewBB(period uint, priceType string) *bb {
	return &bb{
		period:    period,
		prev:      math.NaN(),
		priceType: priceType,
		long:      NewSMA(period, priceType),
		buf:       utils.NewOHLCVBuffer(period),
	}
}

//Calculate : method to Calculate bb and return results as float array
//Return	: SMA result(basis), upper and lower bands (3 values in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *bb) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	basis := ins.long.Calculate(newData)[0]

	if ins.buf.Pushes < ins.buf.Size {
		return []float64{math.NaN()}
	}

	mean := ins.buf.Sum.GetByType(ins.priceType) / float64(ins.buf.Capacity)
	stdev := 0.0
	for i := 0; i < int(ins.buf.Capacity); i++ {
		stdev += math.Pow(ins.buf.Vals[i].GetByType(ins.priceType)-mean, 2)
	}
	stdev = math.Sqrt(stdev / float64(ins.buf.Capacity))

	dev := ins.mult * stdev
	upper := basis + dev
	lower := basis - dev

	ins.prev = basis

	return []float64{ins.prev, upper, lower}
}
