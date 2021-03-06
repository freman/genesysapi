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

// PostTelephonyProvidersEdgesReader is a Reader for the PostTelephonyProvidersEdges structure.
type PostTelephonyProvidersEdgesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostTelephonyProvidersEdgesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostTelephonyProvidersEdgesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostTelephonyProvidersEdgesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostTelephonyProvidersEdgesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostTelephonyProvidersEdgesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostTelephonyProvidersEdgesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostTelephonyProvidersEdgesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostTelephonyProvidersEdgesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostTelephonyProvidersEdgesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostTelephonyProvidersEdgesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostTelephonyProvidersEdgesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostTelephonyProvidersEdgesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostTelephonyProvidersEdgesOK creates a PostTelephonyProvidersEdgesOK with default headers values
func NewPostTelephonyProvidersEdgesOK() *PostTelephonyProvidersEdgesOK {
	return &PostTelephonyProvidersEdgesOK{}
}

/*PostTelephonyProvidersEdgesOK handles this case with default header values.

successful operation
*/
type PostTelephonyProvidersEdgesOK struct {
	Payload *models.Edge
}

func (o *PostTelephonyProvidersEdgesOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesOK  %+v", 200, o.Payload)
}

func (o *PostTelephonyProvidersEdgesOK) GetPayload() *models.Edge {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Edge)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesBadRequest creates a PostTelephonyProvidersEdgesBadRequest with default headers values
func NewPostTelephonyProvidersEdgesBadRequest() *PostTelephonyProvidersEdgesBadRequest {
	return &PostTelephonyProvidersEdgesBadRequest{}
}

/*PostTelephonyProvidersEdgesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostTelephonyProvidersEdgesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesBadRequest  %+v", 400, o.Payload)
}

func (o *PostTelephonyProvidersEdgesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesUnauthorized creates a PostTelephonyProvidersEdgesUnauthorized with default headers values
func NewPostTelephonyProvidersEdgesUnauthorized() *PostTelephonyProvidersEdgesUnauthorized {
	return &PostTelephonyProvidersEdgesUnauthorized{}
}

/*PostTelephonyProvidersEdgesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostTelephonyProvidersEdgesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostTelephonyProvidersEdgesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesForbidden creates a PostTelephonyProvidersEdgesForbidden with default headers values
func NewPostTelephonyProvidersEdgesForbidden() *PostTelephonyProvidersEdgesForbidden {
	return &PostTelephonyProvidersEdgesForbidden{}
}

/*PostTelephonyProvidersEdgesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostTelephonyProvidersEdgesForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesForbidden  %+v", 403, o.Payload)
}

func (o *PostTelephonyProvidersEdgesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesNotFound creates a PostTelephonyProvidersEdgesNotFound with default headers values
func NewPostTelephonyProvidersEdgesNotFound() *PostTelephonyProvidersEdgesNotFound {
	return &PostTelephonyProvidersEdgesNotFound{}
}

/*PostTelephonyProvidersEdgesNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostTelephonyProvidersEdgesNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesNotFound  %+v", 404, o.Payload)
}

func (o *PostTelephonyProvidersEdgesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesRequestEntityTooLarge creates a PostTelephonyProvidersEdgesRequestEntityTooLarge with default headers values
func NewPostTelephonyProvidersEdgesRequestEntityTooLarge() *PostTelephonyProvidersEdgesRequestEntityTooLarge {
	return &PostTelephonyProvidersEdgesRequestEntityTooLarge{}
}

/*PostTelephonyProvidersEdgesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostTelephonyProvidersEdgesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostTelephonyProvidersEdgesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesUnsupportedMediaType creates a PostTelephonyProvidersEdgesUnsupportedMediaType with default headers values
func NewPostTelephonyProvidersEdgesUnsupportedMediaType() *PostTelephonyProvidersEdgesUnsupportedMediaType {
	return &PostTelephonyProvidersEdgesUnsupportedMediaType{}
}

/*PostTelephonyProvidersEdgesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostTelephonyProvidersEdgesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostTelephonyProvidersEdgesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesTooManyRequests creates a PostTelephonyProvidersEdgesTooManyRequests with default headers values
func NewPostTelephonyProvidersEdgesTooManyRequests() *PostTelephonyProvidersEdgesTooManyRequests {
	return &PostTelephonyProvidersEdgesTooManyRequests{}
}

/*PostTelephonyProvidersEdgesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostTelephonyProvidersEdgesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostTelephonyProvidersEdgesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesInternalServerError creates a PostTelephonyProvidersEdgesInternalServerError with default headers values
func NewPostTelephonyProvidersEdgesInternalServerError() *PostTelephonyProvidersEdgesInternalServerError {
	return &PostTelephonyProvidersEdgesInternalServerError{}
}

/*PostTelephonyProvidersEdgesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostTelephonyProvidersEdgesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostTelephonyProvidersEdgesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesServiceUnavailable creates a PostTelephonyProvidersEdgesServiceUnavailable with default headers values
func NewPostTelephonyProvidersEdgesServiceUnavailable() *PostTelephonyProvidersEdgesServiceUnavailable {
	return &PostTelephonyProvidersEdgesServiceUnavailable{}
}

/*PostTelephonyProvidersEdgesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostTelephonyProvidersEdgesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostTelephonyProvidersEdgesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostTelephonyProvidersEdgesGatewayTimeout creates a PostTelephonyProvidersEdgesGatewayTimeout with default headers values
func NewPostTelephonyProvidersEdgesGatewayTimeout() *PostTelephonyProvidersEdgesGatewayTimeout {
	return &PostTelephonyProvidersEdgesGatewayTimeout{}
}

/*PostTelephonyProvidersEdgesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostTelephonyProvidersEdgesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostTelephonyProvidersEdgesGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/telephony/providers/edges][%d] postTelephonyProvidersEdgesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostTelephonyProvidersEdgesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostTelephonyProvidersEdgesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
