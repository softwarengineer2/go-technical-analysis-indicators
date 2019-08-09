package utils

import "strings"

//OHLCV is a data structure that holds Open, High, Low, Close and Volume values
type OHLCV struct {
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

//NewOHLCV : To return new OHLCV instance
//Params
//priceType : set specific price for initialization
//price : for initialization of price
func NewOHLCV(priceType string, price float64) OHLCV {
	oneData := OHLCV{}
	oneData.SetByType(priceType, price)

	return oneData
}

//GetByType : To return specific price type
func (ins OHLCV) GetByType(priceType string) float64 {

	result := ins.Close
	if strings.ToLower(priceType) == "high" {
		result = ins.High
	} else if strings.ToLower(priceType) == "low" {
		result = ins.Low
	} else if strings.ToLower(priceType) == "open" {
		result = ins.Open
	}

	return result
}

//SetByType : To set specific price type
func (ins *OHLCV) SetByType(priceType string, price float64) {
	if strings.ToLower(priceType) == "high" {
		ins.High = price
	} else if strings.ToLower(priceType) == "low" {
		ins.Low = price
	} else if strings.ToLower(priceType) == "open" {
		ins.Open = price
	} else {
		ins.Close = price
	}
}

//HL2 : To return (ins.High + ins.Low) / 2.0
func (ins *OHLCV) HL2() float64 {
	return (ins.High + ins.Low) / 2.0
}

//HLC3 : To return (ins.High + ins.Low + ins.Close) / 3.0
func (ins *OHLCV) HLC3() float64 {
	return (ins.High + ins.Low + ins.Close) / 3.0
}
