// Code generated by go-swagger; DO NOT EDIT.

package user_management

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

// CreateUserReader is a Reader for the CreateUser structure.
type CreateUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewCreateUserTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /users/entities/users/v1] CreateUser", response, response.Code())
	}
}

// NewCreateUserCreated creates a CreateUserCreated with default headers values
func NewCreateUserCreated() *CreateUserCreated {
	return &CreateUserCreated{}
}

/*
CreateUserCreated describes a response with status code 201, with default header values.

Created
*/
type CreateUserCreated struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.APIUserMetadataResponse
}

// IsSuccess returns true when this create user created response has a 2xx status code
func (o *CreateUserCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create user created response has a 3xx status code
func (o *CreateUserCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user created response has a 4xx status code
func (o *CreateUserCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user created response has a 5xx status code
func (o *CreateUserCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create user created response a status code equal to that given
func (o *CreateUserCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create user created response
func (o *CreateUserCreated) Code() int {
	return 201
}

func (o *CreateUserCreated) Error() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) String() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserCreated  %+v", 201, o.Payload)
}

func (o *CreateUserCreated) GetPayload() *models.APIUserMetadataResponse {
	return o.Payload
}

func (o *CreateUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.APIUserMetadataResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserBadRequest creates a CreateUserBadRequest with default headers values
func NewCreateUserBadRequest() *CreateUserBadRequest {
	return &CreateUserBadRequest{}
}

/*
CreateUserBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type CreateUserBadRequest struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this create user bad request response has a 2xx status code
func (o *CreateUserBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user bad request response has a 3xx status code
func (o *CreateUserBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user bad request response has a 4xx status code
func (o *CreateUserBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user bad request response has a 5xx status code
func (o *CreateUserBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create user bad request response a status code equal to that given
func (o *CreateUserBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the create user bad request response
func (o *CreateUserBadRequest) Code() int {
	return 400
}

func (o *CreateUserBadRequest) Error() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) String() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserBadRequest  %+v", 400, o.Payload)
}

func (o *CreateUserBadRequest) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *CreateUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserForbidden creates a CreateUserForbidden with default headers values
func NewCreateUserForbidden() *CreateUserForbidden {
	return &CreateUserForbidden{}
}

/*
CreateUserForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateUserForbidden struct {

	/* Trace-ID: submit to support if resolving an issue
	 */
	XCSTRACEID string

	/* Request limit per minute.
	 */
	XRateLimitLimit int64

	/* The number of requests remaining for the sliding one minute window.
	 */
	XRateLimitRemaining int64

	Payload *models.MsaEntitiesResponse
}

// IsSuccess returns true when this create user forbidden response has a 2xx status code
func (o *CreateUserForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user forbidden response has a 3xx status code
func (o *CreateUserForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user forbidden response has a 4xx status code
func (o *CreateUserForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user forbidden response has a 5xx status code
func (o *CreateUserForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create user forbidden response a status code equal to that given
func (o *CreateUserForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the create user forbidden response
func (o *CreateUserForbidden) Code() int {
	return 403
}

func (o *CreateUserForbidden) Error() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserForbidden) String() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserForbidden  %+v", 403, o.Payload)
}

func (o *CreateUserForbidden) GetPayload() *models.MsaEntitiesResponse {
	return o.Payload
}

func (o *CreateUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

	o.Payload = new(models.MsaEntitiesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserTooManyRequests creates a CreateUserTooManyRequests with default headers values
func NewCreateUserTooManyRequests() *CreateUserTooManyRequests {
	return &CreateUserTooManyRequests{}
}

/*
CreateUserTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type CreateUserTooManyRequests struct {

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

// IsSuccess returns true when this create user too many requests response has a 2xx status code
func (o *CreateUserTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user too many requests response has a 3xx status code
func (o *CreateUserTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user too many requests response has a 4xx status code
func (o *CreateUserTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this create user too many requests response has a 5xx status code
func (o *CreateUserTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this create user too many requests response a status code equal to that given
func (o *CreateUserTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the create user too many requests response
func (o *CreateUserTooManyRequests) Code() int {
	return 429
}

func (o *CreateUserTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUserTooManyRequests) String() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserTooManyRequests  %+v", 429, o.Payload)
}

func (o *CreateUserTooManyRequests) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateUserTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateUserInternalServerError creates a CreateUserInternalServerError with default headers values
func NewCreateUserInternalServerError() *CreateUserInternalServerError {
	return &CreateUserInternalServerError{}
}

/*
CreateUserInternalServerError describes a response with status code 500, with default header values.

Unexpected Error
*/
type CreateUserInternalServerError struct {

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

// IsSuccess returns true when this create user internal server error response has a 2xx status code
func (o *CreateUserInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create user internal server error response has a 3xx status code
func (o *CreateUserInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create user internal server error response has a 4xx status code
func (o *CreateUserInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create user internal server error response has a 5xx status code
func (o *CreateUserInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create user internal server error response a status code equal to that given
func (o *CreateUserInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the create user internal server error response
func (o *CreateUserInternalServerError) Code() int {
	return 500
}

func (o *CreateUserInternalServerError) Error() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserInternalServerError) String() string {
	return fmt.Sprintf("[POST /users/entities/users/v1][%d] createUserInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserInternalServerError) GetPayload() *models.MsaReplyMetaOnly {
	return o.Payload
}

func (o *CreateUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
