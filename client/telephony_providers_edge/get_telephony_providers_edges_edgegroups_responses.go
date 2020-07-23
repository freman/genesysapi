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

// GetTelephonyProvidersEdgesEdgegroupsReader is a Reader for the GetTelephonyProvidersEdgesEdgegroups structure.
type GetTelephonyProvidersEdgesEdgegroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTelephonyProvidersEdgesEdgegroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTelephonyProvidersEdgesEdgegroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTelephonyProvidersEdgesEdgegroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetTelephonyProvidersEdgesEdgegroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetTelephonyProvidersEdgesEdgegroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetTelephonyProvidersEdgesEdgegroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetTelephonyProvidersEdgesEdgegroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetTelephonyProvidersEdgesEdgegroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetTelephonyProvidersEdgesEdgegroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetTelephonyProvidersEdgesEdgegroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetTelephonyProvidersEdgesEdgegroupsOK creates a GetTelephonyProvidersEdgesEdgegroupsOK with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsOK() *GetTelephonyProvidersEdgesEdgegroupsOK {
	return &GetTelephonyProvidersEdgesEdgegroupsOK{}
}

/*GetTelephonyProvidersEdgesEdgegroupsOK handles this case with default header values.

successful operation
*/
type GetTelephonyProvidersEdgesEdgegroupsOK struct {
	Payload *models.EdgeGroupEntityListing
}

func (o *GetTelephonyProvidersEdgesEdgegroupsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsOK  %+v", 200, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsOK) GetPayload() *models.EdgeGroupEntityListing {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EdgeGroupEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsBadRequest creates a GetTelephonyProvidersEdgesEdgegroupsBadRequest with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsBadRequest() *GetTelephonyProvidersEdgesEdgegroupsBadRequest {
	return &GetTelephonyProvidersEdgesEdgegroupsBadRequest{}
}

/*GetTelephonyProvidersEdgesEdgegroupsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetTelephonyProvidersEdgesEdgegroupsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsUnauthorized creates a GetTelephonyProvidersEdgesEdgegroupsUnauthorized with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsUnauthorized() *GetTelephonyProvidersEdgesEdgegroupsUnauthorized {
	return &GetTelephonyProvidersEdgesEdgegroupsUnauthorized{}
}

/*GetTelephonyProvidersEdgesEdgegroupsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetTelephonyProvidersEdgesEdgegroupsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsForbidden creates a GetTelephonyProvidersEdgesEdgegroupsForbidden with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsForbidden() *GetTelephonyProvidersEdgesEdgegroupsForbidden {
	return &GetTelephonyProvidersEdgesEdgegroupsForbidden{}
}

/*GetTelephonyProvidersEdgesEdgegroupsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetTelephonyProvidersEdgesEdgegroupsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsForbidden  %+v", 403, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsNotFound creates a GetTelephonyProvidersEdgesEdgegroupsNotFound with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsNotFound() *GetTelephonyProvidersEdgesEdgegroupsNotFound {
	return &GetTelephonyProvidersEdgesEdgegroupsNotFound{}
}

/*GetTelephonyProvidersEdgesEdgegroupsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetTelephonyProvidersEdgesEdgegroupsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsNotFound  %+v", 404, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge creates a GetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge() *GetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge {
	return &GetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge{}
}

/*GetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType creates a GetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType() *GetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType {
	return &GetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType{}
}

/*GetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsTooManyRequests creates a GetTelephonyProvidersEdgesEdgegroupsTooManyRequests with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsTooManyRequests() *GetTelephonyProvidersEdgesEdgegroupsTooManyRequests {
	return &GetTelephonyProvidersEdgesEdgegroupsTooManyRequests{}
}

/*GetTelephonyProvidersEdgesEdgegroupsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetTelephonyProvidersEdgesEdgegroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsInternalServerError creates a GetTelephonyProvidersEdgesEdgegroupsInternalServerError with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsInternalServerError() *GetTelephonyProvidersEdgesEdgegroupsInternalServerError {
	return &GetTelephonyProvidersEdgesEdgegroupsInternalServerError{}
}

/*GetTelephonyProvidersEdgesEdgegroupsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetTelephonyProvidersEdgesEdgegroupsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsServiceUnavailable creates a GetTelephonyProvidersEdgesEdgegroupsServiceUnavailable with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsServiceUnavailable() *GetTelephonyProvidersEdgesEdgegroupsServiceUnavailable {
	return &GetTelephonyProvidersEdgesEdgegroupsServiceUnavailable{}
}

/*GetTelephonyProvidersEdgesEdgegroupsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetTelephonyProvidersEdgesEdgegroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTelephonyProvidersEdgesEdgegroupsGatewayTimeout creates a GetTelephonyProvidersEdgesEdgegroupsGatewayTimeout with default headers values
func NewGetTelephonyProvidersEdgesEdgegroupsGatewayTimeout() *GetTelephonyProvidersEdgesEdgegroupsGatewayTimeout {
	return &GetTelephonyProvidersEdgesEdgegroupsGatewayTimeout{}
}

/*GetTelephonyProvidersEdgesEdgegroupsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetTelephonyProvidersEdgesEdgegroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetTelephonyProvidersEdgesEdgegroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/telephony/providers/edges/edgegroups][%d] getTelephonyProvidersEdgesEdgegroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetTelephonyProvidersEdgesEdgegroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetTelephonyProvidersEdgesEdgegroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
