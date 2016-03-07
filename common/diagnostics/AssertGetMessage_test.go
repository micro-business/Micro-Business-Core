package diagnostics

import (
	"strings"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var _ = Describe("getMessage", func() {
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

	Describe("when no value name and message provided", func() {
		Context("with value name and message both empty", func() {
			It("Should return empty string message", func() {
				Expect(getMessage(emptyString, emptyString)).To(Equal(emptyString))
			})
		})

		Context("with value name equals to empty and message contains whitespace", func() {
			It("Should return empty string message", func() {
				Expect(getMessage(emptyString, whitespace)).To(Equal(emptyString))
			})
		})

		Context("with value name contains whitespace and message is empty", func() {
			It("Should return empty string message", func() {
				Expect(getMessage(whitespace, emptyString)).To(Equal(emptyString))
			})
		})

		Context("with value name and message both contians whitespace", func() {
			It("Should return empty string message", func() {
				Expect(getMessage(whitespace, whitespace)).To(Equal(emptyString))
			})
		})
	})

	Describe("when value name contains no whitespace", func() {
		Context("with message is empty", func() {
			It("Should return provided value name as message", func() {
				Expect(getMessage(testValueName, emptyString)).To(Equal(testValueName))
			})
		})

		Context("with message contains whitespace", func() {
			It("Should return provided value name as message", func() {
				Expect(getMessage(testValueName, whitespace)).To(Equal(testValueName))
			})
		})
	})

	Describe("when value name contains whitespace", func() {
		Context("with message is empty", func() {
			It("Should trim provided value name and return it as message", func() {
				Expect(getMessage(testValueNameWithWhitespace, emptyString)).To(Equal(strings.TrimSpace(testValueNameWithWhitespace)))
			})
		})

		Context("with message contains whitespace", func() {
			It("Should trim provided value name and return it as message", func() {
				Expect(getMessage(testValueNameWithWhitespace, whitespace)).To(Equal(strings.TrimSpace(testValueNameWithWhitespace)))

			})
		})
	})

	Describe("when message contains no whitespace", func() {
		Context("with value name is empty", func() {
			It("Should return provided message as message", func() {
				Expect(getMessage(emptyString, testMessage)).To(Equal(testMessage))
			})
		})

		Context("with value name contains whitespace", func() {
			It("Should return provided message as message", func() {
				Expect(getMessage(whitespace, testMessage)).To(Equal(testMessage))
			})
		})
	})

	Describe("when message contains whitespace", func() {
		Context("with value name is empty", func() {
			It("Should trim provided message and return it as message", func() {
				Expect(getMessage(emptyString, testMessageWithWhitespace)).To(Equal(strings.TrimSpace(testMessageWithWhitespace)))
			})
		})

		Context("with value name contains whitespace", func() {
			It("Should trim provided message and return it as message", func() {
				Expect(getMessage(whitespace, testMessageWithWhitespace)).To(Equal(strings.TrimSpace(testMessageWithWhitespace)))

			})
		})
	})

	Describe("when value name and message both contain valid message", func() {
		Context("with message contains no whitespace", func() {
			It("Should ignore value name and return message as message", func() {
				Expect(getMessage(testValueName, testMessage)).To(Equal(testMessage))
			})
		})

		Context("with message contains whitespace", func() {
			It("Should ignore value name, trim provided message and return it as message", func() {
				Expect(getMessage(testValueName, testMessageWithWhitespace)).To(Equal(strings.TrimSpace(testMessageWithWhitespace)))
			})
		})
	})
})

func TestGetMessage(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "getMessage")
}
