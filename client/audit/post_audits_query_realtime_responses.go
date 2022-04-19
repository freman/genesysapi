// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostAuditsQueryRealtimeReader is a Reader for the PostAuditsQueryRealtime structure.
type PostAuditsQueryRealtimeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostAuditsQueryRealtimeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostAuditsQueryRealtimeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostAuditsQueryRealtimeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostAuditsQueryRealtimeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostAuditsQueryRealtimeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostAuditsQueryRealtimeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostAuditsQueryRealtimeRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostAuditsQueryRealtimeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostAuditsQueryRealtimeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostAuditsQueryRealtimeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostAuditsQueryRealtimeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostAuditsQueryRealtimeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostAuditsQueryRealtimeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostAuditsQueryRealtimeOK creates a PostAuditsQueryRealtimeOK with default headers values
func NewPostAuditsQueryRealtimeOK() *PostAuditsQueryRealtimeOK {
	return &PostAuditsQueryRealtimeOK{}
}

/*PostAuditsQueryRealtimeOK handles this case with default header values.

successful operation
*/
type PostAuditsQueryRealtimeOK struct {
	Payload *models.AuditRealtimeQueryResultsResponse
}

func (o *PostAuditsQueryRealtimeOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeOK  %+v", 200, o.Payload)
}

func (o *PostAuditsQueryRealtimeOK) GetPayload() *models.AuditRealtimeQueryResultsResponse {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditRealtimeQueryResultsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeBadRequest creates a PostAuditsQueryRealtimeBadRequest with default headers values
func NewPostAuditsQueryRealtimeBadRequest() *PostAuditsQueryRealtimeBadRequest {
	return &PostAuditsQueryRealtimeBadRequest{}
}

/*PostAuditsQueryRealtimeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostAuditsQueryRealtimeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeBadRequest  %+v", 400, o.Payload)
}

func (o *PostAuditsQueryRealtimeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeUnauthorized creates a PostAuditsQueryRealtimeUnauthorized with default headers values
func NewPostAuditsQueryRealtimeUnauthorized() *PostAuditsQueryRealtimeUnauthorized {
	return &PostAuditsQueryRealtimeUnauthorized{}
}

/*PostAuditsQueryRealtimeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostAuditsQueryRealtimeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeUnauthorized  %+v", 401, o.Payload)
}

func (o *PostAuditsQueryRealtimeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeForbidden creates a PostAuditsQueryRealtimeForbidden with default headers values
func NewPostAuditsQueryRealtimeForbidden() *PostAuditsQueryRealtimeForbidden {
	return &PostAuditsQueryRealtimeForbidden{}
}

/*PostAuditsQueryRealtimeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostAuditsQueryRealtimeForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeForbidden  %+v", 403, o.Payload)
}

func (o *PostAuditsQueryRealtimeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeNotFound creates a PostAuditsQueryRealtimeNotFound with default headers values
func NewPostAuditsQueryRealtimeNotFound() *PostAuditsQueryRealtimeNotFound {
	return &PostAuditsQueryRealtimeNotFound{}
}

/*PostAuditsQueryRealtimeNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostAuditsQueryRealtimeNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeNotFound  %+v", 404, o.Payload)
}

func (o *PostAuditsQueryRealtimeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeRequestTimeout creates a PostAuditsQueryRealtimeRequestTimeout with default headers values
func NewPostAuditsQueryRealtimeRequestTimeout() *PostAuditsQueryRealtimeRequestTimeout {
	return &PostAuditsQueryRealtimeRequestTimeout{}
}

/*PostAuditsQueryRealtimeRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostAuditsQueryRealtimeRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostAuditsQueryRealtimeRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeRequestEntityTooLarge creates a PostAuditsQueryRealtimeRequestEntityTooLarge with default headers values
func NewPostAuditsQueryRealtimeRequestEntityTooLarge() *PostAuditsQueryRealtimeRequestEntityTooLarge {
	return &PostAuditsQueryRealtimeRequestEntityTooLarge{}
}

/*PostAuditsQueryRealtimeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PostAuditsQueryRealtimeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeUnsupportedMediaType creates a PostAuditsQueryRealtimeUnsupportedMediaType with default headers values
func NewPostAuditsQueryRealtimeUnsupportedMediaType() *PostAuditsQueryRealtimeUnsupportedMediaType {
	return &PostAuditsQueryRealtimeUnsupportedMediaType{}
}

/*PostAuditsQueryRealtimeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostAuditsQueryRealtimeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostAuditsQueryRealtimeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeTooManyRequests creates a PostAuditsQueryRealtimeTooManyRequests with default headers values
func NewPostAuditsQueryRealtimeTooManyRequests() *PostAuditsQueryRealtimeTooManyRequests {
	return &PostAuditsQueryRealtimeTooManyRequests{}
}

/*PostAuditsQueryRealtimeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostAuditsQueryRealtimeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostAuditsQueryRealtimeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeInternalServerError creates a PostAuditsQueryRealtimeInternalServerError with default headers values
func NewPostAuditsQueryRealtimeInternalServerError() *PostAuditsQueryRealtimeInternalServerError {
	return &PostAuditsQueryRealtimeInternalServerError{}
}

/*PostAuditsQueryRealtimeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostAuditsQueryRealtimeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeInternalServerError  %+v", 500, o.Payload)
}

func (o *PostAuditsQueryRealtimeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeServiceUnavailable creates a PostAuditsQueryRealtimeServiceUnavailable with default headers values
func NewPostAuditsQueryRealtimeServiceUnavailable() *PostAuditsQueryRealtimeServiceUnavailable {
	return &PostAuditsQueryRealtimeServiceUnavailable{}
}

/*PostAuditsQueryRealtimeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostAuditsQueryRealtimeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostAuditsQueryRealtimeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostAuditsQueryRealtimeGatewayTimeout creates a PostAuditsQueryRealtimeGatewayTimeout with default headers values
func NewPostAuditsQueryRealtimeGatewayTimeout() *PostAuditsQueryRealtimeGatewayTimeout {
	return &PostAuditsQueryRealtimeGatewayTimeout{}
}

/*PostAuditsQueryRealtimeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostAuditsQueryRealtimeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostAuditsQueryRealtimeGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/audits/query/realtime][%d] postAuditsQueryRealtimeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostAuditsQueryRealtimeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostAuditsQueryRealtimeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
