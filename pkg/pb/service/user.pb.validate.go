// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: service/user.proto

package service

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"

	enums "github.com/akubi0w1/sge-tech-book-vol3-sample/pkg/pb/enums"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort

	_ = enums.Platform(0)
)

// define the regex for a UUID once up-front
var _user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on ListUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserRequestMultiError, or nil if none found.
func (m *ListUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ListUserRequestMultiError(errors)
	}

	return nil
}

// ListUserRequestMultiError is an error wrapping multiple validation errors
// returned by ListUserRequest.ValidateAll() if the designated constraints
// aren't met.
type ListUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserRequestMultiError) AllErrors() []error { return m }

// ListUserRequestValidationError is the validation error returned by
// ListUserRequest.Validate if the designated constraints aren't met.
type ListUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserRequestValidationError) ErrorName() string { return "ListUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserRequestValidationError{}

// Validate checks the field values on ListUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListUserResponseMultiError, or nil if none found.
func (m *ListUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetUsers() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListUserResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListUserResponseValidationError{
						field:  fmt.Sprintf("Users[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListUserResponseValidationError{
					field:  fmt.Sprintf("Users[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListUserResponseMultiError(errors)
	}

	return nil
}

// ListUserResponseMultiError is an error wrapping multiple validation errors
// returned by ListUserResponse.ValidateAll() if the designated constraints
// aren't met.
type ListUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListUserResponseMultiError) AllErrors() []error { return m }

// ListUserResponseValidationError is the validation error returned by
// ListUserResponse.Validate if the designated constraints aren't met.
type ListUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUserResponseValidationError) ErrorName() string { return "ListUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUserResponseValidationError{}

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserRequestMultiError,
// or nil if none found.
func (m *GetUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUserId()); err != nil {
		err = GetUserRequestValidationError{
			field:  "UserId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return GetUserRequestMultiError(errors)
	}

	return nil
}

func (m *GetUserRequest) _validateUuid(uuid string) error {
	if matched := _user_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// GetUserRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserRequestMultiError) AllErrors() []error { return m }

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserResponseMultiError, or nil if none found.
func (m *GetUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserResponseMultiError(errors)
	}

	return nil
}

// GetUserResponseMultiError is an error wrapping multiple validation errors
// returned by GetUserResponse.ValidateAll() if the designated constraints
// aren't met.
type GetUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserResponseMultiError) AllErrors() []error { return m }

// GetUserResponseValidationError is the validation error returned by
// GetUserResponse.Validate if the designated constraints aren't met.
type GetUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserResponseValidationError) ErrorName() string { return "GetUserResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserResponseValidationError{}

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserRequestMultiError, or nil if none found.
func (m *CreateUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 20 {
		err := CreateUserRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if _, ok := enums.Platform_name[int32(m.GetPlatform())]; !ok {
		err := CreateUserRequestValidationError{
			field:  "Platform",
			reason: "value must be one of the defined enum values",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return CreateUserRequestMultiError(errors)
	}

	return nil
}

// CreateUserRequestMultiError is an error wrapping multiple validation errors
// returned by CreateUserRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserRequestMultiError) AllErrors() []error { return m }

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

// Validate checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateUserResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateUserResponseMultiError, or nil if none found.
func (m *CreateUserResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateUserResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateUserResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateUserResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateUserResponseMultiError(errors)
	}

	return nil
}

// CreateUserResponseMultiError is an error wrapping multiple validation errors
// returned by CreateUserResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateUserResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateUserResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateUserResponseMultiError) AllErrors() []error { return m }

// CreateUserResponseValidationError is the validation error returned by
// CreateUserResponse.Validate if the designated constraints aren't met.
type CreateUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserResponseValidationError) ErrorName() string {
	return "CreateUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserResponseValidationError{}

// Validate checks the field values on UpdateUserNameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserNameRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserNameRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserNameRequestMultiError, or nil if none found.
func (m *UpdateUserNameRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserNameRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUserId()); err != nil {
		err = UpdateUserNameRequestValidationError{
			field:  "UserId",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if l := utf8.RuneCountInString(m.GetName()); l < 1 || l > 20 {
		err := UpdateUserNameRequestValidationError{
			field:  "Name",
			reason: "value length must be between 1 and 20 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateUserNameRequestMultiError(errors)
	}

	return nil
}

func (m *UpdateUserNameRequest) _validateUuid(uuid string) error {
	if matched := _user_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UpdateUserNameRequestMultiError is an error wrapping multiple validation
// errors returned by UpdateUserNameRequest.ValidateAll() if the designated
// constraints aren't met.
type UpdateUserNameRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserNameRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserNameRequestMultiError) AllErrors() []error { return m }

// UpdateUserNameRequestValidationError is the validation error returned by
// UpdateUserNameRequest.Validate if the designated constraints aren't met.
type UpdateUserNameRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserNameRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserNameRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserNameRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserNameRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserNameRequestValidationError) ErrorName() string {
	return "UpdateUserNameRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserNameRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserNameRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserNameRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserNameRequestValidationError{}

// Validate checks the field values on UpdateUserNameResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateUserNameResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateUserNameResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateUserNameResponseMultiError, or nil if none found.
func (m *UpdateUserNameResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateUserNameResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, UpdateUserNameResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, UpdateUserNameResponseValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateUserNameResponseValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return UpdateUserNameResponseMultiError(errors)
	}

	return nil
}

// UpdateUserNameResponseMultiError is an error wrapping multiple validation
// errors returned by UpdateUserNameResponse.ValidateAll() if the designated
// constraints aren't met.
type UpdateUserNameResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateUserNameResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateUserNameResponseMultiError) AllErrors() []error { return m }

// UpdateUserNameResponseValidationError is the validation error returned by
// UpdateUserNameResponse.Validate if the designated constraints aren't met.
type UpdateUserNameResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserNameResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserNameResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserNameResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserNameResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserNameResponseValidationError) ErrorName() string {
	return "UpdateUserNameResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserNameResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserNameResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserNameResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserNameResponseValidationError{}
