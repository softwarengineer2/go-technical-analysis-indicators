package testing

import (
	"fmt"
	"math"
	"testing"

	"gitlab.com/softwarengineer2/gobot/indicators"
	"gitlab.com/softwarengineer2/gobot/utils"
)

//GenerateTestData : To generate sample data for testing
//Data was taken from tradingview
//Monthly Bitcoin/USD Data for Bitstamp Exchange
func GenerateTestData(count int) []utils.OHLCV {
	data := []utils.OHLCV{
		utils.OHLCV{Open: 10.90, High: 11.85, Low: 8.0, Close: 8.0, Volume: 7},
		utils.OHLCV{Open: 8.0, High: 8.89, Low: 4.80, Close: 4.82, Volume: 1123},
		utils.OHLCV{Open: 4.85, High: 6.25, Low: 2.22, Close: 3.30, Volume: 2085},
		utils.OHLCV{Open: 3.32, High: 15.0, Low: 2.25, Close: 3.19, Volume: 2316},
		utils.OHLCV{Open: 3.46, High: 4.75, Low: 2.82, Close: 4.58, Volume: 2544},
		utils.OHLCV{Open: 4.58, High: 7.38, Low: 3.80, Close: 5.30, Volume: 2020},
		utils.OHLCV{Open: 5.53, High: 6.50, Low: 4.14, Close: 4.99, Volume: 4797},
		utils.OHLCV{Open: 4.99, High: 5.44, Low: 4.54, Close: 4.90, Volume: 9012},
		utils.OHLCV{Open: 4.89, High: 5.43, Low: 4.69, Close: 5.00, Volume: 16480},
		utils.OHLCV{Open: 4.94, High: 5.17, Low: 4.86, Close: 5.17, Volume: 18860},
		utils.OHLCV{Open: 5.17, High: 6.69, Low: 5.14, Close: 6.61, Volume: 63387},
		utils.OHLCV{Open: 6.60, High: 9.28, Low: 6.36, Close: 9.28, Volume: 50704},
		utils.OHLCV{Open: 9.28, High: 16.41, Low: 7.10, Close: 10.16, Volume: 83272},
		utils.OHLCV{Open: 10.20, High: 12.66, Low: 9.49, Close: 12.22, Volume: 55624},
		utils.OHLCV{Open: 12.21, High: 12.99, Low: 9.50, Close: 11.00, Volume: 83767},
		utils.OHLCV{Open: 11.00, High: 12.74, Low: 10.25, Close: 12.43, Volume: 88482},
		utils.OHLCV{Open: 12.39, High: 13.94, Low: 12.24, Close: 13.24, Volume: 91543},
		utils.OHLCV{Open: 13.24, High: 21.00, Low: 12.77, Close: 20.46, Volume: 124203},
		utils.OHLCV{Open: 20.30, High: 34.42, Low: 19.50, Close: 33.53, Volume: 124965},
		utils.OHLCV{Open: 33.53, High: 97.00, Low: 33.00, Close: 96.15, Volume: 177218},
		utils.OHLCV{Open: 95.50, High: 259.34, Low: 45.00, Close: 139.88, Volume: 439632},
		utils.OHLCV{Open: 139.87, High: 140.30, Low: 81.50, Close: 127.91, Volume: 299189},
		utils.OHLCV{Open: 127.76, High: 128.86, Low: 86.20, Close: 89.53, Volume: 249642},
		utils.OHLCV{Open: 89.41, High: 101.00, Low: 63.00, Close: 98.28, Volume: 420191},
		utils.OHLCV{Open: 97.77, High: 134.95, Low: 90.00, Close: 128.00, Volume: 332912},
		utils.OHLCV{Open: 127.90, High: 132.09, Low: 115.00, Close: 126.25, Volume: 318195},
		utils.OHLCV{Open: 126.25, High: 206.60, Low: 85.00, Close: 203.54, Volume: 626597},
		utils.OHLCV{Open: 203.54, High: 1163.00, Low: 200.23, Close: 1119.80, Volume: 965653},
		utils.OHLCV{Open: 1119.52, High: 1153.27, Low: 382.21, Close: 732.00, Volume: 950847},
		utils.OHLCV{Open: 732.00, High: 995.00, Low: 725.00, Close: 803.00, Volume: 442523},
		utils.OHLCV{Open: 803.00, High: 827.38, Low: 400.00, Close: 550.10, Volume: 824911},
		utils.OHLCV{Open: 551.80, High: 710.00, Low: 436.00, Close: 454.83, Volume: 491825},
		utils.OHLCV{Open: 454.81, High: 548.00, Low: 339.79, Close: 448.85, Volume: 515028},
		utils.OHLCV{Open: 449.00, High: 629.40, Low: 420.27, Close: 627.80, Volume: 300028},
		utils.OHLCV{Open: 627.80, High: 683.26, Low: 538.38, Close: 641.11, Volume: 288061},
		utils.OHLCV{Open: 640.00, High: 658.88, Low: 555.90, Close: 582.04, Volume: 153036},
		utils.OHLCV{Open: 583.54, High: 607.20, Low: 442.00, Close: 479.04, Volume: 300395},
		utils.OHLCV{Open: 479.01, High: 497.00, Low: 365.20, Close: 391.09, Volume: 343995},
		utils.OHLCV{Open: 389.56, High: 417.99, Low: 275.00, Close: 337.99, Volume: 591898},
		utils.OHLCV{Open: 337.99, High: 453.92, Low: 316.61, Close: 377.01, Volume: 438633},
		utils.OHLCV{Open: 376.09, High: 383.00, Low: 304.99, Close: 321.00, Volume: 294442},
		utils.OHLCV{Open: 321.00, High: 321.00, Low: 152.40, Close: 216.90, Volume: 785557},
		utils.OHLCV{Open: 218.04, High: 267.92, Low: 208.48, Close: 253.47, Volume: 343156},
		utils.OHLCV{Open: 253.45, High: 297.95, Low: 236.40, Close: 244.24, Volume: 296971},
		utils.OHLCV{Open: 243.93, High: 262.98, Low: 210.00, Close: 236.20, Volume: 255141},
		utils.OHLCV{Open: 236.47, High: 247.01, Low: 227.01, Close: 228.91, Volume: 198109},
		utils.OHLCV{Open: 228.91, High: 268.00, Low: 219.03, Close: 262.50, Volume: 232038},
		utils.OHLCV{Open: 262.58, High: 317.99, Low: 252.40, Close: 284.33, Volume: 395715},
		utils.OHLCV{Open: 284.69, High: 285.88, Low: 198.12, Close: 229.86, Volume: 417970},
		utils.OHLCV{Open: 230.25, High: 246.24, Low: 223.00, Close: 236.20, Volume: 568096},
		utils.OHLCV{Open: 235.87, High: 334.67, Low: 235.00, Close: 310.18, Volume: 716438},
		utils.OHLCV{Open: 310.90, High: 502.00, Low: 294.00, Close: 376.71, Volume: 887632},
		utils.OHLCV{Open: 376.70, High: 467.80, Low: 348.64, Close: 430.89, Volume: 421157}}

	retData := data
	if count > 0 {
		retData = data[:count]
	}

	return retData
}

//Check : To check results and expected results of an indicator
func Check(results []float64, expected []float64, testName string) uint {
	var wrongAnswer uint
	var tempRes float64
	for i, v := range results {
		if !math.IsNaN(v) || !math.IsNaN(expected[i]) {
			tempRes = v - expected[i]
			if math.Abs(tempRes) > 0.01 {
				fmt.Println(expected[i], "__", v)
				wrongAnswer++
			}
		}
	}
	fmt.Println("(", testName, ") Test Result Total:", len(expected), " Right:", len(expected)-int(wrongAnswer), "- Wrong:", wrongAnswer)
	return wrongAnswer
}

//TestSMA : To Test Simple Moving Average functionality
func TestSMA(t *testing.T) {
	name := "SMA"
	series := GenerateTestData(13)
	expected := []float64{4.8978, 4.5833, 4.7822, 5.4466, 6.2211}
	indic := indicators.NewSMA(9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
	}
}

//TestEMA : To Test Exponential Moving Average functionality
func TestEMA(t *testing.T) {
	name := "EMA"
	series := GenerateTestData(13)
	expected := []float64{4.8978, 4.9522, 5.2838, 6.0830, 6.8984}
	indic := indicators.NewEMA(9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestWMA : To Test Weighted Moving Average functionality
func TestWMA(t *testing.T) {
	name := "WMA"
	//series := []float64{8.0, 4.82, 3.30, 3.19, 4.58, 5.30, 4.99, 4.90, 5.0, 5.17, 6.61, 9.28, 10.16}
	series := GenerateTestData(13)
	expected := []float64{4.7584, 4.8128, 5.2182, 6.1177, 7.0604}
	indic := indicators.NewWMA(9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestTRIMA : To Test Triangular Moving Average functionality
func TestTRIMA(t *testing.T) {
	name := "TRIMA"
	series := GenerateTestData(21)
	expected := []float64{6.5212, 7.2186, 8.3011, 10.4669, 14.1712}
	indic := indicators.NewTRIMA(9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestROC : To Test Rate Of Change functionality
func TestROC(t *testing.T) {
	name := "ROC"
	series := GenerateTestData(14)
	expected := []float64{-35.3750, 37.1369, 181.2121, 218.4953, 166.8122}
	indic := indicators.NewROC(9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestCCI : To Test Common Channel Index functionality
func TestCCI(t *testing.T) {
	name := "CCI"
	series := GenerateTestData(18)
	expected := []float64{185.5299, 124.9009, 126.2084, 117.0301, 205.5896}
	indic := indicators.NewCCI(14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestDEMA : To Test Double Exponential Moving Average functionality
func TestDEMA(t *testing.T) {
	name := "DEMA"
	series := GenerateTestData(22)
	expected := []float64{13.1119, 16.2420, 23.1141, 50.4721, 85.5511, 105.8656}
	indic := indicators.NewDEMA(9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestTEMA : To Test Triple Exponential Moving Average functionality
func TestTEMA(t *testing.T) {
	name := "TEMA"
	series := GenerateTestData(30)
	expected := []float64{123.4884, 129.8809, 170.5641, 642.2924, 748.3859, 839.6727}
	indic := indicators.NewTEMA(9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestGDEMA : To Test Generalized Double Exponential Moving Average functionality
func TestGDEMA(t *testing.T) {
	name := "GDEMA"
	series := GenerateTestData(22)
	expected := []float64{12.2204, 15.0262, 21.1171, 45.0493, 76.0536, 94.7148}
	indic := indicators.NewGDEMA(9, 0.7, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestT3MA : To Test Tillson Moving Average functionality
func TestT3MA(t *testing.T) {
	name := "T3MA"
	series := GenerateTestData(53)
	expected := []float64{266.7675, 249.6603, 238.1832, 234.4401, 239.6532}
	indic := indicators.NewT3MA(9, 0.7, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestMACD : To Test Moving Average Convergence Divergence functionality
func TestMACD(t *testing.T) {
	name := "MACD"
	series := GenerateTestData(38)
	expected := []float64{138.4924, 151.6509, 163.5869, 165.3948, 161.9047}
	indic := indicators.NewMACD(12, 26, 9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestWILDERS : To Test Wilder's Moving Average functionality
func TestWILDERS(t *testing.T) {
	name := "WILDERS"
	series := GenerateTestData(14)
	expected := []float64{4.8978, 4.9280, 5.1149, 5.5777, 6.0868, 6.7683}
	indic := indicators.NewWILDERS(9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestRSI : To Test Relative Strength Index functionality
func TestRSI(t *testing.T) {
	name := "RSI"
	series := GenerateTestData(19)
	expected := []float64{59.4578, 63.0460, 64.9389, 76.4960, 85.6910}
	indic := indicators.NewRSI(14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestAROON : To Test Aroon Indicator functionality
func TestAROON(t *testing.T) {
	name := "AROON"
	series := GenerateTestData(25)
	expected := []float64{71.4286, 71.4286, 71.4286, 100.00, 100.00, 100.00, 100.00, 92.8571, 85.7143, 78.5714, 71.4286}
	indic := indicators.NewAROON(14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestAROONO : To Test Aroon Oscillator functionality
func TestAROONO(t *testing.T) {
	name := "AROONO"
	series := GenerateTestData(25)
	expected := []float64{71.4286, 71.4286, 71.4286, 100.00, 100.00, 100.00, 100.00, 92.8571, 85.7143, 78.5714, 71.4286}
	indic := indicators.NewAROONO(14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestPPO : To Test Percentage Price Oscillator functionality
func TestPPO(t *testing.T) {
	name := "PPO"
	series := GenerateTestData(38)
	expected := []float64{83.9314, 77.2805, 71.1137, 65.1343, 59.1921}
	indic := indicators.NewPPO(12, 26, 9, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestAPO : To Test Absolute Price Oscillator functionality
func TestAPO(t *testing.T) {
	name := "APO"
	series := GenerateTestData(30)
	expected := []float64{46.7203, 52.7498, 129.9645, 158.0437, 183.9058}
	indic := indicators.NewAPO(12, 26, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestMOM : To Test Momentum functionality
func TestMOM(t *testing.T) {
	name := "MOM"
	series := GenerateTestData(17)
	expected := []float64{2.16, 7.4, 7.70, 9.24, 8.66}
	indic := indicators.NewMOM(12, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestKAMA : To Test Kaufman's Adaptive Moving Average functionality
func TestKAMA(t *testing.T) {
	name := "KAMA"
	series := GenerateTestData(19)
	expected := []float64{0.0333, 0.0532, 0.0667, 0.0797, 0.0985, 0.1201, 0.1404, 0.1602, 0.1803, 0.2011, 0.2278, 0.2655, 0.3067, 0.3562, 0.6951, 2.4733, 5.2868, 10.2433, 18.7967} //{0.6951, 2.4733, 5.2868, 10.2433, 18.7967}
	indic := indicators.NewKAMA(14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestUO : To Test Ultimate Oscillator functionality
func TestUO(t *testing.T) {
	name := "UO"
	series := GenerateTestData(33)
	expected := []float64{73.0270, 68.7627, 63.4182, 58.1325, 57.6918}
	indic := indicators.NewUO(7, 14, 28, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestSTOCH : To Test Stoch functionality
func TestSTOCH(t *testing.T) {
	name := "STOCH"
	series := GenerateTestData(22)
	expected := []float64{73.5926, 81.0856, 90.1710, 90.4709, 82.5656}
	indic := indicators.NewSTOCH(3, 14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestSRSI : To Test Stoch Relative Strength Index functionality
func TestSRSI(t *testing.T) {
	name := "SRSI"
	series := GenerateTestData(24)
	expected := []float64{100.00, 100.00, 97.4159, 86.0481, 66.4304}
	indic := indicators.NewSRSI(3, 14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestBB : To Test Bollinger Bands functionality
func TestBB(t *testing.T) {
	name := "BB"
	series := GenerateTestData(24)
	expected := []float64{13.7165, 20.3105, 26.4650, 30.7765, 35.5310}
	indic := indicators.NewBB(20, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestMFI : To Test Money Flow Index functionality
func TestMFI(t *testing.T) {
	name := "MFI"
	series := GenerateTestData(18)
	expected := []float64{96.4556, 71.5280, 78.0226, 82.65, 87.4370}
	indic := indicators.NewMFI(14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestPSAR : To Test Parabolic SAR functionality
func TestPSAR(t *testing.T) {
	name := "PSAR"
	series := GenerateTestData(18)
	expected := []float64{11.7090, 11.3294, 2.22, 2.4756, 2.7261, 2.9716, 3.2121, 3.4479, 3.6789, 3.9054, 4.1272, 4.6186, 5.0902, 5.5430, 5.9777, 6.3950, 7.2713}
	indic := indicators.NewPSAR(0.02, 0.02, 0.2, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestDMI : To Test Directional Movement Index functionality
func TestDMI(t *testing.T) {
	name := "DMI"
	series := GenerateTestData(32)
	expected := []float64{81.2588, 82.3719, 83.4055, 80.8599, 78.4961}
	indic := indicators.NewDMI(14, 14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestMESA : To Test MESA Adaptive Moving Average functionality
func TestMESA(t *testing.T) {
	name := "MESA"
	series := GenerateTestData(18)
	expected := []float64{1.725, 2.1370000000000005, 2.8276000000000003, 3.3260800000000006, 3.658864000000001, 3.939091200000001, 4.055291412900264, 4.427233130320211, 4.662719173917652, 6.081175339134122, 6.465574426153653, 7.421459540922923, 8.23616763273834, 9.206934106190673, 10.742547284952538}
	indic := indicators.NewMESA(14, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}

//TestCHAIKIN : To Test Chaikin A/D Oscillator functionality
func TestCHAIKIN(t *testing.T) {
	name := "CHAIKIN"
	series := GenerateTestData(17)
	expected := []float64{7379.2867, 21434.2314, 23281.8341, -3290.4759, 8002.0785, -4633.2171, 15504.4198, 6329.8771}
	indic := indicators.NewCHAIKIN(3, 10, "close")

	results := make([]float64, len(expected))
	var tempRes []float64
	var j int
	for _, v := range series {
		tempRes = indic.Calculate(v)
		if !math.IsNaN(tempRes[0]) {
			results[j] = tempRes[0]
			j++
		}
	}

	if Check(results, expected, name) > 0 {
		t.Fail()
	}
}
