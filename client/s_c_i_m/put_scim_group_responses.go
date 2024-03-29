// Code generated by go-swagger; DO NOT EDIT.

package s_c_i_m

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PutScimGroupReader is a Reader for the PutScimGroup structure.
type PutScimGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutScimGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutScimGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutScimGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutScimGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutScimGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutScimGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutScimGroupRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutScimGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutScimGroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutScimGroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutScimGroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutScimGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutScimGroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutScimGroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutScimGroupOK creates a PutScimGroupOK with default headers values
func NewPutScimGroupOK() *PutScimGroupOK {
	return &PutScimGroupOK{}
}

/*
PutScimGroupOK describes a response with status code 200, with default header values.

successful operation
*/
type PutScimGroupOK struct {
	Payload *models.ScimV2Group
}

// IsSuccess returns true when this put scim group o k response has a 2xx status code
func (o *PutScimGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this put scim group o k response has a 3xx status code
func (o *PutScimGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group o k response has a 4xx status code
func (o *PutScimGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this put scim group o k response has a 5xx status code
func (o *PutScimGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group o k response a status code equal to that given
func (o *PutScimGroupOK) IsCode(code int) bool {
	return code == 200
}

func (o *PutScimGroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupOK  %+v", 200, o.Payload)
}

func (o *PutScimGroupOK) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupOK  %+v", 200, o.Payload)
}

func (o *PutScimGroupOK) GetPayload() *models.ScimV2Group {
	return o.Payload
}

func (o *PutScimGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimV2Group)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupBadRequest creates a PutScimGroupBadRequest with default headers values
func NewPutScimGroupBadRequest() *PutScimGroupBadRequest {
	return &PutScimGroupBadRequest{}
}

/*
PutScimGroupBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutScimGroupBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group bad request response has a 2xx status code
func (o *PutScimGroupBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group bad request response has a 3xx status code
func (o *PutScimGroupBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group bad request response has a 4xx status code
func (o *PutScimGroupBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this put scim group bad request response has a 5xx status code
func (o *PutScimGroupBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group bad request response a status code equal to that given
func (o *PutScimGroupBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PutScimGroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupBadRequest  %+v", 400, o.Payload)
}

func (o *PutScimGroupBadRequest) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupBadRequest  %+v", 400, o.Payload)
}

func (o *PutScimGroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupUnauthorized creates a PutScimGroupUnauthorized with default headers values
func NewPutScimGroupUnauthorized() *PutScimGroupUnauthorized {
	return &PutScimGroupUnauthorized{}
}

/*
PutScimGroupUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PutScimGroupUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group unauthorized response has a 2xx status code
func (o *PutScimGroupUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group unauthorized response has a 3xx status code
func (o *PutScimGroupUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group unauthorized response has a 4xx status code
func (o *PutScimGroupUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this put scim group unauthorized response has a 5xx status code
func (o *PutScimGroupUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group unauthorized response a status code equal to that given
func (o *PutScimGroupUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PutScimGroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *PutScimGroupUnauthorized) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *PutScimGroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupForbidden creates a PutScimGroupForbidden with default headers values
func NewPutScimGroupForbidden() *PutScimGroupForbidden {
	return &PutScimGroupForbidden{}
}

/*
PutScimGroupForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PutScimGroupForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group forbidden response has a 2xx status code
func (o *PutScimGroupForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group forbidden response has a 3xx status code
func (o *PutScimGroupForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group forbidden response has a 4xx status code
func (o *PutScimGroupForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this put scim group forbidden response has a 5xx status code
func (o *PutScimGroupForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group forbidden response a status code equal to that given
func (o *PutScimGroupForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PutScimGroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupForbidden  %+v", 403, o.Payload)
}

func (o *PutScimGroupForbidden) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupForbidden  %+v", 403, o.Payload)
}

func (o *PutScimGroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupNotFound creates a PutScimGroupNotFound with default headers values
func NewPutScimGroupNotFound() *PutScimGroupNotFound {
	return &PutScimGroupNotFound{}
}

/*
PutScimGroupNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PutScimGroupNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group not found response has a 2xx status code
func (o *PutScimGroupNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group not found response has a 3xx status code
func (o *PutScimGroupNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group not found response has a 4xx status code
func (o *PutScimGroupNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this put scim group not found response has a 5xx status code
func (o *PutScimGroupNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group not found response a status code equal to that given
func (o *PutScimGroupNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PutScimGroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupNotFound  %+v", 404, o.Payload)
}

func (o *PutScimGroupNotFound) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupNotFound  %+v", 404, o.Payload)
}

func (o *PutScimGroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupRequestTimeout creates a PutScimGroupRequestTimeout with default headers values
func NewPutScimGroupRequestTimeout() *PutScimGroupRequestTimeout {
	return &PutScimGroupRequestTimeout{}
}

/*
PutScimGroupRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutScimGroupRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group request timeout response has a 2xx status code
func (o *PutScimGroupRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group request timeout response has a 3xx status code
func (o *PutScimGroupRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group request timeout response has a 4xx status code
func (o *PutScimGroupRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this put scim group request timeout response has a 5xx status code
func (o *PutScimGroupRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group request timeout response a status code equal to that given
func (o *PutScimGroupRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PutScimGroupRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutScimGroupRequestTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutScimGroupRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupConflict creates a PutScimGroupConflict with default headers values
func NewPutScimGroupConflict() *PutScimGroupConflict {
	return &PutScimGroupConflict{}
}

/*
PutScimGroupConflict describes a response with status code 409, with default header values.

Version does not match current version.
*/
type PutScimGroupConflict struct {
	Payload *models.ScimError
}

// IsSuccess returns true when this put scim group conflict response has a 2xx status code
func (o *PutScimGroupConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group conflict response has a 3xx status code
func (o *PutScimGroupConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group conflict response has a 4xx status code
func (o *PutScimGroupConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this put scim group conflict response has a 5xx status code
func (o *PutScimGroupConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group conflict response a status code equal to that given
func (o *PutScimGroupConflict) IsCode(code int) bool {
	return code == 409
}

func (o *PutScimGroupConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupConflict  %+v", 409, o.Payload)
}

func (o *PutScimGroupConflict) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupConflict  %+v", 409, o.Payload)
}

func (o *PutScimGroupConflict) GetPayload() *models.ScimError {
	return o.Payload
}

func (o *PutScimGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupRequestEntityTooLarge creates a PutScimGroupRequestEntityTooLarge with default headers values
func NewPutScimGroupRequestEntityTooLarge() *PutScimGroupRequestEntityTooLarge {
	return &PutScimGroupRequestEntityTooLarge{}
}

/*
PutScimGroupRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PutScimGroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group request entity too large response has a 2xx status code
func (o *PutScimGroupRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group request entity too large response has a 3xx status code
func (o *PutScimGroupRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group request entity too large response has a 4xx status code
func (o *PutScimGroupRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this put scim group request entity too large response has a 5xx status code
func (o *PutScimGroupRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group request entity too large response a status code equal to that given
func (o *PutScimGroupRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PutScimGroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutScimGroupRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutScimGroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupUnsupportedMediaType creates a PutScimGroupUnsupportedMediaType with default headers values
func NewPutScimGroupUnsupportedMediaType() *PutScimGroupUnsupportedMediaType {
	return &PutScimGroupUnsupportedMediaType{}
}

/*
PutScimGroupUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutScimGroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group unsupported media type response has a 2xx status code
func (o *PutScimGroupUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group unsupported media type response has a 3xx status code
func (o *PutScimGroupUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group unsupported media type response has a 4xx status code
func (o *PutScimGroupUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this put scim group unsupported media type response has a 5xx status code
func (o *PutScimGroupUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group unsupported media type response a status code equal to that given
func (o *PutScimGroupUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PutScimGroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutScimGroupUnsupportedMediaType) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutScimGroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupTooManyRequests creates a PutScimGroupTooManyRequests with default headers values
func NewPutScimGroupTooManyRequests() *PutScimGroupTooManyRequests {
	return &PutScimGroupTooManyRequests{}
}

/*
PutScimGroupTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutScimGroupTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group too many requests response has a 2xx status code
func (o *PutScimGroupTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group too many requests response has a 3xx status code
func (o *PutScimGroupTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group too many requests response has a 4xx status code
func (o *PutScimGroupTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this put scim group too many requests response has a 5xx status code
func (o *PutScimGroupTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this put scim group too many requests response a status code equal to that given
func (o *PutScimGroupTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PutScimGroupTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutScimGroupTooManyRequests) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutScimGroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupInternalServerError creates a PutScimGroupInternalServerError with default headers values
func NewPutScimGroupInternalServerError() *PutScimGroupInternalServerError {
	return &PutScimGroupInternalServerError{}
}

/*
PutScimGroupInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutScimGroupInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group internal server error response has a 2xx status code
func (o *PutScimGroupInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group internal server error response has a 3xx status code
func (o *PutScimGroupInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group internal server error response has a 4xx status code
func (o *PutScimGroupInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this put scim group internal server error response has a 5xx status code
func (o *PutScimGroupInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this put scim group internal server error response a status code equal to that given
func (o *PutScimGroupInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PutScimGroupInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *PutScimGroupInternalServerError) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *PutScimGroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupServiceUnavailable creates a PutScimGroupServiceUnavailable with default headers values
func NewPutScimGroupServiceUnavailable() *PutScimGroupServiceUnavailable {
	return &PutScimGroupServiceUnavailable{}
}

/*
PutScimGroupServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutScimGroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group service unavailable response has a 2xx status code
func (o *PutScimGroupServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group service unavailable response has a 3xx status code
func (o *PutScimGroupServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group service unavailable response has a 4xx status code
func (o *PutScimGroupServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this put scim group service unavailable response has a 5xx status code
func (o *PutScimGroupServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this put scim group service unavailable response a status code equal to that given
func (o *PutScimGroupServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PutScimGroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutScimGroupServiceUnavailable) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutScimGroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutScimGroupGatewayTimeout creates a PutScimGroupGatewayTimeout with default headers values
func NewPutScimGroupGatewayTimeout() *PutScimGroupGatewayTimeout {
	return &PutScimGroupGatewayTimeout{}
}

/*
PutScimGroupGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PutScimGroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this put scim group gateway timeout response has a 2xx status code
func (o *PutScimGroupGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this put scim group gateway timeout response has a 3xx status code
func (o *PutScimGroupGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this put scim group gateway timeout response has a 4xx status code
func (o *PutScimGroupGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this put scim group gateway timeout response has a 5xx status code
func (o *PutScimGroupGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this put scim group gateway timeout response a status code equal to that given
func (o *PutScimGroupGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PutScimGroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutScimGroupGatewayTimeout) String() string {
	return fmt.Sprintf("[PUT /api/v2/scim/groups/{groupId}][%d] putScimGroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutScimGroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutScimGroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
