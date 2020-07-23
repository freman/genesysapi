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

// PutTelephonyProvidersEdgeReader is a Reader for the PutTelephonyProvidersEdge structure.
type PutTelephonyProvidersEdgeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTelephonyProvidersEdgeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTelephonyProvidersEdgeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTelephonyProvidersEdgeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutTelephonyProvidersEdgeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutTelephonyProvidersEdgeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutTelephonyProvidersEdgeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutTelephonyProvidersEdgeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutTelephonyProvidersEdgeRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutTelephonyProvidersEdgeUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTelephonyProvidersEdgeTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutTelephonyProvidersEdgeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutTelephonyProvidersEdgeServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutTelephonyProvidersEdgeGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutTelephonyProvidersEdgeOK creates a PutTelephonyProvidersEdgeOK with default headers values
func NewPutTelephonyProvidersEdgeOK() *PutTelephonyProvidersEdgeOK {
	return &PutTelephonyProvidersEdgeOK{}
}

/*PutTelephonyProvidersEdgeOK handles this case with default header values.

successful operation
*/
type PutTelephonyProvidersEdgeOK struct {
	Payload *models.Edge
}

func (o *PutTelephonyProvidersEdgeOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeOK  %+v", 200, o.Payload)
}

func (o *PutTelephonyProvidersEdgeOK) GetPayload() *models.Edge {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Edge)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeBadRequest creates a PutTelephonyProvidersEdgeBadRequest with default headers values
func NewPutTelephonyProvidersEdgeBadRequest() *PutTelephonyProvidersEdgeBadRequest {
	return &PutTelephonyProvidersEdgeBadRequest{}
}

/*PutTelephonyProvidersEdgeBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutTelephonyProvidersEdgeBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeBadRequest  %+v", 400, o.Payload)
}

func (o *PutTelephonyProvidersEdgeBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeUnauthorized creates a PutTelephonyProvidersEdgeUnauthorized with default headers values
func NewPutTelephonyProvidersEdgeUnauthorized() *PutTelephonyProvidersEdgeUnauthorized {
	return &PutTelephonyProvidersEdgeUnauthorized{}
}

/*PutTelephonyProvidersEdgeUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutTelephonyProvidersEdgeUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeUnauthorized  %+v", 401, o.Payload)
}

func (o *PutTelephonyProvidersEdgeUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeForbidden creates a PutTelephonyProvidersEdgeForbidden with default headers values
func NewPutTelephonyProvidersEdgeForbidden() *PutTelephonyProvidersEdgeForbidden {
	return &PutTelephonyProvidersEdgeForbidden{}
}

/*PutTelephonyProvidersEdgeForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutTelephonyProvidersEdgeForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeForbidden  %+v", 403, o.Payload)
}

func (o *PutTelephonyProvidersEdgeForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeNotFound creates a PutTelephonyProvidersEdgeNotFound with default headers values
func NewPutTelephonyProvidersEdgeNotFound() *PutTelephonyProvidersEdgeNotFound {
	return &PutTelephonyProvidersEdgeNotFound{}
}

/*PutTelephonyProvidersEdgeNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutTelephonyProvidersEdgeNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeNotFound  %+v", 404, o.Payload)
}

func (o *PutTelephonyProvidersEdgeNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeConflict creates a PutTelephonyProvidersEdgeConflict with default headers values
func NewPutTelephonyProvidersEdgeConflict() *PutTelephonyProvidersEdgeConflict {
	return &PutTelephonyProvidersEdgeConflict{}
}

/*PutTelephonyProvidersEdgeConflict handles this case with default header values.

Conflict
*/
type PutTelephonyProvidersEdgeConflict struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeConflict) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeConflict  %+v", 409, o.Payload)
}

func (o *PutTelephonyProvidersEdgeConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeRequestEntityTooLarge creates a PutTelephonyProvidersEdgeRequestEntityTooLarge with default headers values
func NewPutTelephonyProvidersEdgeRequestEntityTooLarge() *PutTelephonyProvidersEdgeRequestEntityTooLarge {
	return &PutTelephonyProvidersEdgeRequestEntityTooLarge{}
}

/*PutTelephonyProvidersEdgeRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutTelephonyProvidersEdgeRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutTelephonyProvidersEdgeRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeUnsupportedMediaType creates a PutTelephonyProvidersEdgeUnsupportedMediaType with default headers values
func NewPutTelephonyProvidersEdgeUnsupportedMediaType() *PutTelephonyProvidersEdgeUnsupportedMediaType {
	return &PutTelephonyProvidersEdgeUnsupportedMediaType{}
}

/*PutTelephonyProvidersEdgeUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutTelephonyProvidersEdgeUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutTelephonyProvidersEdgeUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeTooManyRequests creates a PutTelephonyProvidersEdgeTooManyRequests with default headers values
func NewPutTelephonyProvidersEdgeTooManyRequests() *PutTelephonyProvidersEdgeTooManyRequests {
	return &PutTelephonyProvidersEdgeTooManyRequests{}
}

/*PutTelephonyProvidersEdgeTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutTelephonyProvidersEdgeTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutTelephonyProvidersEdgeTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeInternalServerError creates a PutTelephonyProvidersEdgeInternalServerError with default headers values
func NewPutTelephonyProvidersEdgeInternalServerError() *PutTelephonyProvidersEdgeInternalServerError {
	return &PutTelephonyProvidersEdgeInternalServerError{}
}

/*PutTelephonyProvidersEdgeInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutTelephonyProvidersEdgeInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeInternalServerError  %+v", 500, o.Payload)
}

func (o *PutTelephonyProvidersEdgeInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeServiceUnavailable creates a PutTelephonyProvidersEdgeServiceUnavailable with default headers values
func NewPutTelephonyProvidersEdgeServiceUnavailable() *PutTelephonyProvidersEdgeServiceUnavailable {
	return &PutTelephonyProvidersEdgeServiceUnavailable{}
}

/*PutTelephonyProvidersEdgeServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutTelephonyProvidersEdgeServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutTelephonyProvidersEdgeServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeGatewayTimeout creates a PutTelephonyProvidersEdgeGatewayTimeout with default headers values
func NewPutTelephonyProvidersEdgeGatewayTimeout() *PutTelephonyProvidersEdgeGatewayTimeout {
	return &PutTelephonyProvidersEdgeGatewayTimeout{}
}

/*PutTelephonyProvidersEdgeGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutTelephonyProvidersEdgeGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}][%d] putTelephonyProvidersEdgeGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutTelephonyProvidersEdgeGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
