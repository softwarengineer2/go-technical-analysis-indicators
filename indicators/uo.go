package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Ultimate Oscillator
//Source : https://www.tradingview.com/wiki/Ultimate_Oscillator_(UO)
type uo struct {
	short        uint
	intermediate uint
	long         uint
	period       uint
	prev         float64
	priceType    string
	buf          *utils.OHLCVBuffer
}

//NewUO 		: To return uo struct instance
//Params
//short			: short period
//intermediate	: intermediate period
//long			: long period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewUO(short uint, intermediate uint, long uint, priceType string) *uo {
	return &uo{
		period:       long,
		short:        short,
		intermediate: intermediate,
		long:         long,
		prev:         math.NaN(),
		priceType:    priceType,
		buf:          utils.NewOHLCVBuffer(long + 1),
	}
}

//Calculate : method to Calculate uo and return results as float array
//Return	: uo result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *uo) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.buf.Add(newData)

	bpSumShort := 0.0
	trSumShort := 0.0
	j := 0
	for i := ins.buf.Capacity - 1; i > ins.buf.Capacity-ins.short-1; i-- {
		bpSumShort += ins.buf.Vals[i].Close - math.Min(ins.buf.Vals[i].Low, ins.buf.Vals[i-1].Close)
		trSumShort += math.Max(ins.buf.Vals[i].High, ins.buf.Vals[i-1].Close) - math.Min(ins.buf.Vals[i].Low, ins.buf.Vals[i-1].Close)
		j++
	}
	avgShort := (bpSumShort / trSumShort)

	bpSumIntermediate := 0.0
	trSumIntermediate := 0.0
	for i := ins.buf.Capacity - 1; i > ins.buf.Capacity-ins.intermediate-1; i-- {
		bpSumIntermediate += ins.buf.Vals[i].Close - math.Min(ins.buf.Vals[i].Low, ins.buf.Vals[i-1].Close)
		trSumIntermediate += math.Max(ins.buf.Vals[i].High, ins.buf.Vals[i-1].Close) - math.Min(ins.buf.Vals[i].Low, ins.buf.Vals[i-1].Close)
	}
	avgIntermediate := (bpSumIntermediate / trSumIntermediate)

	bpSumLong := 0.0
	trSumLong := 0.0
	for i := ins.buf.Capacity - 1; i > ins.buf.Capacity-ins.long-1; i-- {
		bpSumLong += ins.buf.Vals[i].Close - math.Min(ins.buf.Vals[i].Low, ins.buf.Vals[i-1].Close)
		trSumLong += math.Max(ins.buf.Vals[i].High, ins.buf.Vals[i-1].Close) - math.Min(ins.buf.Vals[i].Low, ins.buf.Vals[i-1].Close)
	}
	avgLong := (bpSumLong / trSumLong)

	ins.prev = 100.0 * ((4 * avgShort) + (2 * avgIntermediate) + avgLong) / (7.0)

	if ins.buf.Pushes < ins.buf.Size {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
