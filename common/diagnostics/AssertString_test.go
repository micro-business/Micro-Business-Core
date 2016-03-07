package diagnostics

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var _ = Describe("Assert String", func() {
	var (
		emptyString                 string
		whitespace                  string
		testValueName               string
		testValueNameWithWhitespace string
		testMessage                 string
		testMessageWithWhitespace   string
	)

	BeforeEach(func() {
		emptyString = ""
		whitespace = "    "
		testValueName = "Test Value without whitepsace"
		testValueNameWithWhitespace = "     Test Value with whitspace   "
		testMessage = "Test Message without whitespace"
		testMessageWithWhitespace = "     Test Message with whitspace   "
	})

	Describe("when empty string value provided", func() {
		Context("with value name and message both empty", func() {
			It("Should panic with empty string as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with empty as message!!!")
					} else if r != emptyString {
						Fail("Should have paniced with empty as message!!!")
					}
				}()

				StringIsNotEmpty(emptyString, emptyString, emptyString)
			})
		})

		Context("with value name only provided", func() {
			Context("with value name contains no whitespace", func() {
				It("Should panic with provided value as message", func() {
					defer func() {
						if r := recover(); r == nil {
							Fail("Should have paniced with provided value as message!!!")
						} else if r != testValueName {
							Fail("Should have paniced with provided value as message!!!")
						}
					}()

					StringIsNotEmpty(emptyString, testValueName, emptyString)
				})
			})

			Context("with value name contains whitespace", func() {
				It("Should panic with provided trimmed value as message", func() {
					defer func() {
						if r := recover(); r == nil {
							Fail("Should have paniced with provided trimmed value as message!!!")
						} else if r != strings.TrimSpace(testValueNameWithWhitespace) {
							Fail("Should have paniced with provided trimmed value as message!!!")
						}
					}()

					StringIsNotEmpty(emptyString, testValueNameWithWhitespace, emptyString)
				})
			})
		})

		Context("with message only provided", func() {
			Context("with message contains no whitespace", func() {
				It("Should panic with provided message as message", func() {
					defer func() {
						if r := recover(); r == nil {
							Fail("Should have paniced with provided message as message!!!")
						} else if r != testMessage {
							Fail("Should have paniced with provided message as message!!!")
						}
					}()

					StringIsNotEmpty(emptyString, emptyString, testMessage)
				})
			})

			Context("with message contains whitespace", func() {
				It("Should panic with provided trimmed message as message", func() {
					defer func() {
						if r := recover(); r == nil {
							Fail("Should have paniced with provided trimmed message as message!!!")
						} else if r != strings.TrimSpace(testMessageWithWhitespace) {
							Fail("Should have paniced with provided trimmed message as message!!!")
						}
					}()

					StringIsNotEmpty(emptyString, emptyString, testMessageWithWhitespace)
				})
			})
		})

		Context("with value name and message both provided", func() {
			Context("with message contains no whitespace", func() {
				It("Should panic with provided message as message and ignore value name", func() {
					defer func() {
						if r := recover(); r == nil {
							Fail("Should have paniced with provided message as message!!!")
						} else if r != testMessage {
							Fail("Should have paniced with provided message as message!!!")
						}
					}()

					StringIsNotEmpty(emptyString, testValueName, testMessage)
				})
			})

			Context("with message contains whitespace", func() {
				It("Should panic with provided trimmed message as message and ignore value name", func() {
					defer func() {
						if r := recover(); r == nil {
							Fail("Should have paniced with provided trimmed message as message!!!")
						} else if r != strings.TrimSpace(testMessageWithWhitespace) {
							Fail("Should have paniced with provided trimmed message as message!!!")
						}
					}()

					StringIsNotEmpty(emptyString, testValueName, testMessageWithWhitespace)
				})
			})
		})
	})
})

func TestAssertString(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assert String")
}
