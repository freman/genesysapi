// Code generated by go-swagger; DO NOT EDIT.

package process_automation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetProcessautomationTriggersTopicsReader is a Reader for the GetProcessautomationTriggersTopics structure.
type GetProcessautomationTriggersTopicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProcessautomationTriggersTopicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProcessautomationTriggersTopicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProcessautomationTriggersTopicsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProcessautomationTriggersTopicsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProcessautomationTriggersTopicsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetProcessautomationTriggersTopicsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetProcessautomationTriggersTopicsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetProcessautomationTriggersTopicsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetProcessautomationTriggersTopicsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetProcessautomationTriggersTopicsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProcessautomationTriggersTopicsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetProcessautomationTriggersTopicsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetProcessautomationTriggersTopicsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProcessautomationTriggersTopicsOK creates a GetProcessautomationTriggersTopicsOK with default headers values
func NewGetProcessautomationTriggersTopicsOK() *GetProcessautomationTriggersTopicsOK {
	return &GetProcessautomationTriggersTopicsOK{}
}

/*
GetProcessautomationTriggersTopicsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetProcessautomationTriggersTopicsOK struct {
	Payload *models.TopicCursorEntityListing
}

// IsSuccess returns true when this get processautomation triggers topics o k response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get processautomation triggers topics o k response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics o k response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processautomation triggers topics o k response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation triggers topics o k response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetProcessautomationTriggersTopicsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsOK  %+v", 200, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsOK  %+v", 200, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsOK) GetPayload() *models.TopicCursorEntityListing {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TopicCursorEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsBadRequest creates a GetProcessautomationTriggersTopicsBadRequest with default headers values
func NewGetProcessautomationTriggersTopicsBadRequest() *GetProcessautomationTriggersTopicsBadRequest {
	return &GetProcessautomationTriggersTopicsBadRequest{}
}

/*
GetProcessautomationTriggersTopicsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetProcessautomationTriggersTopicsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics bad request response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics bad request response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics bad request response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation triggers topics bad request response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation triggers topics bad request response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetProcessautomationTriggersTopicsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsBadRequest  %+v", 400, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsBadRequest  %+v", 400, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsUnauthorized creates a GetProcessautomationTriggersTopicsUnauthorized with default headers values
func NewGetProcessautomationTriggersTopicsUnauthorized() *GetProcessautomationTriggersTopicsUnauthorized {
	return &GetProcessautomationTriggersTopicsUnauthorized{}
}

/*
GetProcessautomationTriggersTopicsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetProcessautomationTriggersTopicsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics unauthorized response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics unauthorized response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics unauthorized response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation triggers topics unauthorized response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation triggers topics unauthorized response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetProcessautomationTriggersTopicsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsForbidden creates a GetProcessautomationTriggersTopicsForbidden with default headers values
func NewGetProcessautomationTriggersTopicsForbidden() *GetProcessautomationTriggersTopicsForbidden {
	return &GetProcessautomationTriggersTopicsForbidden{}
}

/*
GetProcessautomationTriggersTopicsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetProcessautomationTriggersTopicsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics forbidden response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics forbidden response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics forbidden response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation triggers topics forbidden response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation triggers topics forbidden response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetProcessautomationTriggersTopicsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsForbidden  %+v", 403, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsForbidden  %+v", 403, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsNotFound creates a GetProcessautomationTriggersTopicsNotFound with default headers values
func NewGetProcessautomationTriggersTopicsNotFound() *GetProcessautomationTriggersTopicsNotFound {
	return &GetProcessautomationTriggersTopicsNotFound{}
}

/*
GetProcessautomationTriggersTopicsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetProcessautomationTriggersTopicsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics not found response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics not found response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics not found response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation triggers topics not found response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation triggers topics not found response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetProcessautomationTriggersTopicsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsNotFound  %+v", 404, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsNotFound  %+v", 404, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsRequestTimeout creates a GetProcessautomationTriggersTopicsRequestTimeout with default headers values
func NewGetProcessautomationTriggersTopicsRequestTimeout() *GetProcessautomationTriggersTopicsRequestTimeout {
	return &GetProcessautomationTriggersTopicsRequestTimeout{}
}

/*
GetProcessautomationTriggersTopicsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetProcessautomationTriggersTopicsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics request timeout response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics request timeout response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics request timeout response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation triggers topics request timeout response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation triggers topics request timeout response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetProcessautomationTriggersTopicsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsRequestEntityTooLarge creates a GetProcessautomationTriggersTopicsRequestEntityTooLarge with default headers values
func NewGetProcessautomationTriggersTopicsRequestEntityTooLarge() *GetProcessautomationTriggersTopicsRequestEntityTooLarge {
	return &GetProcessautomationTriggersTopicsRequestEntityTooLarge{}
}

/*
GetProcessautomationTriggersTopicsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetProcessautomationTriggersTopicsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics request entity too large response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics request entity too large response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics request entity too large response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation triggers topics request entity too large response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation triggers topics request entity too large response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetProcessautomationTriggersTopicsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsUnsupportedMediaType creates a GetProcessautomationTriggersTopicsUnsupportedMediaType with default headers values
func NewGetProcessautomationTriggersTopicsUnsupportedMediaType() *GetProcessautomationTriggersTopicsUnsupportedMediaType {
	return &GetProcessautomationTriggersTopicsUnsupportedMediaType{}
}

/*
GetProcessautomationTriggersTopicsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetProcessautomationTriggersTopicsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics unsupported media type response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics unsupported media type response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics unsupported media type response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation triggers topics unsupported media type response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation triggers topics unsupported media type response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetProcessautomationTriggersTopicsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsTooManyRequests creates a GetProcessautomationTriggersTopicsTooManyRequests with default headers values
func NewGetProcessautomationTriggersTopicsTooManyRequests() *GetProcessautomationTriggersTopicsTooManyRequests {
	return &GetProcessautomationTriggersTopicsTooManyRequests{}
}

/*
GetProcessautomationTriggersTopicsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetProcessautomationTriggersTopicsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics too many requests response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics too many requests response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics too many requests response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get processautomation triggers topics too many requests response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get processautomation triggers topics too many requests response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetProcessautomationTriggersTopicsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsInternalServerError creates a GetProcessautomationTriggersTopicsInternalServerError with default headers values
func NewGetProcessautomationTriggersTopicsInternalServerError() *GetProcessautomationTriggersTopicsInternalServerError {
	return &GetProcessautomationTriggersTopicsInternalServerError{}
}

/*
GetProcessautomationTriggersTopicsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetProcessautomationTriggersTopicsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics internal server error response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics internal server error response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics internal server error response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processautomation triggers topics internal server error response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get processautomation triggers topics internal server error response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetProcessautomationTriggersTopicsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsServiceUnavailable creates a GetProcessautomationTriggersTopicsServiceUnavailable with default headers values
func NewGetProcessautomationTriggersTopicsServiceUnavailable() *GetProcessautomationTriggersTopicsServiceUnavailable {
	return &GetProcessautomationTriggersTopicsServiceUnavailable{}
}

/*
GetProcessautomationTriggersTopicsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetProcessautomationTriggersTopicsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics service unavailable response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics service unavailable response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics service unavailable response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processautomation triggers topics service unavailable response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get processautomation triggers topics service unavailable response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetProcessautomationTriggersTopicsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetProcessautomationTriggersTopicsGatewayTimeout creates a GetProcessautomationTriggersTopicsGatewayTimeout with default headers values
func NewGetProcessautomationTriggersTopicsGatewayTimeout() *GetProcessautomationTriggersTopicsGatewayTimeout {
	return &GetProcessautomationTriggersTopicsGatewayTimeout{}
}

/*
GetProcessautomationTriggersTopicsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetProcessautomationTriggersTopicsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get processautomation triggers topics gateway timeout response has a 2xx status code
func (o *GetProcessautomationTriggersTopicsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get processautomation triggers topics gateway timeout response has a 3xx status code
func (o *GetProcessautomationTriggersTopicsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get processautomation triggers topics gateway timeout response has a 4xx status code
func (o *GetProcessautomationTriggersTopicsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get processautomation triggers topics gateway timeout response has a 5xx status code
func (o *GetProcessautomationTriggersTopicsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get processautomation triggers topics gateway timeout response a status code equal to that given
func (o *GetProcessautomationTriggersTopicsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetProcessautomationTriggersTopicsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/processautomation/triggers/topics][%d] getProcessautomationTriggersTopicsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetProcessautomationTriggersTopicsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetProcessautomationTriggersTopicsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
