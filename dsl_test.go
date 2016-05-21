// Copyright 2015 Peter Goetz
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pegomock_test

import (
	"fmt"

	. "github.com/petergtz/pegomock"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/petergtz/pegomock/internal/types"
)

var _ = Describe("MockDisplay", func() {
	var display *MockDisplay

	BeforeEach(func() {
		display = NewMockDisplay()
	})

	Context("Calling SomeValue() with no stubbing", func() {
		It("returns zero value", func() {
			Expect(display.SomeValue()).To(Equal(""))
		})
	})

	Context("Stubbing MultipleParamsAndReturnValue() with matchers", func() {
		BeforeEach(func() {
			When(display.MultipleParamsAndReturnValue(EqString("Hello"), EqInt(333))).ThenReturn("Bla")
		})

		It("fails during verification when mock was not called", func() {
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).To(Panic())
		})

		It("succeeds verification when mock was called", func() {
			display.MultipleParamsAndReturnValue("Hello", 333)
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
		})

		It("succeeds verification when verification and invocation are mixed", func() {
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).To(Panic())
			display.MultipleParamsAndReturnValue("Hello", 333)
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
		})
	})

	Context("Calling MultipleParamsAndReturnValue() with \"Any\"-matchers", func() {
		It("succeeds all verifications that match", func() {
			When(display.MultipleParamsAndReturnValue(AnyString(), EqInt(333))).ThenReturn("Bla")

			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).To(Panic())

			display.MultipleParamsAndReturnValue("Hello", 333)
			display.MultipleParamsAndReturnValue("Hello again", 333)
			display.MultipleParamsAndReturnValue("And again", 333)

			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello", 333) }).NotTo(Panic())
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("Hello again", 333) }).NotTo(Panic())
			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("And again", 333) }).NotTo(Panic())

			Expect(func() { display.VerifyWasCalledOnce().MultipleParamsAndReturnValue("And again", 444) }).To(Panic())

		})
	})

	Context("Calling MultipleParamsAndReturnValue() only with matchers on some parameters", func() {
		It("panics", func() {
			Expect(func() { When(display.MultipleParamsAndReturnValue(EqString("Hello"), 333)) }).To(Panic())
		})
	})

	Context("Stubbing with consecutive return values", func() {
		BeforeEach(func() {
			When(display.SomeValue()).ThenReturn("Hello").ThenReturn("again")
		})

		It("returns stubbed values when calling mock", func() {
			Expect(display.SomeValue()).To(Equal("Hello"))
			Expect(display.SomeValue()).To(Equal("again"))
		})

		It("returns last stubbed value repeatedly", func() {
			Expect(display.SomeValue()).To(Equal("Hello"))
			Expect(display.SomeValue()).To(Equal("again"))
			Expect(display.SomeValue()).To(Equal("again"))
			Expect(display.SomeValue()).To(Equal("again"))
			Expect(display.SomeValue()).To(Equal("again"))
			Expect(display.SomeValue()).To(Equal("again"))
		})

		It("can be verified that mock was called", func() {
			display.SomeValue()
			Expect(func() { display.VerifyWasCalledOnce().SomeValue() }).NotTo(Panic())
		})

		It("fails if verify is called on mock that was not invoked.", func() {
			Expect(func() { display.VerifyWasCalledOnce().Show("param") }).To(Panic())
		})

		It("fails if verify is called on mock that was invoked more than once.", func() {
			display.Show("param")
			display.Show("param")
			Expect(func() { display.VerifyWasCalledOnce().Show("param") }).To(Panic())

		})
	})

	Context("Stubbing with invalid return type", func() {
		It("panics", func() {
			Expect(func() { When(display.SomeValue()).ThenReturn("Hello").ThenReturn(0) }).To(Panic())
		})
	})

	Context("Stubbed method, but no invocation takes place", func() {
		It("fails during verification", func() {
			When(display.SomeValue()).ThenReturn("Hello")
			Expect(func() { display.VerifyWasCalledOnce().SomeValue() }).To(Panic())
		})
	})

	Context("Calling Flash() with specific arguments", func() {

		BeforeEach(func() { display.Flash("Hello", 333) })

		It("succeeds verification if values are matching", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash("Hello", 333) }).NotTo(Panic())
		})

		It("fails during verification if values are not matching", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash("Hello", 666) }).To(Panic())
		})

		It("succeeds during verification when using Any-matchers ", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash(AnyString(), AnyInt()) }).NotTo(Panic())
		})

		It("succeeds during verification when using valid Eq-matchers ", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash(EqString("Hello"), EqInt(333)) }).NotTo(Panic())
		})

		It("fails during verification when using invalid Eq-matchers ", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash(EqString("Invalid"), EqInt(-1)) }).To(Panic())
		})
	})

	Context("Calling Flash() twice", func() {

		BeforeEach(func() {
			display.Flash("Hello", 333)
			display.Flash("Hello", 333)
		})

		It("succeeds verification if verifying with Times(2)", func() {
			Expect(func() { display.VerifyWasCalled(Times(2)).Flash("Hello", 333) }).NotTo(Panic())
		})

		It("fails during verification if verifying with VerifyWasCalledOnce", func() {
			Expect(func() { display.VerifyWasCalledOnce().Flash("Hello", 333) }).To(Panic())
		})

		It("fails during verification if verifying with Times(1)", func() {
			Expect(func() { display.VerifyWasCalled(Times(1)).Flash("Hello", 333) }).To(Panic())
		})

		It("succeeds during verification when using AtLeast(1)", func() {
			Expect(func() { display.VerifyWasCalled(AtLeast(1)).Flash("Hello", 333) }).NotTo(Panic())
		})

		It("succeeds during verification when using AtLeast(2)", func() {
			Expect(func() { display.VerifyWasCalled(AtLeast(2)).Flash("Hello", 333) }).NotTo(Panic())
		})

		It("fails during verification when using AtLeast(3)", func() {
			Expect(func() { display.VerifyWasCalled(AtLeast(3)).Flash("Hello", 333) }).To(Panic())
		})

		It("succeeds during verification when using Never()", func() {
			Expect(func() { display.VerifyWasCalled(Never()).Flash("Other value", 333) }).NotTo(Panic())
		})

		It("fails during verification when using Never()", func() {
			Expect(func() { display.VerifyWasCalled(Never()).Flash("Hello", 333) }).To(Panic())
		})

	})

	Context("Calling MultipleParamsAndReturnValue()", func() {

		It("panics when stubbed to panic", func() {
			When(display.MultipleParamsAndReturnValue(AnyString(), AnyInt())).
				ThenPanic("I'm panicking")
			Expect(func() {
				display.MultipleParamsAndReturnValue("Some string", 123)
			}).To(Panic())
		})

		It("calls back when stubbed to call back", func() {
			When(display.MultipleParamsAndReturnValue(AnyString(), AnyInt())).Then(
				func(params []Param) ReturnValues {
					return []ReturnValue{fmt.Sprintf("%v%v", params[0], params[1])}
				},
			)
			Expect(display.MultipleParamsAndReturnValue("string and ", 123)).
				To(Equal("string and 123"))
		})

	})

	Context("Making calls in a specific order", func() {

		BeforeEach(func() {
			display.Flash("Hello", 111)
			display.Flash("again", 222)
			display.Flash("and again", 333)
		})

		It("succeeds during InOrder verification when order is correct", func() {
			Expect(func() {
				inOrderContext := new(InOrderContext)
				display.VerifyWasCalledInOrder(Once(), inOrderContext).Flash("Hello", 111)
				display.VerifyWasCalledInOrder(Once(), inOrderContext).Flash("again", 222)
				display.VerifyWasCalledInOrder(Once(), inOrderContext).Flash("and again", 333)
			}).NotTo(Panic())
		})

		It("succeeds during InOrder verification when order is correct, but not all invocations are verified", func() {
			Expect(func() {
				inOrder := new(InOrderContext)
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("Hello", 111)
				// not checking for the 2nd call here
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("and again", 333)
			}).NotTo(Panic())
		})

		It("fails during InOrder verification when order is not correct", func() {
			Expect(func() {
				inOrder := new(InOrderContext)
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("again", 222)
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("Hello", 111)
				display.VerifyWasCalledInOrder(Once(), inOrder).Flash("and again", 333)
			}).To(Panic())
		})

	})

	Context("Capturing arguments", func() {
		It("Returns arguments when verifying with argument capture", func() {
			display.Flash("Hello", 111)

			arg1, arg2 := display.VerifyWasCalledOnce().Flash(AnyString(), AnyInt()).getCapturedArguments()

			Expect(arg1).To(Equal("Hello"))
			Expect(arg2).To(Equal(111))
		})

		It("Returns arguments of last invocation when verifying with argument capture", func() {
			display.Flash("Hello", 111)
			display.Flash("Again", 222)

			arg1, arg2 := display.VerifyWasCalled(AtLeast(1)).Flash(AnyString(), AnyInt()).getCapturedArguments()

			Expect(arg1).To(Equal("Again"))
			Expect(arg2).To(Equal(222))
		})

		It("Returns arguments of all invocations when verifying with \"all\" argument capture", func() {
			display.Flash("Hello", 111)
			display.Flash("Again", 222)

			args1, args2 := display.VerifyWasCalled(AtLeast(1)).Flash(AnyString(), AnyInt()).getAllCapturedArguments()

			Expect(args1).To(ConsistOf("Hello", "Again"))
			Expect(args2).To(ConsistOf(111, 222))
		})
	})

})
