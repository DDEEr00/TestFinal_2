package validation

import (
	"testing"

	."github.com/onsi/gomega"
)

func TestAllFull(t *testing.T){
	g := NewGomegaWithT(t)

	s := Student{
		FullName: "deer",
		Age: 20,
		Email: "deer@gmail.com",
		GPA: 3.00
	}

	ok, err := ValidateStudent(s)

	g.Expact(ok).To(BeTrue())
	g.Expact(err).To(BeNil())
}

func TestFullNameRequired(t *testing.T)  {
	g := NewGomegaWithT(t)

	s := Student{
		FullName: "",
		Age: 20,
		Email: "deer@gmail.com",
		GPA: 3.00
	}

	ok, err := ValidateStudent(s)

	g.Expact(ok).To(BeFalse)
	g.Expact(err.Error()).To(Equal("Fullname is required"))
}

func TestAgeMustBeAtLeast18(t *testing.T) {
	g := NewGomegaWithT(t)

	s := Student{
		Fullname: "John Doe",
		Age:      17,
		Email:    "john@example.com",
		GPA:      3.0,
	}

	ok, err := ValidateStudent(s)

	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Age must be at least 18"))
}

func TestEmailInvalidFormat(t *testing.T) {
	g := NewGomegaWithT(t)

	s := Student{
		Fullname: "John Doe",
		Age:      20,
		Email:    "invalid-email",
		GPA:      3.0,
	}

	ok, err := ValidateStudent(s)

	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("Email is invalid"))
}

func TestGPAOutOfRange(t *testing.T) {
	g := NewGomegaWithT(t)

	s := Student{
		Fullname: "John Doe",
		Age:      20,
		Email:    "john@example.com",
		GPA:      5.0,
	}

	ok, err := ValidateStudent(s)

	g.Expect(ok).To(BeFalse())
	g.Expect(err.Error()).To(Equal("GPA must be between 0.00 and 4.00"))
}

