package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Moving Average Convergence/Divergence
//Source : https://www.tradingview.com/wiki/MACD_(Moving_Average_Convergence/Divergence)
type macd struct {
	count     uint
	period    uint
	prev      float64
	priceType string
	short     *ema
	long      *ema
	signal    *sma
}

//NewMACD 	: To return macd struct instance
//Params
//short		: short EMA period
//long		: long EMA period
//signal	: signal EMA period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewMACD(short, long, signal uint, priceType string) *macd {
	return &macd{
		period:    long,
		prev:      math.NaN(),
		priceType: priceType,
		short:     NewEMA(short, priceType),
		long:      NewEMA(long, priceType),
		signal:    NewSMA(signal, priceType),
	}
}

//Calculate : method to Calculate macd and return results as float array
//Return	: macd result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *macd) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	resultFast := ins.short.Calculate(newData)[0]
	resultSlow := ins.long.Calculate(newData)[0]
	EMAResult := resultFast - resultSlow

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}
	data := utils.NewOHLCV(ins.priceType, EMAResult)

	ins.prev = ins.signal.Calculate(data)[0]

	return []float64{ins.prev}
}
