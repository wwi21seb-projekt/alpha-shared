// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: server_alpha/chat/chat.proto

package chat

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

// Validate checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ChatMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ChatMessageMultiError, or
// nil if none found.
func (m *ChatMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Message

	// no validation rules for CreatedAt

	if len(errors) > 0 {
		return ChatMessageMultiError(errors)
	}

	return nil
}

// ChatMessageMultiError is an error wrapping multiple validation errors
// returned by ChatMessage.ValidateAll() if the designated constraints aren't met.
type ChatMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatMessageMultiError) AllErrors() []error { return m }

// ChatMessageValidationError is the validation error returned by
// ChatMessage.Validate if the designated constraints aren't met.
type ChatMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatMessageValidationError) ErrorName() string { return "ChatMessageValidationError" }

// Error satisfies the builtin error interface
func (e ChatMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatMessageValidationError{}

// Validate checks the field values on Chat with the rules defined in the proto
// definition for this message. If any rules are violated, the first error
// encountered is returned, or nil if there are no violations.
func (m *Chat) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Chat with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ChatMultiError, or nil if none found.
func (m *Chat) ValidateAll() error {
	return m.validate(true)
}

func (m *Chat) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if all {
		switch v := interface{}(m.GetUser()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChatValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatValidationError{
					field:  "User",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ChatMultiError(errors)
	}

	return nil
}

// ChatMultiError is an error wrapping multiple validation errors returned by
// Chat.ValidateAll() if the designated constraints aren't met.
type ChatMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatMultiError) AllErrors() []error { return m }

// ChatValidationError is the validation error returned by Chat.Validate if the
// designated constraints aren't met.
type ChatValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatValidationError) ErrorName() string { return "ChatValidationError" }

// Error satisfies the builtin error interface
func (e ChatValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChat.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatValidationError{}

// Validate checks the field values on CreateChatRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CreateChatRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateChatRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateChatRequestMultiError, or nil if none found.
func (m *CreateChatRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateChatRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Username

	// no validation rules for Message

	if len(errors) > 0 {
		return CreateChatRequestMultiError(errors)
	}

	return nil
}

// CreateChatRequestMultiError is an error wrapping multiple validation errors
// returned by CreateChatRequest.ValidateAll() if the designated constraints
// aren't met.
type CreateChatRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateChatRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateChatRequestMultiError) AllErrors() []error { return m }

// CreateChatRequestValidationError is the validation error returned by
// CreateChatRequest.Validate if the designated constraints aren't met.
type CreateChatRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateChatRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateChatRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateChatRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateChatRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateChatRequestValidationError) ErrorName() string {
	return "CreateChatRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateChatRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateChatRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateChatRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateChatRequestValidationError{}

// Validate checks the field values on CreateChatResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *CreateChatResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CreateChatResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CreateChatResponseMultiError, or nil if none found.
func (m *CreateChatResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CreateChatResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ChatId

	if all {
		switch v := interface{}(m.GetMessage()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CreateChatResponseValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CreateChatResponseValidationError{
					field:  "Message",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMessage()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateChatResponseValidationError{
				field:  "Message",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CreateChatResponseMultiError(errors)
	}

	return nil
}

// CreateChatResponseMultiError is an error wrapping multiple validation errors
// returned by CreateChatResponse.ValidateAll() if the designated constraints
// aren't met.
type CreateChatResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CreateChatResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CreateChatResponseMultiError) AllErrors() []error { return m }

// CreateChatResponseValidationError is the validation error returned by
// CreateChatResponse.Validate if the designated constraints aren't met.
type CreateChatResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateChatResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateChatResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateChatResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateChatResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateChatResponseValidationError) ErrorName() string {
	return "CreateChatResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateChatResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateChatResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateChatResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateChatResponseValidationError{}

// Validate checks the field values on GetChatRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetChatRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetChatRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetChatRequestMultiError,
// or nil if none found.
func (m *GetChatRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetChatRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ChatId

	if all {
		switch v := interface{}(m.GetPagination()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetChatRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetChatRequestValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetChatRequestValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetChatRequestMultiError(errors)
	}

	return nil
}

// GetChatRequestMultiError is an error wrapping multiple validation errors
// returned by GetChatRequest.ValidateAll() if the designated constraints
// aren't met.
type GetChatRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetChatRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetChatRequestMultiError) AllErrors() []error { return m }

// GetChatRequestValidationError is the validation error returned by
// GetChatRequest.Validate if the designated constraints aren't met.
type GetChatRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetChatRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetChatRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetChatRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetChatRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetChatRequestValidationError) ErrorName() string { return "GetChatRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetChatRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetChatRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetChatRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetChatRequestValidationError{}

// Validate checks the field values on GetChatResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetChatResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetChatResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetChatResponseMultiError, or nil if none found.
func (m *GetChatResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetChatResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetMessages() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetChatResponseValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetChatResponseValidationError{
						field:  fmt.Sprintf("Messages[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetChatResponseValidationError{
					field:  fmt.Sprintf("Messages[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if all {
		switch v := interface{}(m.GetPagination()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetChatResponseValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetChatResponseValidationError{
					field:  "Pagination",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPagination()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetChatResponseValidationError{
				field:  "Pagination",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetChatResponseMultiError(errors)
	}

	return nil
}

// GetChatResponseMultiError is an error wrapping multiple validation errors
// returned by GetChatResponse.ValidateAll() if the designated constraints
// aren't met.
type GetChatResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetChatResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetChatResponseMultiError) AllErrors() []error { return m }

// GetChatResponseValidationError is the validation error returned by
// GetChatResponse.Validate if the designated constraints aren't met.
type GetChatResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetChatResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetChatResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetChatResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetChatResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetChatResponseValidationError) ErrorName() string { return "GetChatResponseValidationError" }

// Error satisfies the builtin error interface
func (e GetChatResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetChatResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetChatResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetChatResponseValidationError{}

// Validate checks the field values on ListChatsResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ListChatsResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListChatsResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ListChatsResponseMultiError, or nil if none found.
func (m *ListChatsResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListChatsResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetChats() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListChatsResponseValidationError{
						field:  fmt.Sprintf("Chats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListChatsResponseValidationError{
						field:  fmt.Sprintf("Chats[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListChatsResponseValidationError{
					field:  fmt.Sprintf("Chats[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return ListChatsResponseMultiError(errors)
	}

	return nil
}

// ListChatsResponseMultiError is an error wrapping multiple validation errors
// returned by ListChatsResponse.ValidateAll() if the designated constraints
// aren't met.
type ListChatsResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListChatsResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListChatsResponseMultiError) AllErrors() []error { return m }

// ListChatsResponseValidationError is the validation error returned by
// ListChatsResponse.Validate if the designated constraints aren't met.
type ListChatsResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListChatsResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListChatsResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListChatsResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListChatsResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListChatsResponseValidationError) ErrorName() string {
	return "ListChatsResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ListChatsResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListChatsResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListChatsResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListChatsResponseValidationError{}

// Validate checks the field values on PrepareChatStreamRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *PrepareChatStreamRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PrepareChatStreamRequest with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// PrepareChatStreamRequestMultiError, or nil if none found.
func (m *PrepareChatStreamRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PrepareChatStreamRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for ChatId

	if len(errors) > 0 {
		return PrepareChatStreamRequestMultiError(errors)
	}

	return nil
}

// PrepareChatStreamRequestMultiError is an error wrapping multiple validation
// errors returned by PrepareChatStreamRequest.ValidateAll() if the designated
// constraints aren't met.
type PrepareChatStreamRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PrepareChatStreamRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PrepareChatStreamRequestMultiError) AllErrors() []error { return m }

// PrepareChatStreamRequestValidationError is the validation error returned by
// PrepareChatStreamRequest.Validate if the designated constraints aren't met.
type PrepareChatStreamRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PrepareChatStreamRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PrepareChatStreamRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PrepareChatStreamRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PrepareChatStreamRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PrepareChatStreamRequestValidationError) ErrorName() string {
	return "PrepareChatStreamRequestValidationError"
}

// Error satisfies the builtin error interface
func (e PrepareChatStreamRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPrepareChatStreamRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PrepareChatStreamRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PrepareChatStreamRequestValidationError{}