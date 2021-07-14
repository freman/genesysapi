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

// PutRoutingSettingsTranscriptionReader is a Reader for the PutRoutingSettingsTranscription structure.
type PutRoutingSettingsTranscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutRoutingSettingsTranscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutRoutingSettingsTranscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPutRoutingSettingsTranscriptionAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutRoutingSettingsTranscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutRoutingSettingsTranscriptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutRoutingSettingsTranscriptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutRoutingSettingsTranscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutRoutingSettingsTranscriptionRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutRoutingSettingsTranscriptionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutRoutingSettingsTranscriptionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutRoutingSettingsTranscriptionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutRoutingSettingsTranscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutRoutingSettingsTranscriptionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutRoutingSettingsTranscriptionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutRoutingSettingsTranscriptionOK creates a PutRoutingSettingsTranscriptionOK with default headers values
func NewPutRoutingSettingsTranscriptionOK() *PutRoutingSettingsTranscriptionOK {
	return &PutRoutingSettingsTranscriptionOK{}
}

/*PutRoutingSettingsTranscriptionOK handles this case with default header values.

successful operation
*/
type PutRoutingSettingsTranscriptionOK struct {
	Payload *models.TranscriptionSettings
}

func (o *PutRoutingSettingsTranscriptionOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionOK  %+v", 200, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionOK) GetPayload() *models.TranscriptionSettings {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TranscriptionSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionAccepted creates a PutRoutingSettingsTranscriptionAccepted with default headers values
func NewPutRoutingSettingsTranscriptionAccepted() *PutRoutingSettingsTranscriptionAccepted {
	return &PutRoutingSettingsTranscriptionAccepted{}
}

/*PutRoutingSettingsTranscriptionAccepted handles this case with default header values.

Request to update transcription settings has been accepted
*/
type PutRoutingSettingsTranscriptionAccepted struct {
	Payload *models.TranscriptionSettings
}

func (o *PutRoutingSettingsTranscriptionAccepted) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionAccepted  %+v", 202, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionAccepted) GetPayload() *models.TranscriptionSettings {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TranscriptionSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionBadRequest creates a PutRoutingSettingsTranscriptionBadRequest with default headers values
func NewPutRoutingSettingsTranscriptionBadRequest() *PutRoutingSettingsTranscriptionBadRequest {
	return &PutRoutingSettingsTranscriptionBadRequest{}
}

/*PutRoutingSettingsTranscriptionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutRoutingSettingsTranscriptionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionUnauthorized creates a PutRoutingSettingsTranscriptionUnauthorized with default headers values
func NewPutRoutingSettingsTranscriptionUnauthorized() *PutRoutingSettingsTranscriptionUnauthorized {
	return &PutRoutingSettingsTranscriptionUnauthorized{}
}

/*PutRoutingSettingsTranscriptionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutRoutingSettingsTranscriptionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionForbidden creates a PutRoutingSettingsTranscriptionForbidden with default headers values
func NewPutRoutingSettingsTranscriptionForbidden() *PutRoutingSettingsTranscriptionForbidden {
	return &PutRoutingSettingsTranscriptionForbidden{}
}

/*PutRoutingSettingsTranscriptionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutRoutingSettingsTranscriptionForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionForbidden  %+v", 403, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionNotFound creates a PutRoutingSettingsTranscriptionNotFound with default headers values
func NewPutRoutingSettingsTranscriptionNotFound() *PutRoutingSettingsTranscriptionNotFound {
	return &PutRoutingSettingsTranscriptionNotFound{}
}

/*PutRoutingSettingsTranscriptionNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutRoutingSettingsTranscriptionNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionNotFound  %+v", 404, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionRequestTimeout creates a PutRoutingSettingsTranscriptionRequestTimeout with default headers values
func NewPutRoutingSettingsTranscriptionRequestTimeout() *PutRoutingSettingsTranscriptionRequestTimeout {
	return &PutRoutingSettingsTranscriptionRequestTimeout{}
}

/*PutRoutingSettingsTranscriptionRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutRoutingSettingsTranscriptionRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionRequestEntityTooLarge creates a PutRoutingSettingsTranscriptionRequestEntityTooLarge with default headers values
func NewPutRoutingSettingsTranscriptionRequestEntityTooLarge() *PutRoutingSettingsTranscriptionRequestEntityTooLarge {
	return &PutRoutingSettingsTranscriptionRequestEntityTooLarge{}
}

/*PutRoutingSettingsTranscriptionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutRoutingSettingsTranscriptionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionUnsupportedMediaType creates a PutRoutingSettingsTranscriptionUnsupportedMediaType with default headers values
func NewPutRoutingSettingsTranscriptionUnsupportedMediaType() *PutRoutingSettingsTranscriptionUnsupportedMediaType {
	return &PutRoutingSettingsTranscriptionUnsupportedMediaType{}
}

/*PutRoutingSettingsTranscriptionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutRoutingSettingsTranscriptionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionTooManyRequests creates a PutRoutingSettingsTranscriptionTooManyRequests with default headers values
func NewPutRoutingSettingsTranscriptionTooManyRequests() *PutRoutingSettingsTranscriptionTooManyRequests {
	return &PutRoutingSettingsTranscriptionTooManyRequests{}
}

/*PutRoutingSettingsTranscriptionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutRoutingSettingsTranscriptionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionInternalServerError creates a PutRoutingSettingsTranscriptionInternalServerError with default headers values
func NewPutRoutingSettingsTranscriptionInternalServerError() *PutRoutingSettingsTranscriptionInternalServerError {
	return &PutRoutingSettingsTranscriptionInternalServerError{}
}

/*PutRoutingSettingsTranscriptionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutRoutingSettingsTranscriptionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionServiceUnavailable creates a PutRoutingSettingsTranscriptionServiceUnavailable with default headers values
func NewPutRoutingSettingsTranscriptionServiceUnavailable() *PutRoutingSettingsTranscriptionServiceUnavailable {
	return &PutRoutingSettingsTranscriptionServiceUnavailable{}
}

/*PutRoutingSettingsTranscriptionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutRoutingSettingsTranscriptionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutRoutingSettingsTranscriptionGatewayTimeout creates a PutRoutingSettingsTranscriptionGatewayTimeout with default headers values
func NewPutRoutingSettingsTranscriptionGatewayTimeout() *PutRoutingSettingsTranscriptionGatewayTimeout {
	return &PutRoutingSettingsTranscriptionGatewayTimeout{}
}

/*PutRoutingSettingsTranscriptionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutRoutingSettingsTranscriptionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutRoutingSettingsTranscriptionGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/routing/settings/transcription][%d] putRoutingSettingsTranscriptionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutRoutingSettingsTranscriptionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutRoutingSettingsTranscriptionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
