// Code generated by go-swagger; DO NOT EDIT.

package speech_and_text_analytics

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetSpeechandtextanalyticsConversationCommunicationTranscripturlReader is a Reader for the GetSpeechandtextanalyticsConversationCommunicationTranscripturl structure.
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlOK creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlOK with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlOK() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlOK {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlOK{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlOK handles this case with default header values.

successful operation
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlOK struct {
	Payload *models.TranscriptURL
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlOK) GetPayload() *models.TranscriptURL {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TranscriptURL)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout creates a GetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout() *GetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout {
	return &GetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout{}
}

/*GetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/conversations/{conversationId}/communications/{communicationId}/transcripturl][%d] getSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsConversationCommunicationTranscripturlGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
