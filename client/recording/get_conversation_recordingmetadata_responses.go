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

// GetConversationRecordingmetadataReader is a Reader for the GetConversationRecordingmetadata structure.
type GetConversationRecordingmetadataReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConversationRecordingmetadataReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConversationRecordingmetadataOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConversationRecordingmetadataBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConversationRecordingmetadataUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConversationRecordingmetadataForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConversationRecordingmetadataNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetConversationRecordingmetadataRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetConversationRecordingmetadataRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetConversationRecordingmetadataUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetConversationRecordingmetadataTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConversationRecordingmetadataInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetConversationRecordingmetadataServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetConversationRecordingmetadataGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConversationRecordingmetadataOK creates a GetConversationRecordingmetadataOK with default headers values
func NewGetConversationRecordingmetadataOK() *GetConversationRecordingmetadataOK {
	return &GetConversationRecordingmetadataOK{}
}

/*
GetConversationRecordingmetadataOK describes a response with status code 200, with default header values.

successful operation
*/
type GetConversationRecordingmetadataOK struct {
	Payload []*models.RecordingMetadata
}

// IsSuccess returns true when this get conversation recordingmetadata o k response has a 2xx status code
func (o *GetConversationRecordingmetadataOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get conversation recordingmetadata o k response has a 3xx status code
func (o *GetConversationRecordingmetadataOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata o k response has a 4xx status code
func (o *GetConversationRecordingmetadataOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recordingmetadata o k response has a 5xx status code
func (o *GetConversationRecordingmetadataOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata o k response a status code equal to that given
func (o *GetConversationRecordingmetadataOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConversationRecordingmetadataOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataOK  %+v", 200, o.Payload)
}

func (o *GetConversationRecordingmetadataOK) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataOK  %+v", 200, o.Payload)
}

func (o *GetConversationRecordingmetadataOK) GetPayload() []*models.RecordingMetadata {
	return o.Payload
}

func (o *GetConversationRecordingmetadataOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataBadRequest creates a GetConversationRecordingmetadataBadRequest with default headers values
func NewGetConversationRecordingmetadataBadRequest() *GetConversationRecordingmetadataBadRequest {
	return &GetConversationRecordingmetadataBadRequest{}
}

/*
GetConversationRecordingmetadataBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetConversationRecordingmetadataBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata bad request response has a 2xx status code
func (o *GetConversationRecordingmetadataBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata bad request response has a 3xx status code
func (o *GetConversationRecordingmetadataBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata bad request response has a 4xx status code
func (o *GetConversationRecordingmetadataBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata bad request response has a 5xx status code
func (o *GetConversationRecordingmetadataBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata bad request response a status code equal to that given
func (o *GetConversationRecordingmetadataBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetConversationRecordingmetadataBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationRecordingmetadataBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataBadRequest  %+v", 400, o.Payload)
}

func (o *GetConversationRecordingmetadataBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataUnauthorized creates a GetConversationRecordingmetadataUnauthorized with default headers values
func NewGetConversationRecordingmetadataUnauthorized() *GetConversationRecordingmetadataUnauthorized {
	return &GetConversationRecordingmetadataUnauthorized{}
}

/*
GetConversationRecordingmetadataUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetConversationRecordingmetadataUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata unauthorized response has a 2xx status code
func (o *GetConversationRecordingmetadataUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata unauthorized response has a 3xx status code
func (o *GetConversationRecordingmetadataUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata unauthorized response has a 4xx status code
func (o *GetConversationRecordingmetadataUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata unauthorized response has a 5xx status code
func (o *GetConversationRecordingmetadataUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata unauthorized response a status code equal to that given
func (o *GetConversationRecordingmetadataUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConversationRecordingmetadataUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationRecordingmetadataUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConversationRecordingmetadataUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataForbidden creates a GetConversationRecordingmetadataForbidden with default headers values
func NewGetConversationRecordingmetadataForbidden() *GetConversationRecordingmetadataForbidden {
	return &GetConversationRecordingmetadataForbidden{}
}

/*
GetConversationRecordingmetadataForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetConversationRecordingmetadataForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata forbidden response has a 2xx status code
func (o *GetConversationRecordingmetadataForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata forbidden response has a 3xx status code
func (o *GetConversationRecordingmetadataForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata forbidden response has a 4xx status code
func (o *GetConversationRecordingmetadataForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata forbidden response has a 5xx status code
func (o *GetConversationRecordingmetadataForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata forbidden response a status code equal to that given
func (o *GetConversationRecordingmetadataForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConversationRecordingmetadataForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationRecordingmetadataForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataForbidden  %+v", 403, o.Payload)
}

func (o *GetConversationRecordingmetadataForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataNotFound creates a GetConversationRecordingmetadataNotFound with default headers values
func NewGetConversationRecordingmetadataNotFound() *GetConversationRecordingmetadataNotFound {
	return &GetConversationRecordingmetadataNotFound{}
}

/*
GetConversationRecordingmetadataNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetConversationRecordingmetadataNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata not found response has a 2xx status code
func (o *GetConversationRecordingmetadataNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata not found response has a 3xx status code
func (o *GetConversationRecordingmetadataNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata not found response has a 4xx status code
func (o *GetConversationRecordingmetadataNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata not found response has a 5xx status code
func (o *GetConversationRecordingmetadataNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata not found response a status code equal to that given
func (o *GetConversationRecordingmetadataNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetConversationRecordingmetadataNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationRecordingmetadataNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataNotFound  %+v", 404, o.Payload)
}

func (o *GetConversationRecordingmetadataNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRequestTimeout creates a GetConversationRecordingmetadataRequestTimeout with default headers values
func NewGetConversationRecordingmetadataRequestTimeout() *GetConversationRecordingmetadataRequestTimeout {
	return &GetConversationRecordingmetadataRequestTimeout{}
}

/*
GetConversationRecordingmetadataRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetConversationRecordingmetadataRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata request timeout response has a 2xx status code
func (o *GetConversationRecordingmetadataRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata request timeout response has a 3xx status code
func (o *GetConversationRecordingmetadataRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata request timeout response has a 4xx status code
func (o *GetConversationRecordingmetadataRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata request timeout response has a 5xx status code
func (o *GetConversationRecordingmetadataRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata request timeout response a status code equal to that given
func (o *GetConversationRecordingmetadataRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetConversationRecordingmetadataRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRecordingmetadataRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetConversationRecordingmetadataRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataRequestEntityTooLarge creates a GetConversationRecordingmetadataRequestEntityTooLarge with default headers values
func NewGetConversationRecordingmetadataRequestEntityTooLarge() *GetConversationRecordingmetadataRequestEntityTooLarge {
	return &GetConversationRecordingmetadataRequestEntityTooLarge{}
}

/*
GetConversationRecordingmetadataRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetConversationRecordingmetadataRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata request entity too large response has a 2xx status code
func (o *GetConversationRecordingmetadataRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata request entity too large response has a 3xx status code
func (o *GetConversationRecordingmetadataRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata request entity too large response has a 4xx status code
func (o *GetConversationRecordingmetadataRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata request entity too large response has a 5xx status code
func (o *GetConversationRecordingmetadataRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata request entity too large response a status code equal to that given
func (o *GetConversationRecordingmetadataRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetConversationRecordingmetadataRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRecordingmetadataRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetConversationRecordingmetadataRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataUnsupportedMediaType creates a GetConversationRecordingmetadataUnsupportedMediaType with default headers values
func NewGetConversationRecordingmetadataUnsupportedMediaType() *GetConversationRecordingmetadataUnsupportedMediaType {
	return &GetConversationRecordingmetadataUnsupportedMediaType{}
}

/*
GetConversationRecordingmetadataUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetConversationRecordingmetadataUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata unsupported media type response has a 2xx status code
func (o *GetConversationRecordingmetadataUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata unsupported media type response has a 3xx status code
func (o *GetConversationRecordingmetadataUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata unsupported media type response has a 4xx status code
func (o *GetConversationRecordingmetadataUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata unsupported media type response has a 5xx status code
func (o *GetConversationRecordingmetadataUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata unsupported media type response a status code equal to that given
func (o *GetConversationRecordingmetadataUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetConversationRecordingmetadataUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationRecordingmetadataUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetConversationRecordingmetadataUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataTooManyRequests creates a GetConversationRecordingmetadataTooManyRequests with default headers values
func NewGetConversationRecordingmetadataTooManyRequests() *GetConversationRecordingmetadataTooManyRequests {
	return &GetConversationRecordingmetadataTooManyRequests{}
}

/*
GetConversationRecordingmetadataTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetConversationRecordingmetadataTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata too many requests response has a 2xx status code
func (o *GetConversationRecordingmetadataTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata too many requests response has a 3xx status code
func (o *GetConversationRecordingmetadataTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata too many requests response has a 4xx status code
func (o *GetConversationRecordingmetadataTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get conversation recordingmetadata too many requests response has a 5xx status code
func (o *GetConversationRecordingmetadataTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get conversation recordingmetadata too many requests response a status code equal to that given
func (o *GetConversationRecordingmetadataTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetConversationRecordingmetadataTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationRecordingmetadataTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetConversationRecordingmetadataTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataInternalServerError creates a GetConversationRecordingmetadataInternalServerError with default headers values
func NewGetConversationRecordingmetadataInternalServerError() *GetConversationRecordingmetadataInternalServerError {
	return &GetConversationRecordingmetadataInternalServerError{}
}

/*
GetConversationRecordingmetadataInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetConversationRecordingmetadataInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata internal server error response has a 2xx status code
func (o *GetConversationRecordingmetadataInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata internal server error response has a 3xx status code
func (o *GetConversationRecordingmetadataInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata internal server error response has a 4xx status code
func (o *GetConversationRecordingmetadataInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recordingmetadata internal server error response has a 5xx status code
func (o *GetConversationRecordingmetadataInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation recordingmetadata internal server error response a status code equal to that given
func (o *GetConversationRecordingmetadataInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConversationRecordingmetadataInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationRecordingmetadataInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConversationRecordingmetadataInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataServiceUnavailable creates a GetConversationRecordingmetadataServiceUnavailable with default headers values
func NewGetConversationRecordingmetadataServiceUnavailable() *GetConversationRecordingmetadataServiceUnavailable {
	return &GetConversationRecordingmetadataServiceUnavailable{}
}

/*
GetConversationRecordingmetadataServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetConversationRecordingmetadataServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata service unavailable response has a 2xx status code
func (o *GetConversationRecordingmetadataServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata service unavailable response has a 3xx status code
func (o *GetConversationRecordingmetadataServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata service unavailable response has a 4xx status code
func (o *GetConversationRecordingmetadataServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recordingmetadata service unavailable response has a 5xx status code
func (o *GetConversationRecordingmetadataServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation recordingmetadata service unavailable response a status code equal to that given
func (o *GetConversationRecordingmetadataServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetConversationRecordingmetadataServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationRecordingmetadataServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetConversationRecordingmetadataServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConversationRecordingmetadataGatewayTimeout creates a GetConversationRecordingmetadataGatewayTimeout with default headers values
func NewGetConversationRecordingmetadataGatewayTimeout() *GetConversationRecordingmetadataGatewayTimeout {
	return &GetConversationRecordingmetadataGatewayTimeout{}
}

/*
GetConversationRecordingmetadataGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetConversationRecordingmetadataGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get conversation recordingmetadata gateway timeout response has a 2xx status code
func (o *GetConversationRecordingmetadataGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get conversation recordingmetadata gateway timeout response has a 3xx status code
func (o *GetConversationRecordingmetadataGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get conversation recordingmetadata gateway timeout response has a 4xx status code
func (o *GetConversationRecordingmetadataGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get conversation recordingmetadata gateway timeout response has a 5xx status code
func (o *GetConversationRecordingmetadataGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get conversation recordingmetadata gateway timeout response a status code equal to that given
func (o *GetConversationRecordingmetadataGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetConversationRecordingmetadataGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationRecordingmetadataGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/conversations/{conversationId}/recordingmetadata][%d] getConversationRecordingmetadataGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetConversationRecordingmetadataGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetConversationRecordingmetadataGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
