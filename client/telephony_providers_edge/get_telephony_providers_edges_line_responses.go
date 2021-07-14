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

// GetTelephonyProvidersEdgesLineReader is a Reader for the GetTelephonyProvidersEdgesLine structure.
type GetTelephonyProvidersEdgesLineReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesLineReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesLineOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesLineBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesLineUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesLineForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesLineNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesLineRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesLineRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesLineUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesLineTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesLineInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesLineServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesLineGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesLineOK creates a GetTelephonyProvidersEdgesLineOK with default headers values
func NewGetTelephonyProvidersEdgesLineOK() *GetTelephonyProvidersEdgesLineOK {
	return &GetTelephonyProvidersEdgesLineOK{}
}

/*GetTelephonyProvidersEdgesLineOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesLineOK struct {
	Payload *models.Line
}

func (o *GetTelephonyProvidersEdgesLineOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineOK) GetPayload() *models.Line {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Line)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineBadRequest creates a GetTelephonyProvidersEdgesLineBadRequest with default headers values
func NewGetTelephonyProvidersEdgesLineBadRequest() *GetTelephonyProvidersEdgesLineBadRequest {
	return &GetTelephonyProvidersEdgesLineBadRequest{}
}

/*GetTelephonyProvidersEdgesLineBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesLineBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineUnauthorized creates a GetTelephonyProvidersEdgesLineUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesLineUnauthorized() *GetTelephonyProvidersEdgesLineUnauthorized {
	return &GetTelephonyProvidersEdgesLineUnauthorized{}
}

/*GetTelephonyProvidersEdgesLineUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesLineUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineForbidden creates a GetTelephonyProvidersEdgesLineForbidden with default headers values
func NewGetTelephonyProvidersEdgesLineForbidden() *GetTelephonyProvidersEdgesLineForbidden {
	return &GetTelephonyProvidersEdgesLineForbidden{}
}

/*GetTelephonyProvidersEdgesLineForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesLineForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineNotFound creates a GetTelephonyProvidersEdgesLineNotFound with default headers values
func NewGetTelephonyProvidersEdgesLineNotFound() *GetTelephonyProvidersEdgesLineNotFound {
	return &GetTelephonyProvidersEdgesLineNotFound{}
}

/*GetTelephonyProvidersEdgesLineNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesLineNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineRequestTimeout creates a GetTelephonyProvidersEdgesLineRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesLineRequestTimeout() *GetTelephonyProvidersEdgesLineRequestTimeout {
	return &GetTelephonyProvidersEdgesLineRequestTimeout{}
}

/*GetTelephonyProvidersEdgesLineRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesLineRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineRequestEntityTooLarge creates a GetTelephonyProvidersEdgesLineRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesLineRequestEntityTooLarge() *GetTelephonyProvidersEdgesLineRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesLineRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesLineRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesLineRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineUnsupportedMediaType creates a GetTelephonyProvidersEdgesLineUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesLineUnsupportedMediaType() *GetTelephonyProvidersEdgesLineUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesLineUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesLineUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesLineUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineTooManyRequests creates a GetTelephonyProvidersEdgesLineTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesLineTooManyRequests() *GetTelephonyProvidersEdgesLineTooManyRequests {
	return &GetTelephonyProvidersEdgesLineTooManyRequests{}
}

/*GetTelephonyProvidersEdgesLineTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesLineTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineInternalServerError creates a GetTelephonyProvidersEdgesLineInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesLineInternalServerError() *GetTelephonyProvidersEdgesLineInternalServerError {
	return &GetTelephonyProvidersEdgesLineInternalServerError{}
}

/*GetTelephonyProvidersEdgesLineInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesLineInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineServiceUnavailable creates a GetTelephonyProvidersEdgesLineServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesLineServiceUnavailable() *GetTelephonyProvidersEdgesLineServiceUnavailable {
	return &GetTelephonyProvidersEdgesLineServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesLineServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesLineServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesLineGatewayTimeout creates a GetTelephonyProvidersEdgesLineGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesLineGatewayTimeout() *GetTelephonyProvidersEdgesLineGatewayTimeout {
	return &GetTelephonyProvidersEdgesLineGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesLineGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesLineGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesLineGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/lines/{lineId}][%d] getTelephonyProvidersEdgesLineGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesLineGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesLineGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
