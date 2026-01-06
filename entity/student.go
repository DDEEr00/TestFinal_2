package entity

import (
	"github.com/asaskevich/govalidator"
)

type Student struct {
	FullName string  `valid:"required~Fullname is required"`
	Age      int     `valid:"range(18|200)~Age must be at least 18"`
	Email    string  `valid:"email~Email is invalid,required~Email is required"`
	GPA      float32 `valid:"float,range(0|4)~GPA must be between 0.00 and 4.00"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(false)
}

func ValidateStudent(s Student) (bool, error) {
	return govalidator.ValidateStruct(s)
}
