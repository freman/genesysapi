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

// GetTelephonyProvidersEdgesReader is a Reader for the GetTelephonyProvidersEdges structure.
type GetTelephonyProvidersEdgesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesOK creates a GetTelephonyProvidersEdgesOK with default headers values
func NewGetTelephonyProvidersEdgesOK() *GetTelephonyProvidersEdgesOK {
	return &GetTelephonyProvidersEdgesOK{}
}

/*GetTelephonyProvidersEdgesOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesOK struct {
	Payload *models.EdgeEntityListing
}

func (o *GetTelephonyProvidersEdgesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesOK) GetPayload() *models.EdgeEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesBadRequest creates a GetTelephonyProvidersEdgesBadRequest with default headers values
func NewGetTelephonyProvidersEdgesBadRequest() *GetTelephonyProvidersEdgesBadRequest {
	return &GetTelephonyProvidersEdgesBadRequest{}
}

/*GetTelephonyProvidersEdgesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesUnauthorized creates a GetTelephonyProvidersEdgesUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesUnauthorized() *GetTelephonyProvidersEdgesUnauthorized {
	return &GetTelephonyProvidersEdgesUnauthorized{}
}

/*GetTelephonyProvidersEdgesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesForbidden creates a GetTelephonyProvidersEdgesForbidden with default headers values
func NewGetTelephonyProvidersEdgesForbidden() *GetTelephonyProvidersEdgesForbidden {
	return &GetTelephonyProvidersEdgesForbidden{}
}

/*GetTelephonyProvidersEdgesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesNotFound creates a GetTelephonyProvidersEdgesNotFound with default headers values
func NewGetTelephonyProvidersEdgesNotFound() *GetTelephonyProvidersEdgesNotFound {
	return &GetTelephonyProvidersEdgesNotFound{}
}

/*GetTelephonyProvidersEdgesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesRequestTimeout creates a GetTelephonyProvidersEdgesRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesRequestTimeout() *GetTelephonyProvidersEdgesRequestTimeout {
	return &GetTelephonyProvidersEdgesRequestTimeout{}
}

/*GetTelephonyProvidersEdgesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesRequestEntityTooLarge creates a GetTelephonyProvidersEdgesRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesRequestEntityTooLarge() *GetTelephonyProvidersEdgesRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesUnsupportedMediaType creates a GetTelephonyProvidersEdgesUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesUnsupportedMediaType() *GetTelephonyProvidersEdgesUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesTooManyRequests creates a GetTelephonyProvidersEdgesTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesTooManyRequests() *GetTelephonyProvidersEdgesTooManyRequests {
	return &GetTelephonyProvidersEdgesTooManyRequests{}
}

/*GetTelephonyProvidersEdgesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesInternalServerError creates a GetTelephonyProvidersEdgesInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesInternalServerError() *GetTelephonyProvidersEdgesInternalServerError {
	return &GetTelephonyProvidersEdgesInternalServerError{}
}

/*GetTelephonyProvidersEdgesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesServiceUnavailable creates a GetTelephonyProvidersEdgesServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesServiceUnavailable() *GetTelephonyProvidersEdgesServiceUnavailable {
	return &GetTelephonyProvidersEdgesServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesGatewayTimeout creates a GetTelephonyProvidersEdgesGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesGatewayTimeout() *GetTelephonyProvidersEdgesGatewayTimeout {
	return &GetTelephonyProvidersEdgesGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges][%d] getTelephonyProvidersEdgesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
