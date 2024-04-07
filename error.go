package modbus

import (
	"fmt"
)

type ExceptionCodeError interface {
	GetExceptionCode() uint8
	error
}
type ConfigurationError struct {
	msg string
}

func (e *ConfigurationError) Error() string {
	result := "configuration error"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func NewConfigurationError(details string, args ...interface{}) *ConfigurationError {
	return &ConfigurationError{msg: fmt.Sprintf(details, args...)}
}

type RequestTimedOutError struct {
	msg string
}

func (e *RequestTimedOutError) Error() string {
	result := "request timed out"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func NewRequestTimedOutError(details string, args ...interface{}) *RequestTimedOutError {
	return &RequestTimedOutError{msg: fmt.Sprintf(details, args...)}
}

type IllegalFunctionError struct {
	msg string
}

func (e *IllegalFunctionError) Error() string {
	result := "illegal function code"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *IllegalFunctionError) GetExceptionCode() uint8 {
	return exIllegalFunction
}

func NewIllegalFunctionError(details string, args ...interface{}) *IllegalFunctionError {
	return &IllegalFunctionError{msg: fmt.Sprintf(details, args...)}
}

type IllegalDataAddressError struct {
	msg string
}

func (e *IllegalDataAddressError) Error() string {
	result := "illegal data address"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *IllegalDataAddressError) GetExceptionCode() uint8 {
	return exIllegalDataAddress
}

func NewIllegalDataAddressError(details string, args ...interface{}) *IllegalDataAddressError {
	return &IllegalDataAddressError{msg: fmt.Sprintf(details, args...)}
}

type IllegalDataValueError struct {
	msg string
}

func (e *IllegalDataValueError) Error() string {
	result := "illegal data value"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *IllegalDataValueError) GetExceptionCode() uint8 {
	return exIllegalDataValue
}

func NewIllegalDataValueError(details string, args ...interface{}) *IllegalDataValueError {
	return &IllegalDataValueError{msg: fmt.Sprintf(details, args...)}
}

type ServerDeviceFailureError struct {
	msg string
}

func (e *ServerDeviceFailureError) Error() string {
	result := "server device failure"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *ServerDeviceFailureError) GetExceptionCode() uint8 {
	return exServerDeviceFailure
}

func NewServerDeviceFailureError(details string, args ...interface{}) *ServerDeviceFailureError {
	return &ServerDeviceFailureError{msg: fmt.Sprintf(details, args...)}
}

type AcknowledgeError struct {
	msg string
}

func (e *AcknowledgeError) Error() string {
	result := "request not acknowledged"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *AcknowledgeError) GetExceptionCode() uint8 {
	return exAcknowledge
}

func NewAcknowledgeError(details string, args ...interface{}) *AcknowledgeError {
	return &AcknowledgeError{msg: fmt.Sprintf(details, args...)}
}

type ServerDeviceBusyError struct {
	msg string
}

func (e *ServerDeviceBusyError) Error() string {
	result := "server device busy"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *ServerDeviceBusyError) GetExceptionCode() uint8 {
	return exServerDeviceBusy
}

func NewServerDeviceBusyError(details string, args ...interface{}) *ServerDeviceBusyError {
	return &ServerDeviceBusyError{msg: fmt.Sprintf(details, args...)}
}

type MemoryParityError struct {
	msg string
}

func (e *MemoryParityError) Error() string {
	result := "memory parity error"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *MemoryParityError) GetExceptionCode() uint8 {
	return exMemoryParityError
}

func NewMemoryParityError(details string, args ...interface{}) *MemoryParityError {
	return &MemoryParityError{msg: fmt.Sprintf(details, args...)}
}

type GatewayPathUnavailableError struct {
	msg string
}

func (e *GatewayPathUnavailableError) Error() string {
	result := "gateway path unavailable"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *GatewayPathUnavailableError) GetExceptionCode() uint8 {
	return exGWPathUnavailable
}

func NewGatewayPathUnavailableError(details string, args ...interface{}) *GatewayPathUnavailableError {
	return &GatewayPathUnavailableError{msg: fmt.Sprintf(details, args...)}
}

type GatewayTargetFailedToRespondError struct {
	msg string
}

func (e *GatewayTargetFailedToRespondError) Error() string {
	result := "gateway target device failed to respond"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *ServerDeviceBusyError) GatewayTargetFailedToRespondError() uint8 {
	return exGWTargetFailedToRespond
}

func NewGatewayTargetFailedToRespondError(details string, args ...interface{}) *ServerDeviceBusyError {
	return &ServerDeviceBusyError{msg: fmt.Sprintf(details, args...)}
}

type BadCRCError struct {
	msg string
}

func (e *BadCRCError) Error() string {
	result := "bad crc"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *BadCRCError) GetExceptionCode() uint8 {
	return exIllegalDataAddress
}

func NewBadCRCError(details string, args ...interface{}) *BadCRCError {
	return &BadCRCError{msg: fmt.Sprintf(details, args...)}
}

type ShortFrameError struct {
	msg string
}

func (e *ShortFrameError) Error() string {
	result := "short frame"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *ShortFrameError) GetExceptionCode() uint8 {
	return exIllegalDataAddress
}

func NewShortFrameError(details string, args ...interface{}) *ShortFrameError {
	return &ShortFrameError{msg: fmt.Sprintf(details, args...)}
}

type ProtocolError struct {
	msg string
}

func (e *ProtocolError) Error() string {
	result := "protocol error"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *ProtocolError) GetExceptionCode() uint8 {
	return exIllegalDataAddress
}

func NewProtocolError(details string, args ...interface{}) *ProtocolError {
	return &ProtocolError{msg: fmt.Sprintf(details, args...)}
}

type BadUnitIdError struct {
	msg string
}

func (e *BadUnitIdError) Error() string {
	result := "bad unit id"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *BadUnitIdError) GetExceptionCode() uint8 {
	return exIllegalDataAddress
}

func NewBadUnitIdError(details string, args ...interface{}) *BadUnitIdError {
	return &BadUnitIdError{msg: fmt.Sprintf(details, args...)}
}

type BadTransactionIdError struct {
	msg string
}

func (e *BadTransactionIdError) Error() string {
	result := "bad transaction id"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *BadTransactionIdError) GetExceptionCode() uint8 {
	return exIllegalDataAddress
}

func NewBadTransactionIdError(details string, args ...interface{}) *BadTransactionIdError {
	return &BadTransactionIdError{msg: fmt.Sprintf(details, args...)}
}

type UnknownProtocolIdError struct {
	msg string
}

func (e *UnknownProtocolIdError) Error() string {
	result := "unknown protocol identifier"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *UnknownProtocolIdError) GetExceptionCode() uint8 {
	return exIllegalDataAddress
}

func NewUnknownProtocolIdError(details string, args ...interface{}) *UnknownProtocolIdError {
	return &UnknownProtocolIdError{msg: fmt.Sprintf(details, args...)}
}

type UnexpectedParametersError struct {
	msg string
}

func (e *UnexpectedParametersError) Error() string {
	result := "unexpected parameters"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *UnexpectedParametersError) GetExceptionCode() uint8 {
	return exIllegalDataAddress
}

func NewUnexpectedParametersError(details string, args ...interface{}) *UnexpectedParametersError {
	return &UnexpectedParametersError{msg: fmt.Sprintf(details, args...)}
}

type UnknownExceptionCodeError struct {
	msg string
}

func (e *UnknownExceptionCodeError) Error() string {
	result := "unknwon exception code"
	if e.msg != "" {
		result += ", " + e.msg
	}
	return result
}

func (e *UnknownExceptionCodeError) GetExceptionCode() uint8 {
	return 255
}

func NewUnknownExceptionCodeError(details string, args ...interface{}) *UnknownExceptionCodeError {
	return &UnknownExceptionCodeError{msg: fmt.Sprintf(details, args...)}
}

// mapExceptionCodeToError turns a modbus exception code into a higher level Error object.
func mapExceptionCodeToError(exceptionCode uint8, details ...interface{}) (err ExceptionCodeError) {
	formatString := "no further details available"
	if len(details) > 0 {
		formatString = fmt.Sprintf("%v", details[0])
		details = details[1:]
	}
	switch exceptionCode {
	case exIllegalFunction:
		err = NewIllegalFunctionError(formatString, details)
	case exIllegalDataAddress:
		err = NewIllegalDataAddressError(formatString, details)
	case exIllegalDataValue:
		err = NewIllegalDataValueError(formatString, details)
	case exServerDeviceFailure:
		err = NewServerDeviceFailureError(formatString, details)
	case exAcknowledge:
		err = NewAcknowledgeError(formatString, details)
	case exMemoryParityError:
		err = NewMemoryParityError(formatString, details)
	case exServerDeviceBusy:
		err = NewServerDeviceBusyError(formatString, details)
	case exGWPathUnavailable:
		err = NewGatewayPathUnavailableError(formatString, details)
	case exGWTargetFailedToRespond:
		err = NewGatewayTargetFailedToRespondError(formatString, details)
	default:
		err = NewUnknownExceptionCodeError("unknown exception code (%v)", exceptionCode)
	}

	return
}
