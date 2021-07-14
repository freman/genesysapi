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

// GetSpeechandtextanalyticsTopicReader is a Reader for the GetSpeechandtextanalyticsTopic structure.
type GetSpeechandtextanalyticsTopicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSpeechandtextanalyticsTopicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSpeechandtextanalyticsTopicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetSpeechandtextanalyticsTopicBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetSpeechandtextanalyticsTopicUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSpeechandtextanalyticsTopicForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSpeechandtextanalyticsTopicNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetSpeechandtextanalyticsTopicRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetSpeechandtextanalyticsTopicRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetSpeechandtextanalyticsTopicUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetSpeechandtextanalyticsTopicTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSpeechandtextanalyticsTopicInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetSpeechandtextanalyticsTopicServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetSpeechandtextanalyticsTopicGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSpeechandtextanalyticsTopicOK creates a GetSpeechandtextanalyticsTopicOK with default headers values
func NewGetSpeechandtextanalyticsTopicOK() *GetSpeechandtextanalyticsTopicOK {
	return &GetSpeechandtextanalyticsTopicOK{}
}

/*GetSpeechandtextanalyticsTopicOK handles this case with default header values.

successful operation
*/
type GetSpeechandtextanalyticsTopicOK struct {
	Payload *models.Topic
}

func (o *GetSpeechandtextanalyticsTopicOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicOK  %+v", 200, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicOK) GetPayload() *models.Topic {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Topic)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicBadRequest creates a GetSpeechandtextanalyticsTopicBadRequest with default headers values
func NewGetSpeechandtextanalyticsTopicBadRequest() *GetSpeechandtextanalyticsTopicBadRequest {
	return &GetSpeechandtextanalyticsTopicBadRequest{}
}

/*GetSpeechandtextanalyticsTopicBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetSpeechandtextanalyticsTopicBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicBadRequest  %+v", 400, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicUnauthorized creates a GetSpeechandtextanalyticsTopicUnauthorized with default headers values
func NewGetSpeechandtextanalyticsTopicUnauthorized() *GetSpeechandtextanalyticsTopicUnauthorized {
	return &GetSpeechandtextanalyticsTopicUnauthorized{}
}

/*GetSpeechandtextanalyticsTopicUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetSpeechandtextanalyticsTopicUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicUnauthorized  %+v", 401, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicForbidden creates a GetSpeechandtextanalyticsTopicForbidden with default headers values
func NewGetSpeechandtextanalyticsTopicForbidden() *GetSpeechandtextanalyticsTopicForbidden {
	return &GetSpeechandtextanalyticsTopicForbidden{}
}

/*GetSpeechandtextanalyticsTopicForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetSpeechandtextanalyticsTopicForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicForbidden  %+v", 403, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicNotFound creates a GetSpeechandtextanalyticsTopicNotFound with default headers values
func NewGetSpeechandtextanalyticsTopicNotFound() *GetSpeechandtextanalyticsTopicNotFound {
	return &GetSpeechandtextanalyticsTopicNotFound{}
}

/*GetSpeechandtextanalyticsTopicNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetSpeechandtextanalyticsTopicNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicNotFound  %+v", 404, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicRequestTimeout creates a GetSpeechandtextanalyticsTopicRequestTimeout with default headers values
func NewGetSpeechandtextanalyticsTopicRequestTimeout() *GetSpeechandtextanalyticsTopicRequestTimeout {
	return &GetSpeechandtextanalyticsTopicRequestTimeout{}
}

/*GetSpeechandtextanalyticsTopicRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetSpeechandtextanalyticsTopicRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicRequestEntityTooLarge creates a GetSpeechandtextanalyticsTopicRequestEntityTooLarge with default headers values
func NewGetSpeechandtextanalyticsTopicRequestEntityTooLarge() *GetSpeechandtextanalyticsTopicRequestEntityTooLarge {
	return &GetSpeechandtextanalyticsTopicRequestEntityTooLarge{}
}

/*GetSpeechandtextanalyticsTopicRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetSpeechandtextanalyticsTopicRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicUnsupportedMediaType creates a GetSpeechandtextanalyticsTopicUnsupportedMediaType with default headers values
func NewGetSpeechandtextanalyticsTopicUnsupportedMediaType() *GetSpeechandtextanalyticsTopicUnsupportedMediaType {
	return &GetSpeechandtextanalyticsTopicUnsupportedMediaType{}
}

/*GetSpeechandtextanalyticsTopicUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetSpeechandtextanalyticsTopicUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicTooManyRequests creates a GetSpeechandtextanalyticsTopicTooManyRequests with default headers values
func NewGetSpeechandtextanalyticsTopicTooManyRequests() *GetSpeechandtextanalyticsTopicTooManyRequests {
	return &GetSpeechandtextanalyticsTopicTooManyRequests{}
}

/*GetSpeechandtextanalyticsTopicTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetSpeechandtextanalyticsTopicTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicInternalServerError creates a GetSpeechandtextanalyticsTopicInternalServerError with default headers values
func NewGetSpeechandtextanalyticsTopicInternalServerError() *GetSpeechandtextanalyticsTopicInternalServerError {
	return &GetSpeechandtextanalyticsTopicInternalServerError{}
}

/*GetSpeechandtextanalyticsTopicInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetSpeechandtextanalyticsTopicInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicServiceUnavailable creates a GetSpeechandtextanalyticsTopicServiceUnavailable with default headers values
func NewGetSpeechandtextanalyticsTopicServiceUnavailable() *GetSpeechandtextanalyticsTopicServiceUnavailable {
	return &GetSpeechandtextanalyticsTopicServiceUnavailable{}
}

/*GetSpeechandtextanalyticsTopicServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetSpeechandtextanalyticsTopicServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSpeechandtextanalyticsTopicGatewayTimeout creates a GetSpeechandtextanalyticsTopicGatewayTimeout with default headers values
func NewGetSpeechandtextanalyticsTopicGatewayTimeout() *GetSpeechandtextanalyticsTopicGatewayTimeout {
	return &GetSpeechandtextanalyticsTopicGatewayTimeout{}
}

/*GetSpeechandtextanalyticsTopicGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetSpeechandtextanalyticsTopicGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetSpeechandtextanalyticsTopicGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/speechandtextanalytics/topics/{topicId}][%d] getSpeechandtextanalyticsTopicGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetSpeechandtextanalyticsTopicGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetSpeechandtextanalyticsTopicGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
