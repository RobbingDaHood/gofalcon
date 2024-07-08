// Code generated by go-swagger; DO NOT EDIT.

package host_migration

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

// GetMigrationsV1Reader is a Reader for the GetMigrationsV1 structure.
type GetMigrationsV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMigrationsV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetMigrationsV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetMigrationsV1BadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetMigrationsV1Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetMigrationsV1NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetMigrationsV1TooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetMigrationsV1InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /host-migration/entities/migrations/v1] GetMigrationsV1", response, response.Code())
	}
}

// NewGetMigrationsV1OK creates a GetMigrationsV1OK with default headers values
func NewGetMigrationsV1OK() *GetMigrationsV1OK {
	return &GetMigrationsV1OK{}
}

/*
GetMigrationsV1OK describes a response with status code 200, with default header values.

OK
*/
type GetMigrationsV1OK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIGetMigrationsResponseV1
}

// IsSuccess returns true when this get migrations v1 o k response has a 2xx status code
func (o *GetMigrationsV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get migrations v1 o k response has a 3xx status code
func (o *GetMigrationsV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migrations v1 o k response has a 4xx status code
func (o *GetMigrationsV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get migrations v1 o k response has a 5xx status code
func (o *GetMigrationsV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this get migrations v1 o k response a status code equal to that given
func (o *GetMigrationsV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get migrations v1 o k response
func (o *GetMigrationsV1OK) Code() int {
	return 200
}

func (o *GetMigrationsV1OK) Error() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1OK  %+v", 200, o.Payload)
}

func (o *GetMigrationsV1OK) String() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1OK  %+v", 200, o.Payload)
}

func (o *GetMigrationsV1OK) GetPayload() *models.APIGetMigrationsResponseV1 {
	return o.Payload
}

func (o *GetMigrationsV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIGetMigrationsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationsV1BadRequest creates a GetMigrationsV1BadRequest with default headers values
func NewGetMigrationsV1BadRequest() *GetMigrationsV1BadRequest {
	return &GetMigrationsV1BadRequest{}
}

/*
GetMigrationsV1BadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type GetMigrationsV1BadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIGetMigrationsResponseV1
}

// IsSuccess returns true when this get migrations v1 bad request response has a 2xx status code
func (o *GetMigrationsV1BadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migrations v1 bad request response has a 3xx status code
func (o *GetMigrationsV1BadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migrations v1 bad request response has a 4xx status code
func (o *GetMigrationsV1BadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migrations v1 bad request response has a 5xx status code
func (o *GetMigrationsV1BadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get migrations v1 bad request response a status code equal to that given
func (o *GetMigrationsV1BadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get migrations v1 bad request response
func (o *GetMigrationsV1BadRequest) Code() int {
	return 400
}

func (o *GetMigrationsV1BadRequest) Error() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMigrationsV1BadRequest) String() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1BadRequest  %+v", 400, o.Payload)
}

func (o *GetMigrationsV1BadRequest) GetPayload() *models.APIGetMigrationsResponseV1 {
	return o.Payload
}

func (o *GetMigrationsV1BadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIGetMigrationsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationsV1Forbidden creates a GetMigrationsV1Forbidden with default headers values
func NewGetMigrationsV1Forbidden() *GetMigrationsV1Forbidden {
	return &GetMigrationsV1Forbidden{}
}

/*
GetMigrationsV1Forbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetMigrationsV1Forbidden struct {

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

// IsSuccess returns true when this get migrations v1 forbidden response has a 2xx status code
func (o *GetMigrationsV1Forbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migrations v1 forbidden response has a 3xx status code
func (o *GetMigrationsV1Forbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migrations v1 forbidden response has a 4xx status code
func (o *GetMigrationsV1Forbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migrations v1 forbidden response has a 5xx status code
func (o *GetMigrationsV1Forbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get migrations v1 forbidden response a status code equal to that given
func (o *GetMigrationsV1Forbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get migrations v1 forbidden response
func (o *GetMigrationsV1Forbidden) Code() int {
	return 403
}

func (o *GetMigrationsV1Forbidden) Error() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMigrationsV1Forbidden) String() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1Forbidden  %+v", 403, o.Payload)
}

func (o *GetMigrationsV1Forbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMigrationsV1Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMigrationsV1NotFound creates a GetMigrationsV1NotFound with default headers values
func NewGetMigrationsV1NotFound() *GetMigrationsV1NotFound {
	return &GetMigrationsV1NotFound{}
}

/*
GetMigrationsV1NotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetMigrationsV1NotFound struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIGetMigrationsResponseV1
}

// IsSuccess returns true when this get migrations v1 not found response has a 2xx status code
func (o *GetMigrationsV1NotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migrations v1 not found response has a 3xx status code
func (o *GetMigrationsV1NotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migrations v1 not found response has a 4xx status code
func (o *GetMigrationsV1NotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migrations v1 not found response has a 5xx status code
func (o *GetMigrationsV1NotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get migrations v1 not found response a status code equal to that given
func (o *GetMigrationsV1NotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get migrations v1 not found response
func (o *GetMigrationsV1NotFound) Code() int {
	return 404
}

func (o *GetMigrationsV1NotFound) Error() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1NotFound  %+v", 404, o.Payload)
}

func (o *GetMigrationsV1NotFound) String() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1NotFound  %+v", 404, o.Payload)
}

func (o *GetMigrationsV1NotFound) GetPayload() *models.APIGetMigrationsResponseV1 {
	return o.Payload
}

func (o *GetMigrationsV1NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIGetMigrationsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetMigrationsV1TooManyRequests creates a GetMigrationsV1TooManyRequests with default headers values
func NewGetMigrationsV1TooManyRequests() *GetMigrationsV1TooManyRequests {
	return &GetMigrationsV1TooManyRequests{}
}

/*
GetMigrationsV1TooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type GetMigrationsV1TooManyRequests struct {

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

// IsSuccess returns true when this get migrations v1 too many requests response has a 2xx status code
func (o *GetMigrationsV1TooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migrations v1 too many requests response has a 3xx status code
func (o *GetMigrationsV1TooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migrations v1 too many requests response has a 4xx status code
func (o *GetMigrationsV1TooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get migrations v1 too many requests response has a 5xx status code
func (o *GetMigrationsV1TooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get migrations v1 too many requests response a status code equal to that given
func (o *GetMigrationsV1TooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the get migrations v1 too many requests response
func (o *GetMigrationsV1TooManyRequests) Code() int {
	return 429
}

func (o *GetMigrationsV1TooManyRequests) Error() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMigrationsV1TooManyRequests) String() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1TooManyRequests  %+v", 429, o.Payload)
}

func (o *GetMigrationsV1TooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *GetMigrationsV1TooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetMigrationsV1InternalServerError creates a GetMigrationsV1InternalServerError with default headers values
func NewGetMigrationsV1InternalServerError() *GetMigrationsV1InternalServerError {
	return &GetMigrationsV1InternalServerError{}
}

/*
GetMigrationsV1InternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetMigrationsV1InternalServerError struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIGetMigrationsResponseV1
}

// IsSuccess returns true when this get migrations v1 internal server error response has a 2xx status code
func (o *GetMigrationsV1InternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get migrations v1 internal server error response has a 3xx status code
func (o *GetMigrationsV1InternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get migrations v1 internal server error response has a 4xx status code
func (o *GetMigrationsV1InternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get migrations v1 internal server error response has a 5xx status code
func (o *GetMigrationsV1InternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get migrations v1 internal server error response a status code equal to that given
func (o *GetMigrationsV1InternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get migrations v1 internal server error response
func (o *GetMigrationsV1InternalServerError) Code() int {
	return 500
}

func (o *GetMigrationsV1InternalServerError) Error() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMigrationsV1InternalServerError) String() string {
	return fmt.Sprintf("[GET /host-migration/entities/migrations/v1][%d] getMigrationsV1InternalServerError  %+v", 500, o.Payload)
}

func (o *GetMigrationsV1InternalServerError) GetPayload() *models.APIGetMigrationsResponseV1 {
	return o.Payload
}

func (o *GetMigrationsV1InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIGetMigrationsResponseV1)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
