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

// PostGamificationProfileDeactivateReader is a Reader for the PostGamificationProfileDeactivate structure.
type PostGamificationProfileDeactivateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGamificationProfileDeactivateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGamificationProfileDeactivateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGamificationProfileDeactivateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGamificationProfileDeactivateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGamificationProfileDeactivateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGamificationProfileDeactivateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostGamificationProfileDeactivateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGamificationProfileDeactivateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGamificationProfileDeactivateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGamificationProfileDeactivateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGamificationProfileDeactivateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGamificationProfileDeactivateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGamificationProfileDeactivateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGamificationProfileDeactivateOK creates a PostGamificationProfileDeactivateOK with default headers values
func NewPostGamificationProfileDeactivateOK() *PostGamificationProfileDeactivateOK {
	return &PostGamificationProfileDeactivateOK{}
}

/*PostGamificationProfileDeactivateOK handles this case with default header values.

successful operation
*/
type PostGamificationProfileDeactivateOK struct {
	Payload *models.PerformanceProfile
}

func (o *PostGamificationProfileDeactivateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateOK  %+v", 200, o.Payload)
}

func (o *PostGamificationProfileDeactivateOK) GetPayload() *models.PerformanceProfile {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateBadRequest creates a PostGamificationProfileDeactivateBadRequest with default headers values
func NewPostGamificationProfileDeactivateBadRequest() *PostGamificationProfileDeactivateBadRequest {
	return &PostGamificationProfileDeactivateBadRequest{}
}

/*PostGamificationProfileDeactivateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGamificationProfileDeactivateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateBadRequest  %+v", 400, o.Payload)
}

func (o *PostGamificationProfileDeactivateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateUnauthorized creates a PostGamificationProfileDeactivateUnauthorized with default headers values
func NewPostGamificationProfileDeactivateUnauthorized() *PostGamificationProfileDeactivateUnauthorized {
	return &PostGamificationProfileDeactivateUnauthorized{}
}

/*PostGamificationProfileDeactivateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGamificationProfileDeactivateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGamificationProfileDeactivateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateForbidden creates a PostGamificationProfileDeactivateForbidden with default headers values
func NewPostGamificationProfileDeactivateForbidden() *PostGamificationProfileDeactivateForbidden {
	return &PostGamificationProfileDeactivateForbidden{}
}

/*PostGamificationProfileDeactivateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostGamificationProfileDeactivateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateForbidden  %+v", 403, o.Payload)
}

func (o *PostGamificationProfileDeactivateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateNotFound creates a PostGamificationProfileDeactivateNotFound with default headers values
func NewPostGamificationProfileDeactivateNotFound() *PostGamificationProfileDeactivateNotFound {
	return &PostGamificationProfileDeactivateNotFound{}
}

/*PostGamificationProfileDeactivateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostGamificationProfileDeactivateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateNotFound  %+v", 404, o.Payload)
}

func (o *PostGamificationProfileDeactivateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateRequestTimeout creates a PostGamificationProfileDeactivateRequestTimeout with default headers values
func NewPostGamificationProfileDeactivateRequestTimeout() *PostGamificationProfileDeactivateRequestTimeout {
	return &PostGamificationProfileDeactivateRequestTimeout{}
}

/*PostGamificationProfileDeactivateRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGamificationProfileDeactivateRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGamificationProfileDeactivateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateRequestEntityTooLarge creates a PostGamificationProfileDeactivateRequestEntityTooLarge with default headers values
func NewPostGamificationProfileDeactivateRequestEntityTooLarge() *PostGamificationProfileDeactivateRequestEntityTooLarge {
	return &PostGamificationProfileDeactivateRequestEntityTooLarge{}
}

/*PostGamificationProfileDeactivateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostGamificationProfileDeactivateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGamificationProfileDeactivateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateUnsupportedMediaType creates a PostGamificationProfileDeactivateUnsupportedMediaType with default headers values
func NewPostGamificationProfileDeactivateUnsupportedMediaType() *PostGamificationProfileDeactivateUnsupportedMediaType {
	return &PostGamificationProfileDeactivateUnsupportedMediaType{}
}

/*PostGamificationProfileDeactivateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGamificationProfileDeactivateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGamificationProfileDeactivateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateTooManyRequests creates a PostGamificationProfileDeactivateTooManyRequests with default headers values
func NewPostGamificationProfileDeactivateTooManyRequests() *PostGamificationProfileDeactivateTooManyRequests {
	return &PostGamificationProfileDeactivateTooManyRequests{}
}

/*PostGamificationProfileDeactivateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGamificationProfileDeactivateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGamificationProfileDeactivateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateInternalServerError creates a PostGamificationProfileDeactivateInternalServerError with default headers values
func NewPostGamificationProfileDeactivateInternalServerError() *PostGamificationProfileDeactivateInternalServerError {
	return &PostGamificationProfileDeactivateInternalServerError{}
}

/*PostGamificationProfileDeactivateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGamificationProfileDeactivateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGamificationProfileDeactivateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateServiceUnavailable creates a PostGamificationProfileDeactivateServiceUnavailable with default headers values
func NewPostGamificationProfileDeactivateServiceUnavailable() *PostGamificationProfileDeactivateServiceUnavailable {
	return &PostGamificationProfileDeactivateServiceUnavailable{}
}

/*PostGamificationProfileDeactivateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGamificationProfileDeactivateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGamificationProfileDeactivateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileDeactivateGatewayTimeout creates a PostGamificationProfileDeactivateGatewayTimeout with default headers values
func NewPostGamificationProfileDeactivateGatewayTimeout() *PostGamificationProfileDeactivateGatewayTimeout {
	return &PostGamificationProfileDeactivateGatewayTimeout{}
}

/*PostGamificationProfileDeactivateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostGamificationProfileDeactivateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileDeactivateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/deactivate][%d] postGamificationProfileDeactivateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGamificationProfileDeactivateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileDeactivateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
