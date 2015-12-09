// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/petergtz/pegomock/test_interface (interfaces: Display)

package pegomock_test

import (
	pegomock "github.com/petergtz/pegomock/pegomock"
)

// Mock of Display interface
type MockDisplay struct {
	//http://dave.cheney.net/2014/03/25/the-empty-struct
	fieldToMakeSureEveryInstanceHasItsOwnIdentity int
}

func NewMockDisplay() *MockDisplay {
	return &MockDisplay{}
}

func (mock *MockDisplay) Flash(_param0 string, _param1 int) {
	pegomock.LastInvocation = &pegomock.Invocation{Mock: mock, MethodName: "Flash", Params: [pegomock.MaxNumParams]interface{}{_param0, _param1}}
	pegomock.Invocations[*pegomock.LastInvocation]++
}

func (mock *MockDisplay) MultipleValues() (string, int, float32) {
	pegomock.LastInvocation = &pegomock.Invocation{Mock: mock, MethodName: "MultipleValues", Params: [pegomock.MaxNumParams]interface{}{}}
	pegomock.Invocations[*pegomock.LastInvocation]++
	if len(pegomock.Stubbings[*pegomock.LastInvocation]) == 0 {
		var ret0 string
		var ret1 int
		var ret2 float32
		return ret0, ret1, ret2
	}
	result := pegomock.Stubbings[*pegomock.LastInvocation][pegomock.StubbingPointer[*pegomock.LastInvocation]]
	if pegomock.StubbingPointer[*pegomock.LastInvocation] < len(pegomock.Stubbings[*pegomock.LastInvocation])-1 {
		pegomock.StubbingPointer[*pegomock.LastInvocation]++
	}
	return result[0].(string), result[1].(int), result[2].(float32)
}

func (mock *MockDisplay) Show(_param0 string) {
	pegomock.LastInvocation = &pegomock.Invocation{Mock: mock, MethodName: "Show", Params: [pegomock.MaxNumParams]interface{}{_param0}}
	pegomock.Invocations[*pegomock.LastInvocation]++
}

func (mock *MockDisplay) SomeValue() string {
	pegomock.LastInvocation = &pegomock.Invocation{Mock: mock, MethodName: "SomeValue", Params: [pegomock.MaxNumParams]interface{}{}}
	pegomock.Invocations[*pegomock.LastInvocation]++
	if len(pegomock.Stubbings[*pegomock.LastInvocation]) == 0 {
		var ret0 string
		return ret0
	}
	result := pegomock.Stubbings[*pegomock.LastInvocation][pegomock.StubbingPointer[*pegomock.LastInvocation]]
	if pegomock.StubbingPointer[*pegomock.LastInvocation] < len(pegomock.Stubbings[*pegomock.LastInvocation])-1 {
		pegomock.StubbingPointer[*pegomock.LastInvocation]++
	}
	return result[0].(string)
}

type VerifierDisplay struct {
	mock *MockDisplay
}

func (mock *MockDisplay) VerifyWasCalled() *VerifierDisplay {
	return &VerifierDisplay{mock}
}
func (verifier *VerifierDisplay) Flash(_param0 string, _param1 int) {
	if pegomock.Invocations[pegomock.Invocation{verifier.mock, "Flash", [pegomock.MaxNumParams]interface{}{_param0, _param1}}] == 0 {
		panic("Mock not called")
	}
}

func (verifier *VerifierDisplay) MultipleValues() (string, int, float32) {
	if pegomock.Invocations[pegomock.Invocation{verifier.mock, "MultipleValues", [pegomock.MaxNumParams]interface{}{}}] == 0 {
		panic("Mock not called")
	}
	var ret0 string
	var ret1 int
	var ret2 float32
	return ret0, ret1, ret2
}

func (verifier *VerifierDisplay) Show(_param0 string) {
	if pegomock.Invocations[pegomock.Invocation{verifier.mock, "Show", [pegomock.MaxNumParams]interface{}{_param0}}] == 0 {
		panic("Mock not called")
	}
}

func (verifier *VerifierDisplay) SomeValue() string {
	if pegomock.Invocations[pegomock.Invocation{verifier.mock, "SomeValue", [pegomock.MaxNumParams]interface{}{}}] == 0 {
		panic("Mock not called")
	}
	var ret0 string
	return ret0
}