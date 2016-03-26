package diagnostics_test

import (
	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	"github.com/microbusinesses/Micro-Businesses-Core/system"
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

		uuid, _ := system.RandomUUID()
		testString = uuid.String()

		testWhitespace = "      "

		uuid, _ = system.RandomUUID()
		testValue = uuid.String()

		uuid, _ = system.RandomUUID()
		testValueName = uuid.String()

		uuid, _ = system.RandomUUID()
		testMessage = uuid.String()
	})

	Describe("IsNil", func() {
		Context("when nil value provided", func() {
			It("Should return nil", func() {
				Expect(diagnostics.IsNil(nil, emptyString, emptyString)).To(BeNil())
			})
		})

		Context("when non-nill value provided", func() {
			Context("with value name and message both emptyd", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					diagnostics.IsNil(testValue, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					diagnostics.IsNil(testValue, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNil(testValue, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNil(testValue, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNilOrEmpty", func() {
		Context("when nil value provided", func() {
			It("Should return nil", func() {
				Expect(diagnostics.IsNilOrEmpty(nil, emptyString, emptyString)).To(BeNil())
			})
		})

		Context("when empty string value provided", func() {
			It("Should return empty string", func() {
				Expect(diagnostics.IsNilOrEmpty(emptyString, emptyString, emptyString)).To(Equal(emptyString))
			})
		})

		Context("when empty UUID value provided", func() {
			It("Should return empty UUID", func() {
				Expect(diagnostics.IsNilOrEmpty(system.EmptyUUID, emptyString, emptyString)).To(Equal(system.EmptyUUID))
			})
		})

		Context("when non-nill or empty value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					diagnostics.IsNilOrEmpty(testValue, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					diagnostics.IsNilOrEmpty(testValue, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNilOrEmpty(testValue, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNilOrEmpty(testValue, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNilOrEmptyOrWhitespace", func() {
		Context("when nil value provided", func() {
			It("Should return nil", func() {
				Expect(diagnostics.IsNilOrEmptyOrWhitespace(nil, emptyString, emptyString)).To(BeNil())
			})
		})

		Context("when empty string value provided", func() {
			It("Should return empty string", func() {
				Expect(diagnostics.IsNilOrEmptyOrWhitespace(emptyString, emptyString, emptyString)).To(Equal(emptyString))
			})
		})

		Context("when string contains whitespace only as value provided", func() {
			It("Should return provided value", func() {
				Expect(diagnostics.IsNilOrEmptyOrWhitespace(testWhitespace, emptyString, emptyString)).To(Equal(testWhitespace))
			})
		})

		Context("when empty UUID value provided", func() {
			It("Should return empty UUID", func() {
				Expect(diagnostics.IsNilOrEmptyOrWhitespace(system.EmptyUUID, emptyString, emptyString)).To(Equal(system.EmptyUUID))
			})
		})

		Context("when non-nill, empty or contains whitespace only provided as value", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					diagnostics.IsNilOrEmptyOrWhitespace(testValue, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					diagnostics.IsNilOrEmptyOrWhitespace(testValue, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNilOrEmptyOrWhitespace(testValue, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNilOrEmptyOrWhitespace(testValue, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNotNil", func() {
		Context("when non-nil value provided", func() {
			It("Should return provided value", func() {
				Expect(diagnostics.IsNotNil(testValue, emptyString, emptyString)).To(Equal(testValue))
			})
		})

		Context("when nill value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					diagnostics.IsNotNil(nil, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					diagnostics.IsNotNil(nil, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNil(nil, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNil(nil, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNotNilOrEmpty", func() {
		Context("when non-nil value provided", func() {
			It("Should return provided value", func() {
				Expect(diagnostics.IsNotNilOrEmpty(testValue, emptyString, emptyString)).To(Equal(testValue))
			})
		})

		Context("when nill value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					diagnostics.IsNotNilOrEmpty(nil, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					diagnostics.IsNotNilOrEmpty(nil, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmpty(nil, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmpty(nil, testValueName, testMessage)
				})
			})
		})

		Context("when empty string value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					diagnostics.IsNotNilOrEmpty(emptyString, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					diagnostics.IsNotNilOrEmpty(emptyString, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmpty(emptyString, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmpty(emptyString, testValueName, testMessage)
				})
			})
		})
	})

	Describe("IsNotNilOrEmptyOrWhitespace", func() {
		Context("when non-nil value provided", func() {
			It("Should return provided value", func() {
				Expect(diagnostics.IsNotNilOrEmptyOrWhitespace(testValue, emptyString, emptyString)).To(Equal(testValue))
			})
		})

		Context("when nill value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(nil, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(nil, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(nil, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(nil, testValueName, testMessage)
				})
			})
		})

		Context("when empty string value provided", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(emptyString, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(emptyString, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(emptyString, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(emptyString, testValueName, testMessage)
				})
			})
		})

		Context("when string contains whitespace only provided as value", func() {
			Context("with value name and message both empty", func() {
				It("Should panic with empty string as message", func() {
					defer assertPanic(emptyString, "Should have paniced with empty as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(testWhitespace, emptyString, emptyString)
				})
			})

			Context("with value name provided", func() {
				It("Should panic with provided value name as message", func() {
					defer assertPanic(testValueName, "Should have paniced with value name as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(testWhitespace, testValueName, emptyString)
				})
			})

			Context("with message provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(testWhitespace, emptyString, testMessage)
				})
			})

			Context("with name and message both provided", func() {
				It("Should panic with provided message as message", func() {
					defer assertPanic(testMessage, "Should have paniced with message as message!!!")

					diagnostics.IsNotNilOrEmptyOrWhitespace(testWhitespace, testValueName, testMessage)
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
