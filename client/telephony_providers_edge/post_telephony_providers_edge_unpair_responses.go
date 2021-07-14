// Code generated by go-swagger; DO NOT EDIT.

package telephony_providers_edge

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostTelephonyProvidersEdgeUnpairReader is a Reader for the PostTelephonyProvidersEdgeUnpair structure.
type PostTelephonyProvidersEdgeUnpairReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgeUnpairReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgeUnpairOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgeUnpairBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgeUnpairUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgeUnpairForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgeUnpairNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostTelephonyProvidersEdgeUnpairRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostTelephonyProvidersEdgeUnpairConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgeUnpairRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgeUnpairUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgeUnpairTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgeUnpairInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgeUnpairServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgeUnpairGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgeUnpairOK creates a PostTelephonyProvidersEdgeUnpairOK with default headers values
func NewPostTelephonyProvidersEdgeUnpairOK() *PostTelephonyProvidersEdgeUnpairOK {
	return &PostTelephonyProvidersEdgeUnpairOK{}
}

/*PostTelephonyProvidersEdgeUnpairOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgeUnpairOK struct {
	Payload string
}

func (o *PostTelephonyProvidersEdgeUnpairOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairOK) GetPayload() string {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairBadRequest creates a PostTelephonyProvidersEdgeUnpairBadRequest with default headers values
func NewPostTelephonyProvidersEdgeUnpairBadRequest() *PostTelephonyProvidersEdgeUnpairBadRequest {
	return &PostTelephonyProvidersEdgeUnpairBadRequest{}
}

/*PostTelephonyProvidersEdgeUnpairBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgeUnpairBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairUnauthorized creates a PostTelephonyProvidersEdgeUnpairUnauthorized with default headers values
func NewPostTelephonyProvidersEdgeUnpairUnauthorized() *PostTelephonyProvidersEdgeUnpairUnauthorized {
	return &PostTelephonyProvidersEdgeUnpairUnauthorized{}
}

/*PostTelephonyProvidersEdgeUnpairUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgeUnpairUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairForbidden creates a PostTelephonyProvidersEdgeUnpairForbidden with default headers values
func NewPostTelephonyProvidersEdgeUnpairForbidden() *PostTelephonyProvidersEdgeUnpairForbidden {
	return &PostTelephonyProvidersEdgeUnpairForbidden{}
}

/*PostTelephonyProvidersEdgeUnpairForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgeUnpairForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairNotFound creates a PostTelephonyProvidersEdgeUnpairNotFound with default headers values
func NewPostTelephonyProvidersEdgeUnpairNotFound() *PostTelephonyProvidersEdgeUnpairNotFound {
	return &PostTelephonyProvidersEdgeUnpairNotFound{}
}

/*PostTelephonyProvidersEdgeUnpairNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgeUnpairNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairRequestTimeout creates a PostTelephonyProvidersEdgeUnpairRequestTimeout with default headers values
func NewPostTelephonyProvidersEdgeUnpairRequestTimeout() *PostTelephonyProvidersEdgeUnpairRequestTimeout {
	return &PostTelephonyProvidersEdgeUnpairRequestTimeout{}
}

/*PostTelephonyProvidersEdgeUnpairRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostTelephonyProvidersEdgeUnpairRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairConflict creates a PostTelephonyProvidersEdgeUnpairConflict with default headers values
func NewPostTelephonyProvidersEdgeUnpairConflict() *PostTelephonyProvidersEdgeUnpairConflict {
	return &PostTelephonyProvidersEdgeUnpairConflict{}
}

/*PostTelephonyProvidersEdgeUnpairConflict handles this case with default header values.

Conflict
*/
type PostTelephonyProvidersEdgeUnpairConflict struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairConflict  %+v", 409, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairRequestEntityTooLarge creates a PostTelephonyProvidersEdgeUnpairRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgeUnpairRequestEntityTooLarge() *PostTelephonyProvidersEdgeUnpairRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgeUnpairRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgeUnpairRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgeUnpairRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairUnsupportedMediaType creates a PostTelephonyProvidersEdgeUnpairUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgeUnpairUnsupportedMediaType() *PostTelephonyProvidersEdgeUnpairUnsupportedMediaType {
	return &PostTelephonyProvidersEdgeUnpairUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgeUnpairUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgeUnpairUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairTooManyRequests creates a PostTelephonyProvidersEdgeUnpairTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgeUnpairTooManyRequests() *PostTelephonyProvidersEdgeUnpairTooManyRequests {
	return &PostTelephonyProvidersEdgeUnpairTooManyRequests{}
}

/*PostTelephonyProvidersEdgeUnpairTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostTelephonyProvidersEdgeUnpairTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairInternalServerError creates a PostTelephonyProvidersEdgeUnpairInternalServerError with default headers values
func NewPostTelephonyProvidersEdgeUnpairInternalServerError() *PostTelephonyProvidersEdgeUnpairInternalServerError {
	return &PostTelephonyProvidersEdgeUnpairInternalServerError{}
}

/*PostTelephonyProvidersEdgeUnpairInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgeUnpairInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairServiceUnavailable creates a PostTelephonyProvidersEdgeUnpairServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgeUnpairServiceUnavailable() *PostTelephonyProvidersEdgeUnpairServiceUnavailable {
	return &PostTelephonyProvidersEdgeUnpairServiceUnavailable{}
}

/*PostTelephonyProvidersEdgeUnpairServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgeUnpairServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgeUnpairGatewayTimeout creates a PostTelephonyProvidersEdgeUnpairGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgeUnpairGatewayTimeout() *PostTelephonyProvidersEdgeUnpairGatewayTimeout {
	return &PostTelephonyProvidersEdgeUnpairGatewayTimeout{}
}

/*PostTelephonyProvidersEdgeUnpairGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgeUnpairGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgeUnpairGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges/{edgeId}/unpair][%d] postTelephonyProvidersEdgeUnpairGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgeUnpairGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgeUnpairGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
