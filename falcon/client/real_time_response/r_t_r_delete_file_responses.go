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

// RTRDeleteFileReader is a Reader for the RTRDeleteFile structure.
type RTRDeleteFileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RTRDeleteFileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewRTRDeleteFileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRTRDeleteFileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRTRDeleteFileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRTRDeleteFileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewRTRDeleteFileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewRTRDeleteFileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /real-time-response/entities/file/v1] RTR-DeleteFile", response, response.Code())
	}
}

// NewRTRDeleteFileNoContent creates a RTRDeleteFileNoContent with default headers values
func NewRTRDeleteFileNoContent() *RTRDeleteFileNoContent {
	return &RTRDeleteFileNoContent{}
}

/*
RTRDeleteFileNoContent describes a response with status code 204, with default header values.

No Content
*/
type RTRDeleteFileNoContent struct {

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

// IsSuccess returns true when this r t r delete file no content response has a 2xx status code
func (o *RTRDeleteFileNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this r t r delete file no content response has a 3xx status code
func (o *RTRDeleteFileNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file no content response has a 4xx status code
func (o *RTRDeleteFileNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r delete file no content response has a 5xx status code
func (o *RTRDeleteFileNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file no content response a status code equal to that given
func (o *RTRDeleteFileNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the r t r delete file no content response
func (o *RTRDeleteFileNoContent) Code() int {
	return 204
}

func (o *RTRDeleteFileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileNoContent  %+v", 204, o.Payload)
}

func (o *RTRDeleteFileNoContent) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileNoContent  %+v", 204, o.Payload)
}

func (o *RTRDeleteFileNoContent) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteFileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteFileBadRequest creates a RTRDeleteFileBadRequest with default headers values
func NewRTRDeleteFileBadRequest() *RTRDeleteFileBadRequest {
	return &RTRDeleteFileBadRequest{}
}

/*
RTRDeleteFileBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type RTRDeleteFileBadRequest struct {

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

// IsSuccess returns true when this r t r delete file bad request response has a 2xx status code
func (o *RTRDeleteFileBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete file bad request response has a 3xx status code
func (o *RTRDeleteFileBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file bad request response has a 4xx status code
func (o *RTRDeleteFileBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete file bad request response has a 5xx status code
func (o *RTRDeleteFileBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file bad request response a status code equal to that given
func (o *RTRDeleteFileBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the r t r delete file bad request response
func (o *RTRDeleteFileBadRequest) Code() int {
	return 400
}

func (o *RTRDeleteFileBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileBadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeleteFileBadRequest) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileBadRequest  %+v", 400, o.Payload)
}

func (o *RTRDeleteFileBadRequest) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRDeleteFileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteFileForbidden creates a RTRDeleteFileForbidden with default headers values
func NewRTRDeleteFileForbidden() *RTRDeleteFileForbidden {
	return &RTRDeleteFileForbidden{}
}

/*
RTRDeleteFileForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type RTRDeleteFileForbidden struct {

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

// IsSuccess returns true when this r t r delete file forbidden response has a 2xx status code
func (o *RTRDeleteFileForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete file forbidden response has a 3xx status code
func (o *RTRDeleteFileForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file forbidden response has a 4xx status code
func (o *RTRDeleteFileForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete file forbidden response has a 5xx status code
func (o *RTRDeleteFileForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file forbidden response a status code equal to that given
func (o *RTRDeleteFileForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the r t r delete file forbidden response
func (o *RTRDeleteFileForbidden) Code() int {
	return 403
}

func (o *RTRDeleteFileForbidden) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileForbidden  %+v", 403, o.Payload)
}

func (o *RTRDeleteFileForbidden) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileForbidden  %+v", 403, o.Payload)
}

func (o *RTRDeleteFileForbidden) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteFileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteFileNotFound creates a RTRDeleteFileNotFound with default headers values
func NewRTRDeleteFileNotFound() *RTRDeleteFileNotFound {
	return &RTRDeleteFileNotFound{}
}

/*
RTRDeleteFileNotFound describes a response with status code 404, with default header values.

Not Found
*/
type RTRDeleteFileNotFound struct {

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

// IsSuccess returns true when this r t r delete file not found response has a 2xx status code
func (o *RTRDeleteFileNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete file not found response has a 3xx status code
func (o *RTRDeleteFileNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file not found response has a 4xx status code
func (o *RTRDeleteFileNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete file not found response has a 5xx status code
func (o *RTRDeleteFileNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file not found response a status code equal to that given
func (o *RTRDeleteFileNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the r t r delete file not found response
func (o *RTRDeleteFileNotFound) Code() int {
	return 404
}

func (o *RTRDeleteFileNotFound) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileNotFound  %+v", 404, o.Payload)
}

func (o *RTRDeleteFileNotFound) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileNotFound  %+v", 404, o.Payload)
}

func (o *RTRDeleteFileNotFound) GetPayload() *models.DomainAPIError {
	return o.Payload
}

func (o *RTRDeleteFileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteFileTooManyRequests creates a RTRDeleteFileTooManyRequests with default headers values
func NewRTRDeleteFileTooManyRequests() *RTRDeleteFileTooManyRequests {
	return &RTRDeleteFileTooManyRequests{}
}

/*
RTRDeleteFileTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type RTRDeleteFileTooManyRequests struct {

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

// IsSuccess returns true when this r t r delete file too many requests response has a 2xx status code
func (o *RTRDeleteFileTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete file too many requests response has a 3xx status code
func (o *RTRDeleteFileTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file too many requests response has a 4xx status code
func (o *RTRDeleteFileTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this r t r delete file too many requests response has a 5xx status code
func (o *RTRDeleteFileTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this r t r delete file too many requests response a status code equal to that given
func (o *RTRDeleteFileTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the r t r delete file too many requests response
func (o *RTRDeleteFileTooManyRequests) Code() int {
	return 429
}

func (o *RTRDeleteFileTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeleteFileTooManyRequests) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileTooManyRequests  %+v", 429, o.Payload)
}

func (o *RTRDeleteFileTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteFileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewRTRDeleteFileInternalServerError creates a RTRDeleteFileInternalServerError with default headers values
func NewRTRDeleteFileInternalServerError() *RTRDeleteFileInternalServerError {
	return &RTRDeleteFileInternalServerError{}
}

/*
RTRDeleteFileInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type RTRDeleteFileInternalServerError struct {

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

// IsSuccess returns true when this r t r delete file internal server error response has a 2xx status code
func (o *RTRDeleteFileInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this r t r delete file internal server error response has a 3xx status code
func (o *RTRDeleteFileInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this r t r delete file internal server error response has a 4xx status code
func (o *RTRDeleteFileInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this r t r delete file internal server error response has a 5xx status code
func (o *RTRDeleteFileInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this r t r delete file internal server error response a status code equal to that given
func (o *RTRDeleteFileInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the r t r delete file internal server error response
func (o *RTRDeleteFileInternalServerError) Code() int {
	return 500
}

func (o *RTRDeleteFileInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRDeleteFileInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /real-time-response/entities/file/v1][%d] rTRDeleteFileInternalServerError  %+v", 500, o.Payload)
}

func (o *RTRDeleteFileInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *RTRDeleteFileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
