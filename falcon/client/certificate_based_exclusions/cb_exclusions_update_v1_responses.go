// Code generated by go-swagger; DO NOT EDIT.

package certificate_based_exclusions

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

// CbExclusionsUpdateV1Reader is a Reader for the CbExclusionsUpdateV1 structure.
type CbExclusionsUpdateV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CbExclusionsUpdateV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCbExclusionsUpdateV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCbExclusionsUpdateV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCbExclusionsUpdateV1Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCbExclusionsUpdateV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCbExclusionsUpdateV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /exclusions/entities/cert-based-exclusions/v1] cb-exclusions.update.v1", response, response.Code())
	}
}

// NewCbExclusionsUpdateV1OK creates a CbExclusionsUpdateV1OK with default headers values
func NewCbExclusionsUpdateV1OK() *CbExclusionsUpdateV1OK {
	return &CbExclusionsUpdateV1OK{}
}

/*
CbExclusionsUpdateV1OK describes a response with status code 200, with default header values.

OK
*/
type CbExclusionsUpdateV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APICertBasedExclusionRespV1
}

// IsSuccess returns true when this cb exclusions update v1 o k response has a 2xx status code
func (o *CbExclusionsUpdateV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cb exclusions update v1 o k response has a 3xx status code
func (o *CbExclusionsUpdateV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions update v1 o k response has a 4xx status code
func (o *CbExclusionsUpdateV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cb exclusions update v1 o k response has a 5xx status code
func (o *CbExclusionsUpdateV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions update v1 o k response a status code equal to that given
func (o *CbExclusionsUpdateV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cb exclusions update v1 o k response
func (o *CbExclusionsUpdateV1OK) Code() int {
	return 200
}

func (o *CbExclusionsUpdateV1OK) Error() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1OK  %+v", 200, o.Payload)
}

func (o *CbExclusionsUpdateV1OK) String() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1OK  %+v", 200, o.Payload)
}

func (o *CbExclusionsUpdateV1OK) GetPayload() *models.APICertBasedExclusionRespV1 {
	return o.Payload
}

func (o *CbExclusionsUpdateV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APICertBasedExclusionRespV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCbExclusionsUpdateV1BadRequest creates a CbExclusionsUpdateV1BadRequest with default headers values
func NewCbExclusionsUpdateV1BadRequest() *CbExclusionsUpdateV1BadRequest {
	return &CbExclusionsUpdateV1BadRequest{}
}

/*
CbExclusionsUpdateV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CbExclusionsUpdateV1BadRequest struct {

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

// IsSuccess returns true when this cb exclusions update v1 bad request response has a 2xx status code
func (o *CbExclusionsUpdateV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cb exclusions update v1 bad request response has a 3xx status code
func (o *CbExclusionsUpdateV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions update v1 bad request response has a 4xx status code
func (o *CbExclusionsUpdateV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this cb exclusions update v1 bad request response has a 5xx status code
func (o *CbExclusionsUpdateV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions update v1 bad request response a status code equal to that given
func (o *CbExclusionsUpdateV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the cb exclusions update v1 bad request response
func (o *CbExclusionsUpdateV1BadRequest) Code() int {
	return 400
}

func (o *CbExclusionsUpdateV1BadRequest) Error() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1BadRequest  %+v", 400, o.Payload)
}

func (o *CbExclusionsUpdateV1BadRequest) String() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1BadRequest  %+v", 400, o.Payload)
}

func (o *CbExclusionsUpdateV1BadRequest) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CbExclusionsUpdateV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCbExclusionsUpdateV1Unauthorized creates a CbExclusionsUpdateV1Unauthorized with default headers values
func NewCbExclusionsUpdateV1Unauthorized() *CbExclusionsUpdateV1Unauthorized {
	return &CbExclusionsUpdateV1Unauthorized{}
}

/*
CbExclusionsUpdateV1Unauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CbExclusionsUpdateV1Unauthorized struct {

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

// IsSuccess returns true when this cb exclusions update v1 unauthorized response has a 2xx status code
func (o *CbExclusionsUpdateV1Unauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cb exclusions update v1 unauthorized response has a 3xx status code
func (o *CbExclusionsUpdateV1Unauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions update v1 unauthorized response has a 4xx status code
func (o *CbExclusionsUpdateV1Unauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this cb exclusions update v1 unauthorized response has a 5xx status code
func (o *CbExclusionsUpdateV1Unauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions update v1 unauthorized response a status code equal to that given
func (o *CbExclusionsUpdateV1Unauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the cb exclusions update v1 unauthorized response
func (o *CbExclusionsUpdateV1Unauthorized) Code() int {
	return 401
}

func (o *CbExclusionsUpdateV1Unauthorized) Error() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1Unauthorized  %+v", 401, o.Payload)
}

func (o *CbExclusionsUpdateV1Unauthorized) String() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1Unauthorized  %+v", 401, o.Payload)
}

func (o *CbExclusionsUpdateV1Unauthorized) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CbExclusionsUpdateV1Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCbExclusionsUpdateV1Forbidden creates a CbExclusionsUpdateV1Forbidden with default headers values
func NewCbExclusionsUpdateV1Forbidden() *CbExclusionsUpdateV1Forbidden {
	return &CbExclusionsUpdateV1Forbidden{}
}

/*
CbExclusionsUpdateV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CbExclusionsUpdateV1Forbidden struct {

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

// IsSuccess returns true when this cb exclusions update v1 forbidden response has a 2xx status code
func (o *CbExclusionsUpdateV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cb exclusions update v1 forbidden response has a 3xx status code
func (o *CbExclusionsUpdateV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions update v1 forbidden response has a 4xx status code
func (o *CbExclusionsUpdateV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this cb exclusions update v1 forbidden response has a 5xx status code
func (o *CbExclusionsUpdateV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions update v1 forbidden response a status code equal to that given
func (o *CbExclusionsUpdateV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the cb exclusions update v1 forbidden response
func (o *CbExclusionsUpdateV1Forbidden) Code() int {
	return 403
}

func (o *CbExclusionsUpdateV1Forbidden) Error() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1Forbidden  %+v", 403, o.Payload)
}

func (o *CbExclusionsUpdateV1Forbidden) String() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1Forbidden  %+v", 403, o.Payload)
}

func (o *CbExclusionsUpdateV1Forbidden) GetPayload() *models.MsaspecResponseFields {
	return o.Payload
}

func (o *CbExclusionsUpdateV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCbExclusionsUpdateV1TooManyRequests creates a CbExclusionsUpdateV1TooManyRequests with default headers values
func NewCbExclusionsUpdateV1TooManyRequests() *CbExclusionsUpdateV1TooManyRequests {
	return &CbExclusionsUpdateV1TooManyRequests{}
}

/*
CbExclusionsUpdateV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CbExclusionsUpdateV1TooManyRequests struct {

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

// IsSuccess returns true when this cb exclusions update v1 too many requests response has a 2xx status code
func (o *CbExclusionsUpdateV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this cb exclusions update v1 too many requests response has a 3xx status code
func (o *CbExclusionsUpdateV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cb exclusions update v1 too many requests response has a 4xx status code
func (o *CbExclusionsUpdateV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this cb exclusions update v1 too many requests response has a 5xx status code
func (o *CbExclusionsUpdateV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this cb exclusions update v1 too many requests response a status code equal to that given
func (o *CbExclusionsUpdateV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the cb exclusions update v1 too many requests response
func (o *CbExclusionsUpdateV1TooManyRequests) Code() int {
	return 429
}

func (o *CbExclusionsUpdateV1TooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CbExclusionsUpdateV1TooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /exclusions/entities/cert-based-exclusions/v1][%d] cbExclusionsUpdateV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *CbExclusionsUpdateV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CbExclusionsUpdateV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
