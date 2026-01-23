package utils

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

/*
Validation use package
*/
func HandleValidationErrors(err error) gin.H {
	if validationErrpr, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)
		for _, e := range validationErrpr {

			//CreatePostV1Body.ProductAttributes[2].AttributeName
			// split index 0 (from .) to delete it
			root := strings.Split(e.Namespace(), ".")[0]
			// TrimPrefix CreatePostV1Body.ProductAttributes[2].AttributeName
			// from index 0 to .
			// => ProductAttributes[2].AttributeName
			fieldPath := strings.TrimPrefix(e.Namespace(), root+".")

			switch e.Tag() {
			case "gt":
				errors[fieldPath] = fmt.Sprintf("%s phải nhiều hơn %s", fieldPath, e.Param())
			case "lt":
				errors[fieldPath] = fmt.Sprintf("%s phải ít hơn %s", fieldPath, e.Param())
			case "slug":
				errors[fieldPath] = fieldPath + " phải là số, chữ thường và dấu gạch chân hoặc dấu chấm"
			case "uuid":
				errors[fieldPath] = fieldPath + " phải là UUID hợp lệ"
			case "min":
				errors[fieldPath] = fmt.Sprintf("%s phải nhiều hơn %s ký tự", fieldPath, e.Param())
			case "max":
				errors[fieldPath] = fmt.Sprintf("%s phải ít hơn %s ký tự", fieldPath, e.Param())
			case "min_int":
				errors[fieldPath] = fmt.Sprintf("%s phải có giá trị lớn hơn %s", fieldPath, e.Param())
			case "max_int":
				errors[fieldPath] = fmt.Sprintf("%s phải có giá trị nhỏ hơn %s", fieldPath, e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[fieldPath] = fmt.Sprintf("%s phải là 1 trong những giá trị: %s", fieldPath, allowedValues)
			case "search":
				errors[fieldPath] = fmt.Sprintf("%s phải là chữ thường, chữ hoa, số và khoảng trắng", fieldPath)
			case "required":
				errors[fieldPath] = fmt.Sprintf("%s là bắt buộc", fieldPath)

			}
		}
		return gin.H{"error": errors}
	}
	return gin.H{"error": "Yêu cầu không hợp lệ"}
}

/*
*
manual Validation
*/
func ValidationRequired(field, value string) error {
	if value == "" {
		return fmt.Errorf("%s is required", field)
	}
	return nil
}

func ValidationStringLength(field, value string, min, max int) error {
	l := len(value)
	if l < min || l > max {
		return fmt.Errorf("%s must be between %d and %d characters", field, min, max)
	}
	return nil
}

func ValidationRegex(value string, re *regexp.Regexp, errMsg string) error {
	if !re.MatchString(value) {
		return fmt.Errorf("%s", errMsg)
	}
	return nil
}

func ValidationPositiveInt(field, value string) (int, error) {
	v, err := strconv.Atoi(value)
	if err != nil {
		return 0, fmt.Errorf("%s must be a number", field)
	}

	if v <= 0 {
		return 0, fmt.Errorf("%s must be a positive number", field)
	}

	return v, nil
}

func ValidationUUID(field, value string) (uuid.UUID, error) {
	uid, err := uuid.Parse(value)
	if err != nil {
		return uuid.Nil, fmt.Errorf("%s must be a valid UUID", field)
	}
	return uid, nil
}

func ValidationInList(field, value string, allowed map[string]bool) error {
	if !allowed[value] {
		return fmt.Errorf("%s must be one of: %v", field, keys(allowed))
	}
	return nil
}

// convert map to slice
func keys(m map[string]bool) []string {
	var k []string
	for key := range m {
		k = append(k, key)
	}
	return k
}

func RegisterValidator() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}
	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	var searchRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchRegex.MatchString(fl.Field().String())
	})

	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		minValue, err := strconv.Atoi(minStr)
		if err != nil {
			return false
		}
		return int(fl.Field().Int()) >= minValue
	})

	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		maxValue, err := strconv.Atoi(maxStr)
		if err != nil {
			return false
		}
		return int(fl.Field().Int()) <= maxValue
	})

	return nil
}
