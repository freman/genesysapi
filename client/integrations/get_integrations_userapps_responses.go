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

// GetIntegrationsUserappsReader is a Reader for the GetIntegrationsUserapps structure.
type GetIntegrationsUserappsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetIntegrationsUserappsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetIntegrationsUserappsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetIntegrationsUserappsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetIntegrationsUserappsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetIntegrationsUserappsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetIntegrationsUserappsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetIntegrationsUserappsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetIntegrationsUserappsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetIntegrationsUserappsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetIntegrationsUserappsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetIntegrationsUserappsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetIntegrationsUserappsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetIntegrationsUserappsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetIntegrationsUserappsOK creates a GetIntegrationsUserappsOK with default headers values
func NewGetIntegrationsUserappsOK() *GetIntegrationsUserappsOK {
	return &GetIntegrationsUserappsOK{}
}

/*
GetIntegrationsUserappsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetIntegrationsUserappsOK struct {
	Payload *models.UserAppEntityListing
}

// IsSuccess returns true when this get integrations userapps o k response has a 2xx status code
func (o *GetIntegrationsUserappsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get integrations userapps o k response has a 3xx status code
func (o *GetIntegrationsUserappsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps o k response has a 4xx status code
func (o *GetIntegrationsUserappsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations userapps o k response has a 5xx status code
func (o *GetIntegrationsUserappsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations userapps o k response a status code equal to that given
func (o *GetIntegrationsUserappsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetIntegrationsUserappsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsUserappsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsOK  %+v", 200, o.Payload)
}

func (o *GetIntegrationsUserappsOK) GetPayload() *models.UserAppEntityListing {
	return o.Payload
}

func (o *GetIntegrationsUserappsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAppEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsBadRequest creates a GetIntegrationsUserappsBadRequest with default headers values
func NewGetIntegrationsUserappsBadRequest() *GetIntegrationsUserappsBadRequest {
	return &GetIntegrationsUserappsBadRequest{}
}

/*
GetIntegrationsUserappsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetIntegrationsUserappsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps bad request response has a 2xx status code
func (o *GetIntegrationsUserappsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps bad request response has a 3xx status code
func (o *GetIntegrationsUserappsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps bad request response has a 4xx status code
func (o *GetIntegrationsUserappsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations userapps bad request response has a 5xx status code
func (o *GetIntegrationsUserappsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations userapps bad request response a status code equal to that given
func (o *GetIntegrationsUserappsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetIntegrationsUserappsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsUserappsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsBadRequest  %+v", 400, o.Payload)
}

func (o *GetIntegrationsUserappsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsUnauthorized creates a GetIntegrationsUserappsUnauthorized with default headers values
func NewGetIntegrationsUserappsUnauthorized() *GetIntegrationsUserappsUnauthorized {
	return &GetIntegrationsUserappsUnauthorized{}
}

/*
GetIntegrationsUserappsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetIntegrationsUserappsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps unauthorized response has a 2xx status code
func (o *GetIntegrationsUserappsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps unauthorized response has a 3xx status code
func (o *GetIntegrationsUserappsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps unauthorized response has a 4xx status code
func (o *GetIntegrationsUserappsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations userapps unauthorized response has a 5xx status code
func (o *GetIntegrationsUserappsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations userapps unauthorized response a status code equal to that given
func (o *GetIntegrationsUserappsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetIntegrationsUserappsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsUserappsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetIntegrationsUserappsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsForbidden creates a GetIntegrationsUserappsForbidden with default headers values
func NewGetIntegrationsUserappsForbidden() *GetIntegrationsUserappsForbidden {
	return &GetIntegrationsUserappsForbidden{}
}

/*
GetIntegrationsUserappsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetIntegrationsUserappsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps forbidden response has a 2xx status code
func (o *GetIntegrationsUserappsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps forbidden response has a 3xx status code
func (o *GetIntegrationsUserappsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps forbidden response has a 4xx status code
func (o *GetIntegrationsUserappsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations userapps forbidden response has a 5xx status code
func (o *GetIntegrationsUserappsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations userapps forbidden response a status code equal to that given
func (o *GetIntegrationsUserappsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetIntegrationsUserappsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsUserappsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsForbidden  %+v", 403, o.Payload)
}

func (o *GetIntegrationsUserappsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsNotFound creates a GetIntegrationsUserappsNotFound with default headers values
func NewGetIntegrationsUserappsNotFound() *GetIntegrationsUserappsNotFound {
	return &GetIntegrationsUserappsNotFound{}
}

/*
GetIntegrationsUserappsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetIntegrationsUserappsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps not found response has a 2xx status code
func (o *GetIntegrationsUserappsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps not found response has a 3xx status code
func (o *GetIntegrationsUserappsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps not found response has a 4xx status code
func (o *GetIntegrationsUserappsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations userapps not found response has a 5xx status code
func (o *GetIntegrationsUserappsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations userapps not found response a status code equal to that given
func (o *GetIntegrationsUserappsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetIntegrationsUserappsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsUserappsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsNotFound  %+v", 404, o.Payload)
}

func (o *GetIntegrationsUserappsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsRequestTimeout creates a GetIntegrationsUserappsRequestTimeout with default headers values
func NewGetIntegrationsUserappsRequestTimeout() *GetIntegrationsUserappsRequestTimeout {
	return &GetIntegrationsUserappsRequestTimeout{}
}

/*
GetIntegrationsUserappsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetIntegrationsUserappsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps request timeout response has a 2xx status code
func (o *GetIntegrationsUserappsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps request timeout response has a 3xx status code
func (o *GetIntegrationsUserappsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps request timeout response has a 4xx status code
func (o *GetIntegrationsUserappsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations userapps request timeout response has a 5xx status code
func (o *GetIntegrationsUserappsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations userapps request timeout response a status code equal to that given
func (o *GetIntegrationsUserappsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetIntegrationsUserappsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsUserappsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetIntegrationsUserappsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsRequestEntityTooLarge creates a GetIntegrationsUserappsRequestEntityTooLarge with default headers values
func NewGetIntegrationsUserappsRequestEntityTooLarge() *GetIntegrationsUserappsRequestEntityTooLarge {
	return &GetIntegrationsUserappsRequestEntityTooLarge{}
}

/*
GetIntegrationsUserappsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetIntegrationsUserappsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps request entity too large response has a 2xx status code
func (o *GetIntegrationsUserappsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps request entity too large response has a 3xx status code
func (o *GetIntegrationsUserappsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps request entity too large response has a 4xx status code
func (o *GetIntegrationsUserappsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations userapps request entity too large response has a 5xx status code
func (o *GetIntegrationsUserappsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations userapps request entity too large response a status code equal to that given
func (o *GetIntegrationsUserappsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetIntegrationsUserappsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsUserappsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetIntegrationsUserappsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsUnsupportedMediaType creates a GetIntegrationsUserappsUnsupportedMediaType with default headers values
func NewGetIntegrationsUserappsUnsupportedMediaType() *GetIntegrationsUserappsUnsupportedMediaType {
	return &GetIntegrationsUserappsUnsupportedMediaType{}
}

/*
GetIntegrationsUserappsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetIntegrationsUserappsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps unsupported media type response has a 2xx status code
func (o *GetIntegrationsUserappsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps unsupported media type response has a 3xx status code
func (o *GetIntegrationsUserappsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps unsupported media type response has a 4xx status code
func (o *GetIntegrationsUserappsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations userapps unsupported media type response has a 5xx status code
func (o *GetIntegrationsUserappsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations userapps unsupported media type response a status code equal to that given
func (o *GetIntegrationsUserappsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetIntegrationsUserappsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsUserappsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetIntegrationsUserappsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsTooManyRequests creates a GetIntegrationsUserappsTooManyRequests with default headers values
func NewGetIntegrationsUserappsTooManyRequests() *GetIntegrationsUserappsTooManyRequests {
	return &GetIntegrationsUserappsTooManyRequests{}
}

/*
GetIntegrationsUserappsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetIntegrationsUserappsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps too many requests response has a 2xx status code
func (o *GetIntegrationsUserappsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps too many requests response has a 3xx status code
func (o *GetIntegrationsUserappsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps too many requests response has a 4xx status code
func (o *GetIntegrationsUserappsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get integrations userapps too many requests response has a 5xx status code
func (o *GetIntegrationsUserappsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get integrations userapps too many requests response a status code equal to that given
func (o *GetIntegrationsUserappsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetIntegrationsUserappsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsUserappsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetIntegrationsUserappsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsInternalServerError creates a GetIntegrationsUserappsInternalServerError with default headers values
func NewGetIntegrationsUserappsInternalServerError() *GetIntegrationsUserappsInternalServerError {
	return &GetIntegrationsUserappsInternalServerError{}
}

/*
GetIntegrationsUserappsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetIntegrationsUserappsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps internal server error response has a 2xx status code
func (o *GetIntegrationsUserappsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps internal server error response has a 3xx status code
func (o *GetIntegrationsUserappsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps internal server error response has a 4xx status code
func (o *GetIntegrationsUserappsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations userapps internal server error response has a 5xx status code
func (o *GetIntegrationsUserappsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations userapps internal server error response a status code equal to that given
func (o *GetIntegrationsUserappsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetIntegrationsUserappsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsUserappsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetIntegrationsUserappsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsServiceUnavailable creates a GetIntegrationsUserappsServiceUnavailable with default headers values
func NewGetIntegrationsUserappsServiceUnavailable() *GetIntegrationsUserappsServiceUnavailable {
	return &GetIntegrationsUserappsServiceUnavailable{}
}

/*
GetIntegrationsUserappsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetIntegrationsUserappsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps service unavailable response has a 2xx status code
func (o *GetIntegrationsUserappsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps service unavailable response has a 3xx status code
func (o *GetIntegrationsUserappsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps service unavailable response has a 4xx status code
func (o *GetIntegrationsUserappsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations userapps service unavailable response has a 5xx status code
func (o *GetIntegrationsUserappsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations userapps service unavailable response a status code equal to that given
func (o *GetIntegrationsUserappsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetIntegrationsUserappsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsUserappsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetIntegrationsUserappsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetIntegrationsUserappsGatewayTimeout creates a GetIntegrationsUserappsGatewayTimeout with default headers values
func NewGetIntegrationsUserappsGatewayTimeout() *GetIntegrationsUserappsGatewayTimeout {
	return &GetIntegrationsUserappsGatewayTimeout{}
}

/*
GetIntegrationsUserappsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetIntegrationsUserappsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get integrations userapps gateway timeout response has a 2xx status code
func (o *GetIntegrationsUserappsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get integrations userapps gateway timeout response has a 3xx status code
func (o *GetIntegrationsUserappsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get integrations userapps gateway timeout response has a 4xx status code
func (o *GetIntegrationsUserappsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get integrations userapps gateway timeout response has a 5xx status code
func (o *GetIntegrationsUserappsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get integrations userapps gateway timeout response a status code equal to that given
func (o *GetIntegrationsUserappsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetIntegrationsUserappsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsUserappsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/integrations/userapps][%d] getIntegrationsUserappsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetIntegrationsUserappsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetIntegrationsUserappsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
