package indicators

import (
	"gitlab.com/softwarengineer2/gobot/utils"
)

//Indicator : Interface for all indicators
type Indicator interface {
	Calculate(newData utils.OHLCV) []float64
}
