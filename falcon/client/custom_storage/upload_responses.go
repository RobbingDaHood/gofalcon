// Code generated by go-swagger; DO NOT EDIT.

package custom_storage

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

// UploadReader is a Reader for the Upload structure.
type UploadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUploadForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewUploadTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PUT /customobjects/v1/collections/{collection_name}/objects/{object_key}] upload", response, response.Code())
	}
}

// NewUploadOK creates a UploadOK with default headers values
func NewUploadOK() *UploadOK {
	return &UploadOK{}
}

/*
UploadOK describes a response with status code 200, with default header values.

OK
*/
type UploadOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.CustomStorageResponse
}

// IsSuccess returns true when this upload o k response has a 2xx status code
func (o *UploadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this upload o k response has a 3xx status code
func (o *UploadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload o k response has a 4xx status code
func (o *UploadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload o k response has a 5xx status code
func (o *UploadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this upload o k response a status code equal to that given
func (o *UploadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the upload o k response
func (o *UploadOK) Code() int {
	return 200
}

func (o *UploadOK) Error() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections/{collection_name}/objects/{object_key}][%d] uploadOK  %+v", 200, o.Payload)
}

func (o *UploadOK) String() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections/{collection_name}/objects/{object_key}][%d] uploadOK  %+v", 200, o.Payload)
}

func (o *UploadOK) GetPayload() *models.CustomStorageResponse {
	return o.Payload
}

func (o *UploadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.CustomStorageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadForbidden creates a UploadForbidden with default headers values
func NewUploadForbidden() *UploadForbidden {
	return &UploadForbidden{}
}

/*
UploadForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UploadForbidden struct {

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

// IsSuccess returns true when this upload forbidden response has a 2xx status code
func (o *UploadForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload forbidden response has a 3xx status code
func (o *UploadForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload forbidden response has a 4xx status code
func (o *UploadForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload forbidden response has a 5xx status code
func (o *UploadForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this upload forbidden response a status code equal to that given
func (o *UploadForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the upload forbidden response
func (o *UploadForbidden) Code() int {
	return 403
}

func (o *UploadForbidden) Error() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections/{collection_name}/objects/{object_key}][%d] uploadForbidden  %+v", 403, o.Payload)
}

func (o *UploadForbidden) String() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections/{collection_name}/objects/{object_key}][%d] uploadForbidden  %+v", 403, o.Payload)
}

func (o *UploadForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UploadForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUploadTooManyRequests creates a UploadTooManyRequests with default headers values
func NewUploadTooManyRequests() *UploadTooManyRequests {
	return &UploadTooManyRequests{}
}

/*
UploadTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type UploadTooManyRequests struct {

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

// IsSuccess returns true when this upload too many requests response has a 2xx status code
func (o *UploadTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload too many requests response has a 3xx status code
func (o *UploadTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload too many requests response has a 4xx status code
func (o *UploadTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this upload too many requests response has a 5xx status code
func (o *UploadTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this upload too many requests response a status code equal to that given
func (o *UploadTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the upload too many requests response
func (o *UploadTooManyRequests) Code() int {
	return 429
}

func (o *UploadTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections/{collection_name}/objects/{object_key}][%d] uploadTooManyRequests  %+v", 429, o.Payload)
}

func (o *UploadTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections/{collection_name}/objects/{object_key}][%d] uploadTooManyRequests  %+v", 429, o.Payload)
}

func (o *UploadTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UploadTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUploadInternalServerError creates a UploadInternalServerError with default headers values
func NewUploadInternalServerError() *UploadInternalServerError {
	return &UploadInternalServerError{}
}

/*
UploadInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type UploadInternalServerError struct {

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

// IsSuccess returns true when this upload internal server error response has a 2xx status code
func (o *UploadInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this upload internal server error response has a 3xx status code
func (o *UploadInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this upload internal server error response has a 4xx status code
func (o *UploadInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this upload internal server error response has a 5xx status code
func (o *UploadInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this upload internal server error response a status code equal to that given
func (o *UploadInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the upload internal server error response
func (o *UploadInternalServerError) Code() int {
	return 500
}

func (o *UploadInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections/{collection_name}/objects/{object_key}][%d] uploadInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadInternalServerError) String() string {
	return fmt.Sprintf("[PUT /customobjects/v1/collections/{collection_name}/objects/{object_key}][%d] uploadInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *UploadInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
