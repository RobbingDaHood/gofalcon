// Code generated by go-swagger; DO NOT EDIT.

package falcon_complete_dashboard

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

// QueryIncidentIdsByFilterReader is a Reader for the QueryIncidentIdsByFilter structure.
type QueryIncidentIdsByFilterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryIncidentIdsByFilterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryIncidentIdsByFilterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewQueryIncidentIdsByFilterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewQueryIncidentIdsByFilterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewQueryIncidentIdsByFilterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /falcon-complete-dashboards/queries/incidents/v1] QueryIncidentIdsByFilter", response, response.Code())
	}
}

// NewQueryIncidentIdsByFilterOK creates a QueryIncidentIdsByFilterOK with default headers values
func NewQueryIncidentIdsByFilterOK() *QueryIncidentIdsByFilterOK {
	return &QueryIncidentIdsByFilterOK{}
}

/*
QueryIncidentIdsByFilterOK describes a response with status code 200, with default header values.

OK
*/
type QueryIncidentIdsByFilterOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaQueryResponse
}

// IsSuccess returns true when this query incident ids by filter o k response has a 2xx status code
func (o *QueryIncidentIdsByFilterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query incident ids by filter o k response has a 3xx status code
func (o *QueryIncidentIdsByFilterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query incident ids by filter o k response has a 4xx status code
func (o *QueryIncidentIdsByFilterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query incident ids by filter o k response has a 5xx status code
func (o *QueryIncidentIdsByFilterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query incident ids by filter o k response a status code equal to that given
func (o *QueryIncidentIdsByFilterOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query incident ids by filter o k response
func (o *QueryIncidentIdsByFilterOK) Code() int {
	return 200
}

func (o *QueryIncidentIdsByFilterOK) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/incidents/v1][%d] queryIncidentIdsByFilterOK  %+v", 200, o.Payload)
}

func (o *QueryIncidentIdsByFilterOK) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/incidents/v1][%d] queryIncidentIdsByFilterOK  %+v", 200, o.Payload)
}

func (o *QueryIncidentIdsByFilterOK) GetPayload() *models.MsaQueryResponse {
	return o.Payload
}

func (o *QueryIncidentIdsByFilterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaQueryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryIncidentIdsByFilterForbidden creates a QueryIncidentIdsByFilterForbidden with default headers values
func NewQueryIncidentIdsByFilterForbidden() *QueryIncidentIdsByFilterForbidden {
	return &QueryIncidentIdsByFilterForbidden{}
}

/*
QueryIncidentIdsByFilterForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type QueryIncidentIdsByFilterForbidden struct {

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

// IsSuccess returns true when this query incident ids by filter forbidden response has a 2xx status code
func (o *QueryIncidentIdsByFilterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query incident ids by filter forbidden response has a 3xx status code
func (o *QueryIncidentIdsByFilterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query incident ids by filter forbidden response has a 4xx status code
func (o *QueryIncidentIdsByFilterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this query incident ids by filter forbidden response has a 5xx status code
func (o *QueryIncidentIdsByFilterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this query incident ids by filter forbidden response a status code equal to that given
func (o *QueryIncidentIdsByFilterForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the query incident ids by filter forbidden response
func (o *QueryIncidentIdsByFilterForbidden) Code() int {
	return 403
}

func (o *QueryIncidentIdsByFilterForbidden) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/incidents/v1][%d] queryIncidentIdsByFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryIncidentIdsByFilterForbidden) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/incidents/v1][%d] queryIncidentIdsByFilterForbidden  %+v", 403, o.Payload)
}

func (o *QueryIncidentIdsByFilterForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIncidentIdsByFilterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIncidentIdsByFilterTooManyRequests creates a QueryIncidentIdsByFilterTooManyRequests with default headers values
func NewQueryIncidentIdsByFilterTooManyRequests() *QueryIncidentIdsByFilterTooManyRequests {
	return &QueryIncidentIdsByFilterTooManyRequests{}
}

/*
QueryIncidentIdsByFilterTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type QueryIncidentIdsByFilterTooManyRequests struct {

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

// IsSuccess returns true when this query incident ids by filter too many requests response has a 2xx status code
func (o *QueryIncidentIdsByFilterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query incident ids by filter too many requests response has a 3xx status code
func (o *QueryIncidentIdsByFilterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query incident ids by filter too many requests response has a 4xx status code
func (o *QueryIncidentIdsByFilterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this query incident ids by filter too many requests response has a 5xx status code
func (o *QueryIncidentIdsByFilterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this query incident ids by filter too many requests response a status code equal to that given
func (o *QueryIncidentIdsByFilterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the query incident ids by filter too many requests response
func (o *QueryIncidentIdsByFilterTooManyRequests) Code() int {
	return 429
}

func (o *QueryIncidentIdsByFilterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/incidents/v1][%d] queryIncidentIdsByFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryIncidentIdsByFilterTooManyRequests) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/incidents/v1][%d] queryIncidentIdsByFilterTooManyRequests  %+v", 429, o.Payload)
}

func (o *QueryIncidentIdsByFilterTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIncidentIdsByFilterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewQueryIncidentIdsByFilterInternalServerError creates a QueryIncidentIdsByFilterInternalServerError with default headers values
func NewQueryIncidentIdsByFilterInternalServerError() *QueryIncidentIdsByFilterInternalServerError {
	return &QueryIncidentIdsByFilterInternalServerError{}
}

/*
QueryIncidentIdsByFilterInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type QueryIncidentIdsByFilterInternalServerError struct {

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

// IsSuccess returns true when this query incident ids by filter internal server error response has a 2xx status code
func (o *QueryIncidentIdsByFilterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this query incident ids by filter internal server error response has a 3xx status code
func (o *QueryIncidentIdsByFilterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query incident ids by filter internal server error response has a 4xx status code
func (o *QueryIncidentIdsByFilterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this query incident ids by filter internal server error response has a 5xx status code
func (o *QueryIncidentIdsByFilterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this query incident ids by filter internal server error response a status code equal to that given
func (o *QueryIncidentIdsByFilterInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the query incident ids by filter internal server error response
func (o *QueryIncidentIdsByFilterInternalServerError) Code() int {
	return 500
}

func (o *QueryIncidentIdsByFilterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/incidents/v1][%d] queryIncidentIdsByFilterInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryIncidentIdsByFilterInternalServerError) String() string {
	return fmt.Sprintf("[GET /falcon-complete-dashboards/queries/incidents/v1][%d] queryIncidentIdsByFilterInternalServerError  %+v", 500, o.Payload)
}

func (o *QueryIncidentIdsByFilterInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *QueryIncidentIdsByFilterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
