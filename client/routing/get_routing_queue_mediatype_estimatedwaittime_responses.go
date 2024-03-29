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

// GetRoutingQueueMediatypeEstimatedwaittimeReader is a Reader for the GetRoutingQueueMediatypeEstimatedwaittime structure.
type GetRoutingQueueMediatypeEstimatedwaittimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingQueueMediatypeEstimatedwaittimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeOK creates a GetRoutingQueueMediatypeEstimatedwaittimeOK with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeOK() *GetRoutingQueueMediatypeEstimatedwaittimeOK {
	return &GetRoutingQueueMediatypeEstimatedwaittimeOK{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingQueueMediatypeEstimatedwaittimeOK struct {
	Payload *models.EstimatedWaitTimePredictions
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime o k response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime o k response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime o k response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime o k response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime o k response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeOK  %+v", 200, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeOK) GetPayload() *models.EstimatedWaitTimePredictions {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EstimatedWaitTimePredictions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeBadRequest creates a GetRoutingQueueMediatypeEstimatedwaittimeBadRequest with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeBadRequest() *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest {
	return &GetRoutingQueueMediatypeEstimatedwaittimeBadRequest{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingQueueMediatypeEstimatedwaittimeBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime bad request response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime bad request response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime bad request response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime bad request response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime bad request response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeUnauthorized creates a GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeUnauthorized() *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized {
	return &GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime unauthorized response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime unauthorized response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime unauthorized response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime unauthorized response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime unauthorized response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeForbidden creates a GetRoutingQueueMediatypeEstimatedwaittimeForbidden with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeForbidden() *GetRoutingQueueMediatypeEstimatedwaittimeForbidden {
	return &GetRoutingQueueMediatypeEstimatedwaittimeForbidden{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingQueueMediatypeEstimatedwaittimeForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime forbidden response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime forbidden response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime forbidden response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime forbidden response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime forbidden response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeNotFound creates a GetRoutingQueueMediatypeEstimatedwaittimeNotFound with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeNotFound() *GetRoutingQueueMediatypeEstimatedwaittimeNotFound {
	return &GetRoutingQueueMediatypeEstimatedwaittimeNotFound{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingQueueMediatypeEstimatedwaittimeNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime not found response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime not found response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime not found response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime not found response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime not found response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout creates a GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout() *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout {
	return &GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime request timeout response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime request timeout response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime request timeout response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime request timeout response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime request timeout response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge creates a GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge() *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge {
	return &GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime request entity too large response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime request entity too large response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime request entity too large response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime request entity too large response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime request entity too large response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType creates a GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType() *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType {
	return &GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime unsupported media type response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime unsupported media type response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime unsupported media type response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime unsupported media type response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime unsupported media type response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests creates a GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests() *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests {
	return &GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime too many requests response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime too many requests response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime too many requests response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime too many requests response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime too many requests response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeInternalServerError creates a GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeInternalServerError() *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError {
	return &GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime internal server error response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime internal server error response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime internal server error response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime internal server error response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime internal server error response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable creates a GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable() *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable {
	return &GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime service unavailable response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime service unavailable response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime service unavailable response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime service unavailable response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime service unavailable response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout creates a GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout with default headers values
func NewGetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout() *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout {
	return &GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout{}
}

/*
GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing queue mediatype estimatedwaittime gateway timeout response has a 2xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing queue mediatype estimatedwaittime gateway timeout response has a 3xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing queue mediatype estimatedwaittime gateway timeout response has a 4xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing queue mediatype estimatedwaittime gateway timeout response has a 5xx status code
func (o *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing queue mediatype estimatedwaittime gateway timeout response a status code equal to that given
func (o *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/queues/{queueId}/mediatypes/{mediaType}/estimatedwaittime][%d] getRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingQueueMediatypeEstimatedwaittimeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
