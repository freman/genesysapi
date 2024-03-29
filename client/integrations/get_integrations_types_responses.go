// Code generated by go-swagger; DO NOT EDIT.

package integrations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetIntegrationsTypesReader is a Reader for the GetIntegrationsTypes structure.
type GetIntegrationsTypesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsTypesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsTypesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsTypesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsTypesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsTypesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsTypesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsTypesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsTypesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsTypesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsTypesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsTypesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsTypesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsTypesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsTypesOK creates a GetIntegrationsTypesOK with default headers values
func NewGetIntegrationsTypesOK() *GetIntegrationsTypesOK {
	return &GetIntegrationsTypesOK{}
}

/*
GetIntegrationsTypesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIntegrationsTypesOK struct {
	Payload *models.IntegrationTypeEntityListing
}

// IsSuccess returns true when this get integrations types o k response has a 2xx status code
func (o *GetIntegrationsTypesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrations types o k response has a 3xx status code
func (o *GetIntegrationsTypesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types o k response has a 4xx status code
func (o *GetIntegrationsTypesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations types o k response has a 5xx status code
func (o *GetIntegrationsTypesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations types o k response a status code equal to that given
func (o *GetIntegrationsTypesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegrationsTypesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsTypesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsTypesOK) GetPayload() *models.IntegrationTypeEntityListing {
	return o.Payload
}

func (o *GetIntegrationsTypesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IntegrationTypeEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesBadRequest creates a GetIntegrationsTypesBadRequest with default headers values
func NewGetIntegrationsTypesBadRequest() *GetIntegrationsTypesBadRequest {
	return &GetIntegrationsTypesBadRequest{}
}

/*
GetIntegrationsTypesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsTypesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types bad request response has a 2xx status code
func (o *GetIntegrationsTypesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types bad request response has a 3xx status code
func (o *GetIntegrationsTypesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types bad request response has a 4xx status code
func (o *GetIntegrationsTypesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations types bad request response has a 5xx status code
func (o *GetIntegrationsTypesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations types bad request response a status code equal to that given
func (o *GetIntegrationsTypesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIntegrationsTypesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsTypesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsTypesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesUnauthorized creates a GetIntegrationsTypesUnauthorized with default headers values
func NewGetIntegrationsTypesUnauthorized() *GetIntegrationsTypesUnauthorized {
	return &GetIntegrationsTypesUnauthorized{}
}

/*
GetIntegrationsTypesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsTypesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types unauthorized response has a 2xx status code
func (o *GetIntegrationsTypesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types unauthorized response has a 3xx status code
func (o *GetIntegrationsTypesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types unauthorized response has a 4xx status code
func (o *GetIntegrationsTypesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations types unauthorized response has a 5xx status code
func (o *GetIntegrationsTypesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations types unauthorized response a status code equal to that given
func (o *GetIntegrationsTypesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIntegrationsTypesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsTypesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsTypesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesForbidden creates a GetIntegrationsTypesForbidden with default headers values
func NewGetIntegrationsTypesForbidden() *GetIntegrationsTypesForbidden {
	return &GetIntegrationsTypesForbidden{}
}

/*
GetIntegrationsTypesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsTypesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types forbidden response has a 2xx status code
func (o *GetIntegrationsTypesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types forbidden response has a 3xx status code
func (o *GetIntegrationsTypesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types forbidden response has a 4xx status code
func (o *GetIntegrationsTypesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations types forbidden response has a 5xx status code
func (o *GetIntegrationsTypesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations types forbidden response a status code equal to that given
func (o *GetIntegrationsTypesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntegrationsTypesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsTypesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsTypesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesNotFound creates a GetIntegrationsTypesNotFound with default headers values
func NewGetIntegrationsTypesNotFound() *GetIntegrationsTypesNotFound {
	return &GetIntegrationsTypesNotFound{}
}

/*
GetIntegrationsTypesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetIntegrationsTypesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types not found response has a 2xx status code
func (o *GetIntegrationsTypesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types not found response has a 3xx status code
func (o *GetIntegrationsTypesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types not found response has a 4xx status code
func (o *GetIntegrationsTypesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations types not found response has a 5xx status code
func (o *GetIntegrationsTypesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations types not found response a status code equal to that given
func (o *GetIntegrationsTypesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIntegrationsTypesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsTypesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsTypesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesRequestTimeout creates a GetIntegrationsTypesRequestTimeout with default headers values
func NewGetIntegrationsTypesRequestTimeout() *GetIntegrationsTypesRequestTimeout {
	return &GetIntegrationsTypesRequestTimeout{}
}

/*
GetIntegrationsTypesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsTypesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types request timeout response has a 2xx status code
func (o *GetIntegrationsTypesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types request timeout response has a 3xx status code
func (o *GetIntegrationsTypesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types request timeout response has a 4xx status code
func (o *GetIntegrationsTypesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations types request timeout response has a 5xx status code
func (o *GetIntegrationsTypesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations types request timeout response a status code equal to that given
func (o *GetIntegrationsTypesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetIntegrationsTypesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsTypesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsTypesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesRequestEntityTooLarge creates a GetIntegrationsTypesRequestEntityTooLarge with default headers values
func NewGetIntegrationsTypesRequestEntityTooLarge() *GetIntegrationsTypesRequestEntityTooLarge {
	return &GetIntegrationsTypesRequestEntityTooLarge{}
}

/*
GetIntegrationsTypesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetIntegrationsTypesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types request entity too large response has a 2xx status code
func (o *GetIntegrationsTypesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types request entity too large response has a 3xx status code
func (o *GetIntegrationsTypesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types request entity too large response has a 4xx status code
func (o *GetIntegrationsTypesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations types request entity too large response has a 5xx status code
func (o *GetIntegrationsTypesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations types request entity too large response a status code equal to that given
func (o *GetIntegrationsTypesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetIntegrationsTypesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsTypesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsTypesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesUnsupportedMediaType creates a GetIntegrationsTypesUnsupportedMediaType with default headers values
func NewGetIntegrationsTypesUnsupportedMediaType() *GetIntegrationsTypesUnsupportedMediaType {
	return &GetIntegrationsTypesUnsupportedMediaType{}
}

/*
GetIntegrationsTypesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsTypesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types unsupported media type response has a 2xx status code
func (o *GetIntegrationsTypesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types unsupported media type response has a 3xx status code
func (o *GetIntegrationsTypesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types unsupported media type response has a 4xx status code
func (o *GetIntegrationsTypesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations types unsupported media type response has a 5xx status code
func (o *GetIntegrationsTypesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations types unsupported media type response a status code equal to that given
func (o *GetIntegrationsTypesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetIntegrationsTypesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsTypesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsTypesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesTooManyRequests creates a GetIntegrationsTypesTooManyRequests with default headers values
func NewGetIntegrationsTypesTooManyRequests() *GetIntegrationsTypesTooManyRequests {
	return &GetIntegrationsTypesTooManyRequests{}
}

/*
GetIntegrationsTypesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsTypesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types too many requests response has a 2xx status code
func (o *GetIntegrationsTypesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types too many requests response has a 3xx status code
func (o *GetIntegrationsTypesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types too many requests response has a 4xx status code
func (o *GetIntegrationsTypesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations types too many requests response has a 5xx status code
func (o *GetIntegrationsTypesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations types too many requests response a status code equal to that given
func (o *GetIntegrationsTypesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntegrationsTypesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsTypesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsTypesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesInternalServerError creates a GetIntegrationsTypesInternalServerError with default headers values
func NewGetIntegrationsTypesInternalServerError() *GetIntegrationsTypesInternalServerError {
	return &GetIntegrationsTypesInternalServerError{}
}

/*
GetIntegrationsTypesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsTypesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types internal server error response has a 2xx status code
func (o *GetIntegrationsTypesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types internal server error response has a 3xx status code
func (o *GetIntegrationsTypesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types internal server error response has a 4xx status code
func (o *GetIntegrationsTypesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations types internal server error response has a 5xx status code
func (o *GetIntegrationsTypesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations types internal server error response a status code equal to that given
func (o *GetIntegrationsTypesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntegrationsTypesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsTypesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsTypesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesServiceUnavailable creates a GetIntegrationsTypesServiceUnavailable with default headers values
func NewGetIntegrationsTypesServiceUnavailable() *GetIntegrationsTypesServiceUnavailable {
	return &GetIntegrationsTypesServiceUnavailable{}
}

/*
GetIntegrationsTypesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsTypesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types service unavailable response has a 2xx status code
func (o *GetIntegrationsTypesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types service unavailable response has a 3xx status code
func (o *GetIntegrationsTypesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types service unavailable response has a 4xx status code
func (o *GetIntegrationsTypesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations types service unavailable response has a 5xx status code
func (o *GetIntegrationsTypesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations types service unavailable response a status code equal to that given
func (o *GetIntegrationsTypesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIntegrationsTypesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsTypesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsTypesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsTypesGatewayTimeout creates a GetIntegrationsTypesGatewayTimeout with default headers values
func NewGetIntegrationsTypesGatewayTimeout() *GetIntegrationsTypesGatewayTimeout {
	return &GetIntegrationsTypesGatewayTimeout{}
}

/*
GetIntegrationsTypesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetIntegrationsTypesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations types gateway timeout response has a 2xx status code
func (o *GetIntegrationsTypesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations types gateway timeout response has a 3xx status code
func (o *GetIntegrationsTypesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations types gateway timeout response has a 4xx status code
func (o *GetIntegrationsTypesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations types gateway timeout response has a 5xx status code
func (o *GetIntegrationsTypesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations types gateway timeout response a status code equal to that given
func (o *GetIntegrationsTypesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIntegrationsTypesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsTypesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/types][%d] getIntegrationsTypesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsTypesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsTypesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
