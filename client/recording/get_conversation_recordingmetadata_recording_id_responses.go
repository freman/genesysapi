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

// GetConversationRecordingmetadataRecordingIDReader is a Reader for the GetConversationRecordingmetadataRecordingID structure.
type GetConversationRecordingmetadataRecordingIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationRecordingmetadataRecordingIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationRecordingmetadataRecordingIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationRecordingmetadataRecordingIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationRecordingmetadataRecordingIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationRecordingmetadataRecordingIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationRecordingmetadataRecordingIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationRecordingmetadataRecordingIDRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationRecordingmetadataRecordingIDUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationRecordingmetadataRecordingIDTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationRecordingmetadataRecordingIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationRecordingmetadataRecordingIDServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationRecordingmetadataRecordingIDGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationRecordingmetadataRecordingIDOK creates a GetConversationRecordingmetadataRecordingIDOK with default headers values
func NewGetConversationRecordingmetadataRecordingIDOK() *GetConversationRecordingmetadataRecordingIDOK {
	return &GetConversationRecordingmetadataRecordingIDOK{}
}

/*GetConversationRecordingmetadataRecordingIDOK handles this case with default header values.

successful operation
*/
type GetConversationRecordingmetadataRecordingIDOK struct {
	Payload *models.RecordingMetadata
}

func (o *GetConversationRecordingmetadataRecordingIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdOK  %+v", 200, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDOK) GetPayload() *models.RecordingMetadata {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RecordingMetadata)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDBadRequest creates a GetConversationRecordingmetadataRecordingIDBadRequest with default headers values
func NewGetConversationRecordingmetadataRecordingIDBadRequest() *GetConversationRecordingmetadataRecordingIDBadRequest {
	return &GetConversationRecordingmetadataRecordingIDBadRequest{}
}

/*GetConversationRecordingmetadataRecordingIDBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationRecordingmetadataRecordingIDBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDUnauthorized creates a GetConversationRecordingmetadataRecordingIDUnauthorized with default headers values
func NewGetConversationRecordingmetadataRecordingIDUnauthorized() *GetConversationRecordingmetadataRecordingIDUnauthorized {
	return &GetConversationRecordingmetadataRecordingIDUnauthorized{}
}

/*GetConversationRecordingmetadataRecordingIDUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationRecordingmetadataRecordingIDUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDForbidden creates a GetConversationRecordingmetadataRecordingIDForbidden with default headers values
func NewGetConversationRecordingmetadataRecordingIDForbidden() *GetConversationRecordingmetadataRecordingIDForbidden {
	return &GetConversationRecordingmetadataRecordingIDForbidden{}
}

/*GetConversationRecordingmetadataRecordingIDForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationRecordingmetadataRecordingIDForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDNotFound creates a GetConversationRecordingmetadataRecordingIDNotFound with default headers values
func NewGetConversationRecordingmetadataRecordingIDNotFound() *GetConversationRecordingmetadataRecordingIDNotFound {
	return &GetConversationRecordingmetadataRecordingIDNotFound{}
}

/*GetConversationRecordingmetadataRecordingIDNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetConversationRecordingmetadataRecordingIDNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDRequestEntityTooLarge creates a GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge with default headers values
func NewGetConversationRecordingmetadataRecordingIDRequestEntityTooLarge() *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge {
	return &GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge{}
}

/*GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDUnsupportedMediaType creates a GetConversationRecordingmetadataRecordingIDUnsupportedMediaType with default headers values
func NewGetConversationRecordingmetadataRecordingIDUnsupportedMediaType() *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType {
	return &GetConversationRecordingmetadataRecordingIDUnsupportedMediaType{}
}

/*GetConversationRecordingmetadataRecordingIDUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationRecordingmetadataRecordingIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDTooManyRequests creates a GetConversationRecordingmetadataRecordingIDTooManyRequests with default headers values
func NewGetConversationRecordingmetadataRecordingIDTooManyRequests() *GetConversationRecordingmetadataRecordingIDTooManyRequests {
	return &GetConversationRecordingmetadataRecordingIDTooManyRequests{}
}

/*GetConversationRecordingmetadataRecordingIDTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetConversationRecordingmetadataRecordingIDTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDInternalServerError creates a GetConversationRecordingmetadataRecordingIDInternalServerError with default headers values
func NewGetConversationRecordingmetadataRecordingIDInternalServerError() *GetConversationRecordingmetadataRecordingIDInternalServerError {
	return &GetConversationRecordingmetadataRecordingIDInternalServerError{}
}

/*GetConversationRecordingmetadataRecordingIDInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationRecordingmetadataRecordingIDInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDServiceUnavailable creates a GetConversationRecordingmetadataRecordingIDServiceUnavailable with default headers values
func NewGetConversationRecordingmetadataRecordingIDServiceUnavailable() *GetConversationRecordingmetadataRecordingIDServiceUnavailable {
	return &GetConversationRecordingmetadataRecordingIDServiceUnavailable{}
}

/*GetConversationRecordingmetadataRecordingIDServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationRecordingmetadataRecordingIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRecordingIDGatewayTimeout creates a GetConversationRecordingmetadataRecordingIDGatewayTimeout with default headers values
func NewGetConversationRecordingmetadataRecordingIDGatewayTimeout() *GetConversationRecordingmetadataRecordingIDGatewayTimeout {
	return &GetConversationRecordingmetadataRecordingIDGatewayTimeout{}
}

/*GetConversationRecordingmetadataRecordingIDGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetConversationRecordingmetadataRecordingIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
