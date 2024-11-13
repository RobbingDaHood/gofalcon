// Code generated by go-swagger; DO NOT EDIT.

package real_time_response

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/crowdstrike/gofalcon/falcon/models"
)

// RTRCheckCommandStatusReader is a Reader for the RTRCheckCommandStatus structure.
type RTRCheckCommandStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRCheckCommandStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRCheckCommandStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRTRCheckCommandStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRCheckCommandStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRCheckCommandStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRTRCheckCommandStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /real-time-response/entities/command/v1] RTR-CheckCommandStatus", response, response.Code())
	}
}

// NewRTRCheckCommandStatusOK creates a RTRCheckCommandStatusOK with default headers values
func NewRTRCheckCommandStatusOK() *RTRCheckCommandStatusOK {
	return &RTRCheckCommandStatusOK{}
}

/*
RTRCheckCommandStatusOK describes a response with status code 200, with default header values.

success
*/
type RTRCheckCommandStatusOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainStatusResponseWrapper
}

// IsSuccess returns true when this r t r check command status o k response has a 2xx status code
func (o *RTRCheckCommandStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r check command status o k response has a 3xx status code
func (o *RTRCheckCommandStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check command status o k response has a 4xx status code
func (o *RTRCheckCommandStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r check command status o k response has a 5xx status code
func (o *RTRCheckCommandStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r check command status o k response a status code equal to that given
func (o *RTRCheckCommandStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r check command status o k response
func (o *RTRCheckCommandStatusOK) Code() int {
	return 200
}

func (o *RTRCheckCommandStatusOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusOK  %+v", 200, o.Payload)
}

func (o *RTRCheckCommandStatusOK) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusOK  %+v", 200, o.Payload)
}

func (o *RTRCheckCommandStatusOK) GetPayload() *models.DomainStatusResponseWrapper {
	return o.Payload
}

func (o *RTRCheckCommandStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainStatusResponseWrapper)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckCommandStatusUnauthorized creates a RTRCheckCommandStatusUnauthorized with default headers values
func NewRTRCheckCommandStatusUnauthorized() *RTRCheckCommandStatusUnauthorized {
	return &RTRCheckCommandStatusUnauthorized{}
}

/*
RTRCheckCommandStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RTRCheckCommandStatusUnauthorized struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.DomainAPIError
}

// IsSuccess returns true when this r t r check command status unauthorized response has a 2xx status code
func (o *RTRCheckCommandStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r check command status unauthorized response has a 3xx status code
func (o *RTRCheckCommandStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check command status unauthorized response has a 4xx status code
func (o *RTRCheckCommandStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r check command status unauthorized response has a 5xx status code
func (o *RTRCheckCommandStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r check command status unauthorized response a status code equal to that given
func (o *RTRCheckCommandStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the r t r check command status unauthorized response
func (o *RTRCheckCommandStatusUnauthorized) Code() int {
	return 401
}

func (o *RTRCheckCommandStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *RTRCheckCommandStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *RTRCheckCommandStatusUnauthorized) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRCheckCommandStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.DomainAPIError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckCommandStatusForbidden creates a RTRCheckCommandStatusForbidden with default headers values
func NewRTRCheckCommandStatusForbidden() *RTRCheckCommandStatusForbidden {
	return &RTRCheckCommandStatusForbidden{}
}

/*
RTRCheckCommandStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRCheckCommandStatusForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r check command status forbidden response has a 2xx status code
func (o *RTRCheckCommandStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r check command status forbidden response has a 3xx status code
func (o *RTRCheckCommandStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check command status forbidden response has a 4xx status code
func (o *RTRCheckCommandStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r check command status forbidden response has a 5xx status code
func (o *RTRCheckCommandStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r check command status forbidden response a status code equal to that given
func (o *RTRCheckCommandStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r check command status forbidden response
func (o *RTRCheckCommandStatusForbidden) Code() int {
	return 403
}

func (o *RTRCheckCommandStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusForbidden  %+v", 403, o.Payload)
}

func (o *RTRCheckCommandStatusForbidden) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusForbidden  %+v", 403, o.Payload)
}

func (o *RTRCheckCommandStatusForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckCommandStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckCommandStatusTooManyRequests creates a RTRCheckCommandStatusTooManyRequests with default headers values
func NewRTRCheckCommandStatusTooManyRequests() *RTRCheckCommandStatusTooManyRequests {
	return &RTRCheckCommandStatusTooManyRequests{}
}

/*
RTRCheckCommandStatusTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRCheckCommandStatusTooManyRequests struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	/* Too many requests, retry after this time (as milliseconds since epoch)
	 */
	XRateLimitRetryAfter int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r check command status too many requests response has a 2xx status code
func (o *RTRCheckCommandStatusTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r check command status too many requests response has a 3xx status code
func (o *RTRCheckCommandStatusTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check command status too many requests response has a 4xx status code
func (o *RTRCheckCommandStatusTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r check command status too many requests response has a 5xx status code
func (o *RTRCheckCommandStatusTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r check command status too many requests response a status code equal to that given
func (o *RTRCheckCommandStatusTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r check command status too many requests response
func (o *RTRCheckCommandStatusTooManyRequests) Code() int {
	return 429
}

func (o *RTRCheckCommandStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRCheckCommandStatusTooManyRequests) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRCheckCommandStatusTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckCommandStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	// hydrates response header X-RateLimit-RetryAfter
	hdrXRateLimitRetryAfter := response.GetHeader("X-RateLimit-RetryAfter")

	if hdrXRateLimitRetryAfter != "" {
		valxRateLimitRetryAfter, err := swag.ConvertInt64(hdrXRateLimitRetryAfter)
		if err != nil {
			return errors.InvalidType("X-RateLimit-RetryAfter", "header", "int64", hdrXRateLimitRetryAfter)
		}
		o.XRateLimitRetryAfter = valxRateLimitRetryAfter
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRTRCheckCommandStatusInternalServerError creates a RTRCheckCommandStatusInternalServerError with default headers values
func NewRTRCheckCommandStatusInternalServerError() *RTRCheckCommandStatusInternalServerError {
	return &RTRCheckCommandStatusInternalServerError{}
}

/*
RTRCheckCommandStatusInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type RTRCheckCommandStatusInternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaReplyMetaOnly
}

// IsSuccess returns true when this r t r check command status internal server error response has a 2xx status code
func (o *RTRCheckCommandStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r check command status internal server error response has a 3xx status code
func (o *RTRCheckCommandStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check command status internal server error response has a 4xx status code
func (o *RTRCheckCommandStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r check command status internal server error response has a 5xx status code
func (o *RTRCheckCommandStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this r t r check command status internal server error response a status code equal to that given
func (o *RTRCheckCommandStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the r t r check command status internal server error response
func (o *RTRCheckCommandStatusInternalServerError) Code() int {
	return 500
}

func (o *RTRCheckCommandStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRCheckCommandStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/command/v1][%d] rTRCheckCommandStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRCheckCommandStatusInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckCommandStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-CS-TRACEID
	hdrXCSTRACEID := response.GetHeader("X-CS-TRACEID")

	if hdrXCSTRACEID != "" {
		o.XCSTRACEID = hdrXCSTRACEID
	}

	// hydrates response header X-RateLimit-Limit
	hdrXRateLimitLimit := response.GetHeader("X-RateLimit-Limit")

	if hdrXRateLimitLimit != "" {
		valxRateLimitLimit, err := swag.ConvertInt64(hdrXRateLimitLimit)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Limit", "header", "int64", hdrXRateLimitLimit)
		}
		o.XRateLimitLimit = valxRateLimitLimit
	}

	// hydrates response header X-RateLimit-Remaining
	hdrXRateLimitRemaining := response.GetHeader("X-RateLimit-Remaining")

	if hdrXRateLimitRemaining != "" {
		valxRateLimitRemaining, err := swag.ConvertInt64(hdrXRateLimitRemaining)
		if err != nil {
			return errors.InvalidType("X-RateLimit-Remaining", "header", "int64", hdrXRateLimitRemaining)
		}
		o.XRateLimitRemaining = valxRateLimitRemaining
	}

	o.Payload = new(models.MsaReplyMetaOnly)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
