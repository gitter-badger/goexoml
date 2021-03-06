// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// Autogenerated by buildergenerator

package goexoml

import (
	"errors"
)

var _ = errors.New("_")

//SetVoice sets Voice for Say struct instance
func (__say__ *Say) SetVoice(voice string) *Say {
	__say__.Voice = voice
	return __say__
}

//SetLanguage sets Language for Say struct instance
func (__say__ *Say) SetLanguage(language string) *Say {
	__say__.Language = language
	return __say__
}

//SetLoop sets Loop for Say struct instance
func (__say__ *Say) SetLoop(loop int) *Say {
	__say__.Loop = loop
	return __say__
}

//SetText sets Text for Say struct instance
func (__say__ *Say) SetText(text string) *Say {
	__say__.Text = text
	return __say__
}

//Setter returns setter function for the field given
func (__say__ *Say) Setter(field string) (setter func(interface{}) (*Say, error)) {
	switch field {
	case "Voice":
		setter = func(VoiceField interface{}) (*Say, error) {
			if VoiceValue, ok := VoiceField.(string); ok {
				return __say__.SetVoice(VoiceValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	case "Language":
		setter = func(LanguageField interface{}) (*Say, error) {
			if LanguageValue, ok := LanguageField.(string); ok {
				return __say__.SetLanguage(LanguageValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	case "Loop":
		setter = func(LoopField interface{}) (*Say, error) {
			if LoopValue, ok := LoopField.(int); ok {
				return __say__.SetLoop(LoopValue), nil
			}
			return nil, errors.New("Invalid type Expected int ")
		}
	case "Text":
		setter = func(TextField interface{}) (*Say, error) {
			if TextValue, ok := TextField.(string); ok {
				return __say__.SetText(TextValue), nil
			}
			return nil, errors.New("Invalid type Expected string ")
		}
	}
	return
}

//NewSay return a new Say pointer
func NewSay() *Say {
	return new(Say)
}

//ISay The interface that satisfies all the methods for this struct
//ISay asserts implementation of setters for all the fields of Say
type ISay interface {
	SetVoice(voice string) *Say
	SetLanguage(language string) *Say
	SetLoop(loop int) *Say
	SetText(text string) *Say
	Setter(string) func(interface{}) (*Say, error)
}

//AddSay appends the verb to response
func (r *Response) AddSay(say ISay) *Response {
	r.Response = append(r.Response, say)
	return r
}
