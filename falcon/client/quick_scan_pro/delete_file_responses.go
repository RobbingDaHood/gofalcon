// Code generated by go-swagger; DO NOT EDIT.

package quick_scan_pro

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

// DeleteFileReader is a Reader for the DeleteFile structure.
type DeleteFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteFileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewDeleteFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteFileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /quickscanpro/entities/files/v1] DeleteFile", response, response.Code())
	}
}

// NewDeleteFileOK creates a DeleteFileOK with default headers values
func NewDeleteFileOK() *DeleteFileOK {
	return &DeleteFileOK{}
}

/*
DeleteFileOK describes a response with status code 200, with default header values.

OK
*/
type DeleteFileOK struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.QuickscanproDeleteFileResponse
}

// IsSuccess returns true when this delete file o k response has a 2xx status code
func (o *DeleteFileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete file o k response has a 3xx status code
func (o *DeleteFileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete file o k response has a 4xx status code
func (o *DeleteFileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete file o k response has a 5xx status code
func (o *DeleteFileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delete file o k response a status code equal to that given
func (o *DeleteFileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delete file o k response
func (o *DeleteFileOK) Code() int {
	return 200
}

func (o *DeleteFileOK) Error() string {
	return fmt.Sprintf("[DELETE /quickscanpro/entities/files/v1][%d] deleteFileOK  %+v", 200, o.Payload)
}

func (o *DeleteFileOK) String() string {
	return fmt.Sprintf("[DELETE /quickscanpro/entities/files/v1][%d] deleteFileOK  %+v", 200, o.Payload)
}

func (o *DeleteFileOK) GetPayload() *models.QuickscanproDeleteFileResponse {
	return o.Payload
}

func (o *DeleteFileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.QuickscanproDeleteFileResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteFileForbidden creates a DeleteFileForbidden with default headers values
func NewDeleteFileForbidden() *DeleteFileForbidden {
	return &DeleteFileForbidden{}
}

/*
DeleteFileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteFileForbidden struct {

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

// IsSuccess returns true when this delete file forbidden response has a 2xx status code
func (o *DeleteFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete file forbidden response has a 3xx status code
func (o *DeleteFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete file forbidden response has a 4xx status code
func (o *DeleteFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete file forbidden response has a 5xx status code
func (o *DeleteFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete file forbidden response a status code equal to that given
func (o *DeleteFileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete file forbidden response
func (o *DeleteFileForbidden) Code() int {
	return 403
}

func (o *DeleteFileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /quickscanpro/entities/files/v1][%d] deleteFileForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFileForbidden) String() string {
	return fmt.Sprintf("[DELETE /quickscanpro/entities/files/v1][%d] deleteFileForbidden  %+v", 403, o.Payload)
}

func (o *DeleteFileForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteFileTooManyRequests creates a DeleteFileTooManyRequests with default headers values
func NewDeleteFileTooManyRequests() *DeleteFileTooManyRequests {
	return &DeleteFileTooManyRequests{}
}

/*
DeleteFileTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type DeleteFileTooManyRequests struct {

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

// IsSuccess returns true when this delete file too many requests response has a 2xx status code
func (o *DeleteFileTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete file too many requests response has a 3xx status code
func (o *DeleteFileTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete file too many requests response has a 4xx status code
func (o *DeleteFileTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete file too many requests response has a 5xx status code
func (o *DeleteFileTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this delete file too many requests response a status code equal to that given
func (o *DeleteFileTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the delete file too many requests response
func (o *DeleteFileTooManyRequests) Code() int {
	return 429
}

func (o *DeleteFileTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /quickscanpro/entities/files/v1][%d] deleteFileTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFileTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /quickscanpro/entities/files/v1][%d] deleteFileTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteFileTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteFileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteFileInternalServerError creates a DeleteFileInternalServerError with default headers values
func NewDeleteFileInternalServerError() *DeleteFileInternalServerError {
	return &DeleteFileInternalServerError{}
}

/*
DeleteFileInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type DeleteFileInternalServerError struct {

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

// IsSuccess returns true when this delete file internal server error response has a 2xx status code
func (o *DeleteFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete file internal server error response has a 3xx status code
func (o *DeleteFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete file internal server error response has a 4xx status code
func (o *DeleteFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete file internal server error response has a 5xx status code
func (o *DeleteFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete file internal server error response a status code equal to that given
func (o *DeleteFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete file internal server error response
func (o *DeleteFileInternalServerError) Code() int {
	return 500
}

func (o *DeleteFileInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /quickscanpro/entities/files/v1][%d] deleteFileInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFileInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /quickscanpro/entities/files/v1][%d] deleteFileInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteFileInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *DeleteFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
