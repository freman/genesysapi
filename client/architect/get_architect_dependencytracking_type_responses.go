// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetArchitectDependencytrackingTypeReader is a Reader for the GetArchitectDependencytrackingType structure.
type GetArchitectDependencytrackingTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectDependencytrackingTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectDependencytrackingTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectDependencytrackingTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectDependencytrackingTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectDependencytrackingTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectDependencytrackingTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewGetArchitectDependencytrackingTypeMethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectDependencytrackingTypeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectDependencytrackingTypeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectDependencytrackingTypeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectDependencytrackingTypeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectDependencytrackingTypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectDependencytrackingTypeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectDependencytrackingTypeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectDependencytrackingTypeOK creates a GetArchitectDependencytrackingTypeOK with default headers values
func NewGetArchitectDependencytrackingTypeOK() *GetArchitectDependencytrackingTypeOK {
	return &GetArchitectDependencytrackingTypeOK{}
}

/*
GetArchitectDependencytrackingTypeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetArchitectDependencytrackingTypeOK struct {
	Payload *models.DependencyType
}

// IsSuccess returns true when this get architect dependencytracking type o k response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get architect dependencytracking type o k response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type o k response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect dependencytracking type o k response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type o k response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetArchitectDependencytrackingTypeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeOK  %+v", 200, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeOK) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeOK  %+v", 200, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeOK) GetPayload() *models.DependencyType {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DependencyType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeBadRequest creates a GetArchitectDependencytrackingTypeBadRequest with default headers values
func NewGetArchitectDependencytrackingTypeBadRequest() *GetArchitectDependencytrackingTypeBadRequest {
	return &GetArchitectDependencytrackingTypeBadRequest{}
}

/*
GetArchitectDependencytrackingTypeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectDependencytrackingTypeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type bad request response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type bad request response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type bad request response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking type bad request response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type bad request response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetArchitectDependencytrackingTypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeUnauthorized creates a GetArchitectDependencytrackingTypeUnauthorized with default headers values
func NewGetArchitectDependencytrackingTypeUnauthorized() *GetArchitectDependencytrackingTypeUnauthorized {
	return &GetArchitectDependencytrackingTypeUnauthorized{}
}

/*
GetArchitectDependencytrackingTypeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectDependencytrackingTypeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type unauthorized response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type unauthorized response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type unauthorized response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking type unauthorized response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type unauthorized response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetArchitectDependencytrackingTypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeForbidden creates a GetArchitectDependencytrackingTypeForbidden with default headers values
func NewGetArchitectDependencytrackingTypeForbidden() *GetArchitectDependencytrackingTypeForbidden {
	return &GetArchitectDependencytrackingTypeForbidden{}
}

/*
GetArchitectDependencytrackingTypeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectDependencytrackingTypeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type forbidden response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type forbidden response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type forbidden response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking type forbidden response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type forbidden response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetArchitectDependencytrackingTypeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeNotFound creates a GetArchitectDependencytrackingTypeNotFound with default headers values
func NewGetArchitectDependencytrackingTypeNotFound() *GetArchitectDependencytrackingTypeNotFound {
	return &GetArchitectDependencytrackingTypeNotFound{}
}

/*
GetArchitectDependencytrackingTypeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetArchitectDependencytrackingTypeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type not found response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type not found response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type not found response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking type not found response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type not found response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetArchitectDependencytrackingTypeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeMethodNotAllowed creates a GetArchitectDependencytrackingTypeMethodNotAllowed with default headers values
func NewGetArchitectDependencytrackingTypeMethodNotAllowed() *GetArchitectDependencytrackingTypeMethodNotAllowed {
	return &GetArchitectDependencytrackingTypeMethodNotAllowed{}
}

/*
GetArchitectDependencytrackingTypeMethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed
*/
type GetArchitectDependencytrackingTypeMethodNotAllowed struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type method not allowed response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeMethodNotAllowed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type method not allowed response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeMethodNotAllowed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type method not allowed response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeMethodNotAllowed) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking type method not allowed response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeMethodNotAllowed) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type method not allowed response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeMethodNotAllowed) IsCode(code int) bool {
	return code == 405
}

func (o *GetArchitectDependencytrackingTypeMethodNotAllowed) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeMethodNotAllowed) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeMethodNotAllowed  %+v", 405, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeMethodNotAllowed) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeMethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeRequestTimeout creates a GetArchitectDependencytrackingTypeRequestTimeout with default headers values
func NewGetArchitectDependencytrackingTypeRequestTimeout() *GetArchitectDependencytrackingTypeRequestTimeout {
	return &GetArchitectDependencytrackingTypeRequestTimeout{}
}

/*
GetArchitectDependencytrackingTypeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectDependencytrackingTypeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type request timeout response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type request timeout response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type request timeout response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking type request timeout response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type request timeout response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetArchitectDependencytrackingTypeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeRequestEntityTooLarge creates a GetArchitectDependencytrackingTypeRequestEntityTooLarge with default headers values
func NewGetArchitectDependencytrackingTypeRequestEntityTooLarge() *GetArchitectDependencytrackingTypeRequestEntityTooLarge {
	return &GetArchitectDependencytrackingTypeRequestEntityTooLarge{}
}

/*
GetArchitectDependencytrackingTypeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetArchitectDependencytrackingTypeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type request entity too large response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type request entity too large response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type request entity too large response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking type request entity too large response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type request entity too large response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetArchitectDependencytrackingTypeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeUnsupportedMediaType creates a GetArchitectDependencytrackingTypeUnsupportedMediaType with default headers values
func NewGetArchitectDependencytrackingTypeUnsupportedMediaType() *GetArchitectDependencytrackingTypeUnsupportedMediaType {
	return &GetArchitectDependencytrackingTypeUnsupportedMediaType{}
}

/*
GetArchitectDependencytrackingTypeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectDependencytrackingTypeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type unsupported media type response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type unsupported media type response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type unsupported media type response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking type unsupported media type response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type unsupported media type response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetArchitectDependencytrackingTypeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeTooManyRequests creates a GetArchitectDependencytrackingTypeTooManyRequests with default headers values
func NewGetArchitectDependencytrackingTypeTooManyRequests() *GetArchitectDependencytrackingTypeTooManyRequests {
	return &GetArchitectDependencytrackingTypeTooManyRequests{}
}

/*
GetArchitectDependencytrackingTypeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectDependencytrackingTypeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type too many requests response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type too many requests response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type too many requests response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking type too many requests response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking type too many requests response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetArchitectDependencytrackingTypeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeInternalServerError creates a GetArchitectDependencytrackingTypeInternalServerError with default headers values
func NewGetArchitectDependencytrackingTypeInternalServerError() *GetArchitectDependencytrackingTypeInternalServerError {
	return &GetArchitectDependencytrackingTypeInternalServerError{}
}

/*
GetArchitectDependencytrackingTypeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectDependencytrackingTypeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type internal server error response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type internal server error response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type internal server error response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect dependencytracking type internal server error response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect dependencytracking type internal server error response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetArchitectDependencytrackingTypeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeServiceUnavailable creates a GetArchitectDependencytrackingTypeServiceUnavailable with default headers values
func NewGetArchitectDependencytrackingTypeServiceUnavailable() *GetArchitectDependencytrackingTypeServiceUnavailable {
	return &GetArchitectDependencytrackingTypeServiceUnavailable{}
}

/*
GetArchitectDependencytrackingTypeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectDependencytrackingTypeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type service unavailable response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type service unavailable response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type service unavailable response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect dependencytracking type service unavailable response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect dependencytracking type service unavailable response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetArchitectDependencytrackingTypeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingTypeGatewayTimeout creates a GetArchitectDependencytrackingTypeGatewayTimeout with default headers values
func NewGetArchitectDependencytrackingTypeGatewayTimeout() *GetArchitectDependencytrackingTypeGatewayTimeout {
	return &GetArchitectDependencytrackingTypeGatewayTimeout{}
}

/*
GetArchitectDependencytrackingTypeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetArchitectDependencytrackingTypeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking type gateway timeout response has a 2xx status code
func (o *GetArchitectDependencytrackingTypeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking type gateway timeout response has a 3xx status code
func (o *GetArchitectDependencytrackingTypeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking type gateway timeout response has a 4xx status code
func (o *GetArchitectDependencytrackingTypeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect dependencytracking type gateway timeout response has a 5xx status code
func (o *GetArchitectDependencytrackingTypeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect dependencytracking type gateway timeout response a status code equal to that given
func (o *GetArchitectDependencytrackingTypeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetArchitectDependencytrackingTypeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/types/{typeId}][%d] getArchitectDependencytrackingTypeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectDependencytrackingTypeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingTypeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
