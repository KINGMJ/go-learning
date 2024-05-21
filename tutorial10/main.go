package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// User contains user information
type User struct {
	FirstName      string       `json:"first_name,optional"`
	LastName       string       `validate:"required"`
	Age            uint8        `json:"age" validate:"gte=0,lte=130"`
	Email          string       `validate:"required,email"`
	Gender         SupplierMode `validate:"oneof=1 2"`
	FavouriteColor string       `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address   `validate:"required,dive,required"` // a person can have a home and cottage...
}

type SupplierMode int

const (
	_                        SupplierMode = iota
	SupplierModePurchaseSale              // 购销
	SupplierModeJoint                     // 联营
)

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	validate = validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})

	validateStruct()
}

func validateStruct() {
	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Gender:         6,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
	}

	// returns nil or ValidationErrors ( []FieldError )
	err := validate.Struct(user)
	if err != nil {
		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}
		validationErrors := make([]string, 0)
		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, fmt.Sprintf("%s 格式错误", err.Field()))
		}
		fmt.Println(strings.Join(validationErrors, ", "))
	}
}
