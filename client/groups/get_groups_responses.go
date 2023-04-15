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

// GetGroupsReader is a Reader for the GetGroups structure.
type GetGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGroupsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupsOK creates a GetGroupsOK with default headers values
func NewGetGroupsOK() *GetGroupsOK {
	return &GetGroupsOK{}
}

/*
GetGroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetGroupsOK struct {
	Payload *models.GroupEntityListing
}

// IsSuccess returns true when this get groups o k response has a 2xx status code
func (o *GetGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get groups o k response has a 3xx status code
func (o *GetGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups o k response has a 4xx status code
func (o *GetGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get groups o k response has a 5xx status code
func (o *GetGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get groups o k response a status code equal to that given
func (o *GetGroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsOK  %+v", 200, o.Payload)
}

func (o *GetGroupsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsOK  %+v", 200, o.Payload)
}

func (o *GetGroupsOK) GetPayload() *models.GroupEntityListing {
	return o.Payload
}

func (o *GetGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GroupEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsBadRequest creates a GetGroupsBadRequest with default headers values
func NewGetGroupsBadRequest() *GetGroupsBadRequest {
	return &GetGroupsBadRequest{}
}

/*
GetGroupsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGroupsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups bad request response has a 2xx status code
func (o *GetGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups bad request response has a 3xx status code
func (o *GetGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups bad request response has a 4xx status code
func (o *GetGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get groups bad request response has a 5xx status code
func (o *GetGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get groups bad request response a status code equal to that given
func (o *GetGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetGroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGroupsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsUnauthorized creates a GetGroupsUnauthorized with default headers values
func NewGetGroupsUnauthorized() *GetGroupsUnauthorized {
	return &GetGroupsUnauthorized{}
}

/*
GetGroupsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGroupsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups unauthorized response has a 2xx status code
func (o *GetGroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups unauthorized response has a 3xx status code
func (o *GetGroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups unauthorized response has a 4xx status code
func (o *GetGroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get groups unauthorized response has a 5xx status code
func (o *GetGroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get groups unauthorized response a status code equal to that given
func (o *GetGroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGroupsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsForbidden creates a GetGroupsForbidden with default headers values
func NewGetGroupsForbidden() *GetGroupsForbidden {
	return &GetGroupsForbidden{}
}

/*
GetGroupsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetGroupsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups forbidden response has a 2xx status code
func (o *GetGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups forbidden response has a 3xx status code
func (o *GetGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups forbidden response has a 4xx status code
func (o *GetGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get groups forbidden response has a 5xx status code
func (o *GetGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get groups forbidden response a status code equal to that given
func (o *GetGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetGroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetGroupsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetGroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsNotFound creates a GetGroupsNotFound with default headers values
func NewGetGroupsNotFound() *GetGroupsNotFound {
	return &GetGroupsNotFound{}
}

/*
GetGroupsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetGroupsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups not found response has a 2xx status code
func (o *GetGroupsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups not found response has a 3xx status code
func (o *GetGroupsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups not found response has a 4xx status code
func (o *GetGroupsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get groups not found response has a 5xx status code
func (o *GetGroupsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get groups not found response a status code equal to that given
func (o *GetGroupsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetGroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetGroupsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetGroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsRequestTimeout creates a GetGroupsRequestTimeout with default headers values
func NewGetGroupsRequestTimeout() *GetGroupsRequestTimeout {
	return &GetGroupsRequestTimeout{}
}

/*
GetGroupsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGroupsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups request timeout response has a 2xx status code
func (o *GetGroupsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups request timeout response has a 3xx status code
func (o *GetGroupsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups request timeout response has a 4xx status code
func (o *GetGroupsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get groups request timeout response has a 5xx status code
func (o *GetGroupsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get groups request timeout response a status code equal to that given
func (o *GetGroupsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetGroupsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGroupsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGroupsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsRequestEntityTooLarge creates a GetGroupsRequestEntityTooLarge with default headers values
func NewGetGroupsRequestEntityTooLarge() *GetGroupsRequestEntityTooLarge {
	return &GetGroupsRequestEntityTooLarge{}
}

/*
GetGroupsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetGroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups request entity too large response has a 2xx status code
func (o *GetGroupsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups request entity too large response has a 3xx status code
func (o *GetGroupsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups request entity too large response has a 4xx status code
func (o *GetGroupsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get groups request entity too large response has a 5xx status code
func (o *GetGroupsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get groups request entity too large response a status code equal to that given
func (o *GetGroupsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetGroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGroupsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsUnsupportedMediaType creates a GetGroupsUnsupportedMediaType with default headers values
func NewGetGroupsUnsupportedMediaType() *GetGroupsUnsupportedMediaType {
	return &GetGroupsUnsupportedMediaType{}
}

/*
GetGroupsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups unsupported media type response has a 2xx status code
func (o *GetGroupsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups unsupported media type response has a 3xx status code
func (o *GetGroupsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups unsupported media type response has a 4xx status code
func (o *GetGroupsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get groups unsupported media type response has a 5xx status code
func (o *GetGroupsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get groups unsupported media type response a status code equal to that given
func (o *GetGroupsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetGroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGroupsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsTooManyRequests creates a GetGroupsTooManyRequests with default headers values
func NewGetGroupsTooManyRequests() *GetGroupsTooManyRequests {
	return &GetGroupsTooManyRequests{}
}

/*
GetGroupsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups too many requests response has a 2xx status code
func (o *GetGroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups too many requests response has a 3xx status code
func (o *GetGroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups too many requests response has a 4xx status code
func (o *GetGroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get groups too many requests response has a 5xx status code
func (o *GetGroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get groups too many requests response a status code equal to that given
func (o *GetGroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGroupsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsInternalServerError creates a GetGroupsInternalServerError with default headers values
func NewGetGroupsInternalServerError() *GetGroupsInternalServerError {
	return &GetGroupsInternalServerError{}
}

/*
GetGroupsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGroupsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups internal server error response has a 2xx status code
func (o *GetGroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups internal server error response has a 3xx status code
func (o *GetGroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups internal server error response has a 4xx status code
func (o *GetGroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get groups internal server error response has a 5xx status code
func (o *GetGroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get groups internal server error response a status code equal to that given
func (o *GetGroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGroupsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsServiceUnavailable creates a GetGroupsServiceUnavailable with default headers values
func NewGetGroupsServiceUnavailable() *GetGroupsServiceUnavailable {
	return &GetGroupsServiceUnavailable{}
}

/*
GetGroupsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups service unavailable response has a 2xx status code
func (o *GetGroupsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups service unavailable response has a 3xx status code
func (o *GetGroupsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups service unavailable response has a 4xx status code
func (o *GetGroupsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get groups service unavailable response has a 5xx status code
func (o *GetGroupsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get groups service unavailable response a status code equal to that given
func (o *GetGroupsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetGroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGroupsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupsGatewayTimeout creates a GetGroupsGatewayTimeout with default headers values
func NewGetGroupsGatewayTimeout() *GetGroupsGatewayTimeout {
	return &GetGroupsGatewayTimeout{}
}

/*
GetGroupsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetGroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get groups gateway timeout response has a 2xx status code
func (o *GetGroupsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get groups gateway timeout response has a 3xx status code
func (o *GetGroupsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get groups gateway timeout response has a 4xx status code
func (o *GetGroupsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get groups gateway timeout response has a 5xx status code
func (o *GetGroupsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get groups gateway timeout response a status code equal to that given
func (o *GetGroupsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetGroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGroupsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/groups][%d] getGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
