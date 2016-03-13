package diagnostics_test

import (
	. "github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	. "github.com/microbusinesses/Micro-Businesses-Core/system"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var _ = Describe("Assert", func() {
	var (
		emptyString    string
		testString     string
		testWhitespace string
		testValue      string
		testValueName  string
		testMessage    string
	)

	BeforeEach(func() {
		emptyString = ""

		uuid, _ := RandomUUID()
		testString = uuid.String()

		testWhitespace = "      "

		uuid, _ = RandomUUID()
		testValue = uuid.String()

		uuid, _ = RandomUUID()
		testValueName = uuid.String()

		uuid, _ = RandomUUID()
		testMessage = uuid.String()
	})

	Describe("IsNil", func() {
		Context("when nil value provided", func() {
			It("Should return nil", func() {
				Expect(IsNil(nil, emptyString, emptyString)).To(BeNil())
			})
		})

		Context("when non-nill value provided", func() {
			Context("with value name and message both emptyd", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					IsNil(testValue, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					IsNil(testValue, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNil(testValue, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNil(testValue, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNilOrEmpty", func() {
		Context("when nil value provided", func() {
			It("Should return nil", func() {
				Expect(IsNilOrEmpty(nil, emptyString, emptyString)).To(BeNil())
			})
		})

		Context("when empty string value provided", func() {
			It("Should return empty string", func() {
				Expect(IsNilOrEmpty(emptyString, emptyString, emptyString)).To(Equal(emptyString))
			})
		})

		Context("when empty UUID value provided", func() {
			It("Should return empty UUID", func() {
				Expect(IsNilOrEmpty(EmptyUUID, emptyString, emptyString)).To(Equal(EmptyUUID))
			})
		})

		Context("when non-nill or empty value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					IsNilOrEmpty(testValue, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					IsNilOrEmpty(testValue, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNilOrEmpty(testValue, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNilOrEmpty(testValue, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNilOrEmptyOrWhitespace", func() {
		Context("when nil value provided", func() {
			It("Should return nil", func() {
				Expect(IsNilOrEmptyOrWhitespace(nil, emptyString, emptyString)).To(BeNil())
			})
		})

		Context("when empty string value provided", func() {
			It("Should return empty string", func() {
				Expect(IsNilOrEmptyOrWhitespace(emptyString, emptyString, emptyString)).To(Equal(emptyString))
			})
		})

		Context("when string contains whitespace only as value provided", func() {
			It("Should return provided value", func() {
				Expect(IsNilOrEmptyOrWhitespace(testWhitespace, emptyString, emptyString)).To(Equal(testWhitespace))
			})
		})

		Context("when empty UUID value provided", func() {
			It("Should return empty UUID", func() {
				Expect(IsNilOrEmptyOrWhitespace(EmptyUUID, emptyString, emptyString)).To(Equal(EmptyUUID))
			})
		})

		Context("when non-nill, empty or contains whitespace only provided as value", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					IsNilOrEmptyOrWhitespace(testValue, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					IsNilOrEmptyOrWhitespace(testValue, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNilOrEmptyOrWhitespace(testValue, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNilOrEmptyOrWhitespace(testValue, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNotNil", func() {
		Context("when non-nil value provided", func() {
			It("Should return provided value", func() {
				Expect(IsNotNil(testValue, emptyString, emptyString)).To(Equal(testValue))
			})
		})

		Context("when nill value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					IsNotNil(nil, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					IsNotNil(nil, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNil(nil, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNil(nil, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNotNilOrEmpty", func() {
		Context("when non-nil value provided", func() {
			It("Should return provided value", func() {
				Expect(IsNotNilOrEmpty(testValue, emptyString, emptyString)).To(Equal(testValue))
			})
		})

		Context("when nill value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					IsNotNilOrEmpty(nil, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					IsNotNilOrEmpty(nil, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmpty(nil, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmpty(nil, testValueName, testMessage)
				})
			})
		})

		Context("when empty string value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					IsNotNilOrEmpty(emptyString, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					IsNotNilOrEmpty(emptyString, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmpty(emptyString, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmpty(emptyString, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNotNilOrEmptyOrWhitespace", func() {
		Context("when non-nil value provided", func() {
			It("Should return provided value", func() {
				Expect(IsNotNilOrEmptyOrWhitespace(testValue, emptyString, emptyString)).To(Equal(testValue))
			})
		})

		Context("when nill value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					IsNotNilOrEmptyOrWhitespace(nil, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					IsNotNilOrEmptyOrWhitespace(nil, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmptyOrWhitespace(nil, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmptyOrWhitespace(nil, testValueName, testMessage)
				})
			})
		})

		Context("when empty string value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					IsNotNilOrEmptyOrWhitespace(emptyString, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					IsNotNilOrEmptyOrWhitespace(emptyString, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmptyOrWhitespace(emptyString, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmptyOrWhitespace(emptyString, testValueName, testMessage)
				})
			})
		})

		Context("when string contains whitespace only provided as value", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					IsNotNilOrEmptyOrWhitespace(testWhitespace, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					IsNotNilOrEmptyOrWhitespace(testWhitespace, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmptyOrWhitespace(testWhitespace, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					IsNotNilOrEmptyOrWhitespace(testWhitespace, testValueName, testMessage)
				})
			})
		})
	})

})

func TestAssertString(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assert")
}

func assertPanic(panicMessage, messageToDisplay string) {
	if r := recover(); r == nil || r != panicMessage {
		Fail(messageToDisplay)
	}
}
