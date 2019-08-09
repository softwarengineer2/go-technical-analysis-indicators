package indicators

import (
	"math"

	"gitlab.com/softwarengineer2/gobot/utils"
)

//Percentage Price Oscillator
//Source : https://www.tradingview.com/wiki/Price_Oscillator_(PPO)
type ppo struct {
	count     uint
	period    uint
	prev      float64
	priceType string
	short     *ema
	long      *ema
	signal    *ema
}

//NewPPO 	: To return ppo struct instance
//Params
//short		: short EMA period
//long		: long EMA period
//signal	: signal EMA period
//priceType : To use it as base price for calculations from OHLCV Buffer (open, high, low, close)
func NewPPO(short, long, signal uint, priceType string) *ppo {
	return &ppo{
		period:    long,
		prev:      math.NaN(),
		priceType: priceType,
		short:     NewEMA(short, priceType),
		long:      NewEMA(long, priceType),
		signal:    NewEMA(signal, priceType),
	}
}

//Calculate : method to Calculate dema and return results as float array
//Return	: ppo result (1 value in return array)
//Params
//newData	: OHLCV instance to use its values for price calculation
func (ins *ppo) Calculate(newData utils.OHLCV) []float64 {
	newPrice := newData.GetByType(ins.priceType)

	if math.IsNaN(newPrice) {
		return []float64{ins.prev}
	}

	ins.count++

	resultFast := ins.short.Calculate(newData)[0]
	resultSlow := ins.long.Calculate(newData)[0]
	EMAResult := ((resultFast - resultSlow) / resultSlow) * 100.0

	if ins.count < ins.period {
		return []float64{math.NaN()}
	}

	data := utils.NewOHLCV(ins.priceType, EMAResult)

	ins.prev = ins.signal.Calculate(data)[0]

	if math.IsNaN(ins.prev) {
		return []float64{math.NaN()}
	}

	return []float64{ins.prev}
}
