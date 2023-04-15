// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetRoutingSettingsContactcenterReader is a Reader for the GetRoutingSettingsContactcenter structure.
type GetRoutingSettingsContactcenterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSettingsContactcenterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSettingsContactcenterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSettingsContactcenterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSettingsContactcenterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSettingsContactcenterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSettingsContactcenterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingSettingsContactcenterRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSettingsContactcenterRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSettingsContactcenterUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSettingsContactcenterTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSettingsContactcenterInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSettingsContactcenterServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSettingsContactcenterGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSettingsContactcenterOK creates a GetRoutingSettingsContactcenterOK with default headers values
func NewGetRoutingSettingsContactcenterOK() *GetRoutingSettingsContactcenterOK {
	return &GetRoutingSettingsContactcenterOK{}
}

/*
GetRoutingSettingsContactcenterOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingSettingsContactcenterOK struct {
	Payload *models.ContactCenterSettings
}

// IsSuccess returns true when this get routing settings contactcenter o k response has a 2xx status code
func (o *GetRoutingSettingsContactcenterOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing settings contactcenter o k response has a 3xx status code
func (o *GetRoutingSettingsContactcenterOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter o k response has a 4xx status code
func (o *GetRoutingSettingsContactcenterOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing settings contactcenter o k response has a 5xx status code
func (o *GetRoutingSettingsContactcenterOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings contactcenter o k response a status code equal to that given
func (o *GetRoutingSettingsContactcenterOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingSettingsContactcenterOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSettingsContactcenterOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSettingsContactcenterOK) GetPayload() *models.ContactCenterSettings {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactCenterSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterBadRequest creates a GetRoutingSettingsContactcenterBadRequest with default headers values
func NewGetRoutingSettingsContactcenterBadRequest() *GetRoutingSettingsContactcenterBadRequest {
	return &GetRoutingSettingsContactcenterBadRequest{}
}

/*
GetRoutingSettingsContactcenterBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSettingsContactcenterBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter bad request response has a 2xx status code
func (o *GetRoutingSettingsContactcenterBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter bad request response has a 3xx status code
func (o *GetRoutingSettingsContactcenterBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter bad request response has a 4xx status code
func (o *GetRoutingSettingsContactcenterBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings contactcenter bad request response has a 5xx status code
func (o *GetRoutingSettingsContactcenterBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings contactcenter bad request response a status code equal to that given
func (o *GetRoutingSettingsContactcenterBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingSettingsContactcenterBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSettingsContactcenterBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSettingsContactcenterBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterUnauthorized creates a GetRoutingSettingsContactcenterUnauthorized with default headers values
func NewGetRoutingSettingsContactcenterUnauthorized() *GetRoutingSettingsContactcenterUnauthorized {
	return &GetRoutingSettingsContactcenterUnauthorized{}
}

/*
GetRoutingSettingsContactcenterUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSettingsContactcenterUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter unauthorized response has a 2xx status code
func (o *GetRoutingSettingsContactcenterUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter unauthorized response has a 3xx status code
func (o *GetRoutingSettingsContactcenterUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter unauthorized response has a 4xx status code
func (o *GetRoutingSettingsContactcenterUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings contactcenter unauthorized response has a 5xx status code
func (o *GetRoutingSettingsContactcenterUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings contactcenter unauthorized response a status code equal to that given
func (o *GetRoutingSettingsContactcenterUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingSettingsContactcenterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSettingsContactcenterUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSettingsContactcenterUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterForbidden creates a GetRoutingSettingsContactcenterForbidden with default headers values
func NewGetRoutingSettingsContactcenterForbidden() *GetRoutingSettingsContactcenterForbidden {
	return &GetRoutingSettingsContactcenterForbidden{}
}

/*
GetRoutingSettingsContactcenterForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSettingsContactcenterForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter forbidden response has a 2xx status code
func (o *GetRoutingSettingsContactcenterForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter forbidden response has a 3xx status code
func (o *GetRoutingSettingsContactcenterForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter forbidden response has a 4xx status code
func (o *GetRoutingSettingsContactcenterForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings contactcenter forbidden response has a 5xx status code
func (o *GetRoutingSettingsContactcenterForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings contactcenter forbidden response a status code equal to that given
func (o *GetRoutingSettingsContactcenterForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingSettingsContactcenterForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSettingsContactcenterForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSettingsContactcenterForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterNotFound creates a GetRoutingSettingsContactcenterNotFound with default headers values
func NewGetRoutingSettingsContactcenterNotFound() *GetRoutingSettingsContactcenterNotFound {
	return &GetRoutingSettingsContactcenterNotFound{}
}

/*
GetRoutingSettingsContactcenterNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingSettingsContactcenterNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter not found response has a 2xx status code
func (o *GetRoutingSettingsContactcenterNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter not found response has a 3xx status code
func (o *GetRoutingSettingsContactcenterNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter not found response has a 4xx status code
func (o *GetRoutingSettingsContactcenterNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings contactcenter not found response has a 5xx status code
func (o *GetRoutingSettingsContactcenterNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings contactcenter not found response a status code equal to that given
func (o *GetRoutingSettingsContactcenterNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingSettingsContactcenterNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSettingsContactcenterNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSettingsContactcenterNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterRequestTimeout creates a GetRoutingSettingsContactcenterRequestTimeout with default headers values
func NewGetRoutingSettingsContactcenterRequestTimeout() *GetRoutingSettingsContactcenterRequestTimeout {
	return &GetRoutingSettingsContactcenterRequestTimeout{}
}

/*
GetRoutingSettingsContactcenterRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingSettingsContactcenterRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter request timeout response has a 2xx status code
func (o *GetRoutingSettingsContactcenterRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter request timeout response has a 3xx status code
func (o *GetRoutingSettingsContactcenterRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter request timeout response has a 4xx status code
func (o *GetRoutingSettingsContactcenterRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings contactcenter request timeout response has a 5xx status code
func (o *GetRoutingSettingsContactcenterRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings contactcenter request timeout response a status code equal to that given
func (o *GetRoutingSettingsContactcenterRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingSettingsContactcenterRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSettingsContactcenterRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingSettingsContactcenterRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterRequestEntityTooLarge creates a GetRoutingSettingsContactcenterRequestEntityTooLarge with default headers values
func NewGetRoutingSettingsContactcenterRequestEntityTooLarge() *GetRoutingSettingsContactcenterRequestEntityTooLarge {
	return &GetRoutingSettingsContactcenterRequestEntityTooLarge{}
}

/*
GetRoutingSettingsContactcenterRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingSettingsContactcenterRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter request entity too large response has a 2xx status code
func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter request entity too large response has a 3xx status code
func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter request entity too large response has a 4xx status code
func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings contactcenter request entity too large response has a 5xx status code
func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings contactcenter request entity too large response a status code equal to that given
func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterUnsupportedMediaType creates a GetRoutingSettingsContactcenterUnsupportedMediaType with default headers values
func NewGetRoutingSettingsContactcenterUnsupportedMediaType() *GetRoutingSettingsContactcenterUnsupportedMediaType {
	return &GetRoutingSettingsContactcenterUnsupportedMediaType{}
}

/*
GetRoutingSettingsContactcenterUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSettingsContactcenterUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter unsupported media type response has a 2xx status code
func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter unsupported media type response has a 3xx status code
func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter unsupported media type response has a 4xx status code
func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings contactcenter unsupported media type response has a 5xx status code
func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings contactcenter unsupported media type response a status code equal to that given
func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterTooManyRequests creates a GetRoutingSettingsContactcenterTooManyRequests with default headers values
func NewGetRoutingSettingsContactcenterTooManyRequests() *GetRoutingSettingsContactcenterTooManyRequests {
	return &GetRoutingSettingsContactcenterTooManyRequests{}
}

/*
GetRoutingSettingsContactcenterTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingSettingsContactcenterTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter too many requests response has a 2xx status code
func (o *GetRoutingSettingsContactcenterTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter too many requests response has a 3xx status code
func (o *GetRoutingSettingsContactcenterTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter too many requests response has a 4xx status code
func (o *GetRoutingSettingsContactcenterTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing settings contactcenter too many requests response has a 5xx status code
func (o *GetRoutingSettingsContactcenterTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing settings contactcenter too many requests response a status code equal to that given
func (o *GetRoutingSettingsContactcenterTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingSettingsContactcenterTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSettingsContactcenterTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSettingsContactcenterTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterInternalServerError creates a GetRoutingSettingsContactcenterInternalServerError with default headers values
func NewGetRoutingSettingsContactcenterInternalServerError() *GetRoutingSettingsContactcenterInternalServerError {
	return &GetRoutingSettingsContactcenterInternalServerError{}
}

/*
GetRoutingSettingsContactcenterInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSettingsContactcenterInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter internal server error response has a 2xx status code
func (o *GetRoutingSettingsContactcenterInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter internal server error response has a 3xx status code
func (o *GetRoutingSettingsContactcenterInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter internal server error response has a 4xx status code
func (o *GetRoutingSettingsContactcenterInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing settings contactcenter internal server error response has a 5xx status code
func (o *GetRoutingSettingsContactcenterInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing settings contactcenter internal server error response a status code equal to that given
func (o *GetRoutingSettingsContactcenterInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingSettingsContactcenterInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSettingsContactcenterInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSettingsContactcenterInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterServiceUnavailable creates a GetRoutingSettingsContactcenterServiceUnavailable with default headers values
func NewGetRoutingSettingsContactcenterServiceUnavailable() *GetRoutingSettingsContactcenterServiceUnavailable {
	return &GetRoutingSettingsContactcenterServiceUnavailable{}
}

/*
GetRoutingSettingsContactcenterServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSettingsContactcenterServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter service unavailable response has a 2xx status code
func (o *GetRoutingSettingsContactcenterServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter service unavailable response has a 3xx status code
func (o *GetRoutingSettingsContactcenterServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter service unavailable response has a 4xx status code
func (o *GetRoutingSettingsContactcenterServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing settings contactcenter service unavailable response has a 5xx status code
func (o *GetRoutingSettingsContactcenterServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing settings contactcenter service unavailable response a status code equal to that given
func (o *GetRoutingSettingsContactcenterServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingSettingsContactcenterServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSettingsContactcenterServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSettingsContactcenterServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsContactcenterGatewayTimeout creates a GetRoutingSettingsContactcenterGatewayTimeout with default headers values
func NewGetRoutingSettingsContactcenterGatewayTimeout() *GetRoutingSettingsContactcenterGatewayTimeout {
	return &GetRoutingSettingsContactcenterGatewayTimeout{}
}

/*
GetRoutingSettingsContactcenterGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingSettingsContactcenterGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing settings contactcenter gateway timeout response has a 2xx status code
func (o *GetRoutingSettingsContactcenterGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing settings contactcenter gateway timeout response has a 3xx status code
func (o *GetRoutingSettingsContactcenterGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing settings contactcenter gateway timeout response has a 4xx status code
func (o *GetRoutingSettingsContactcenterGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing settings contactcenter gateway timeout response has a 5xx status code
func (o *GetRoutingSettingsContactcenterGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing settings contactcenter gateway timeout response a status code equal to that given
func (o *GetRoutingSettingsContactcenterGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingSettingsContactcenterGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSettingsContactcenterGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/contactcenter][%d] getRoutingSettingsContactcenterGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSettingsContactcenterGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsContactcenterGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
