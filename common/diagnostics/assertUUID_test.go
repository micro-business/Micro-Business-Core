package diagnostics_test

import (
	"github.com/microbusinesses/Micro-Businesses-Core/common/diagnostics"
	"github.com/microbusinesses/Micro-Businesses-Core/system"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

var _ = Describe("Assert UUID is empty", func() {
	var (
		emptyString   string
		testUUID      system.UUID
		testValueName string
		testMessage   string
	)

	BeforeEach(func() {
		emptyString = ""

		var uuid system.UUID

		uuid, _ = system.RandomUUID()
		testUUID = uuid

		uuid, _ = system.RandomUUID()
		testValueName = uuid.String()

		uuid, _ = system.RandomUUID()
		testMessage = uuid.String()
	})

	Describe("when empty UUID value provided", func() {
		It("Should return empty UUID", func() {
			Expect(diagnostics.UUIDIsEmpty(system.EmptyUUID, emptyString, emptyString)).To(Equal(system.EmptyUUID))
		})
	})

	Describe("when UUID provided", func() {
		Context("with value name and message both emptyd", func() {
			It("Should panic with empty string as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with empty as message!!!")
					} else if r != emptyString {
						Fail("Should have paniced with empty as message!!!")
					}
				}()

				diagnostics.UUIDIsEmpty(testUUID, emptyString, emptyString)
			})
		})

		Context("with value name provided", func() {
			It("Should panic with provided value name as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with value name as message!!!")
					} else if r != testValueName {
						Fail("Should have paniced with value name as message!!!")
					}
				}()

				diagnostics.UUIDIsEmpty(testUUID, testValueName, emptyString)
			})
		})

		Context("with message provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.UUIDIsEmpty(testUUID, emptyString, testMessage)
			})
		})

		Context("with name and message both provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.UUIDIsEmpty(testUUID, testValueName, testMessage)
			})
		})
	})
})

var _ = Describe("Assert UUID is not empty", func() {
	var (
		emptyString   string
		testUUID      system.UUID
		testValueName string
		testMessage   string
	)

	BeforeEach(func() {
		emptyString = ""

		var uuid system.UUID

		uuid, _ = system.RandomUUID()
		testUUID = uuid

		uuid, _ = system.RandomUUID()
		testValueName = uuid.String()

		uuid, _ = system.RandomUUID()
		testMessage = uuid.String()
	})

	Describe("when UUID value provided", func() {
		It("Should return empty string", func() {
			Expect(diagnostics.UUIDIsNotEmpty(testUUID, emptyString, emptyString)).To(Equal(testUUID))
		})
	})

	Describe("when empty UUID value provided", func() {
		Context("with value name and message both emptyd", func() {
			It("Should panic with empty string as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with empty as message!!!")
					} else if r != emptyString {
						Fail("Should have paniced with empty as message!!!")
					}
				}()

				diagnostics.UUIDIsNotEmpty(system.EmptyUUID, emptyString, emptyString)
			})
		})

		Context("with value name provided", func() {
			It("Should panic with provided value name as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with value name as message!!!")
					} else if r != testValueName {
						Fail("Should have paniced with value name as message!!!")
					}
				}()

				diagnostics.UUIDIsNotEmpty(system.EmptyUUID, testValueName, emptyString)
			})
		})

		Context("with message provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.StringIsNotEmpty(emptyString, emptyString, testMessage)
			})
		})

		Context("with name and message both provided", func() {
			It("Should panic with provided message as message", func() {
				defer func() {
					if r := recover(); r == nil {
						Fail("Should have paniced with message as message!!!")
					} else if r != testMessage {
						Fail("Should have paniced with message as message!!!")
					}
				}()

				diagnostics.UUIDIsNotEmpty(system.EmptyUUID, testValueName, testMessage)
			})
		})
	})
})

func TestAssertUUID(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Assert UUID is empty")
	RunSpecs(t, "Assert UUID is not empty")
}
