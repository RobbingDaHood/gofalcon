// Code generated by go-swagger; DO NOT EDIT.

package custom_ioa

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

// UpdateRulesReader is a Reader for the UpdateRules structure.
type UpdateRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUpdateRulesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRulesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUpdateRulesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRulesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /ioarules/entities/rules/v1] update-rules", response, response.Code())
	}
}

// NewUpdateRulesOK creates a UpdateRulesOK with default headers values
func NewUpdateRulesOK() *UpdateRulesOK {
	return &UpdateRulesOK{}
}

/*
UpdateRulesOK describes a response with status code 200, with default header values.

OK
*/
type UpdateRulesOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIRulesResponse
}

// IsSuccess returns true when this update rules o k response has a 2xx status code
func (o *UpdateRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update rules o k response has a 3xx status code
func (o *UpdateRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rules o k response has a 4xx status code
func (o *UpdateRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update rules o k response has a 5xx status code
func (o *UpdateRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update rules o k response a status code equal to that given
func (o *UpdateRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update rules o k response
func (o *UpdateRulesOK) Code() int {
	return 200
}

func (o *UpdateRulesOK) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateRulesOK) String() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateRulesOK) GetPayload() *models.APIRulesResponse {
	return o.Payload
}

func (o *UpdateRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIRulesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRulesForbidden creates a UpdateRulesForbidden with default headers values
func NewUpdateRulesForbidden() *UpdateRulesForbidden {
	return &UpdateRulesForbidden{}
}

/*
UpdateRulesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRulesForbidden struct {

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

// IsSuccess returns true when this update rules forbidden response has a 2xx status code
func (o *UpdateRulesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rules forbidden response has a 3xx status code
func (o *UpdateRulesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rules forbidden response has a 4xx status code
func (o *UpdateRulesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rules forbidden response has a 5xx status code
func (o *UpdateRulesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update rules forbidden response a status code equal to that given
func (o *UpdateRulesForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the update rules forbidden response
func (o *UpdateRulesForbidden) Code() int {
	return 403
}

func (o *UpdateRulesForbidden) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRulesForbidden) String() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRulesForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRulesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRulesNotFound creates a UpdateRulesNotFound with default headers values
func NewUpdateRulesNotFound() *UpdateRulesNotFound {
	return &UpdateRulesNotFound{}
}

/*
UpdateRulesNotFound describes a response with status code 404, with default header values.

Not Found
*/
type UpdateRulesNotFound struct {

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

// IsSuccess returns true when this update rules not found response has a 2xx status code
func (o *UpdateRulesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rules not found response has a 3xx status code
func (o *UpdateRulesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rules not found response has a 4xx status code
func (o *UpdateRulesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rules not found response has a 5xx status code
func (o *UpdateRulesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update rules not found response a status code equal to that given
func (o *UpdateRulesNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the update rules not found response
func (o *UpdateRulesNotFound) Code() int {
	return 404
}

func (o *UpdateRulesNotFound) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRulesNotFound) String() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRulesNotFound) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRulesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRulesTooManyRequests creates a UpdateRulesTooManyRequests with default headers values
func NewUpdateRulesTooManyRequests() *UpdateRulesTooManyRequests {
	return &UpdateRulesTooManyRequests{}
}

/*
UpdateRulesTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UpdateRulesTooManyRequests struct {

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

// IsSuccess returns true when this update rules too many requests response has a 2xx status code
func (o *UpdateRulesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rules too many requests response has a 3xx status code
func (o *UpdateRulesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rules too many requests response has a 4xx status code
func (o *UpdateRulesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this update rules too many requests response has a 5xx status code
func (o *UpdateRulesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this update rules too many requests response a status code equal to that given
func (o *UpdateRulesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the update rules too many requests response
func (o *UpdateRulesTooManyRequests) Code() int {
	return 429
}

func (o *UpdateRulesTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRulesTooManyRequests) String() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesTooManyRequests  %+v", 429, o.Payload)
}

func (o *UpdateRulesTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRulesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRulesInternalServerError creates a UpdateRulesInternalServerError with default headers values
func NewUpdateRulesInternalServerError() *UpdateRulesInternalServerError {
	return &UpdateRulesInternalServerError{}
}

/*
UpdateRulesInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type UpdateRulesInternalServerError struct {

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

// IsSuccess returns true when this update rules internal server error response has a 2xx status code
func (o *UpdateRulesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update rules internal server error response has a 3xx status code
func (o *UpdateRulesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update rules internal server error response has a 4xx status code
func (o *UpdateRulesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update rules internal server error response has a 5xx status code
func (o *UpdateRulesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update rules internal server error response a status code equal to that given
func (o *UpdateRulesInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the update rules internal server error response
func (o *UpdateRulesInternalServerError) Code() int {
	return 500
}

func (o *UpdateRulesInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRulesInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /ioarules/entities/rules/v1][%d] updateRulesInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRulesInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UpdateRulesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
