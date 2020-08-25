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

// GetRoutingSettingsTranscriptionReader is a Reader for the GetRoutingSettingsTranscription structure.
type GetRoutingSettingsTranscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingSettingsTranscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingSettingsTranscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingSettingsTranscriptionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingSettingsTranscriptionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingSettingsTranscriptionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingSettingsTranscriptionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingSettingsTranscriptionRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingSettingsTranscriptionUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingSettingsTranscriptionTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingSettingsTranscriptionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingSettingsTranscriptionServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingSettingsTranscriptionGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingSettingsTranscriptionOK creates a GetRoutingSettingsTranscriptionOK with default headers values
func NewGetRoutingSettingsTranscriptionOK() *GetRoutingSettingsTranscriptionOK {
	return &GetRoutingSettingsTranscriptionOK{}
}

/*GetRoutingSettingsTranscriptionOK handles this case with default header values.

successful operation
*/
type GetRoutingSettingsTranscriptionOK struct {
	Payload *models.TranscriptionSettings
}

func (o *GetRoutingSettingsTranscriptionOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionOK  %+v", 200, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionOK) GetPayload() *models.TranscriptionSettings {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TranscriptionSettings)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionBadRequest creates a GetRoutingSettingsTranscriptionBadRequest with default headers values
func NewGetRoutingSettingsTranscriptionBadRequest() *GetRoutingSettingsTranscriptionBadRequest {
	return &GetRoutingSettingsTranscriptionBadRequest{}
}

/*GetRoutingSettingsTranscriptionBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingSettingsTranscriptionBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionUnauthorized creates a GetRoutingSettingsTranscriptionUnauthorized with default headers values
func NewGetRoutingSettingsTranscriptionUnauthorized() *GetRoutingSettingsTranscriptionUnauthorized {
	return &GetRoutingSettingsTranscriptionUnauthorized{}
}

/*GetRoutingSettingsTranscriptionUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingSettingsTranscriptionUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionForbidden creates a GetRoutingSettingsTranscriptionForbidden with default headers values
func NewGetRoutingSettingsTranscriptionForbidden() *GetRoutingSettingsTranscriptionForbidden {
	return &GetRoutingSettingsTranscriptionForbidden{}
}

/*GetRoutingSettingsTranscriptionForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingSettingsTranscriptionForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionNotFound creates a GetRoutingSettingsTranscriptionNotFound with default headers values
func NewGetRoutingSettingsTranscriptionNotFound() *GetRoutingSettingsTranscriptionNotFound {
	return &GetRoutingSettingsTranscriptionNotFound{}
}

/*GetRoutingSettingsTranscriptionNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRoutingSettingsTranscriptionNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionRequestEntityTooLarge creates a GetRoutingSettingsTranscriptionRequestEntityTooLarge with default headers values
func NewGetRoutingSettingsTranscriptionRequestEntityTooLarge() *GetRoutingSettingsTranscriptionRequestEntityTooLarge {
	return &GetRoutingSettingsTranscriptionRequestEntityTooLarge{}
}

/*GetRoutingSettingsTranscriptionRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRoutingSettingsTranscriptionRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionUnsupportedMediaType creates a GetRoutingSettingsTranscriptionUnsupportedMediaType with default headers values
func NewGetRoutingSettingsTranscriptionUnsupportedMediaType() *GetRoutingSettingsTranscriptionUnsupportedMediaType {
	return &GetRoutingSettingsTranscriptionUnsupportedMediaType{}
}

/*GetRoutingSettingsTranscriptionUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingSettingsTranscriptionUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionTooManyRequests creates a GetRoutingSettingsTranscriptionTooManyRequests with default headers values
func NewGetRoutingSettingsTranscriptionTooManyRequests() *GetRoutingSettingsTranscriptionTooManyRequests {
	return &GetRoutingSettingsTranscriptionTooManyRequests{}
}

/*GetRoutingSettingsTranscriptionTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRoutingSettingsTranscriptionTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionInternalServerError creates a GetRoutingSettingsTranscriptionInternalServerError with default headers values
func NewGetRoutingSettingsTranscriptionInternalServerError() *GetRoutingSettingsTranscriptionInternalServerError {
	return &GetRoutingSettingsTranscriptionInternalServerError{}
}

/*GetRoutingSettingsTranscriptionInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingSettingsTranscriptionInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionServiceUnavailable creates a GetRoutingSettingsTranscriptionServiceUnavailable with default headers values
func NewGetRoutingSettingsTranscriptionServiceUnavailable() *GetRoutingSettingsTranscriptionServiceUnavailable {
	return &GetRoutingSettingsTranscriptionServiceUnavailable{}
}

/*GetRoutingSettingsTranscriptionServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingSettingsTranscriptionServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingSettingsTranscriptionGatewayTimeout creates a GetRoutingSettingsTranscriptionGatewayTimeout with default headers values
func NewGetRoutingSettingsTranscriptionGatewayTimeout() *GetRoutingSettingsTranscriptionGatewayTimeout {
	return &GetRoutingSettingsTranscriptionGatewayTimeout{}
}

/*GetRoutingSettingsTranscriptionGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRoutingSettingsTranscriptionGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRoutingSettingsTranscriptionGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/settings/transcription][%d] getRoutingSettingsTranscriptionGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingSettingsTranscriptionGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingSettingsTranscriptionGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}