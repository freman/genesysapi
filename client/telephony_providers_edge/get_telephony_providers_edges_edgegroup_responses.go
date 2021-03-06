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

// GetTelephonyProvidersEdgesEdgegroupReader is a Reader for the GetTelephonyProvidersEdgesEdgegroup structure.
type GetTelephonyProvidersEdgesEdgegroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesEdgegroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesEdgegroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesEdgegroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesEdgegroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesEdgegroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesEdgegroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesEdgegroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesEdgegroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesEdgegroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesEdgegroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesEdgegroupOK creates a GetTelephonyProvidersEdgesEdgegroupOK with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupOK() *GetTelephonyProvidersEdgesEdgegroupOK {
	return &GetTelephonyProvidersEdgesEdgegroupOK{}
}

/*GetTelephonyProvidersEdgesEdgegroupOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesEdgegroupOK struct {
	Payload *models.EdgeGroup
}

func (o *GetTelephonyProvidersEdgesEdgegroupOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupOK) GetPayload() *models.EdgeGroup {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupBadRequest creates a GetTelephonyProvidersEdgesEdgegroupBadRequest with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupBadRequest() *GetTelephonyProvidersEdgesEdgegroupBadRequest {
	return &GetTelephonyProvidersEdgesEdgegroupBadRequest{}
}

/*GetTelephonyProvidersEdgesEdgegroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesEdgegroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupUnauthorized creates a GetTelephonyProvidersEdgesEdgegroupUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupUnauthorized() *GetTelephonyProvidersEdgesEdgegroupUnauthorized {
	return &GetTelephonyProvidersEdgesEdgegroupUnauthorized{}
}

/*GetTelephonyProvidersEdgesEdgegroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesEdgegroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupForbidden creates a GetTelephonyProvidersEdgesEdgegroupForbidden with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupForbidden() *GetTelephonyProvidersEdgesEdgegroupForbidden {
	return &GetTelephonyProvidersEdgesEdgegroupForbidden{}
}

/*GetTelephonyProvidersEdgesEdgegroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesEdgegroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupNotFound creates a GetTelephonyProvidersEdgesEdgegroupNotFound with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupNotFound() *GetTelephonyProvidersEdgesEdgegroupNotFound {
	return &GetTelephonyProvidersEdgesEdgegroupNotFound{}
}

/*GetTelephonyProvidersEdgesEdgegroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesEdgegroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge creates a GetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge() *GetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType creates a GetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType() *GetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupTooManyRequests creates a GetTelephonyProvidersEdgesEdgegroupTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupTooManyRequests() *GetTelephonyProvidersEdgesEdgegroupTooManyRequests {
	return &GetTelephonyProvidersEdgesEdgegroupTooManyRequests{}
}

/*GetTelephonyProvidersEdgesEdgegroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesEdgegroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupInternalServerError creates a GetTelephonyProvidersEdgesEdgegroupInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupInternalServerError() *GetTelephonyProvidersEdgesEdgegroupInternalServerError {
	return &GetTelephonyProvidersEdgesEdgegroupInternalServerError{}
}

/*GetTelephonyProvidersEdgesEdgegroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesEdgegroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupServiceUnavailable creates a GetTelephonyProvidersEdgesEdgegroupServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupServiceUnavailable() *GetTelephonyProvidersEdgesEdgegroupServiceUnavailable {
	return &GetTelephonyProvidersEdgesEdgegroupServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesEdgegroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesEdgegroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupGatewayTimeout creates a GetTelephonyProvidersEdgesEdgegroupGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupGatewayTimeout() *GetTelephonyProvidersEdgesEdgegroupGatewayTimeout {
	return &GetTelephonyProvidersEdgesEdgegroupGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesEdgegroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesEdgegroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] getTelephonyProvidersEdgesEdgegroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
