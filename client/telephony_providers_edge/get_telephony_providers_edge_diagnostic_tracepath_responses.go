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

// GetTelephonyProvidersEdgeDiagnosticTracepathReader is a Reader for the GetTelephonyProvidersEdgeDiagnosticTracepath structure.
type GetTelephonyProvidersEdgeDiagnosticTracepathReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgeDiagnosticTracepathReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathOK creates a GetTelephonyProvidersEdgeDiagnosticTracepathOK with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathOK() *GetTelephonyProvidersEdgeDiagnosticTracepathOK {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathOK{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathOK handles this case with default header values.

Request to get network diagnostic was successful.
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathOK struct {
	Payload *models.EdgeNetworkDiagnosticResponse
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathOK) GetPayload() *models.EdgeNetworkDiagnosticResponse {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeNetworkDiagnosticResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathBadRequest creates a GetTelephonyProvidersEdgeDiagnosticTracepathBadRequest with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathBadRequest() *GetTelephonyProvidersEdgeDiagnosticTracepathBadRequest {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathBadRequest{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized creates a GetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized() *GetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathForbidden creates a GetTelephonyProvidersEdgeDiagnosticTracepathForbidden with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathForbidden() *GetTelephonyProvidersEdgeDiagnosticTracepathForbidden {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathForbidden{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathNotFound creates a GetTelephonyProvidersEdgeDiagnosticTracepathNotFound with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathNotFound() *GetTelephonyProvidersEdgeDiagnosticTracepathNotFound {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathNotFound{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge creates a GetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge() *GetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType creates a GetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType() *GetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests creates a GetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests() *GetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError creates a GetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError() *GetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable creates a GetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable() *GetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout creates a GetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout() *GetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout {
	return &GetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout{}
}

/*GetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/diagnostic/tracepath][%d] getTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeDiagnosticTracepathGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
