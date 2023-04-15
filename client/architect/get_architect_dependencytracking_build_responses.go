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

// GetArchitectDependencytrackingBuildReader is a Reader for the GetArchitectDependencytrackingBuild structure.
type GetArchitectDependencytrackingBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectDependencytrackingBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectDependencytrackingBuildOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectDependencytrackingBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectDependencytrackingBuildUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectDependencytrackingBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectDependencytrackingBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetArchitectDependencytrackingBuildRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectDependencytrackingBuildRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectDependencytrackingBuildUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectDependencytrackingBuildTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectDependencytrackingBuildInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectDependencytrackingBuildServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectDependencytrackingBuildGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectDependencytrackingBuildOK creates a GetArchitectDependencytrackingBuildOK with default headers values
func NewGetArchitectDependencytrackingBuildOK() *GetArchitectDependencytrackingBuildOK {
	return &GetArchitectDependencytrackingBuildOK{}
}

/*
GetArchitectDependencytrackingBuildOK describes a response with status code 200, with default header values.

successful operation
*/
type GetArchitectDependencytrackingBuildOK struct {
	Payload *models.DependencyStatus
}

// IsSuccess returns true when this get architect dependencytracking build o k response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get architect dependencytracking build o k response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build o k response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect dependencytracking build o k response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking build o k response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetArchitectDependencytrackingBuildOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildOK  %+v", 200, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildOK) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildOK  %+v", 200, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildOK) GetPayload() *models.DependencyStatus {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DependencyStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildBadRequest creates a GetArchitectDependencytrackingBuildBadRequest with default headers values
func NewGetArchitectDependencytrackingBuildBadRequest() *GetArchitectDependencytrackingBuildBadRequest {
	return &GetArchitectDependencytrackingBuildBadRequest{}
}

/*
GetArchitectDependencytrackingBuildBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectDependencytrackingBuildBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build bad request response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build bad request response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build bad request response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking build bad request response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking build bad request response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetArchitectDependencytrackingBuildBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildUnauthorized creates a GetArchitectDependencytrackingBuildUnauthorized with default headers values
func NewGetArchitectDependencytrackingBuildUnauthorized() *GetArchitectDependencytrackingBuildUnauthorized {
	return &GetArchitectDependencytrackingBuildUnauthorized{}
}

/*
GetArchitectDependencytrackingBuildUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectDependencytrackingBuildUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build unauthorized response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build unauthorized response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build unauthorized response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking build unauthorized response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking build unauthorized response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetArchitectDependencytrackingBuildUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildForbidden creates a GetArchitectDependencytrackingBuildForbidden with default headers values
func NewGetArchitectDependencytrackingBuildForbidden() *GetArchitectDependencytrackingBuildForbidden {
	return &GetArchitectDependencytrackingBuildForbidden{}
}

/*
GetArchitectDependencytrackingBuildForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectDependencytrackingBuildForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build forbidden response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build forbidden response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build forbidden response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking build forbidden response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking build forbidden response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetArchitectDependencytrackingBuildForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildNotFound creates a GetArchitectDependencytrackingBuildNotFound with default headers values
func NewGetArchitectDependencytrackingBuildNotFound() *GetArchitectDependencytrackingBuildNotFound {
	return &GetArchitectDependencytrackingBuildNotFound{}
}

/*
GetArchitectDependencytrackingBuildNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetArchitectDependencytrackingBuildNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build not found response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build not found response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build not found response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking build not found response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking build not found response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetArchitectDependencytrackingBuildNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildRequestTimeout creates a GetArchitectDependencytrackingBuildRequestTimeout with default headers values
func NewGetArchitectDependencytrackingBuildRequestTimeout() *GetArchitectDependencytrackingBuildRequestTimeout {
	return &GetArchitectDependencytrackingBuildRequestTimeout{}
}

/*
GetArchitectDependencytrackingBuildRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetArchitectDependencytrackingBuildRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build request timeout response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build request timeout response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build request timeout response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking build request timeout response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking build request timeout response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetArchitectDependencytrackingBuildRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildRequestEntityTooLarge creates a GetArchitectDependencytrackingBuildRequestEntityTooLarge with default headers values
func NewGetArchitectDependencytrackingBuildRequestEntityTooLarge() *GetArchitectDependencytrackingBuildRequestEntityTooLarge {
	return &GetArchitectDependencytrackingBuildRequestEntityTooLarge{}
}

/*
GetArchitectDependencytrackingBuildRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetArchitectDependencytrackingBuildRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build request entity too large response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build request entity too large response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build request entity too large response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking build request entity too large response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking build request entity too large response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetArchitectDependencytrackingBuildRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildUnsupportedMediaType creates a GetArchitectDependencytrackingBuildUnsupportedMediaType with default headers values
func NewGetArchitectDependencytrackingBuildUnsupportedMediaType() *GetArchitectDependencytrackingBuildUnsupportedMediaType {
	return &GetArchitectDependencytrackingBuildUnsupportedMediaType{}
}

/*
GetArchitectDependencytrackingBuildUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectDependencytrackingBuildUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build unsupported media type response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build unsupported media type response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build unsupported media type response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking build unsupported media type response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking build unsupported media type response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetArchitectDependencytrackingBuildUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildTooManyRequests creates a GetArchitectDependencytrackingBuildTooManyRequests with default headers values
func NewGetArchitectDependencytrackingBuildTooManyRequests() *GetArchitectDependencytrackingBuildTooManyRequests {
	return &GetArchitectDependencytrackingBuildTooManyRequests{}
}

/*
GetArchitectDependencytrackingBuildTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetArchitectDependencytrackingBuildTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build too many requests response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build too many requests response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build too many requests response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get architect dependencytracking build too many requests response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get architect dependencytracking build too many requests response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetArchitectDependencytrackingBuildTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildInternalServerError creates a GetArchitectDependencytrackingBuildInternalServerError with default headers values
func NewGetArchitectDependencytrackingBuildInternalServerError() *GetArchitectDependencytrackingBuildInternalServerError {
	return &GetArchitectDependencytrackingBuildInternalServerError{}
}

/*
GetArchitectDependencytrackingBuildInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectDependencytrackingBuildInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build internal server error response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build internal server error response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build internal server error response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect dependencytracking build internal server error response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect dependencytracking build internal server error response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetArchitectDependencytrackingBuildInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildServiceUnavailable creates a GetArchitectDependencytrackingBuildServiceUnavailable with default headers values
func NewGetArchitectDependencytrackingBuildServiceUnavailable() *GetArchitectDependencytrackingBuildServiceUnavailable {
	return &GetArchitectDependencytrackingBuildServiceUnavailable{}
}

/*
GetArchitectDependencytrackingBuildServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectDependencytrackingBuildServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build service unavailable response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build service unavailable response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build service unavailable response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect dependencytracking build service unavailable response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect dependencytracking build service unavailable response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetArchitectDependencytrackingBuildServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectDependencytrackingBuildGatewayTimeout creates a GetArchitectDependencytrackingBuildGatewayTimeout with default headers values
func NewGetArchitectDependencytrackingBuildGatewayTimeout() *GetArchitectDependencytrackingBuildGatewayTimeout {
	return &GetArchitectDependencytrackingBuildGatewayTimeout{}
}

/*
GetArchitectDependencytrackingBuildGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetArchitectDependencytrackingBuildGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get architect dependencytracking build gateway timeout response has a 2xx status code
func (o *GetArchitectDependencytrackingBuildGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get architect dependencytracking build gateway timeout response has a 3xx status code
func (o *GetArchitectDependencytrackingBuildGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get architect dependencytracking build gateway timeout response has a 4xx status code
func (o *GetArchitectDependencytrackingBuildGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get architect dependencytracking build gateway timeout response has a 5xx status code
func (o *GetArchitectDependencytrackingBuildGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get architect dependencytracking build gateway timeout response a status code equal to that given
func (o *GetArchitectDependencytrackingBuildGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetArchitectDependencytrackingBuildGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/architect/dependencytracking/build][%d] getArchitectDependencytrackingBuildGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectDependencytrackingBuildGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectDependencytrackingBuildGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
