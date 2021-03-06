// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// Autogenerated by buildergenerator

package goexoml

import (
	"errors"
)

var _ = errors.New("_")

//SetAction sets Action for Gather struct instance
func (__gather__ *Gather) SetAction(action string) *Gather {
	__gather__.Action = action
	return __gather__
}

//SetMethod sets Method for Gather struct instance
func (__gather__ *Gather) SetMethod(method string) *Gather {
	__gather__.Method = method
	return __gather__
}

//SetTimeout sets Timeout for Gather struct instance
func (__gather__ *Gather) SetTimeout(timeout int) *Gather {
	__gather__.Timeout = timeout
	return __gather__
}

//SetFinishOnKey sets FinishOnKey for Gather struct instance
func (__gather__ *Gather) SetFinishOnKey(finishonkey string) *Gather {
	__gather__.FinishOnKey = finishonkey
	return __gather__
}

//SetNumDigits sets NumDigits for Gather struct instance
func (__gather__ *Gather) SetNumDigits(numdigits int) *Gather {
	__gather__.NumDigits = numdigits
	return __gather__
}

//SetSay sets Say for Gather struct instance
func (__gather__ *Gather) SetSay(say *Say) *Gather {
	__gather__.Say = say
	return __gather__
}

//SetPlay sets Play for Gather struct instance
func (__gather__ *Gather) SetPlay(play *Play) *Gather {
	__gather__.Play = play
	return __gather__
}

//Setter returns setter function for the field given
func (__gather__ *Gather) Setter(field string) (setter func(interface{}) (*Gather, error)) {
	switch field {
	case "Action":
		setter = func(ActionField interface{}) (*Gather, error) {
			if ActionValue, ok := ActionField.(string); ok {
				return __gather__.SetAction(ActionValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	case "Method":
		setter = func(MethodField interface{}) (*Gather, error) {
			if MethodValue, ok := MethodField.(string); ok {
				return __gather__.SetMethod(MethodValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	case "Timeout":
		setter = func(TimeoutField interface{}) (*Gather, error) {
			if TimeoutValue, ok := TimeoutField.(int); ok {
				return __gather__.SetTimeout(TimeoutValue), nil
			}
			return nil, errors.New("Invalid type Expected int ")
		}
	case "FinishOnKey":
		setter = func(FinishOnKeyField interface{}) (*Gather, error) {
			if FinishOnKeyValue, ok := FinishOnKeyField.(string); ok {
				return __gather__.SetFinishOnKey(FinishOnKeyValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	case "NumDigits":
		setter = func(NumDigitsField interface{}) (*Gather, error) {
			if NumDigitsValue, ok := NumDigitsField.(int); ok {
				return __gather__.SetNumDigits(NumDigitsValue), nil
			}
			return nil, errors.New("Invalid type Expected int ")
		}
	case "Say":
		setter = func(SayField interface{}) (*Gather, error) {
			if SayValue, ok := SayField.(*Say); ok {
				return __gather__.SetSay(SayValue), nil
			}
			return nil, errors.New("Invalid type Expected *Say ")
		}
	case "Play":
		setter = func(PlayField interface{}) (*Gather, error) {
			if PlayValue, ok := PlayField.(*Play); ok {
				return __gather__.SetPlay(PlayValue), nil
			}
			return nil, errors.New("Invalid type Expected *Play ")
		}
	}
	return
}

//NewGather return a new Gather pointer
func NewGather() *Gather {
	return new(Gather)
}

//IGather The interface that satisfies all the methods for this struct
//IGather asserts implementation of setters for all the fields of Gather
type IGather interface {
	SetAction(action string) *Gather
	SetMethod(method string) *Gather
	SetTimeout(timeout int) *Gather
	SetFinishOnKey(finishonkey string) *Gather
	SetNumDigits(numdigits int) *Gather
	SetSay(say *Say) *Gather
	SetPlay(play *Play) *Gather
	Setter(string) func(interface{}) (*Gather, error)
}

//AddGather appends the verb to response
func (r *Response) AddGather(gather IGather) *Response {
	r.Response = append(r.Response, gather)
	return r
}
