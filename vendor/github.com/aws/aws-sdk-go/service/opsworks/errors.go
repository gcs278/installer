// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package opsworks

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Indicates that a resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// Indicates that a request was not valid.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"ValidationException":       newErrorValidationException,
}