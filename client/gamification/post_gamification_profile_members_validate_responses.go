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

// PostGamificationProfileMembersValidateReader is a Reader for the PostGamificationProfileMembersValidate structure.
type PostGamificationProfileMembersValidateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostGamificationProfileMembersValidateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostGamificationProfileMembersValidateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostGamificationProfileMembersValidateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostGamificationProfileMembersValidateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostGamificationProfileMembersValidateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostGamificationProfileMembersValidateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostGamificationProfileMembersValidateRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostGamificationProfileMembersValidateRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostGamificationProfileMembersValidateUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostGamificationProfileMembersValidateTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostGamificationProfileMembersValidateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostGamificationProfileMembersValidateServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostGamificationProfileMembersValidateGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostGamificationProfileMembersValidateOK creates a PostGamificationProfileMembersValidateOK with default headers values
func NewPostGamificationProfileMembersValidateOK() *PostGamificationProfileMembersValidateOK {
	return &PostGamificationProfileMembersValidateOK{}
}

/*PostGamificationProfileMembersValidateOK handles this case with default header values.

successful operation
*/
type PostGamificationProfileMembersValidateOK struct {
	Payload *models.AssignmentValidation
}

func (o *PostGamificationProfileMembersValidateOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateOK  %+v", 200, o.Payload)
}

func (o *PostGamificationProfileMembersValidateOK) GetPayload() *models.AssignmentValidation {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AssignmentValidation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateBadRequest creates a PostGamificationProfileMembersValidateBadRequest with default headers values
func NewPostGamificationProfileMembersValidateBadRequest() *PostGamificationProfileMembersValidateBadRequest {
	return &PostGamificationProfileMembersValidateBadRequest{}
}

/*PostGamificationProfileMembersValidateBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostGamificationProfileMembersValidateBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateBadRequest  %+v", 400, o.Payload)
}

func (o *PostGamificationProfileMembersValidateBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateUnauthorized creates a PostGamificationProfileMembersValidateUnauthorized with default headers values
func NewPostGamificationProfileMembersValidateUnauthorized() *PostGamificationProfileMembersValidateUnauthorized {
	return &PostGamificationProfileMembersValidateUnauthorized{}
}

/*PostGamificationProfileMembersValidateUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostGamificationProfileMembersValidateUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateUnauthorized  %+v", 401, o.Payload)
}

func (o *PostGamificationProfileMembersValidateUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateForbidden creates a PostGamificationProfileMembersValidateForbidden with default headers values
func NewPostGamificationProfileMembersValidateForbidden() *PostGamificationProfileMembersValidateForbidden {
	return &PostGamificationProfileMembersValidateForbidden{}
}

/*PostGamificationProfileMembersValidateForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostGamificationProfileMembersValidateForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateForbidden  %+v", 403, o.Payload)
}

func (o *PostGamificationProfileMembersValidateForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateNotFound creates a PostGamificationProfileMembersValidateNotFound with default headers values
func NewPostGamificationProfileMembersValidateNotFound() *PostGamificationProfileMembersValidateNotFound {
	return &PostGamificationProfileMembersValidateNotFound{}
}

/*PostGamificationProfileMembersValidateNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostGamificationProfileMembersValidateNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateNotFound  %+v", 404, o.Payload)
}

func (o *PostGamificationProfileMembersValidateNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateRequestTimeout creates a PostGamificationProfileMembersValidateRequestTimeout with default headers values
func NewPostGamificationProfileMembersValidateRequestTimeout() *PostGamificationProfileMembersValidateRequestTimeout {
	return &PostGamificationProfileMembersValidateRequestTimeout{}
}

/*PostGamificationProfileMembersValidateRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostGamificationProfileMembersValidateRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostGamificationProfileMembersValidateRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateRequestEntityTooLarge creates a PostGamificationProfileMembersValidateRequestEntityTooLarge with default headers values
func NewPostGamificationProfileMembersValidateRequestEntityTooLarge() *PostGamificationProfileMembersValidateRequestEntityTooLarge {
	return &PostGamificationProfileMembersValidateRequestEntityTooLarge{}
}

/*PostGamificationProfileMembersValidateRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostGamificationProfileMembersValidateRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostGamificationProfileMembersValidateRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateUnsupportedMediaType creates a PostGamificationProfileMembersValidateUnsupportedMediaType with default headers values
func NewPostGamificationProfileMembersValidateUnsupportedMediaType() *PostGamificationProfileMembersValidateUnsupportedMediaType {
	return &PostGamificationProfileMembersValidateUnsupportedMediaType{}
}

/*PostGamificationProfileMembersValidateUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostGamificationProfileMembersValidateUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostGamificationProfileMembersValidateUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateTooManyRequests creates a PostGamificationProfileMembersValidateTooManyRequests with default headers values
func NewPostGamificationProfileMembersValidateTooManyRequests() *PostGamificationProfileMembersValidateTooManyRequests {
	return &PostGamificationProfileMembersValidateTooManyRequests{}
}

/*PostGamificationProfileMembersValidateTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostGamificationProfileMembersValidateTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostGamificationProfileMembersValidateTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateInternalServerError creates a PostGamificationProfileMembersValidateInternalServerError with default headers values
func NewPostGamificationProfileMembersValidateInternalServerError() *PostGamificationProfileMembersValidateInternalServerError {
	return &PostGamificationProfileMembersValidateInternalServerError{}
}

/*PostGamificationProfileMembersValidateInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostGamificationProfileMembersValidateInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateInternalServerError  %+v", 500, o.Payload)
}

func (o *PostGamificationProfileMembersValidateInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateServiceUnavailable creates a PostGamificationProfileMembersValidateServiceUnavailable with default headers values
func NewPostGamificationProfileMembersValidateServiceUnavailable() *PostGamificationProfileMembersValidateServiceUnavailable {
	return &PostGamificationProfileMembersValidateServiceUnavailable{}
}

/*PostGamificationProfileMembersValidateServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostGamificationProfileMembersValidateServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostGamificationProfileMembersValidateServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostGamificationProfileMembersValidateGatewayTimeout creates a PostGamificationProfileMembersValidateGatewayTimeout with default headers values
func NewPostGamificationProfileMembersValidateGatewayTimeout() *PostGamificationProfileMembersValidateGatewayTimeout {
	return &PostGamificationProfileMembersValidateGatewayTimeout{}
}

/*PostGamificationProfileMembersValidateGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostGamificationProfileMembersValidateGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostGamificationProfileMembersValidateGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/gamification/profiles/{profileId}/members/validate][%d] postGamificationProfileMembersValidateGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostGamificationProfileMembersValidateGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostGamificationProfileMembersValidateGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
