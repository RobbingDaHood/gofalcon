// Code generated by go-swagger; DO NOT EDIT.

package ods

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

// AggregateScansReader is a Reader for the AggregateScans structure.
type AggregateScansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AggregateScansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAggregateScansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewAggregateScansForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAggregateScansNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewAggregateScansTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAggregateScansInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /ods/aggregates/scans/v1] aggregate-scans", response, response.Code())
	}
}

// NewAggregateScansOK creates a AggregateScansOK with default headers values
func NewAggregateScansOK() *AggregateScansOK {
	return &AggregateScansOK{}
}

/*
AggregateScansOK describes a response with status code 200, with default header values.

OK
*/
type AggregateScansOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaAggregatesResponse
}

// IsSuccess returns true when this aggregate scans o k response has a 2xx status code
func (o *AggregateScansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this aggregate scans o k response has a 3xx status code
func (o *AggregateScansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate scans o k response has a 4xx status code
func (o *AggregateScansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate scans o k response has a 5xx status code
func (o *AggregateScansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate scans o k response a status code equal to that given
func (o *AggregateScansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the aggregate scans o k response
func (o *AggregateScansOK) Code() int {
	return 200
}

func (o *AggregateScansOK) Error() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansOK  %+v", 200, o.Payload)
}

func (o *AggregateScansOK) String() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansOK  %+v", 200, o.Payload)
}

func (o *AggregateScansOK) GetPayload() *models.MsaAggregatesResponse {
	return o.Payload
}

func (o *AggregateScansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaAggregatesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateScansForbidden creates a AggregateScansForbidden with default headers values
func NewAggregateScansForbidden() *AggregateScansForbidden {
	return &AggregateScansForbidden{}
}

/*
AggregateScansForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AggregateScansForbidden struct {

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

// IsSuccess returns true when this aggregate scans forbidden response has a 2xx status code
func (o *AggregateScansForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate scans forbidden response has a 3xx status code
func (o *AggregateScansForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate scans forbidden response has a 4xx status code
func (o *AggregateScansForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate scans forbidden response has a 5xx status code
func (o *AggregateScansForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate scans forbidden response a status code equal to that given
func (o *AggregateScansForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the aggregate scans forbidden response
func (o *AggregateScansForbidden) Code() int {
	return 403
}

func (o *AggregateScansForbidden) Error() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansForbidden  %+v", 403, o.Payload)
}

func (o *AggregateScansForbidden) String() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansForbidden  %+v", 403, o.Payload)
}

func (o *AggregateScansForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateScansForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateScansNotFound creates a AggregateScansNotFound with default headers values
func NewAggregateScansNotFound() *AggregateScansNotFound {
	return &AggregateScansNotFound{}
}

/*
AggregateScansNotFound describes a response with status code 404, with default header values.

Not Found
*/
type AggregateScansNotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaspecResponseFields
}

// IsSuccess returns true when this aggregate scans not found response has a 2xx status code
func (o *AggregateScansNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate scans not found response has a 3xx status code
func (o *AggregateScansNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate scans not found response has a 4xx status code
func (o *AggregateScansNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate scans not found response has a 5xx status code
func (o *AggregateScansNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate scans not found response a status code equal to that given
func (o *AggregateScansNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the aggregate scans not found response
func (o *AggregateScansNotFound) Code() int {
	return 404
}

func (o *AggregateScansNotFound) Error() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansNotFound  %+v", 404, o.Payload)
}

func (o *AggregateScansNotFound) String() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansNotFound  %+v", 404, o.Payload)
}

func (o *AggregateScansNotFound) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *AggregateScansNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaspecResponseFields)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAggregateScansTooManyRequests creates a AggregateScansTooManyRequests with default headers values
func NewAggregateScansTooManyRequests() *AggregateScansTooManyRequests {
	return &AggregateScansTooManyRequests{}
}

/*
AggregateScansTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type AggregateScansTooManyRequests struct {

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

// IsSuccess returns true when this aggregate scans too many requests response has a 2xx status code
func (o *AggregateScansTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate scans too many requests response has a 3xx status code
func (o *AggregateScansTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate scans too many requests response has a 4xx status code
func (o *AggregateScansTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this aggregate scans too many requests response has a 5xx status code
func (o *AggregateScansTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this aggregate scans too many requests response a status code equal to that given
func (o *AggregateScansTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the aggregate scans too many requests response
func (o *AggregateScansTooManyRequests) Code() int {
	return 429
}

func (o *AggregateScansTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateScansTooManyRequests) String() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansTooManyRequests  %+v", 429, o.Payload)
}

func (o *AggregateScansTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateScansTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAggregateScansInternalServerError creates a AggregateScansInternalServerError with default headers values
func NewAggregateScansInternalServerError() *AggregateScansInternalServerError {
	return &AggregateScansInternalServerError{}
}

/*
AggregateScansInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type AggregateScansInternalServerError struct {

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

// IsSuccess returns true when this aggregate scans internal server error response has a 2xx status code
func (o *AggregateScansInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this aggregate scans internal server error response has a 3xx status code
func (o *AggregateScansInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this aggregate scans internal server error response has a 4xx status code
func (o *AggregateScansInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this aggregate scans internal server error response has a 5xx status code
func (o *AggregateScansInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this aggregate scans internal server error response a status code equal to that given
func (o *AggregateScansInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the aggregate scans internal server error response
func (o *AggregateScansInternalServerError) Code() int {
	return 500
}

func (o *AggregateScansInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateScansInternalServerError) String() string {
	return fmt.Sprintf("[POST /ods/aggregates/scans/v1][%d] aggregateScansInternalServerError  %+v", 500, o.Payload)
}

func (o *AggregateScansInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *AggregateScansInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
