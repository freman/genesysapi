// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGroupIndividualsReader is a Reader for the GetGroupIndividuals structure.
type GetGroupIndividualsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupIndividualsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupIndividualsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupIndividualsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGroupIndividualsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGroupIndividualsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupIndividualsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGroupIndividualsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGroupIndividualsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGroupIndividualsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGroupIndividualsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupIndividualsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGroupIndividualsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGroupIndividualsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupIndividualsOK creates a GetGroupIndividualsOK with default headers values
func NewGetGroupIndividualsOK() *GetGroupIndividualsOK {
	return &GetGroupIndividualsOK{}
}

/*
GetGroupIndividualsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGroupIndividualsOK struct {
	Payload *models.UserEntityListing
}

// IsSuccess returns true when this get group individuals o k response has a 2xx status code
func (o *GetGroupIndividualsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get group individuals o k response has a 3xx status code
func (o *GetGroupIndividualsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals o k response has a 4xx status code
func (o *GetGroupIndividualsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get group individuals o k response has a 5xx status code
func (o *GetGroupIndividualsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get group individuals o k response a status code equal to that given
func (o *GetGroupIndividualsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGroupIndividualsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsOK  %+v", 200, o.Payload)
}

func (o *GetGroupIndividualsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsOK  %+v", 200, o.Payload)
}

func (o *GetGroupIndividualsOK) GetPayload() *models.UserEntityListing {
	return o.Payload
}

func (o *GetGroupIndividualsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsBadRequest creates a GetGroupIndividualsBadRequest with default headers values
func NewGetGroupIndividualsBadRequest() *GetGroupIndividualsBadRequest {
	return &GetGroupIndividualsBadRequest{}
}

/*
GetGroupIndividualsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGroupIndividualsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals bad request response has a 2xx status code
func (o *GetGroupIndividualsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals bad request response has a 3xx status code
func (o *GetGroupIndividualsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals bad request response has a 4xx status code
func (o *GetGroupIndividualsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group individuals bad request response has a 5xx status code
func (o *GetGroupIndividualsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get group individuals bad request response a status code equal to that given
func (o *GetGroupIndividualsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetGroupIndividualsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGroupIndividualsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGroupIndividualsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsUnauthorized creates a GetGroupIndividualsUnauthorized with default headers values
func NewGetGroupIndividualsUnauthorized() *GetGroupIndividualsUnauthorized {
	return &GetGroupIndividualsUnauthorized{}
}

/*
GetGroupIndividualsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGroupIndividualsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals unauthorized response has a 2xx status code
func (o *GetGroupIndividualsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals unauthorized response has a 3xx status code
func (o *GetGroupIndividualsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals unauthorized response has a 4xx status code
func (o *GetGroupIndividualsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group individuals unauthorized response has a 5xx status code
func (o *GetGroupIndividualsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get group individuals unauthorized response a status code equal to that given
func (o *GetGroupIndividualsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGroupIndividualsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGroupIndividualsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGroupIndividualsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsForbidden creates a GetGroupIndividualsForbidden with default headers values
func NewGetGroupIndividualsForbidden() *GetGroupIndividualsForbidden {
	return &GetGroupIndividualsForbidden{}
}

/*
GetGroupIndividualsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetGroupIndividualsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals forbidden response has a 2xx status code
func (o *GetGroupIndividualsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals forbidden response has a 3xx status code
func (o *GetGroupIndividualsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals forbidden response has a 4xx status code
func (o *GetGroupIndividualsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group individuals forbidden response has a 5xx status code
func (o *GetGroupIndividualsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get group individuals forbidden response a status code equal to that given
func (o *GetGroupIndividualsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGroupIndividualsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsForbidden  %+v", 403, o.Payload)
}

func (o *GetGroupIndividualsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsForbidden  %+v", 403, o.Payload)
}

func (o *GetGroupIndividualsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsNotFound creates a GetGroupIndividualsNotFound with default headers values
func NewGetGroupIndividualsNotFound() *GetGroupIndividualsNotFound {
	return &GetGroupIndividualsNotFound{}
}

/*
GetGroupIndividualsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetGroupIndividualsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals not found response has a 2xx status code
func (o *GetGroupIndividualsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals not found response has a 3xx status code
func (o *GetGroupIndividualsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals not found response has a 4xx status code
func (o *GetGroupIndividualsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group individuals not found response has a 5xx status code
func (o *GetGroupIndividualsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get group individuals not found response a status code equal to that given
func (o *GetGroupIndividualsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGroupIndividualsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsNotFound  %+v", 404, o.Payload)
}

func (o *GetGroupIndividualsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsNotFound  %+v", 404, o.Payload)
}

func (o *GetGroupIndividualsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsRequestTimeout creates a GetGroupIndividualsRequestTimeout with default headers values
func NewGetGroupIndividualsRequestTimeout() *GetGroupIndividualsRequestTimeout {
	return &GetGroupIndividualsRequestTimeout{}
}

/*
GetGroupIndividualsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGroupIndividualsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals request timeout response has a 2xx status code
func (o *GetGroupIndividualsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals request timeout response has a 3xx status code
func (o *GetGroupIndividualsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals request timeout response has a 4xx status code
func (o *GetGroupIndividualsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group individuals request timeout response has a 5xx status code
func (o *GetGroupIndividualsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get group individuals request timeout response a status code equal to that given
func (o *GetGroupIndividualsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetGroupIndividualsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGroupIndividualsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGroupIndividualsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsRequestEntityTooLarge creates a GetGroupIndividualsRequestEntityTooLarge with default headers values
func NewGetGroupIndividualsRequestEntityTooLarge() *GetGroupIndividualsRequestEntityTooLarge {
	return &GetGroupIndividualsRequestEntityTooLarge{}
}

/*
GetGroupIndividualsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetGroupIndividualsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals request entity too large response has a 2xx status code
func (o *GetGroupIndividualsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals request entity too large response has a 3xx status code
func (o *GetGroupIndividualsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals request entity too large response has a 4xx status code
func (o *GetGroupIndividualsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group individuals request entity too large response has a 5xx status code
func (o *GetGroupIndividualsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get group individuals request entity too large response a status code equal to that given
func (o *GetGroupIndividualsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetGroupIndividualsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGroupIndividualsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGroupIndividualsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsUnsupportedMediaType creates a GetGroupIndividualsUnsupportedMediaType with default headers values
func NewGetGroupIndividualsUnsupportedMediaType() *GetGroupIndividualsUnsupportedMediaType {
	return &GetGroupIndividualsUnsupportedMediaType{}
}

/*
GetGroupIndividualsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGroupIndividualsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals unsupported media type response has a 2xx status code
func (o *GetGroupIndividualsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals unsupported media type response has a 3xx status code
func (o *GetGroupIndividualsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals unsupported media type response has a 4xx status code
func (o *GetGroupIndividualsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group individuals unsupported media type response has a 5xx status code
func (o *GetGroupIndividualsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get group individuals unsupported media type response a status code equal to that given
func (o *GetGroupIndividualsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetGroupIndividualsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGroupIndividualsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGroupIndividualsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsTooManyRequests creates a GetGroupIndividualsTooManyRequests with default headers values
func NewGetGroupIndividualsTooManyRequests() *GetGroupIndividualsTooManyRequests {
	return &GetGroupIndividualsTooManyRequests{}
}

/*
GetGroupIndividualsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGroupIndividualsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals too many requests response has a 2xx status code
func (o *GetGroupIndividualsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals too many requests response has a 3xx status code
func (o *GetGroupIndividualsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals too many requests response has a 4xx status code
func (o *GetGroupIndividualsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get group individuals too many requests response has a 5xx status code
func (o *GetGroupIndividualsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get group individuals too many requests response a status code equal to that given
func (o *GetGroupIndividualsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetGroupIndividualsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGroupIndividualsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGroupIndividualsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsInternalServerError creates a GetGroupIndividualsInternalServerError with default headers values
func NewGetGroupIndividualsInternalServerError() *GetGroupIndividualsInternalServerError {
	return &GetGroupIndividualsInternalServerError{}
}

/*
GetGroupIndividualsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGroupIndividualsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals internal server error response has a 2xx status code
func (o *GetGroupIndividualsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals internal server error response has a 3xx status code
func (o *GetGroupIndividualsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals internal server error response has a 4xx status code
func (o *GetGroupIndividualsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get group individuals internal server error response has a 5xx status code
func (o *GetGroupIndividualsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get group individuals internal server error response a status code equal to that given
func (o *GetGroupIndividualsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGroupIndividualsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGroupIndividualsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGroupIndividualsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsServiceUnavailable creates a GetGroupIndividualsServiceUnavailable with default headers values
func NewGetGroupIndividualsServiceUnavailable() *GetGroupIndividualsServiceUnavailable {
	return &GetGroupIndividualsServiceUnavailable{}
}

/*
GetGroupIndividualsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGroupIndividualsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals service unavailable response has a 2xx status code
func (o *GetGroupIndividualsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals service unavailable response has a 3xx status code
func (o *GetGroupIndividualsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals service unavailable response has a 4xx status code
func (o *GetGroupIndividualsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get group individuals service unavailable response has a 5xx status code
func (o *GetGroupIndividualsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get group individuals service unavailable response a status code equal to that given
func (o *GetGroupIndividualsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetGroupIndividualsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGroupIndividualsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGroupIndividualsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsGatewayTimeout creates a GetGroupIndividualsGatewayTimeout with default headers values
func NewGetGroupIndividualsGatewayTimeout() *GetGroupIndividualsGatewayTimeout {
	return &GetGroupIndividualsGatewayTimeout{}
}

/*
GetGroupIndividualsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetGroupIndividualsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get group individuals gateway timeout response has a 2xx status code
func (o *GetGroupIndividualsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get group individuals gateway timeout response has a 3xx status code
func (o *GetGroupIndividualsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get group individuals gateway timeout response has a 4xx status code
func (o *GetGroupIndividualsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get group individuals gateway timeout response has a 5xx status code
func (o *GetGroupIndividualsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get group individuals gateway timeout response a status code equal to that given
func (o *GetGroupIndividualsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetGroupIndividualsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGroupIndividualsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGroupIndividualsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
