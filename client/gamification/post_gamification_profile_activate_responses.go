// Code generated by go-swagger; DO NOT EDIT.

package gamification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostGamificationProfileActivateReader is a Reader for the PostGamificationProfileActivate structure.
type PostGamificationProfileActivateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGamificationProfileActivateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGamificationProfileActivateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGamificationProfileActivateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGamificationProfileActivateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGamificationProfileActivateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGamificationProfileActivateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostGamificationProfileActivateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGamificationProfileActivateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGamificationProfileActivateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGamificationProfileActivateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGamificationProfileActivateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGamificationProfileActivateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGamificationProfileActivateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGamificationProfileActivateOK creates a PostGamificationProfileActivateOK with default headers values
func NewPostGamificationProfileActivateOK() *PostGamificationProfileActivateOK {
	return &PostGamificationProfileActivateOK{}
}

/*
PostGamificationProfileActivateOK describes a response with status code 200, with default header values.

successful operation
*/
type PostGamificationProfileActivateOK struct {
	Payload *models.PerformanceProfile
}

// IsSuccess returns true when this post gamification profile activate o k response has a 2xx status code
func (o *PostGamificationProfileActivateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post gamification profile activate o k response has a 3xx status code
func (o *PostGamificationProfileActivateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate o k response has a 4xx status code
func (o *PostGamificationProfileActivateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gamification profile activate o k response has a 5xx status code
func (o *PostGamificationProfileActivateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile activate o k response a status code equal to that given
func (o *PostGamificationProfileActivateOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostGamificationProfileActivateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateOK  %+v", 200, o.Payload)
}

func (o *PostGamificationProfileActivateOK) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateOK  %+v", 200, o.Payload)
}

func (o *PostGamificationProfileActivateOK) GetPayload() *models.PerformanceProfile {
	return o.Payload
}

func (o *PostGamificationProfileActivateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateBadRequest creates a PostGamificationProfileActivateBadRequest with default headers values
func NewPostGamificationProfileActivateBadRequest() *PostGamificationProfileActivateBadRequest {
	return &PostGamificationProfileActivateBadRequest{}
}

/*
PostGamificationProfileActivateBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGamificationProfileActivateBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate bad request response has a 2xx status code
func (o *PostGamificationProfileActivateBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate bad request response has a 3xx status code
func (o *PostGamificationProfileActivateBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate bad request response has a 4xx status code
func (o *PostGamificationProfileActivateBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile activate bad request response has a 5xx status code
func (o *PostGamificationProfileActivateBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile activate bad request response a status code equal to that given
func (o *PostGamificationProfileActivateBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostGamificationProfileActivateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateBadRequest  %+v", 400, o.Payload)
}

func (o *PostGamificationProfileActivateBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateBadRequest  %+v", 400, o.Payload)
}

func (o *PostGamificationProfileActivateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateUnauthorized creates a PostGamificationProfileActivateUnauthorized with default headers values
func NewPostGamificationProfileActivateUnauthorized() *PostGamificationProfileActivateUnauthorized {
	return &PostGamificationProfileActivateUnauthorized{}
}

/*
PostGamificationProfileActivateUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGamificationProfileActivateUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate unauthorized response has a 2xx status code
func (o *PostGamificationProfileActivateUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate unauthorized response has a 3xx status code
func (o *PostGamificationProfileActivateUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate unauthorized response has a 4xx status code
func (o *PostGamificationProfileActivateUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile activate unauthorized response has a 5xx status code
func (o *PostGamificationProfileActivateUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile activate unauthorized response a status code equal to that given
func (o *PostGamificationProfileActivateUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostGamificationProfileActivateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGamificationProfileActivateUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGamificationProfileActivateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateForbidden creates a PostGamificationProfileActivateForbidden with default headers values
func NewPostGamificationProfileActivateForbidden() *PostGamificationProfileActivateForbidden {
	return &PostGamificationProfileActivateForbidden{}
}

/*
PostGamificationProfileActivateForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostGamificationProfileActivateForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate forbidden response has a 2xx status code
func (o *PostGamificationProfileActivateForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate forbidden response has a 3xx status code
func (o *PostGamificationProfileActivateForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate forbidden response has a 4xx status code
func (o *PostGamificationProfileActivateForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile activate forbidden response has a 5xx status code
func (o *PostGamificationProfileActivateForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile activate forbidden response a status code equal to that given
func (o *PostGamificationProfileActivateForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostGamificationProfileActivateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateForbidden  %+v", 403, o.Payload)
}

func (o *PostGamificationProfileActivateForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateForbidden  %+v", 403, o.Payload)
}

func (o *PostGamificationProfileActivateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateNotFound creates a PostGamificationProfileActivateNotFound with default headers values
func NewPostGamificationProfileActivateNotFound() *PostGamificationProfileActivateNotFound {
	return &PostGamificationProfileActivateNotFound{}
}

/*
PostGamificationProfileActivateNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostGamificationProfileActivateNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate not found response has a 2xx status code
func (o *PostGamificationProfileActivateNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate not found response has a 3xx status code
func (o *PostGamificationProfileActivateNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate not found response has a 4xx status code
func (o *PostGamificationProfileActivateNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile activate not found response has a 5xx status code
func (o *PostGamificationProfileActivateNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile activate not found response a status code equal to that given
func (o *PostGamificationProfileActivateNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostGamificationProfileActivateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateNotFound  %+v", 404, o.Payload)
}

func (o *PostGamificationProfileActivateNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateNotFound  %+v", 404, o.Payload)
}

func (o *PostGamificationProfileActivateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateRequestTimeout creates a PostGamificationProfileActivateRequestTimeout with default headers values
func NewPostGamificationProfileActivateRequestTimeout() *PostGamificationProfileActivateRequestTimeout {
	return &PostGamificationProfileActivateRequestTimeout{}
}

/*
PostGamificationProfileActivateRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGamificationProfileActivateRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate request timeout response has a 2xx status code
func (o *PostGamificationProfileActivateRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate request timeout response has a 3xx status code
func (o *PostGamificationProfileActivateRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate request timeout response has a 4xx status code
func (o *PostGamificationProfileActivateRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile activate request timeout response has a 5xx status code
func (o *PostGamificationProfileActivateRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile activate request timeout response a status code equal to that given
func (o *PostGamificationProfileActivateRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostGamificationProfileActivateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGamificationProfileActivateRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGamificationProfileActivateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateRequestEntityTooLarge creates a PostGamificationProfileActivateRequestEntityTooLarge with default headers values
func NewPostGamificationProfileActivateRequestEntityTooLarge() *PostGamificationProfileActivateRequestEntityTooLarge {
	return &PostGamificationProfileActivateRequestEntityTooLarge{}
}

/*
PostGamificationProfileActivateRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostGamificationProfileActivateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate request entity too large response has a 2xx status code
func (o *PostGamificationProfileActivateRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate request entity too large response has a 3xx status code
func (o *PostGamificationProfileActivateRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate request entity too large response has a 4xx status code
func (o *PostGamificationProfileActivateRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile activate request entity too large response has a 5xx status code
func (o *PostGamificationProfileActivateRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile activate request entity too large response a status code equal to that given
func (o *PostGamificationProfileActivateRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostGamificationProfileActivateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGamificationProfileActivateRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGamificationProfileActivateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateUnsupportedMediaType creates a PostGamificationProfileActivateUnsupportedMediaType with default headers values
func NewPostGamificationProfileActivateUnsupportedMediaType() *PostGamificationProfileActivateUnsupportedMediaType {
	return &PostGamificationProfileActivateUnsupportedMediaType{}
}

/*
PostGamificationProfileActivateUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGamificationProfileActivateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate unsupported media type response has a 2xx status code
func (o *PostGamificationProfileActivateUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate unsupported media type response has a 3xx status code
func (o *PostGamificationProfileActivateUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate unsupported media type response has a 4xx status code
func (o *PostGamificationProfileActivateUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile activate unsupported media type response has a 5xx status code
func (o *PostGamificationProfileActivateUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile activate unsupported media type response a status code equal to that given
func (o *PostGamificationProfileActivateUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostGamificationProfileActivateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGamificationProfileActivateUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGamificationProfileActivateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateTooManyRequests creates a PostGamificationProfileActivateTooManyRequests with default headers values
func NewPostGamificationProfileActivateTooManyRequests() *PostGamificationProfileActivateTooManyRequests {
	return &PostGamificationProfileActivateTooManyRequests{}
}

/*
PostGamificationProfileActivateTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGamificationProfileActivateTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate too many requests response has a 2xx status code
func (o *PostGamificationProfileActivateTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate too many requests response has a 3xx status code
func (o *PostGamificationProfileActivateTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate too many requests response has a 4xx status code
func (o *PostGamificationProfileActivateTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post gamification profile activate too many requests response has a 5xx status code
func (o *PostGamificationProfileActivateTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post gamification profile activate too many requests response a status code equal to that given
func (o *PostGamificationProfileActivateTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostGamificationProfileActivateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGamificationProfileActivateTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGamificationProfileActivateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateInternalServerError creates a PostGamificationProfileActivateInternalServerError with default headers values
func NewPostGamificationProfileActivateInternalServerError() *PostGamificationProfileActivateInternalServerError {
	return &PostGamificationProfileActivateInternalServerError{}
}

/*
PostGamificationProfileActivateInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGamificationProfileActivateInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate internal server error response has a 2xx status code
func (o *PostGamificationProfileActivateInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate internal server error response has a 3xx status code
func (o *PostGamificationProfileActivateInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate internal server error response has a 4xx status code
func (o *PostGamificationProfileActivateInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gamification profile activate internal server error response has a 5xx status code
func (o *PostGamificationProfileActivateInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post gamification profile activate internal server error response a status code equal to that given
func (o *PostGamificationProfileActivateInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostGamificationProfileActivateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGamificationProfileActivateInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGamificationProfileActivateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateServiceUnavailable creates a PostGamificationProfileActivateServiceUnavailable with default headers values
func NewPostGamificationProfileActivateServiceUnavailable() *PostGamificationProfileActivateServiceUnavailable {
	return &PostGamificationProfileActivateServiceUnavailable{}
}

/*
PostGamificationProfileActivateServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGamificationProfileActivateServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate service unavailable response has a 2xx status code
func (o *PostGamificationProfileActivateServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate service unavailable response has a 3xx status code
func (o *PostGamificationProfileActivateServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate service unavailable response has a 4xx status code
func (o *PostGamificationProfileActivateServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gamification profile activate service unavailable response has a 5xx status code
func (o *PostGamificationProfileActivateServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post gamification profile activate service unavailable response a status code equal to that given
func (o *PostGamificationProfileActivateServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostGamificationProfileActivateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGamificationProfileActivateServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGamificationProfileActivateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileActivateGatewayTimeout creates a PostGamificationProfileActivateGatewayTimeout with default headers values
func NewPostGamificationProfileActivateGatewayTimeout() *PostGamificationProfileActivateGatewayTimeout {
	return &PostGamificationProfileActivateGatewayTimeout{}
}

/*
PostGamificationProfileActivateGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostGamificationProfileActivateGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post gamification profile activate gateway timeout response has a 2xx status code
func (o *PostGamificationProfileActivateGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post gamification profile activate gateway timeout response has a 3xx status code
func (o *PostGamificationProfileActivateGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post gamification profile activate gateway timeout response has a 4xx status code
func (o *PostGamificationProfileActivateGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post gamification profile activate gateway timeout response has a 5xx status code
func (o *PostGamificationProfileActivateGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post gamification profile activate gateway timeout response a status code equal to that given
func (o *PostGamificationProfileActivateGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostGamificationProfileActivateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGamificationProfileActivateGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/activate][%d] postGamificationProfileActivateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGamificationProfileActivateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileActivateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
