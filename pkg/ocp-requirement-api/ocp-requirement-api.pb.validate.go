// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/ocp-requirement-api/ocp-requirement-api.proto

package ocp_requirement_api

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
)

// Validate checks the field values on ListRequirementsV1Request with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRequirementsV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Limit

	// no validation rules for Offset

	return nil
}

// ListRequirementsV1RequestValidationError is the validation error returned by
// ListRequirementsV1Request.Validate if the designated constraints aren't met.
type ListRequirementsV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequirementsV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequirementsV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequirementsV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequirementsV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequirementsV1RequestValidationError) ErrorName() string {
	return "ListRequirementsV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e ListRequirementsV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequirementsV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequirementsV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequirementsV1RequestValidationError{}

// Validate checks the field values on ListRequirementsV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ListRequirementsV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRequirements() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListRequirementsV1ResponseValidationError{
					field:  fmt.Sprintf("Requirements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ListRequirementsV1ResponseValidationError is the validation error returned
// by ListRequirementsV1Response.Validate if the designated constraints aren't met.
type ListRequirementsV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequirementsV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequirementsV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequirementsV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequirementsV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequirementsV1ResponseValidationError) ErrorName() string {
	return "ListRequirementsV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListRequirementsV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequirementsV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequirementsV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequirementsV1ResponseValidationError{}

// Validate checks the field values on CreateRequirementV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRequirementV1Request) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for UserId

	// no validation rules for Text

	return nil
}

// CreateRequirementV1RequestValidationError is the validation error returned
// by CreateRequirementV1Request.Validate if the designated constraints aren't met.
type CreateRequirementV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequirementV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequirementV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequirementV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequirementV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequirementV1RequestValidationError) ErrorName() string {
	return "CreateRequirementV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRequirementV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequirementV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequirementV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequirementV1RequestValidationError{}

// Validate checks the field values on CreateRequirementV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateRequirementV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for RequirementId

	return nil
}

// CreateRequirementV1ResponseValidationError is the validation error returned
// by CreateRequirementV1Response.Validate if the designated constraints
// aren't met.
type CreateRequirementV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateRequirementV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateRequirementV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateRequirementV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateRequirementV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateRequirementV1ResponseValidationError) ErrorName() string {
	return "CreateRequirementV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateRequirementV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateRequirementV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateRequirementV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateRequirementV1ResponseValidationError{}

// Validate checks the field values on UpdateRequirementV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRequirementV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRequirements()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateRequirementV1RequestValidationError{
				field:  "Requirements",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateRequirementV1RequestValidationError is the validation error returned
// by UpdateRequirementV1Request.Validate if the designated constraints aren't met.
type UpdateRequirementV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRequirementV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRequirementV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRequirementV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRequirementV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRequirementV1RequestValidationError) ErrorName() string {
	return "UpdateRequirementV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRequirementV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRequirementV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRequirementV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRequirementV1RequestValidationError{}

// Validate checks the field values on UpdateRequirementV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateRequirementV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Updated

	return nil
}

// UpdateRequirementV1ResponseValidationError is the validation error returned
// by UpdateRequirementV1Response.Validate if the designated constraints
// aren't met.
type UpdateRequirementV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateRequirementV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateRequirementV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateRequirementV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateRequirementV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateRequirementV1ResponseValidationError) ErrorName() string {
	return "UpdateRequirementV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateRequirementV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateRequirementV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateRequirementV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateRequirementV1ResponseValidationError{}

// Validate checks the field values on MultiCreateRequirementV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *MultiCreateRequirementV1Request) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRequirements() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateRequirementV1RequestValidationError{
					field:  fmt.Sprintf("Requirements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateRequirementV1RequestValidationError is the validation error
// returned by MultiCreateRequirementV1Request.Validate if the designated
// constraints aren't met.
type MultiCreateRequirementV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateRequirementV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateRequirementV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateRequirementV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateRequirementV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateRequirementV1RequestValidationError) ErrorName() string {
	return "MultiCreateRequirementV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateRequirementV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateRequirementV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateRequirementV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateRequirementV1RequestValidationError{}

// Validate checks the field values on MultiCreateRequirementV1Response with
// the rules defined in the proto definition for this message. If any rules
// are violated, an error is returned.
func (m *MultiCreateRequirementV1Response) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetRequirements() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return MultiCreateRequirementV1ResponseValidationError{
					field:  fmt.Sprintf("Requirements[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// MultiCreateRequirementV1ResponseValidationError is the validation error
// returned by MultiCreateRequirementV1Response.Validate if the designated
// constraints aren't met.
type MultiCreateRequirementV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MultiCreateRequirementV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MultiCreateRequirementV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MultiCreateRequirementV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MultiCreateRequirementV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MultiCreateRequirementV1ResponseValidationError) ErrorName() string {
	return "MultiCreateRequirementV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e MultiCreateRequirementV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMultiCreateRequirementV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MultiCreateRequirementV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MultiCreateRequirementV1ResponseValidationError{}

// Validate checks the field values on RemoveRequirementV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRequirementV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRequirementId() <= 0 {
		return RemoveRequirementV1RequestValidationError{
			field:  "RequirementId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// RemoveRequirementV1RequestValidationError is the validation error returned
// by RemoveRequirementV1Request.Validate if the designated constraints aren't met.
type RemoveRequirementV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRequirementV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRequirementV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRequirementV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRequirementV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRequirementV1RequestValidationError) ErrorName() string {
	return "RemoveRequirementV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRequirementV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRequirementV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRequirementV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRequirementV1RequestValidationError{}

// Validate checks the field values on RemoveRequirementV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RemoveRequirementV1Response) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Found

	return nil
}

// RemoveRequirementV1ResponseValidationError is the validation error returned
// by RemoveRequirementV1Response.Validate if the designated constraints
// aren't met.
type RemoveRequirementV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RemoveRequirementV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RemoveRequirementV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RemoveRequirementV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RemoveRequirementV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RemoveRequirementV1ResponseValidationError) ErrorName() string {
	return "RemoveRequirementV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e RemoveRequirementV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRemoveRequirementV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RemoveRequirementV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RemoveRequirementV1ResponseValidationError{}

// Validate checks the field values on DescribeRequirementV1Request with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRequirementV1Request) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetRequirementId() <= 0 {
		return DescribeRequirementV1RequestValidationError{
			field:  "RequirementId",
			reason: "value must be greater than 0",
		}
	}

	return nil
}

// DescribeRequirementV1RequestValidationError is the validation error returned
// by DescribeRequirementV1Request.Validate if the designated constraints
// aren't met.
type DescribeRequirementV1RequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRequirementV1RequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRequirementV1RequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRequirementV1RequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRequirementV1RequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRequirementV1RequestValidationError) ErrorName() string {
	return "DescribeRequirementV1RequestValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRequirementV1RequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRequirementV1Request.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRequirementV1RequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRequirementV1RequestValidationError{}

// Validate checks the field values on DescribeRequirementV1Response with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *DescribeRequirementV1Response) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetRequirement()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return DescribeRequirementV1ResponseValidationError{
				field:  "Requirement",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// DescribeRequirementV1ResponseValidationError is the validation error
// returned by DescribeRequirementV1Response.Validate if the designated
// constraints aren't met.
type DescribeRequirementV1ResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DescribeRequirementV1ResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DescribeRequirementV1ResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DescribeRequirementV1ResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DescribeRequirementV1ResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DescribeRequirementV1ResponseValidationError) ErrorName() string {
	return "DescribeRequirementV1ResponseValidationError"
}

// Error satisfies the builtin error interface
func (e DescribeRequirementV1ResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDescribeRequirementV1Response.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DescribeRequirementV1ResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DescribeRequirementV1ResponseValidationError{}

// Validate checks the field values on Requirement with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *Requirement) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Id

	// no validation rules for UserId

	// no validation rules for Text

	return nil
}

// RequirementValidationError is the validation error returned by
// Requirement.Validate if the designated constraints aren't met.
type RequirementValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequirementValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequirementValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequirementValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequirementValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequirementValidationError) ErrorName() string { return "RequirementValidationError" }

// Error satisfies the builtin error interface
func (e RequirementValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequirement.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequirementValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequirementValidationError{}
