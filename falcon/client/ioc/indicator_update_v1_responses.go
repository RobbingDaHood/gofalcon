// Code generated by go-swagger; DO NOT EDIT.

package ioc

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

// IndicatorUpdateV1Reader is a Reader for the IndicatorUpdateV1 structure.
type IndicatorUpdateV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndicatorUpdateV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIndicatorUpdateV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewIndicatorUpdateV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewIndicatorUpdateV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewIndicatorUpdateV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /iocs/entities/indicators/v1] indicator.update.v1", response, response.Code())
	}
}

// NewIndicatorUpdateV1OK creates a IndicatorUpdateV1OK with default headers values
func NewIndicatorUpdateV1OK() *IndicatorUpdateV1OK {
	return &IndicatorUpdateV1OK{}
}

/*
IndicatorUpdateV1OK describes a response with status code 200, with default header values.

OK
*/
type IndicatorUpdateV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIIndicatorRespV1
}

// IsSuccess returns true when this indicator update v1 o k response has a 2xx status code
func (o *IndicatorUpdateV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this indicator update v1 o k response has a 3xx status code
func (o *IndicatorUpdateV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this indicator update v1 o k response has a 4xx status code
func (o *IndicatorUpdateV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this indicator update v1 o k response has a 5xx status code
func (o *IndicatorUpdateV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this indicator update v1 o k response a status code equal to that given
func (o *IndicatorUpdateV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the indicator update v1 o k response
func (o *IndicatorUpdateV1OK) Code() int {
	return 200
}

func (o *IndicatorUpdateV1OK) Error() string {
	return fmt.Sprintf("[PATCH /iocs/entities/indicators/v1][%d] indicatorUpdateV1OK  %+v", 200, o.Payload)
}

func (o *IndicatorUpdateV1OK) String() string {
	return fmt.Sprintf("[PATCH /iocs/entities/indicators/v1][%d] indicatorUpdateV1OK  %+v", 200, o.Payload)
}

func (o *IndicatorUpdateV1OK) GetPayload() *models.APIIndicatorRespV1 {
	return o.Payload
}

func (o *IndicatorUpdateV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIIndicatorRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIndicatorUpdateV1Forbidden creates a IndicatorUpdateV1Forbidden with default headers values
func NewIndicatorUpdateV1Forbidden() *IndicatorUpdateV1Forbidden {
	return &IndicatorUpdateV1Forbidden{}
}

/*
IndicatorUpdateV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type IndicatorUpdateV1Forbidden struct {

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

// IsSuccess returns true when this indicator update v1 forbidden response has a 2xx status code
func (o *IndicatorUpdateV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this indicator update v1 forbidden response has a 3xx status code
func (o *IndicatorUpdateV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this indicator update v1 forbidden response has a 4xx status code
func (o *IndicatorUpdateV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this indicator update v1 forbidden response has a 5xx status code
func (o *IndicatorUpdateV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this indicator update v1 forbidden response a status code equal to that given
func (o *IndicatorUpdateV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the indicator update v1 forbidden response
func (o *IndicatorUpdateV1Forbidden) Code() int {
	return 403
}

func (o *IndicatorUpdateV1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /iocs/entities/indicators/v1][%d] indicatorUpdateV1Forbidden  %+v", 403, o.Payload)
}

func (o *IndicatorUpdateV1Forbidden) String() string {
	return fmt.Sprintf("[PATCH /iocs/entities/indicators/v1][%d] indicatorUpdateV1Forbidden  %+v", 403, o.Payload)
}

func (o *IndicatorUpdateV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *IndicatorUpdateV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIndicatorUpdateV1TooManyRequests creates a IndicatorUpdateV1TooManyRequests with default headers values
func NewIndicatorUpdateV1TooManyRequests() *IndicatorUpdateV1TooManyRequests {
	return &IndicatorUpdateV1TooManyRequests{}
}

/*
IndicatorUpdateV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type IndicatorUpdateV1TooManyRequests struct {

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

// IsSuccess returns true when this indicator update v1 too many requests response has a 2xx status code
func (o *IndicatorUpdateV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this indicator update v1 too many requests response has a 3xx status code
func (o *IndicatorUpdateV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this indicator update v1 too many requests response has a 4xx status code
func (o *IndicatorUpdateV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this indicator update v1 too many requests response has a 5xx status code
func (o *IndicatorUpdateV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this indicator update v1 too many requests response a status code equal to that given
func (o *IndicatorUpdateV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the indicator update v1 too many requests response
func (o *IndicatorUpdateV1TooManyRequests) Code() int {
	return 429
}

func (o *IndicatorUpdateV1TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /iocs/entities/indicators/v1][%d] indicatorUpdateV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *IndicatorUpdateV1TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /iocs/entities/indicators/v1][%d] indicatorUpdateV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *IndicatorUpdateV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *IndicatorUpdateV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewIndicatorUpdateV1InternalServerError creates a IndicatorUpdateV1InternalServerError with default headers values
func NewIndicatorUpdateV1InternalServerError() *IndicatorUpdateV1InternalServerError {
	return &IndicatorUpdateV1InternalServerError{}
}

/*
IndicatorUpdateV1InternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type IndicatorUpdateV1InternalServerError struct {

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

// IsSuccess returns true when this indicator update v1 internal server error response has a 2xx status code
func (o *IndicatorUpdateV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this indicator update v1 internal server error response has a 3xx status code
func (o *IndicatorUpdateV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this indicator update v1 internal server error response has a 4xx status code
func (o *IndicatorUpdateV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this indicator update v1 internal server error response has a 5xx status code
func (o *IndicatorUpdateV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this indicator update v1 internal server error response a status code equal to that given
func (o *IndicatorUpdateV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the indicator update v1 internal server error response
func (o *IndicatorUpdateV1InternalServerError) Code() int {
	return 500
}

func (o *IndicatorUpdateV1InternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /iocs/entities/indicators/v1][%d] indicatorUpdateV1InternalServerError  %+v", 500, o.Payload)
}

func (o *IndicatorUpdateV1InternalServerError) String() string {
	return fmt.Sprintf("[PATCH /iocs/entities/indicators/v1][%d] indicatorUpdateV1InternalServerError  %+v", 500, o.Payload)
}

func (o *IndicatorUpdateV1InternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *IndicatorUpdateV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
