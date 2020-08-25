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

// GetTelephonyProvidersEdgesEdgeversionreportReader is a Reader for the GetTelephonyProvidersEdgesEdgeversionreport structure.
type GetTelephonyProvidersEdgesEdgeversionreportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesEdgeversionreportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesEdgeversionreportOK creates a GetTelephonyProvidersEdgesEdgeversionreportOK with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportOK() *GetTelephonyProvidersEdgesEdgeversionreportOK {
	return &GetTelephonyProvidersEdgesEdgeversionreportOK{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesEdgeversionreportOK struct {
	Payload *models.EdgeVersionReport
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) GetPayload() *models.EdgeVersionReport {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeVersionReport)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportBadRequest creates a GetTelephonyProvidersEdgesEdgeversionreportBadRequest with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportBadRequest() *GetTelephonyProvidersEdgesEdgeversionreportBadRequest {
	return &GetTelephonyProvidersEdgesEdgeversionreportBadRequest{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesEdgeversionreportBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportUnauthorized creates a GetTelephonyProvidersEdgesEdgeversionreportUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportUnauthorized() *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized {
	return &GetTelephonyProvidersEdgesEdgeversionreportUnauthorized{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesEdgeversionreportUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportForbidden creates a GetTelephonyProvidersEdgesEdgeversionreportForbidden with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportForbidden() *GetTelephonyProvidersEdgesEdgeversionreportForbidden {
	return &GetTelephonyProvidersEdgesEdgeversionreportForbidden{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesEdgeversionreportForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportNotFound creates a GetTelephonyProvidersEdgesEdgeversionreportNotFound with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportNotFound() *GetTelephonyProvidersEdgesEdgeversionreportNotFound {
	return &GetTelephonyProvidersEdgesEdgeversionreportNotFound{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesEdgeversionreportNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge creates a GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge() *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType creates a GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType() *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportTooManyRequests creates a GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportTooManyRequests() *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests {
	return &GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportInternalServerError creates a GetTelephonyProvidersEdgesEdgeversionreportInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportInternalServerError() *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError {
	return &GetTelephonyProvidersEdgesEdgeversionreportInternalServerError{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesEdgeversionreportInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable creates a GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable() *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable {
	return &GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout creates a GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout() *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout {
	return &GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgeversionreport][%d] getTelephonyProvidersEdgesEdgeversionreportGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgeversionreportGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}