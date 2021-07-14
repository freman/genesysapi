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

// GetTelephonyProvidersEdgesDidsReader is a Reader for the GetTelephonyProvidersEdgesDids structure.
type GetTelephonyProvidersEdgesDidsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesDidsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesDidsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesDidsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesDidsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesDidsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesDidsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetTelephonyProvidersEdgesDidsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesDidsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesDidsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesDidsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesDidsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesDidsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesDidsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesDidsOK creates a GetTelephonyProvidersEdgesDidsOK with default headers values
func NewGetTelephonyProvidersEdgesDidsOK() *GetTelephonyProvidersEdgesDidsOK {
	return &GetTelephonyProvidersEdgesDidsOK{}
}

/*GetTelephonyProvidersEdgesDidsOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesDidsOK struct {
	Payload *models.DIDEntityListing
}

func (o *GetTelephonyProvidersEdgesDidsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsOK) GetPayload() *models.DIDEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DIDEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsBadRequest creates a GetTelephonyProvidersEdgesDidsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesDidsBadRequest() *GetTelephonyProvidersEdgesDidsBadRequest {
	return &GetTelephonyProvidersEdgesDidsBadRequest{}
}

/*GetTelephonyProvidersEdgesDidsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesDidsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsUnauthorized creates a GetTelephonyProvidersEdgesDidsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesDidsUnauthorized() *GetTelephonyProvidersEdgesDidsUnauthorized {
	return &GetTelephonyProvidersEdgesDidsUnauthorized{}
}

/*GetTelephonyProvidersEdgesDidsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesDidsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsForbidden creates a GetTelephonyProvidersEdgesDidsForbidden with default headers values
func NewGetTelephonyProvidersEdgesDidsForbidden() *GetTelephonyProvidersEdgesDidsForbidden {
	return &GetTelephonyProvidersEdgesDidsForbidden{}
}

/*GetTelephonyProvidersEdgesDidsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesDidsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsNotFound creates a GetTelephonyProvidersEdgesDidsNotFound with default headers values
func NewGetTelephonyProvidersEdgesDidsNotFound() *GetTelephonyProvidersEdgesDidsNotFound {
	return &GetTelephonyProvidersEdgesDidsNotFound{}
}

/*GetTelephonyProvidersEdgesDidsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesDidsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsRequestTimeout creates a GetTelephonyProvidersEdgesDidsRequestTimeout with default headers values
func NewGetTelephonyProvidersEdgesDidsRequestTimeout() *GetTelephonyProvidersEdgesDidsRequestTimeout {
	return &GetTelephonyProvidersEdgesDidsRequestTimeout{}
}

/*GetTelephonyProvidersEdgesDidsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetTelephonyProvidersEdgesDidsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesDidsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesDidsRequestEntityTooLarge() *GetTelephonyProvidersEdgesDidsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesDidsRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesDidsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesDidsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsUnsupportedMediaType creates a GetTelephonyProvidersEdgesDidsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesDidsUnsupportedMediaType() *GetTelephonyProvidersEdgesDidsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesDidsUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesDidsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesDidsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsTooManyRequests creates a GetTelephonyProvidersEdgesDidsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesDidsTooManyRequests() *GetTelephonyProvidersEdgesDidsTooManyRequests {
	return &GetTelephonyProvidersEdgesDidsTooManyRequests{}
}

/*GetTelephonyProvidersEdgesDidsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetTelephonyProvidersEdgesDidsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsInternalServerError creates a GetTelephonyProvidersEdgesDidsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesDidsInternalServerError() *GetTelephonyProvidersEdgesDidsInternalServerError {
	return &GetTelephonyProvidersEdgesDidsInternalServerError{}
}

/*GetTelephonyProvidersEdgesDidsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesDidsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsServiceUnavailable creates a GetTelephonyProvidersEdgesDidsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesDidsServiceUnavailable() *GetTelephonyProvidersEdgesDidsServiceUnavailable {
	return &GetTelephonyProvidersEdgesDidsServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesDidsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesDidsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesDidsGatewayTimeout creates a GetTelephonyProvidersEdgesDidsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesDidsGatewayTimeout() *GetTelephonyProvidersEdgesDidsGatewayTimeout {
	return &GetTelephonyProvidersEdgesDidsGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesDidsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesDidsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesDidsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/dids][%d] getTelephonyProvidersEdgesDidsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesDidsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesDidsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
