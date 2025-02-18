package models

import "github.com/elisalimli/nextsync/server/validator"

func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", r.Email)
	v.IsEmail("email", r.Email)

	v.Required("password", r.Password)
	v.MinLength("password", r.Password, 6)

	v.Required("confirmPassword", r.ConfirmPassword)
	v.EqualToField("confirmPassword", r.ConfirmPassword, "password", r.Password)

	v.Required("username", r.Username)
	v.MinLength("username", r.Username, 2)

	v.Required("firstName", r.FirstName)
	v.MinLength("firstName", r.FirstName, 2)

	v.Required("lastName", r.LastName)
	v.MinLength("lastName", r.LastName, 2)

	v.Required("phoneNumber", r.PhoneNumber)
	v.IsPhoneNumber("phoneNumber", r.PhoneNumber)

	return v.IsValid(), v.Errors
}

func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("email", l.Email)
	v.IsEmail("email", l.Email)

	v.Required("password", l.Password)

	return v.IsValid(), v.Errors
}

func (l SendOtpInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("to", l.To)
	v.IsPhoneNumber("to", l.To)

	return v.IsValid(), v.Errors
}

func (l VerifyOtpInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("to", l.To)
	v.IsPhoneNumber("to", l.To)

	v.Required("code", l.Code)

	return v.IsValid(), v.Errors
}

func (l CreatePostInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Required("title", l.Title)
	// v.MinLength("lastName", l.LastName, 2)

	return v.IsValid(), v.Errors
}
