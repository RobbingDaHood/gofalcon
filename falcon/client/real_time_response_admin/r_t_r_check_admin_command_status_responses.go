// Code generated by go-swagger; DO NOT EDIT.

package real_time_response_admin

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

// RTRCheckAdminCommandStatusReader is a Reader for the RTRCheckAdminCommandStatus structure.
type RTRCheckAdminCommandStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRCheckAdminCommandStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRTRCheckAdminCommandStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewRTRCheckAdminCommandStatusUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRCheckAdminCommandStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRCheckAdminCommandStatusTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRTRCheckAdminCommandStatusInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /real-time-response/entities/admin-command/v1] RTR-CheckAdminCommandStatus", response, response.Code())
	}
}

// NewRTRCheckAdminCommandStatusOK creates a RTRCheckAdminCommandStatusOK with default headers values
func NewRTRCheckAdminCommandStatusOK() *RTRCheckAdminCommandStatusOK {
	return &RTRCheckAdminCommandStatusOK{}
}

/*
RTRCheckAdminCommandStatusOK describes a response with status code 200, with default header values.

success
*/
type RTRCheckAdminCommandStatusOK struct {

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

// IsSuccess returns true when this r t r check admin command status o k response has a 2xx status code
func (o *RTRCheckAdminCommandStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r check admin command status o k response has a 3xx status code
func (o *RTRCheckAdminCommandStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check admin command status o k response has a 4xx status code
func (o *RTRCheckAdminCommandStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r check admin command status o k response has a 5xx status code
func (o *RTRCheckAdminCommandStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r check admin command status o k response a status code equal to that given
func (o *RTRCheckAdminCommandStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the r t r check admin command status o k response
func (o *RTRCheckAdminCommandStatusOK) Code() int {
	return 200
}

func (o *RTRCheckAdminCommandStatusOK) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusOK  %+v", 200, o.Payload)
}

func (o *RTRCheckAdminCommandStatusOK) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusOK  %+v", 200, o.Payload)
}

func (o *RTRCheckAdminCommandStatusOK) GetPayload() *models.DomainStatusResponseWrapper {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCheckAdminCommandStatusUnauthorized creates a RTRCheckAdminCommandStatusUnauthorized with default headers values
func NewRTRCheckAdminCommandStatusUnauthorized() *RTRCheckAdminCommandStatusUnauthorized {
	return &RTRCheckAdminCommandStatusUnauthorized{}
}

/*
RTRCheckAdminCommandStatusUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type RTRCheckAdminCommandStatusUnauthorized struct {

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

// IsSuccess returns true when this r t r check admin command status unauthorized response has a 2xx status code
func (o *RTRCheckAdminCommandStatusUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r check admin command status unauthorized response has a 3xx status code
func (o *RTRCheckAdminCommandStatusUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check admin command status unauthorized response has a 4xx status code
func (o *RTRCheckAdminCommandStatusUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r check admin command status unauthorized response has a 5xx status code
func (o *RTRCheckAdminCommandStatusUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r check admin command status unauthorized response a status code equal to that given
func (o *RTRCheckAdminCommandStatusUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the r t r check admin command status unauthorized response
func (o *RTRCheckAdminCommandStatusUnauthorized) Code() int {
	return 401
}

func (o *RTRCheckAdminCommandStatusUnauthorized) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *RTRCheckAdminCommandStatusUnauthorized) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusUnauthorized  %+v", 401, o.Payload)
}

func (o *RTRCheckAdminCommandStatusUnauthorized) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCheckAdminCommandStatusForbidden creates a RTRCheckAdminCommandStatusForbidden with default headers values
func NewRTRCheckAdminCommandStatusForbidden() *RTRCheckAdminCommandStatusForbidden {
	return &RTRCheckAdminCommandStatusForbidden{}
}

/*
RTRCheckAdminCommandStatusForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRCheckAdminCommandStatusForbidden struct {

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

// IsSuccess returns true when this r t r check admin command status forbidden response has a 2xx status code
func (o *RTRCheckAdminCommandStatusForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r check admin command status forbidden response has a 3xx status code
func (o *RTRCheckAdminCommandStatusForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check admin command status forbidden response has a 4xx status code
func (o *RTRCheckAdminCommandStatusForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r check admin command status forbidden response has a 5xx status code
func (o *RTRCheckAdminCommandStatusForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r check admin command status forbidden response a status code equal to that given
func (o *RTRCheckAdminCommandStatusForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r check admin command status forbidden response
func (o *RTRCheckAdminCommandStatusForbidden) Code() int {
	return 403
}

func (o *RTRCheckAdminCommandStatusForbidden) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusForbidden  %+v", 403, o.Payload)
}

func (o *RTRCheckAdminCommandStatusForbidden) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusForbidden  %+v", 403, o.Payload)
}

func (o *RTRCheckAdminCommandStatusForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCheckAdminCommandStatusTooManyRequests creates a RTRCheckAdminCommandStatusTooManyRequests with default headers values
func NewRTRCheckAdminCommandStatusTooManyRequests() *RTRCheckAdminCommandStatusTooManyRequests {
	return &RTRCheckAdminCommandStatusTooManyRequests{}
}

/*
RTRCheckAdminCommandStatusTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRCheckAdminCommandStatusTooManyRequests struct {

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

// IsSuccess returns true when this r t r check admin command status too many requests response has a 2xx status code
func (o *RTRCheckAdminCommandStatusTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r check admin command status too many requests response has a 3xx status code
func (o *RTRCheckAdminCommandStatusTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check admin command status too many requests response has a 4xx status code
func (o *RTRCheckAdminCommandStatusTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r check admin command status too many requests response has a 5xx status code
func (o *RTRCheckAdminCommandStatusTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r check admin command status too many requests response a status code equal to that given
func (o *RTRCheckAdminCommandStatusTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r check admin command status too many requests response
func (o *RTRCheckAdminCommandStatusTooManyRequests) Code() int {
	return 429
}

func (o *RTRCheckAdminCommandStatusTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRCheckAdminCommandStatusTooManyRequests) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRCheckAdminCommandStatusTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRCheckAdminCommandStatusInternalServerError creates a RTRCheckAdminCommandStatusInternalServerError with default headers values
func NewRTRCheckAdminCommandStatusInternalServerError() *RTRCheckAdminCommandStatusInternalServerError {
	return &RTRCheckAdminCommandStatusInternalServerError{}
}

/*
RTRCheckAdminCommandStatusInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type RTRCheckAdminCommandStatusInternalServerError struct {

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

// IsSuccess returns true when this r t r check admin command status internal server error response has a 2xx status code
func (o *RTRCheckAdminCommandStatusInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r check admin command status internal server error response has a 3xx status code
func (o *RTRCheckAdminCommandStatusInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r check admin command status internal server error response has a 4xx status code
func (o *RTRCheckAdminCommandStatusInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r check admin command status internal server error response has a 5xx status code
func (o *RTRCheckAdminCommandStatusInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this r t r check admin command status internal server error response a status code equal to that given
func (o *RTRCheckAdminCommandStatusInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the r t r check admin command status internal server error response
func (o *RTRCheckAdminCommandStatusInternalServerError) Code() int {
	return 500
}

func (o *RTRCheckAdminCommandStatusInternalServerError) Error() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRCheckAdminCommandStatusInternalServerError) String() string {
	return fmt.Sprintf("[GET /real-time-response/entities/admin-command/v1][%d] rTRCheckAdminCommandStatusInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRCheckAdminCommandStatusInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRCheckAdminCommandStatusInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
