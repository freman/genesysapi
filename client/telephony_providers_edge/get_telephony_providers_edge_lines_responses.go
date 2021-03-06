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

// GetTelephonyProvidersEdgeLinesReader is a Reader for the GetTelephonyProvidersEdgeLines structure.
type GetTelephonyProvidersEdgeLinesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgeLinesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgeLinesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgeLinesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgeLinesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgeLinesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgeLinesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgeLinesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgeLinesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgeLinesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgeLinesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgeLinesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgeLinesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgeLinesOK creates a GetTelephonyProvidersEdgeLinesOK with default headers values
func NewGetTelephonyProvidersEdgeLinesOK() *GetTelephonyProvidersEdgeLinesOK {
	return &GetTelephonyProvidersEdgeLinesOK{}
}

/*GetTelephonyProvidersEdgeLinesOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgeLinesOK struct {
	Payload *models.EdgeLineEntityListing
}

func (o *GetTelephonyProvidersEdgeLinesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesOK) GetPayload() *models.EdgeLineEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeLineEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesBadRequest creates a GetTelephonyProvidersEdgeLinesBadRequest with default headers values
func NewGetTelephonyProvidersEdgeLinesBadRequest() *GetTelephonyProvidersEdgeLinesBadRequest {
	return &GetTelephonyProvidersEdgeLinesBadRequest{}
}

/*GetTelephonyProvidersEdgeLinesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgeLinesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesUnauthorized creates a GetTelephonyProvidersEdgeLinesUnauthorized with default headers values
func NewGetTelephonyProvidersEdgeLinesUnauthorized() *GetTelephonyProvidersEdgeLinesUnauthorized {
	return &GetTelephonyProvidersEdgeLinesUnauthorized{}
}

/*GetTelephonyProvidersEdgeLinesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgeLinesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesForbidden creates a GetTelephonyProvidersEdgeLinesForbidden with default headers values
func NewGetTelephonyProvidersEdgeLinesForbidden() *GetTelephonyProvidersEdgeLinesForbidden {
	return &GetTelephonyProvidersEdgeLinesForbidden{}
}

/*GetTelephonyProvidersEdgeLinesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgeLinesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesNotFound creates a GetTelephonyProvidersEdgeLinesNotFound with default headers values
func NewGetTelephonyProvidersEdgeLinesNotFound() *GetTelephonyProvidersEdgeLinesNotFound {
	return &GetTelephonyProvidersEdgeLinesNotFound{}
}

/*GetTelephonyProvidersEdgeLinesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgeLinesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesRequestEntityTooLarge creates a GetTelephonyProvidersEdgeLinesRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgeLinesRequestEntityTooLarge() *GetTelephonyProvidersEdgeLinesRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgeLinesRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgeLinesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgeLinesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesUnsupportedMediaType creates a GetTelephonyProvidersEdgeLinesUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgeLinesUnsupportedMediaType() *GetTelephonyProvidersEdgeLinesUnsupportedMediaType {
	return &GetTelephonyProvidersEdgeLinesUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgeLinesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgeLinesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesTooManyRequests creates a GetTelephonyProvidersEdgeLinesTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgeLinesTooManyRequests() *GetTelephonyProvidersEdgeLinesTooManyRequests {
	return &GetTelephonyProvidersEdgeLinesTooManyRequests{}
}

/*GetTelephonyProvidersEdgeLinesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgeLinesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesInternalServerError creates a GetTelephonyProvidersEdgeLinesInternalServerError with default headers values
func NewGetTelephonyProvidersEdgeLinesInternalServerError() *GetTelephonyProvidersEdgeLinesInternalServerError {
	return &GetTelephonyProvidersEdgeLinesInternalServerError{}
}

/*GetTelephonyProvidersEdgeLinesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgeLinesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesServiceUnavailable creates a GetTelephonyProvidersEdgeLinesServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgeLinesServiceUnavailable() *GetTelephonyProvidersEdgeLinesServiceUnavailable {
	return &GetTelephonyProvidersEdgeLinesServiceUnavailable{}
}

/*GetTelephonyProvidersEdgeLinesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgeLinesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgeLinesGatewayTimeout creates a GetTelephonyProvidersEdgeLinesGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgeLinesGatewayTimeout() *GetTelephonyProvidersEdgeLinesGatewayTimeout {
	return &GetTelephonyProvidersEdgeLinesGatewayTimeout{}
}

/*GetTelephonyProvidersEdgeLinesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgeLinesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgeLinesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/{edgeId}/lines][%d] getTelephonyProvidersEdgeLinesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgeLinesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgeLinesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
