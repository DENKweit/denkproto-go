// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: modelfile-v1.proto

package v1

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
)

// Validate checks the field values on RegionFromEdge with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegionFromEdge) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegionFromEdge with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegionFromEdgeMultiError,
// or nil if none found.
func (m *RegionFromEdge) ValidateAll() error {
	return m.validate(true)
}

func (m *RegionFromEdge) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Left

	// no validation rules for Right

	// no validation rules for Top

	// no validation rules for Bottom

	if len(errors) > 0 {
		return RegionFromEdgeMultiError(errors)
	}

	return nil
}

// RegionFromEdgeMultiError is an error wrapping multiple validation errors
// returned by RegionFromEdge.ValidateAll() if the designated constraints
// aren't met.
type RegionFromEdgeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegionFromEdgeMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegionFromEdgeMultiError) AllErrors() []error { return m }

// RegionFromEdgeValidationError is the validation error returned by
// RegionFromEdge.Validate if the designated constraints aren't met.
type RegionFromEdgeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegionFromEdgeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegionFromEdgeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegionFromEdgeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegionFromEdgeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegionFromEdgeValidationError) ErrorName() string { return "RegionFromEdgeValidationError" }

// Error satisfies the builtin error interface
func (e RegionFromEdgeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegionFromEdge.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegionFromEdgeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegionFromEdgeValidationError{}

// Validate checks the field values on FeatureClass with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *FeatureClass) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on FeatureClass with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in FeatureClassMultiError, or
// nil if none found.
func (m *FeatureClass) ValidateAll() error {
	return m.validate(true)
}

func (m *FeatureClass) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for FeatureUid

	// no validation rules for FeatureName

	// no validation rules for FeatureTag

	if len(errors) > 0 {
		return FeatureClassMultiError(errors)
	}

	return nil
}

// FeatureClassMultiError is an error wrapping multiple validation errors
// returned by FeatureClass.ValidateAll() if the designated constraints aren't met.
type FeatureClassMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m FeatureClassMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m FeatureClassMultiError) AllErrors() []error { return m }

// FeatureClassValidationError is the validation error returned by
// FeatureClass.Validate if the designated constraints aren't met.
type FeatureClassValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FeatureClassValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FeatureClassValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FeatureClassValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FeatureClassValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FeatureClassValidationError) ErrorName() string { return "FeatureClassValidationError" }

// Error satisfies the builtin error interface
func (e FeatureClassValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFeatureClass.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FeatureClassValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FeatureClassValidationError{}

// Validate checks the field values on InputField with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *InputField) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on InputField with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in InputFieldMultiError, or
// nil if none found.
func (m *InputField) ValidateAll() error {
	return m.validate(true)
}

func (m *InputField) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Label

	// no validation rules for Datatype

	// no validation rules for ImageW

	// no validation rules for ImageH

	// no validation rules for ImageC

	for idx, item := range m.GetRegionOfInterest() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, InputFieldValidationError{
						field:  fmt.Sprintf("RegionOfInterest[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, InputFieldValidationError{
						field:  fmt.Sprintf("RegionOfInterest[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return InputFieldValidationError{
					field:  fmt.Sprintf("RegionOfInterest[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TrainingImageW

	// no validation rules for TrainingImageH

	// no validation rules for TrainingImageC

	// no validation rules for MovingWindow

	if len(errors) > 0 {
		return InputFieldMultiError(errors)
	}

	return nil
}

// InputFieldMultiError is an error wrapping multiple validation errors
// returned by InputField.ValidateAll() if the designated constraints aren't met.
type InputFieldMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m InputFieldMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m InputFieldMultiError) AllErrors() []error { return m }

// InputFieldValidationError is the validation error returned by
// InputField.Validate if the designated constraints aren't met.
type InputFieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e InputFieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e InputFieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e InputFieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e InputFieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e InputFieldValidationError) ErrorName() string { return "InputFieldValidationError" }

// Error satisfies the builtin error interface
func (e InputFieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sInputField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = InputFieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = InputFieldValidationError{}

// Validate checks the field values on OcrFormatRestrictionBlock with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *OcrFormatRestrictionBlock) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OcrFormatRestrictionBlock with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// OcrFormatRestrictionBlockMultiError, or nil if none found.
func (m *OcrFormatRestrictionBlock) ValidateAll() error {
	return m.validate(true)
}

func (m *OcrFormatRestrictionBlock) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for NumberOfCharacters

	// no validation rules for AllowedCharacters

	if len(errors) > 0 {
		return OcrFormatRestrictionBlockMultiError(errors)
	}

	return nil
}

// OcrFormatRestrictionBlockMultiError is an error wrapping multiple validation
// errors returned by OcrFormatRestrictionBlock.ValidateAll() if the
// designated constraints aren't met.
type OcrFormatRestrictionBlockMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OcrFormatRestrictionBlockMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OcrFormatRestrictionBlockMultiError) AllErrors() []error { return m }

// OcrFormatRestrictionBlockValidationError is the validation error returned by
// OcrFormatRestrictionBlock.Validate if the designated constraints aren't met.
type OcrFormatRestrictionBlockValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OcrFormatRestrictionBlockValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OcrFormatRestrictionBlockValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OcrFormatRestrictionBlockValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OcrFormatRestrictionBlockValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OcrFormatRestrictionBlockValidationError) ErrorName() string {
	return "OcrFormatRestrictionBlockValidationError"
}

// Error satisfies the builtin error interface
func (e OcrFormatRestrictionBlockValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOcrFormatRestrictionBlock.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OcrFormatRestrictionBlockValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OcrFormatRestrictionBlockValidationError{}

// Validate checks the field values on OutputField with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *OutputField) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on OutputField with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in OutputFieldMultiError, or
// nil if none found.
func (m *OutputField) ValidateAll() error {
	return m.validate(true)
}

func (m *OutputField) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Label

	// no validation rules for Datatype

	// no validation rules for ImageW

	// no validation rules for ImageH

	// no validation rules for ImageC

	for idx, item := range m.GetClasses() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OutputFieldValidationError{
						field:  fmt.Sprintf("Classes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OutputFieldValidationError{
						field:  fmt.Sprintf("Classes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OutputFieldValidationError{
					field:  fmt.Sprintf("Classes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for MaxQueries

	// no validation rules for Charset

	// no validation rules for CharsetFilter

	for idx, item := range m.GetOcrFormatRestrictions() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, OutputFieldValidationError{
						field:  fmt.Sprintf("OcrFormatRestrictions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, OutputFieldValidationError{
						field:  fmt.Sprintf("OcrFormatRestrictions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return OutputFieldValidationError{
					field:  fmt.Sprintf("OcrFormatRestrictions[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return OutputFieldMultiError(errors)
	}

	return nil
}

// OutputFieldMultiError is an error wrapping multiple validation errors
// returned by OutputField.ValidateAll() if the designated constraints aren't met.
type OutputFieldMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m OutputFieldMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m OutputFieldMultiError) AllErrors() []error { return m }

// OutputFieldValidationError is the validation error returned by
// OutputField.Validate if the designated constraints aren't met.
type OutputFieldValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e OutputFieldValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e OutputFieldValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e OutputFieldValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e OutputFieldValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e OutputFieldValidationError) ErrorName() string { return "OutputFieldValidationError" }

// Error satisfies the builtin error interface
func (e OutputFieldValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sOutputField.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = OutputFieldValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = OutputFieldValidationError{}

// Validate checks the field values on ModelFile with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ModelFile) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ModelFile with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ModelFileMultiError, or nil
// if none found.
func (m *ModelFile) ValidateAll() error {
	return m.validate(true)
}

func (m *ModelFile) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ProtocolVersion

	// no validation rules for Model

	// no validation rules for CompressionMethod

	// no validation rules for ModelHashBlake2B

	for idx, item := range m.GetInput() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ModelFileValidationError{
						field:  fmt.Sprintf("Input[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ModelFileValidationError{
						field:  fmt.Sprintf("Input[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ModelFileValidationError{
					field:  fmt.Sprintf("Input[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	for idx, item := range m.GetOutput() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ModelFileValidationError{
						field:  fmt.Sprintf("Output[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ModelFileValidationError{
						field:  fmt.Sprintf("Output[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ModelFileValidationError{
					field:  fmt.Sprintf("Output[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for Tenant

	// no validation rules for ModelLabel

	// no validation rules for ModelTag

	// no validation rules for ModelId

	// no validation rules for CreationTimestamp

	// no validation rules for TenantId

	// no validation rules for OnnxVersionMajor

	// no validation rules for OnnxVersionMinor

	// no validation rules for ModelUid

	// no validation rules for ModelTimestamp

	// no validation rules for ModelType

	// no validation rules for ModelOutputType

	// no validation rules for FileType

	// no validation rules for Key

	if len(errors) > 0 {
		return ModelFileMultiError(errors)
	}

	return nil
}

// ModelFileMultiError is an error wrapping multiple validation errors returned
// by ModelFile.ValidateAll() if the designated constraints aren't met.
type ModelFileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ModelFileMultiError) Error() string {
	msgs := make([]string, 0, len(m))
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ModelFileMultiError) AllErrors() []error { return m }

// ModelFileValidationError is the validation error returned by
// ModelFile.Validate if the designated constraints aren't met.
type ModelFileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ModelFileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ModelFileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ModelFileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ModelFileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ModelFileValidationError) ErrorName() string { return "ModelFileValidationError" }

// Error satisfies the builtin error interface
func (e ModelFileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sModelFile.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ModelFileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ModelFileValidationError{}
