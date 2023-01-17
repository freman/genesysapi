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
	case 408:
		result := NewGetConversationRecordingmetadataRecordingIDRequestTimeout()
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

/*
GetConversationRecordingmetadataRecordingIDOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationRecordingmetadataRecordingIDOK struct {
	Payload *models.RecordingMetadata
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id o k response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id o k response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id o k response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recordingmetadata recording Id o k response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata recording Id o k response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationRecordingmetadataRecordingIDOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdOK  %+v", 200, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDOK) String() string {
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

/*
GetConversationRecordingmetadataRecordingIDBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationRecordingmetadataRecordingIDBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id bad request response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id bad request response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id bad request response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata recording Id bad request response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata recording Id bad request response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationRecordingmetadataRecordingIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDBadRequest) String() string {
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

/*
GetConversationRecordingmetadataRecordingIDUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationRecordingmetadataRecordingIDUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id unauthorized response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id unauthorized response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id unauthorized response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata recording Id unauthorized response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata recording Id unauthorized response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDUnauthorized) String() string {
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

/*
GetConversationRecordingmetadataRecordingIDForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationRecordingmetadataRecordingIDForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id forbidden response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id forbidden response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id forbidden response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata recording Id forbidden response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata recording Id forbidden response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationRecordingmetadataRecordingIDForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDForbidden) String() string {
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

/*
GetConversationRecordingmetadataRecordingIDNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationRecordingmetadataRecordingIDNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id not found response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id not found response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id not found response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata recording Id not found response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata recording Id not found response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationRecordingmetadataRecordingIDNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDNotFound) String() string {
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

// NewGetConversationRecordingmetadataRecordingIDRequestTimeout creates a GetConversationRecordingmetadataRecordingIDRequestTimeout with default headers values
func NewGetConversationRecordingmetadataRecordingIDRequestTimeout() *GetConversationRecordingmetadataRecordingIDRequestTimeout {
	return &GetConversationRecordingmetadataRecordingIDRequestTimeout{}
}

/*
GetConversationRecordingmetadataRecordingIDRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationRecordingmetadataRecordingIDRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id request timeout response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id request timeout response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id request timeout response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata recording Id request timeout response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata recording Id request timeout response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationRecordingmetadataRecordingIDRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRecordingIDRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id request entity too large response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id request entity too large response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id request entity too large response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata recording Id request entity too large response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata recording Id request entity too large response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDRequestEntityTooLarge) String() string {
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

/*
GetConversationRecordingmetadataRecordingIDUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationRecordingmetadataRecordingIDUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id unsupported media type response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id unsupported media type response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id unsupported media type response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata recording Id unsupported media type response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata recording Id unsupported media type response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDUnsupportedMediaType) String() string {
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

/*
GetConversationRecordingmetadataRecordingIDTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationRecordingmetadataRecordingIDTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id too many requests response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id too many requests response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id too many requests response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata recording Id too many requests response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata recording Id too many requests response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDTooManyRequests) String() string {
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

/*
GetConversationRecordingmetadataRecordingIDInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationRecordingmetadataRecordingIDInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id internal server error response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id internal server error response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id internal server error response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recordingmetadata recording Id internal server error response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation recordingmetadata recording Id internal server error response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDInternalServerError) String() string {
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

/*
GetConversationRecordingmetadataRecordingIDServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationRecordingmetadataRecordingIDServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id service unavailable response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id service unavailable response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id service unavailable response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recordingmetadata recording Id service unavailable response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation recordingmetadata recording Id service unavailable response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDServiceUnavailable) String() string {
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

/*
GetConversationRecordingmetadataRecordingIDGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationRecordingmetadataRecordingIDGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata recording Id gateway timeout response has a 2xx status code
func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata recording Id gateway timeout response has a 3xx status code
func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata recording Id gateway timeout response has a 4xx status code
func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recordingmetadata recording Id gateway timeout response has a 5xx status code
func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation recordingmetadata recording Id gateway timeout response a status code equal to that given
func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata/{recordingId}][%d] getConversationRecordingmetadataRecordingIdGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationRecordingmetadataRecordingIDGatewayTimeout) String() string {
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
