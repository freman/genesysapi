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

// PutGamificationProfileReader is a Reader for the PutGamificationProfile structure.
type PutGamificationProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutGamificationProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutGamificationProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutGamificationProfileBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutGamificationProfileUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutGamificationProfileForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutGamificationProfileNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutGamificationProfileRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutGamificationProfileUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutGamificationProfileTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutGamificationProfileInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutGamificationProfileServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutGamificationProfileGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutGamificationProfileOK creates a PutGamificationProfileOK with default headers values
func NewPutGamificationProfileOK() *PutGamificationProfileOK {
	return &PutGamificationProfileOK{}
}

/*PutGamificationProfileOK handles this case with default header values.

successful operation
*/
type PutGamificationProfileOK struct {
	Payload *models.PerformanceProfile
}

func (o *PutGamificationProfileOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileOK  %+v", 200, o.Payload)
}

func (o *PutGamificationProfileOK) GetPayload() *models.PerformanceProfile {
	return o.Payload
}

func (o *PutGamificationProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileBadRequest creates a PutGamificationProfileBadRequest with default headers values
func NewPutGamificationProfileBadRequest() *PutGamificationProfileBadRequest {
	return &PutGamificationProfileBadRequest{}
}

/*PutGamificationProfileBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutGamificationProfileBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileBadRequest  %+v", 400, o.Payload)
}

func (o *PutGamificationProfileBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileUnauthorized creates a PutGamificationProfileUnauthorized with default headers values
func NewPutGamificationProfileUnauthorized() *PutGamificationProfileUnauthorized {
	return &PutGamificationProfileUnauthorized{}
}

/*PutGamificationProfileUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutGamificationProfileUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileUnauthorized  %+v", 401, o.Payload)
}

func (o *PutGamificationProfileUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileForbidden creates a PutGamificationProfileForbidden with default headers values
func NewPutGamificationProfileForbidden() *PutGamificationProfileForbidden {
	return &PutGamificationProfileForbidden{}
}

/*PutGamificationProfileForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutGamificationProfileForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileForbidden  %+v", 403, o.Payload)
}

func (o *PutGamificationProfileForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileNotFound creates a PutGamificationProfileNotFound with default headers values
func NewPutGamificationProfileNotFound() *PutGamificationProfileNotFound {
	return &PutGamificationProfileNotFound{}
}

/*PutGamificationProfileNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutGamificationProfileNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileNotFound  %+v", 404, o.Payload)
}

func (o *PutGamificationProfileNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileRequestEntityTooLarge creates a PutGamificationProfileRequestEntityTooLarge with default headers values
func NewPutGamificationProfileRequestEntityTooLarge() *PutGamificationProfileRequestEntityTooLarge {
	return &PutGamificationProfileRequestEntityTooLarge{}
}

/*PutGamificationProfileRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutGamificationProfileRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutGamificationProfileRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileUnsupportedMediaType creates a PutGamificationProfileUnsupportedMediaType with default headers values
func NewPutGamificationProfileUnsupportedMediaType() *PutGamificationProfileUnsupportedMediaType {
	return &PutGamificationProfileUnsupportedMediaType{}
}

/*PutGamificationProfileUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutGamificationProfileUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutGamificationProfileUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileTooManyRequests creates a PutGamificationProfileTooManyRequests with default headers values
func NewPutGamificationProfileTooManyRequests() *PutGamificationProfileTooManyRequests {
	return &PutGamificationProfileTooManyRequests{}
}

/*PutGamificationProfileTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutGamificationProfileTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutGamificationProfileTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileInternalServerError creates a PutGamificationProfileInternalServerError with default headers values
func NewPutGamificationProfileInternalServerError() *PutGamificationProfileInternalServerError {
	return &PutGamificationProfileInternalServerError{}
}

/*PutGamificationProfileInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutGamificationProfileInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileInternalServerError  %+v", 500, o.Payload)
}

func (o *PutGamificationProfileInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileServiceUnavailable creates a PutGamificationProfileServiceUnavailable with default headers values
func NewPutGamificationProfileServiceUnavailable() *PutGamificationProfileServiceUnavailable {
	return &PutGamificationProfileServiceUnavailable{}
}

/*PutGamificationProfileServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutGamificationProfileServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutGamificationProfileServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutGamificationProfileGatewayTimeout creates a PutGamificationProfileGatewayTimeout with default headers values
func NewPutGamificationProfileGatewayTimeout() *PutGamificationProfileGatewayTimeout {
	return &PutGamificationProfileGatewayTimeout{}
}

/*PutGamificationProfileGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutGamificationProfileGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutGamificationProfileGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/gamification/profiles/{performanceProfileId}][%d] putGamificationProfileGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutGamificationProfileGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutGamificationProfileGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
