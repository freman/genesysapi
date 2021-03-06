// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetConversationRecordingsReader is a Reader for the GetConversationRecordings structure.
type GetConversationRecordingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationRecordingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationRecordingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewGetConversationRecordingsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationRecordingsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationRecordingsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationRecordingsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationRecordingsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationRecordingsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationRecordingsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationRecordingsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationRecordingsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationRecordingsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationRecordingsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationRecordingsOK creates a GetConversationRecordingsOK with default headers values
func NewGetConversationRecordingsOK() *GetConversationRecordingsOK {
	return &GetConversationRecordingsOK{}
}

/*GetConversationRecordingsOK handles this case with default header values.

successful operation
*/
type GetConversationRecordingsOK struct {
	Payload []*models.Recording
}

func (o *GetConversationRecordingsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsOK  %+v", 200, o.Payload)
}

func (o *GetConversationRecordingsOK) GetPayload() []*models.Recording {
	return o.Payload
}

func (o *GetConversationRecordingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsAccepted creates a GetConversationRecordingsAccepted with default headers values
func NewGetConversationRecordingsAccepted() *GetConversationRecordingsAccepted {
	return &GetConversationRecordingsAccepted{}
}

/*GetConversationRecordingsAccepted handles this case with default header values.

Success - recording is transcoding
*/
type GetConversationRecordingsAccepted struct {
}

func (o *GetConversationRecordingsAccepted) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsAccepted ", 202)
}

func (o *GetConversationRecordingsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConversationRecordingsBadRequest creates a GetConversationRecordingsBadRequest with default headers values
func NewGetConversationRecordingsBadRequest() *GetConversationRecordingsBadRequest {
	return &GetConversationRecordingsBadRequest{}
}

/*GetConversationRecordingsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationRecordingsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationRecordingsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsUnauthorized creates a GetConversationRecordingsUnauthorized with default headers values
func NewGetConversationRecordingsUnauthorized() *GetConversationRecordingsUnauthorized {
	return &GetConversationRecordingsUnauthorized{}
}

/*GetConversationRecordingsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationRecordingsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationRecordingsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsForbidden creates a GetConversationRecordingsForbidden with default headers values
func NewGetConversationRecordingsForbidden() *GetConversationRecordingsForbidden {
	return &GetConversationRecordingsForbidden{}
}

/*GetConversationRecordingsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationRecordingsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationRecordingsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsNotFound creates a GetConversationRecordingsNotFound with default headers values
func NewGetConversationRecordingsNotFound() *GetConversationRecordingsNotFound {
	return &GetConversationRecordingsNotFound{}
}

/*GetConversationRecordingsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationRecordingsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationRecordingsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsRequestEntityTooLarge creates a GetConversationRecordingsRequestEntityTooLarge with default headers values
func NewGetConversationRecordingsRequestEntityTooLarge() *GetConversationRecordingsRequestEntityTooLarge {
	return &GetConversationRecordingsRequestEntityTooLarge{}
}

/*GetConversationRecordingsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationRecordingsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRecordingsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsUnsupportedMediaType creates a GetConversationRecordingsUnsupportedMediaType with default headers values
func NewGetConversationRecordingsUnsupportedMediaType() *GetConversationRecordingsUnsupportedMediaType {
	return &GetConversationRecordingsUnsupportedMediaType{}
}

/*GetConversationRecordingsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationRecordingsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationRecordingsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsTooManyRequests creates a GetConversationRecordingsTooManyRequests with default headers values
func NewGetConversationRecordingsTooManyRequests() *GetConversationRecordingsTooManyRequests {
	return &GetConversationRecordingsTooManyRequests{}
}

/*GetConversationRecordingsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationRecordingsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationRecordingsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsInternalServerError creates a GetConversationRecordingsInternalServerError with default headers values
func NewGetConversationRecordingsInternalServerError() *GetConversationRecordingsInternalServerError {
	return &GetConversationRecordingsInternalServerError{}
}

/*GetConversationRecordingsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationRecordingsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationRecordingsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsServiceUnavailable creates a GetConversationRecordingsServiceUnavailable with default headers values
func NewGetConversationRecordingsServiceUnavailable() *GetConversationRecordingsServiceUnavailable {
	return &GetConversationRecordingsServiceUnavailable{}
}

/*GetConversationRecordingsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationRecordingsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationRecordingsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingsGatewayTimeout creates a GetConversationRecordingsGatewayTimeout with default headers values
func NewGetConversationRecordingsGatewayTimeout() *GetConversationRecordingsGatewayTimeout {
	return &GetConversationRecordingsGatewayTimeout{}
}

/*GetConversationRecordingsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationRecordingsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordings][%d] getConversationRecordingsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationRecordingsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
