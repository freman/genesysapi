// Code generated by go-swagger; DO NOT EDIT.

package journey

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetJourneyActionmapsReader is a Reader for the GetJourneyActionmaps structure.
type GetJourneyActionmapsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetJourneyActionmapsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetJourneyActionmapsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetJourneyActionmapsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetJourneyActionmapsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetJourneyActionmapsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetJourneyActionmapsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetJourneyActionmapsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetJourneyActionmapsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetJourneyActionmapsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetJourneyActionmapsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetJourneyActionmapsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetJourneyActionmapsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetJourneyActionmapsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetJourneyActionmapsOK creates a GetJourneyActionmapsOK with default headers values
func NewGetJourneyActionmapsOK() *GetJourneyActionmapsOK {
	return &GetJourneyActionmapsOK{}
}

/*
GetJourneyActionmapsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetJourneyActionmapsOK struct {
	Payload *models.ActionMapListing
}

// IsSuccess returns true when this get journey actionmaps o k response has a 2xx status code
func (o *GetJourneyActionmapsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get journey actionmaps o k response has a 3xx status code
func (o *GetJourneyActionmapsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps o k response has a 4xx status code
func (o *GetJourneyActionmapsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actionmaps o k response has a 5xx status code
func (o *GetJourneyActionmapsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps o k response a status code equal to that given
func (o *GetJourneyActionmapsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetJourneyActionmapsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActionmapsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsOK  %+v", 200, o.Payload)
}

func (o *GetJourneyActionmapsOK) GetPayload() *models.ActionMapListing {
	return o.Payload
}

func (o *GetJourneyActionmapsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionMapListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsBadRequest creates a GetJourneyActionmapsBadRequest with default headers values
func NewGetJourneyActionmapsBadRequest() *GetJourneyActionmapsBadRequest {
	return &GetJourneyActionmapsBadRequest{}
}

/*
GetJourneyActionmapsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetJourneyActionmapsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps bad request response has a 2xx status code
func (o *GetJourneyActionmapsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps bad request response has a 3xx status code
func (o *GetJourneyActionmapsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps bad request response has a 4xx status code
func (o *GetJourneyActionmapsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps bad request response has a 5xx status code
func (o *GetJourneyActionmapsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps bad request response a status code equal to that given
func (o *GetJourneyActionmapsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetJourneyActionmapsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActionmapsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsBadRequest  %+v", 400, o.Payload)
}

func (o *GetJourneyActionmapsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsUnauthorized creates a GetJourneyActionmapsUnauthorized with default headers values
func NewGetJourneyActionmapsUnauthorized() *GetJourneyActionmapsUnauthorized {
	return &GetJourneyActionmapsUnauthorized{}
}

/*
GetJourneyActionmapsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetJourneyActionmapsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps unauthorized response has a 2xx status code
func (o *GetJourneyActionmapsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps unauthorized response has a 3xx status code
func (o *GetJourneyActionmapsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps unauthorized response has a 4xx status code
func (o *GetJourneyActionmapsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps unauthorized response has a 5xx status code
func (o *GetJourneyActionmapsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps unauthorized response a status code equal to that given
func (o *GetJourneyActionmapsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetJourneyActionmapsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActionmapsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetJourneyActionmapsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsForbidden creates a GetJourneyActionmapsForbidden with default headers values
func NewGetJourneyActionmapsForbidden() *GetJourneyActionmapsForbidden {
	return &GetJourneyActionmapsForbidden{}
}

/*
GetJourneyActionmapsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetJourneyActionmapsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps forbidden response has a 2xx status code
func (o *GetJourneyActionmapsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps forbidden response has a 3xx status code
func (o *GetJourneyActionmapsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps forbidden response has a 4xx status code
func (o *GetJourneyActionmapsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps forbidden response has a 5xx status code
func (o *GetJourneyActionmapsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps forbidden response a status code equal to that given
func (o *GetJourneyActionmapsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetJourneyActionmapsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActionmapsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsForbidden  %+v", 403, o.Payload)
}

func (o *GetJourneyActionmapsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsNotFound creates a GetJourneyActionmapsNotFound with default headers values
func NewGetJourneyActionmapsNotFound() *GetJourneyActionmapsNotFound {
	return &GetJourneyActionmapsNotFound{}
}

/*
GetJourneyActionmapsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetJourneyActionmapsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps not found response has a 2xx status code
func (o *GetJourneyActionmapsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps not found response has a 3xx status code
func (o *GetJourneyActionmapsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps not found response has a 4xx status code
func (o *GetJourneyActionmapsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps not found response has a 5xx status code
func (o *GetJourneyActionmapsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps not found response a status code equal to that given
func (o *GetJourneyActionmapsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetJourneyActionmapsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActionmapsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsNotFound  %+v", 404, o.Payload)
}

func (o *GetJourneyActionmapsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsRequestTimeout creates a GetJourneyActionmapsRequestTimeout with default headers values
func NewGetJourneyActionmapsRequestTimeout() *GetJourneyActionmapsRequestTimeout {
	return &GetJourneyActionmapsRequestTimeout{}
}

/*
GetJourneyActionmapsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetJourneyActionmapsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps request timeout response has a 2xx status code
func (o *GetJourneyActionmapsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps request timeout response has a 3xx status code
func (o *GetJourneyActionmapsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps request timeout response has a 4xx status code
func (o *GetJourneyActionmapsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps request timeout response has a 5xx status code
func (o *GetJourneyActionmapsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps request timeout response a status code equal to that given
func (o *GetJourneyActionmapsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetJourneyActionmapsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyActionmapsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetJourneyActionmapsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsRequestEntityTooLarge creates a GetJourneyActionmapsRequestEntityTooLarge with default headers values
func NewGetJourneyActionmapsRequestEntityTooLarge() *GetJourneyActionmapsRequestEntityTooLarge {
	return &GetJourneyActionmapsRequestEntityTooLarge{}
}

/*
GetJourneyActionmapsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetJourneyActionmapsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps request entity too large response has a 2xx status code
func (o *GetJourneyActionmapsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps request entity too large response has a 3xx status code
func (o *GetJourneyActionmapsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps request entity too large response has a 4xx status code
func (o *GetJourneyActionmapsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps request entity too large response has a 5xx status code
func (o *GetJourneyActionmapsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps request entity too large response a status code equal to that given
func (o *GetJourneyActionmapsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetJourneyActionmapsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActionmapsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetJourneyActionmapsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsUnsupportedMediaType creates a GetJourneyActionmapsUnsupportedMediaType with default headers values
func NewGetJourneyActionmapsUnsupportedMediaType() *GetJourneyActionmapsUnsupportedMediaType {
	return &GetJourneyActionmapsUnsupportedMediaType{}
}

/*
GetJourneyActionmapsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetJourneyActionmapsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps unsupported media type response has a 2xx status code
func (o *GetJourneyActionmapsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps unsupported media type response has a 3xx status code
func (o *GetJourneyActionmapsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps unsupported media type response has a 4xx status code
func (o *GetJourneyActionmapsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps unsupported media type response has a 5xx status code
func (o *GetJourneyActionmapsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps unsupported media type response a status code equal to that given
func (o *GetJourneyActionmapsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetJourneyActionmapsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActionmapsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetJourneyActionmapsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsTooManyRequests creates a GetJourneyActionmapsTooManyRequests with default headers values
func NewGetJourneyActionmapsTooManyRequests() *GetJourneyActionmapsTooManyRequests {
	return &GetJourneyActionmapsTooManyRequests{}
}

/*
GetJourneyActionmapsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetJourneyActionmapsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps too many requests response has a 2xx status code
func (o *GetJourneyActionmapsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps too many requests response has a 3xx status code
func (o *GetJourneyActionmapsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps too many requests response has a 4xx status code
func (o *GetJourneyActionmapsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get journey actionmaps too many requests response has a 5xx status code
func (o *GetJourneyActionmapsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get journey actionmaps too many requests response a status code equal to that given
func (o *GetJourneyActionmapsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetJourneyActionmapsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActionmapsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetJourneyActionmapsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsInternalServerError creates a GetJourneyActionmapsInternalServerError with default headers values
func NewGetJourneyActionmapsInternalServerError() *GetJourneyActionmapsInternalServerError {
	return &GetJourneyActionmapsInternalServerError{}
}

/*
GetJourneyActionmapsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetJourneyActionmapsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps internal server error response has a 2xx status code
func (o *GetJourneyActionmapsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps internal server error response has a 3xx status code
func (o *GetJourneyActionmapsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps internal server error response has a 4xx status code
func (o *GetJourneyActionmapsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actionmaps internal server error response has a 5xx status code
func (o *GetJourneyActionmapsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actionmaps internal server error response a status code equal to that given
func (o *GetJourneyActionmapsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetJourneyActionmapsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActionmapsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetJourneyActionmapsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsServiceUnavailable creates a GetJourneyActionmapsServiceUnavailable with default headers values
func NewGetJourneyActionmapsServiceUnavailable() *GetJourneyActionmapsServiceUnavailable {
	return &GetJourneyActionmapsServiceUnavailable{}
}

/*
GetJourneyActionmapsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetJourneyActionmapsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps service unavailable response has a 2xx status code
func (o *GetJourneyActionmapsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps service unavailable response has a 3xx status code
func (o *GetJourneyActionmapsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps service unavailable response has a 4xx status code
func (o *GetJourneyActionmapsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actionmaps service unavailable response has a 5xx status code
func (o *GetJourneyActionmapsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actionmaps service unavailable response a status code equal to that given
func (o *GetJourneyActionmapsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetJourneyActionmapsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActionmapsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetJourneyActionmapsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetJourneyActionmapsGatewayTimeout creates a GetJourneyActionmapsGatewayTimeout with default headers values
func NewGetJourneyActionmapsGatewayTimeout() *GetJourneyActionmapsGatewayTimeout {
	return &GetJourneyActionmapsGatewayTimeout{}
}

/*
GetJourneyActionmapsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetJourneyActionmapsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get journey actionmaps gateway timeout response has a 2xx status code
func (o *GetJourneyActionmapsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get journey actionmaps gateway timeout response has a 3xx status code
func (o *GetJourneyActionmapsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get journey actionmaps gateway timeout response has a 4xx status code
func (o *GetJourneyActionmapsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get journey actionmaps gateway timeout response has a 5xx status code
func (o *GetJourneyActionmapsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get journey actionmaps gateway timeout response a status code equal to that given
func (o *GetJourneyActionmapsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetJourneyActionmapsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActionmapsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/journey/actionmaps][%d] getJourneyActionmapsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetJourneyActionmapsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetJourneyActionmapsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
