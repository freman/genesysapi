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

// PutTelephonyProvidersEdgeLineReader is a Reader for the PutTelephonyProvidersEdgeLine structure.
type PutTelephonyProvidersEdgeLineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTelephonyProvidersEdgeLineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTelephonyProvidersEdgeLineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTelephonyProvidersEdgeLineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutTelephonyProvidersEdgeLineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutTelephonyProvidersEdgeLineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutTelephonyProvidersEdgeLineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPutTelephonyProvidersEdgeLineRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutTelephonyProvidersEdgeLineRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutTelephonyProvidersEdgeLineUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTelephonyProvidersEdgeLineTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutTelephonyProvidersEdgeLineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutTelephonyProvidersEdgeLineServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutTelephonyProvidersEdgeLineGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutTelephonyProvidersEdgeLineOK creates a PutTelephonyProvidersEdgeLineOK with default headers values
func NewPutTelephonyProvidersEdgeLineOK() *PutTelephonyProvidersEdgeLineOK {
	return &PutTelephonyProvidersEdgeLineOK{}
}

/*PutTelephonyProvidersEdgeLineOK handles this case with default header values.

successful operation
*/
type PutTelephonyProvidersEdgeLineOK struct {
	Payload *models.EdgeLine
}

func (o *PutTelephonyProvidersEdgeLineOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineOK  %+v", 200, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineOK) GetPayload() *models.EdgeLine {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeLine)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineBadRequest creates a PutTelephonyProvidersEdgeLineBadRequest with default headers values
func NewPutTelephonyProvidersEdgeLineBadRequest() *PutTelephonyProvidersEdgeLineBadRequest {
	return &PutTelephonyProvidersEdgeLineBadRequest{}
}

/*PutTelephonyProvidersEdgeLineBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutTelephonyProvidersEdgeLineBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineBadRequest  %+v", 400, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineUnauthorized creates a PutTelephonyProvidersEdgeLineUnauthorized with default headers values
func NewPutTelephonyProvidersEdgeLineUnauthorized() *PutTelephonyProvidersEdgeLineUnauthorized {
	return &PutTelephonyProvidersEdgeLineUnauthorized{}
}

/*PutTelephonyProvidersEdgeLineUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutTelephonyProvidersEdgeLineUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineUnauthorized  %+v", 401, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineForbidden creates a PutTelephonyProvidersEdgeLineForbidden with default headers values
func NewPutTelephonyProvidersEdgeLineForbidden() *PutTelephonyProvidersEdgeLineForbidden {
	return &PutTelephonyProvidersEdgeLineForbidden{}
}

/*PutTelephonyProvidersEdgeLineForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutTelephonyProvidersEdgeLineForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineForbidden  %+v", 403, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineNotFound creates a PutTelephonyProvidersEdgeLineNotFound with default headers values
func NewPutTelephonyProvidersEdgeLineNotFound() *PutTelephonyProvidersEdgeLineNotFound {
	return &PutTelephonyProvidersEdgeLineNotFound{}
}

/*PutTelephonyProvidersEdgeLineNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutTelephonyProvidersEdgeLineNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineNotFound  %+v", 404, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineRequestTimeout creates a PutTelephonyProvidersEdgeLineRequestTimeout with default headers values
func NewPutTelephonyProvidersEdgeLineRequestTimeout() *PutTelephonyProvidersEdgeLineRequestTimeout {
	return &PutTelephonyProvidersEdgeLineRequestTimeout{}
}

/*PutTelephonyProvidersEdgeLineRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PutTelephonyProvidersEdgeLineRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineRequestTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineRequestTimeout  %+v", 408, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineRequestEntityTooLarge creates a PutTelephonyProvidersEdgeLineRequestEntityTooLarge with default headers values
func NewPutTelephonyProvidersEdgeLineRequestEntityTooLarge() *PutTelephonyProvidersEdgeLineRequestEntityTooLarge {
	return &PutTelephonyProvidersEdgeLineRequestEntityTooLarge{}
}

/*PutTelephonyProvidersEdgeLineRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutTelephonyProvidersEdgeLineRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineUnsupportedMediaType creates a PutTelephonyProvidersEdgeLineUnsupportedMediaType with default headers values
func NewPutTelephonyProvidersEdgeLineUnsupportedMediaType() *PutTelephonyProvidersEdgeLineUnsupportedMediaType {
	return &PutTelephonyProvidersEdgeLineUnsupportedMediaType{}
}

/*PutTelephonyProvidersEdgeLineUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutTelephonyProvidersEdgeLineUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineTooManyRequests creates a PutTelephonyProvidersEdgeLineTooManyRequests with default headers values
func NewPutTelephonyProvidersEdgeLineTooManyRequests() *PutTelephonyProvidersEdgeLineTooManyRequests {
	return &PutTelephonyProvidersEdgeLineTooManyRequests{}
}

/*PutTelephonyProvidersEdgeLineTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PutTelephonyProvidersEdgeLineTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineInternalServerError creates a PutTelephonyProvidersEdgeLineInternalServerError with default headers values
func NewPutTelephonyProvidersEdgeLineInternalServerError() *PutTelephonyProvidersEdgeLineInternalServerError {
	return &PutTelephonyProvidersEdgeLineInternalServerError{}
}

/*PutTelephonyProvidersEdgeLineInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutTelephonyProvidersEdgeLineInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineInternalServerError  %+v", 500, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineServiceUnavailable creates a PutTelephonyProvidersEdgeLineServiceUnavailable with default headers values
func NewPutTelephonyProvidersEdgeLineServiceUnavailable() *PutTelephonyProvidersEdgeLineServiceUnavailable {
	return &PutTelephonyProvidersEdgeLineServiceUnavailable{}
}

/*PutTelephonyProvidersEdgeLineServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutTelephonyProvidersEdgeLineServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgeLineGatewayTimeout creates a PutTelephonyProvidersEdgeLineGatewayTimeout with default headers values
func NewPutTelephonyProvidersEdgeLineGatewayTimeout() *PutTelephonyProvidersEdgeLineGatewayTimeout {
	return &PutTelephonyProvidersEdgeLineGatewayTimeout{}
}

/*PutTelephonyProvidersEdgeLineGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutTelephonyProvidersEdgeLineGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgeLineGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/{edgeId}/lines/{lineId}][%d] putTelephonyProvidersEdgeLineGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutTelephonyProvidersEdgeLineGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgeLineGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
