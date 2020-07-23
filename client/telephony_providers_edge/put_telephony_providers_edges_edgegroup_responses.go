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

// PutTelephonyProvidersEdgesEdgegroupReader is a Reader for the PutTelephonyProvidersEdgesEdgegroup structure.
type PutTelephonyProvidersEdgesEdgegroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutTelephonyProvidersEdgesEdgegroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutTelephonyProvidersEdgesEdgegroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutTelephonyProvidersEdgesEdgegroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutTelephonyProvidersEdgesEdgegroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutTelephonyProvidersEdgesEdgegroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutTelephonyProvidersEdgesEdgegroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPutTelephonyProvidersEdgesEdgegroupTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutTelephonyProvidersEdgesEdgegroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPutTelephonyProvidersEdgesEdgegroupServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPutTelephonyProvidersEdgesEdgegroupGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPutTelephonyProvidersEdgesEdgegroupOK creates a PutTelephonyProvidersEdgesEdgegroupOK with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupOK() *PutTelephonyProvidersEdgesEdgegroupOK {
	return &PutTelephonyProvidersEdgesEdgegroupOK{}
}

/*PutTelephonyProvidersEdgesEdgegroupOK handles this case with default header values.

successful operation
*/
type PutTelephonyProvidersEdgesEdgegroupOK struct {
	Payload *models.EdgeGroup
}

func (o *PutTelephonyProvidersEdgesEdgegroupOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupOK  %+v", 200, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupOK) GetPayload() *models.EdgeGroup {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupBadRequest creates a PutTelephonyProvidersEdgesEdgegroupBadRequest with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupBadRequest() *PutTelephonyProvidersEdgesEdgegroupBadRequest {
	return &PutTelephonyProvidersEdgesEdgegroupBadRequest{}
}

/*PutTelephonyProvidersEdgesEdgegroupBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PutTelephonyProvidersEdgesEdgegroupBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupBadRequest  %+v", 400, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupUnauthorized creates a PutTelephonyProvidersEdgesEdgegroupUnauthorized with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupUnauthorized() *PutTelephonyProvidersEdgesEdgegroupUnauthorized {
	return &PutTelephonyProvidersEdgesEdgegroupUnauthorized{}
}

/*PutTelephonyProvidersEdgesEdgegroupUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PutTelephonyProvidersEdgesEdgegroupUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupUnauthorized  %+v", 401, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupForbidden creates a PutTelephonyProvidersEdgesEdgegroupForbidden with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupForbidden() *PutTelephonyProvidersEdgesEdgegroupForbidden {
	return &PutTelephonyProvidersEdgesEdgegroupForbidden{}
}

/*PutTelephonyProvidersEdgesEdgegroupForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PutTelephonyProvidersEdgesEdgegroupForbidden struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupForbidden  %+v", 403, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupNotFound creates a PutTelephonyProvidersEdgesEdgegroupNotFound with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupNotFound() *PutTelephonyProvidersEdgesEdgegroupNotFound {
	return &PutTelephonyProvidersEdgesEdgegroupNotFound{}
}

/*PutTelephonyProvidersEdgesEdgegroupNotFound handles this case with default header values.

The requested resource was not found.
*/
type PutTelephonyProvidersEdgesEdgegroupNotFound struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupNotFound  %+v", 404, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge creates a PutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge() *PutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge {
	return &PutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge{}
}

/*PutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType creates a PutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType() *PutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType {
	return &PutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType{}
}

/*PutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupTooManyRequests creates a PutTelephonyProvidersEdgesEdgegroupTooManyRequests with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupTooManyRequests() *PutTelephonyProvidersEdgesEdgegroupTooManyRequests {
	return &PutTelephonyProvidersEdgesEdgegroupTooManyRequests{}
}

/*PutTelephonyProvidersEdgesEdgegroupTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PutTelephonyProvidersEdgesEdgegroupTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupTooManyRequests) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupTooManyRequests  %+v", 429, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupInternalServerError creates a PutTelephonyProvidersEdgesEdgegroupInternalServerError with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupInternalServerError() *PutTelephonyProvidersEdgesEdgegroupInternalServerError {
	return &PutTelephonyProvidersEdgesEdgegroupInternalServerError{}
}

/*PutTelephonyProvidersEdgesEdgegroupInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PutTelephonyProvidersEdgesEdgegroupInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupInternalServerError  %+v", 500, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupServiceUnavailable creates a PutTelephonyProvidersEdgesEdgegroupServiceUnavailable with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupServiceUnavailable() *PutTelephonyProvidersEdgesEdgegroupServiceUnavailable {
	return &PutTelephonyProvidersEdgesEdgegroupServiceUnavailable{}
}

/*PutTelephonyProvidersEdgesEdgegroupServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PutTelephonyProvidersEdgesEdgegroupServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupServiceUnavailable) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutTelephonyProvidersEdgesEdgegroupGatewayTimeout creates a PutTelephonyProvidersEdgesEdgegroupGatewayTimeout with default headers values
func NewPutTelephonyProvidersEdgesEdgegroupGatewayTimeout() *PutTelephonyProvidersEdgesEdgegroupGatewayTimeout {
	return &PutTelephonyProvidersEdgesEdgegroupGatewayTimeout{}
}

/*PutTelephonyProvidersEdgesEdgegroupGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PutTelephonyProvidersEdgesEdgegroupGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PutTelephonyProvidersEdgesEdgegroupGatewayTimeout) Error() string {
	return fmt.Sprintf("[PUT /api/v2/telephony/providers/edges/edgegroups/{edgeGroupId}][%d] putTelephonyProvidersEdgesEdgegroupGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PutTelephonyProvidersEdgesEdgegroupGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PutTelephonyProvidersEdgesEdgegroupGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
