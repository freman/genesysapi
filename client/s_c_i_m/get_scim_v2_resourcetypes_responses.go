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

// GetScimV2ResourcetypesReader is a Reader for the GetScimV2Resourcetypes structure.
type GetScimV2ResourcetypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetScimV2ResourcetypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetScimV2ResourcetypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetScimV2ResourcetypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetScimV2ResourcetypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetScimV2ResourcetypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetScimV2ResourcetypesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetScimV2ResourcetypesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetScimV2ResourcetypesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetScimV2ResourcetypesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetScimV2ResourcetypesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetScimV2ResourcetypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetScimV2ResourcetypesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetScimV2ResourcetypesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetScimV2ResourcetypesOK creates a GetScimV2ResourcetypesOK with default headers values
func NewGetScimV2ResourcetypesOK() *GetScimV2ResourcetypesOK {
	return &GetScimV2ResourcetypesOK{}
}

/*
GetScimV2ResourcetypesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetScimV2ResourcetypesOK struct {
	Payload *models.ScimConfigResourceTypesListResponse
}

// IsSuccess returns true when this get scim v2 resourcetypes o k response has a 2xx status code
func (o *GetScimV2ResourcetypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get scim v2 resourcetypes o k response has a 3xx status code
func (o *GetScimV2ResourcetypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes o k response has a 4xx status code
func (o *GetScimV2ResourcetypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 resourcetypes o k response has a 5xx status code
func (o *GetScimV2ResourcetypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 resourcetypes o k response a status code equal to that given
func (o *GetScimV2ResourcetypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetScimV2ResourcetypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesOK  %+v", 200, o.Payload)
}

func (o *GetScimV2ResourcetypesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesOK  %+v", 200, o.Payload)
}

func (o *GetScimV2ResourcetypesOK) GetPayload() *models.ScimConfigResourceTypesListResponse {
	return o.Payload
}

func (o *GetScimV2ResourcetypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScimConfigResourceTypesListResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesBadRequest creates a GetScimV2ResourcetypesBadRequest with default headers values
func NewGetScimV2ResourcetypesBadRequest() *GetScimV2ResourcetypesBadRequest {
	return &GetScimV2ResourcetypesBadRequest{}
}

/*
GetScimV2ResourcetypesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetScimV2ResourcetypesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes bad request response has a 2xx status code
func (o *GetScimV2ResourcetypesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes bad request response has a 3xx status code
func (o *GetScimV2ResourcetypesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes bad request response has a 4xx status code
func (o *GetScimV2ResourcetypesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 resourcetypes bad request response has a 5xx status code
func (o *GetScimV2ResourcetypesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 resourcetypes bad request response a status code equal to that given
func (o *GetScimV2ResourcetypesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetScimV2ResourcetypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimV2ResourcetypesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesBadRequest  %+v", 400, o.Payload)
}

func (o *GetScimV2ResourcetypesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesUnauthorized creates a GetScimV2ResourcetypesUnauthorized with default headers values
func NewGetScimV2ResourcetypesUnauthorized() *GetScimV2ResourcetypesUnauthorized {
	return &GetScimV2ResourcetypesUnauthorized{}
}

/*
GetScimV2ResourcetypesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetScimV2ResourcetypesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes unauthorized response has a 2xx status code
func (o *GetScimV2ResourcetypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes unauthorized response has a 3xx status code
func (o *GetScimV2ResourcetypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes unauthorized response has a 4xx status code
func (o *GetScimV2ResourcetypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 resourcetypes unauthorized response has a 5xx status code
func (o *GetScimV2ResourcetypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 resourcetypes unauthorized response a status code equal to that given
func (o *GetScimV2ResourcetypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetScimV2ResourcetypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimV2ResourcetypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetScimV2ResourcetypesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesForbidden creates a GetScimV2ResourcetypesForbidden with default headers values
func NewGetScimV2ResourcetypesForbidden() *GetScimV2ResourcetypesForbidden {
	return &GetScimV2ResourcetypesForbidden{}
}

/*
GetScimV2ResourcetypesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetScimV2ResourcetypesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes forbidden response has a 2xx status code
func (o *GetScimV2ResourcetypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes forbidden response has a 3xx status code
func (o *GetScimV2ResourcetypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes forbidden response has a 4xx status code
func (o *GetScimV2ResourcetypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 resourcetypes forbidden response has a 5xx status code
func (o *GetScimV2ResourcetypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 resourcetypes forbidden response a status code equal to that given
func (o *GetScimV2ResourcetypesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetScimV2ResourcetypesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesForbidden  %+v", 403, o.Payload)
}

func (o *GetScimV2ResourcetypesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesForbidden  %+v", 403, o.Payload)
}

func (o *GetScimV2ResourcetypesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesNotFound creates a GetScimV2ResourcetypesNotFound with default headers values
func NewGetScimV2ResourcetypesNotFound() *GetScimV2ResourcetypesNotFound {
	return &GetScimV2ResourcetypesNotFound{}
}

/*
GetScimV2ResourcetypesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetScimV2ResourcetypesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes not found response has a 2xx status code
func (o *GetScimV2ResourcetypesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes not found response has a 3xx status code
func (o *GetScimV2ResourcetypesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes not found response has a 4xx status code
func (o *GetScimV2ResourcetypesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 resourcetypes not found response has a 5xx status code
func (o *GetScimV2ResourcetypesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 resourcetypes not found response a status code equal to that given
func (o *GetScimV2ResourcetypesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetScimV2ResourcetypesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesNotFound  %+v", 404, o.Payload)
}

func (o *GetScimV2ResourcetypesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesNotFound  %+v", 404, o.Payload)
}

func (o *GetScimV2ResourcetypesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesRequestTimeout creates a GetScimV2ResourcetypesRequestTimeout with default headers values
func NewGetScimV2ResourcetypesRequestTimeout() *GetScimV2ResourcetypesRequestTimeout {
	return &GetScimV2ResourcetypesRequestTimeout{}
}

/*
GetScimV2ResourcetypesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetScimV2ResourcetypesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes request timeout response has a 2xx status code
func (o *GetScimV2ResourcetypesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes request timeout response has a 3xx status code
func (o *GetScimV2ResourcetypesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes request timeout response has a 4xx status code
func (o *GetScimV2ResourcetypesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 resourcetypes request timeout response has a 5xx status code
func (o *GetScimV2ResourcetypesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 resourcetypes request timeout response a status code equal to that given
func (o *GetScimV2ResourcetypesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetScimV2ResourcetypesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScimV2ResourcetypesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetScimV2ResourcetypesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesRequestEntityTooLarge creates a GetScimV2ResourcetypesRequestEntityTooLarge with default headers values
func NewGetScimV2ResourcetypesRequestEntityTooLarge() *GetScimV2ResourcetypesRequestEntityTooLarge {
	return &GetScimV2ResourcetypesRequestEntityTooLarge{}
}

/*
GetScimV2ResourcetypesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetScimV2ResourcetypesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes request entity too large response has a 2xx status code
func (o *GetScimV2ResourcetypesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes request entity too large response has a 3xx status code
func (o *GetScimV2ResourcetypesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes request entity too large response has a 4xx status code
func (o *GetScimV2ResourcetypesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 resourcetypes request entity too large response has a 5xx status code
func (o *GetScimV2ResourcetypesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 resourcetypes request entity too large response a status code equal to that given
func (o *GetScimV2ResourcetypesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetScimV2ResourcetypesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimV2ResourcetypesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetScimV2ResourcetypesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesUnsupportedMediaType creates a GetScimV2ResourcetypesUnsupportedMediaType with default headers values
func NewGetScimV2ResourcetypesUnsupportedMediaType() *GetScimV2ResourcetypesUnsupportedMediaType {
	return &GetScimV2ResourcetypesUnsupportedMediaType{}
}

/*
GetScimV2ResourcetypesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetScimV2ResourcetypesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes unsupported media type response has a 2xx status code
func (o *GetScimV2ResourcetypesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes unsupported media type response has a 3xx status code
func (o *GetScimV2ResourcetypesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes unsupported media type response has a 4xx status code
func (o *GetScimV2ResourcetypesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 resourcetypes unsupported media type response has a 5xx status code
func (o *GetScimV2ResourcetypesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 resourcetypes unsupported media type response a status code equal to that given
func (o *GetScimV2ResourcetypesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetScimV2ResourcetypesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimV2ResourcetypesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetScimV2ResourcetypesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesTooManyRequests creates a GetScimV2ResourcetypesTooManyRequests with default headers values
func NewGetScimV2ResourcetypesTooManyRequests() *GetScimV2ResourcetypesTooManyRequests {
	return &GetScimV2ResourcetypesTooManyRequests{}
}

/*
GetScimV2ResourcetypesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetScimV2ResourcetypesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes too many requests response has a 2xx status code
func (o *GetScimV2ResourcetypesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes too many requests response has a 3xx status code
func (o *GetScimV2ResourcetypesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes too many requests response has a 4xx status code
func (o *GetScimV2ResourcetypesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get scim v2 resourcetypes too many requests response has a 5xx status code
func (o *GetScimV2ResourcetypesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get scim v2 resourcetypes too many requests response a status code equal to that given
func (o *GetScimV2ResourcetypesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetScimV2ResourcetypesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimV2ResourcetypesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetScimV2ResourcetypesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesInternalServerError creates a GetScimV2ResourcetypesInternalServerError with default headers values
func NewGetScimV2ResourcetypesInternalServerError() *GetScimV2ResourcetypesInternalServerError {
	return &GetScimV2ResourcetypesInternalServerError{}
}

/*
GetScimV2ResourcetypesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetScimV2ResourcetypesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes internal server error response has a 2xx status code
func (o *GetScimV2ResourcetypesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes internal server error response has a 3xx status code
func (o *GetScimV2ResourcetypesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes internal server error response has a 4xx status code
func (o *GetScimV2ResourcetypesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 resourcetypes internal server error response has a 5xx status code
func (o *GetScimV2ResourcetypesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get scim v2 resourcetypes internal server error response a status code equal to that given
func (o *GetScimV2ResourcetypesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetScimV2ResourcetypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimV2ResourcetypesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetScimV2ResourcetypesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesServiceUnavailable creates a GetScimV2ResourcetypesServiceUnavailable with default headers values
func NewGetScimV2ResourcetypesServiceUnavailable() *GetScimV2ResourcetypesServiceUnavailable {
	return &GetScimV2ResourcetypesServiceUnavailable{}
}

/*
GetScimV2ResourcetypesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetScimV2ResourcetypesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes service unavailable response has a 2xx status code
func (o *GetScimV2ResourcetypesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes service unavailable response has a 3xx status code
func (o *GetScimV2ResourcetypesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes service unavailable response has a 4xx status code
func (o *GetScimV2ResourcetypesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 resourcetypes service unavailable response has a 5xx status code
func (o *GetScimV2ResourcetypesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get scim v2 resourcetypes service unavailable response a status code equal to that given
func (o *GetScimV2ResourcetypesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetScimV2ResourcetypesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimV2ResourcetypesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetScimV2ResourcetypesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetScimV2ResourcetypesGatewayTimeout creates a GetScimV2ResourcetypesGatewayTimeout with default headers values
func NewGetScimV2ResourcetypesGatewayTimeout() *GetScimV2ResourcetypesGatewayTimeout {
	return &GetScimV2ResourcetypesGatewayTimeout{}
}

/*
GetScimV2ResourcetypesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetScimV2ResourcetypesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get scim v2 resourcetypes gateway timeout response has a 2xx status code
func (o *GetScimV2ResourcetypesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get scim v2 resourcetypes gateway timeout response has a 3xx status code
func (o *GetScimV2ResourcetypesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get scim v2 resourcetypes gateway timeout response has a 4xx status code
func (o *GetScimV2ResourcetypesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get scim v2 resourcetypes gateway timeout response has a 5xx status code
func (o *GetScimV2ResourcetypesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get scim v2 resourcetypes gateway timeout response a status code equal to that given
func (o *GetScimV2ResourcetypesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetScimV2ResourcetypesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimV2ResourcetypesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/scim/v2/resourcetypes][%d] getScimV2ResourcetypesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetScimV2ResourcetypesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetScimV2ResourcetypesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
