// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// Autogenerated by buildergenerator

package goexoml

import (
	"errors"
)

var _ = errors.New("_")

//SetAction sets Action for Record struct instance
func (__record__ *Record) SetAction(action string) *Record {
	__record__.Action = action
	return __record__
}

//SetMethod sets Method for Record struct instance
func (__record__ *Record) SetMethod(method string) *Record {
	__record__.Method = method
	return __record__
}

//SetTimeout sets Timeout for Record struct instance
func (__record__ *Record) SetTimeout(timeout int) *Record {
	__record__.Timeout = timeout
	return __record__
}

//SetFinishOnKey sets FinishOnKey for Record struct instance
func (__record__ *Record) SetFinishOnKey(finishonkey string) *Record {
	__record__.FinishOnKey = finishonkey
	return __record__
}

//SetMaxLength sets MaxLength for Record struct instance
func (__record__ *Record) SetMaxLength(maxlength int) *Record {
	__record__.MaxLength = maxlength
	return __record__
}

//SetTranscribe sets Transcribe for Record struct instance
func (__record__ *Record) SetTranscribe(transcribe bool) *Record {
	__record__.Transcribe = transcribe
	return __record__
}

//SetTranscribeCallback sets TranscribeCallback for Record struct instance
func (__record__ *Record) SetTranscribeCallback(transcribecallback string) *Record {
	__record__.TranscribeCallback = transcribecallback
	return __record__
}

//SetPlayBeep sets PlayBeep for Record struct instance
func (__record__ *Record) SetPlayBeep(playbeep bool) *Record {
	__record__.PlayBeep = playbeep
	return __record__
}

//SetTrim sets Trim for Record struct instance
func (__record__ *Record) SetTrim(trim string) *Record {
	__record__.Trim = trim
	return __record__
}

//Setter returns setter function for the field given
func (__record__ *Record) Setter(field string) (setter func(interface{}) (*Record, error)) {
	switch field {
	case "Action":
		setter = func(ActionField interface{}) (*Record, error) {
			if ActionValue, ok := ActionField.(string); ok {
				return __record__.SetAction(ActionValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	case "Method":
		setter = func(MethodField interface{}) (*Record, error) {
			if MethodValue, ok := MethodField.(string); ok {
				return __record__.SetMethod(MethodValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	case "Timeout":
		setter = func(TimeoutField interface{}) (*Record, error) {
			if TimeoutValue, ok := TimeoutField.(int); ok {
				return __record__.SetTimeout(TimeoutValue), nil
			}
			return nil, errors.New("Invalid type Expected int ")
		}
	case "FinishOnKey":
		setter = func(FinishOnKeyField interface{}) (*Record, error) {
			if FinishOnKeyValue, ok := FinishOnKeyField.(string); ok {
				return __record__.SetFinishOnKey(FinishOnKeyValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	case "MaxLength":
		setter = func(MaxLengthField interface{}) (*Record, error) {
			if MaxLengthValue, ok := MaxLengthField.(int); ok {
				return __record__.SetMaxLength(MaxLengthValue), nil
			}
			return nil, errors.New("Invalid type Expected int ")
		}
	case "Transcribe":
		setter = func(TranscribeField interface{}) (*Record, error) {
			if TranscribeValue, ok := TranscribeField.(bool); ok {
				return __record__.SetTranscribe(TranscribeValue), nil
			}
			return nil, errors.New("Invalid type Expected bool ")
		}
	case "TranscribeCallback":
		setter = func(TranscribeCallbackField interface{}) (*Record, error) {
			if TranscribeCallbackValue, ok := TranscribeCallbackField.(string); ok {
				return __record__.SetTranscribeCallback(TranscribeCallbackValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	case "PlayBeep":
		setter = func(PlayBeepField interface{}) (*Record, error) {
			if PlayBeepValue, ok := PlayBeepField.(bool); ok {
				return __record__.SetPlayBeep(PlayBeepValue), nil
			}
			return nil, errors.New("Invalid type Expected bool ")
		}
	case "Trim":
		setter = func(TrimField interface{}) (*Record, error) {
			if TrimValue, ok := TrimField.(string); ok {
				return __record__.SetTrim(TrimValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	}
	return
}

//NewRecord return a new Record pointer
func NewRecord() *Record {
	return new(Record)
}

//IRecord The interface that satisfies all the methods for this struct
//IRecord asserts implementation of setters for all the fields of Record
type IRecord interface {
	SetAction(action string) *Record
	SetMethod(method string) *Record
	SetTimeout(timeout int) *Record
	SetFinishOnKey(finishonkey string) *Record
	SetMaxLength(maxlength int) *Record
	SetTranscribe(transcribe bool) *Record
	SetTranscribeCallback(transcribecallback string) *Record
	SetPlayBeep(playbeep bool) *Record
	SetTrim(trim string) *Record
	Setter(string) func(interface{}) (*Record, error)
}

//AddRecord appends the verb to response
func (r *Response) AddRecord(record IRecord) *Response {
	r.Response = append(r.Response, record)
	return r
}
