// Code generated by go-swagger; DO NOT EDIT.

package audit

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAuditsQueryRealtimeServicemappingReader is a Reader for the GetAuditsQueryRealtimeServicemapping structure.
type GetAuditsQueryRealtimeServicemappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditsQueryRealtimeServicemappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditsQueryRealtimeServicemappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuditsQueryRealtimeServicemappingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuditsQueryRealtimeServicemappingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuditsQueryRealtimeServicemappingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuditsQueryRealtimeServicemappingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuditsQueryRealtimeServicemappingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuditsQueryRealtimeServicemappingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuditsQueryRealtimeServicemappingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuditsQueryRealtimeServicemappingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuditsQueryRealtimeServicemappingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuditsQueryRealtimeServicemappingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuditsQueryRealtimeServicemappingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuditsQueryRealtimeServicemappingOK creates a GetAuditsQueryRealtimeServicemappingOK with default headers values
func NewGetAuditsQueryRealtimeServicemappingOK() *GetAuditsQueryRealtimeServicemappingOK {
	return &GetAuditsQueryRealtimeServicemappingOK{}
}

/*GetAuditsQueryRealtimeServicemappingOK handles this case with default header values.

successful operation
*/
type GetAuditsQueryRealtimeServicemappingOK struct {
	Payload *models.AuditQueryServiceMapping
}

func (o *GetAuditsQueryRealtimeServicemappingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingOK  %+v", 200, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingOK) GetPayload() *models.AuditQueryServiceMapping {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditQueryServiceMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingBadRequest creates a GetAuditsQueryRealtimeServicemappingBadRequest with default headers values
func NewGetAuditsQueryRealtimeServicemappingBadRequest() *GetAuditsQueryRealtimeServicemappingBadRequest {
	return &GetAuditsQueryRealtimeServicemappingBadRequest{}
}

/*GetAuditsQueryRealtimeServicemappingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuditsQueryRealtimeServicemappingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingUnauthorized creates a GetAuditsQueryRealtimeServicemappingUnauthorized with default headers values
func NewGetAuditsQueryRealtimeServicemappingUnauthorized() *GetAuditsQueryRealtimeServicemappingUnauthorized {
	return &GetAuditsQueryRealtimeServicemappingUnauthorized{}
}

/*GetAuditsQueryRealtimeServicemappingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuditsQueryRealtimeServicemappingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingForbidden creates a GetAuditsQueryRealtimeServicemappingForbidden with default headers values
func NewGetAuditsQueryRealtimeServicemappingForbidden() *GetAuditsQueryRealtimeServicemappingForbidden {
	return &GetAuditsQueryRealtimeServicemappingForbidden{}
}

/*GetAuditsQueryRealtimeServicemappingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuditsQueryRealtimeServicemappingForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingForbidden  %+v", 403, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingNotFound creates a GetAuditsQueryRealtimeServicemappingNotFound with default headers values
func NewGetAuditsQueryRealtimeServicemappingNotFound() *GetAuditsQueryRealtimeServicemappingNotFound {
	return &GetAuditsQueryRealtimeServicemappingNotFound{}
}

/*GetAuditsQueryRealtimeServicemappingNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuditsQueryRealtimeServicemappingNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingNotFound  %+v", 404, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingRequestTimeout creates a GetAuditsQueryRealtimeServicemappingRequestTimeout with default headers values
func NewGetAuditsQueryRealtimeServicemappingRequestTimeout() *GetAuditsQueryRealtimeServicemappingRequestTimeout {
	return &GetAuditsQueryRealtimeServicemappingRequestTimeout{}
}

/*GetAuditsQueryRealtimeServicemappingRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuditsQueryRealtimeServicemappingRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingRequestEntityTooLarge creates a GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge with default headers values
func NewGetAuditsQueryRealtimeServicemappingRequestEntityTooLarge() *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge {
	return &GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge{}
}

/*GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingUnsupportedMediaType creates a GetAuditsQueryRealtimeServicemappingUnsupportedMediaType with default headers values
func NewGetAuditsQueryRealtimeServicemappingUnsupportedMediaType() *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType {
	return &GetAuditsQueryRealtimeServicemappingUnsupportedMediaType{}
}

/*GetAuditsQueryRealtimeServicemappingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuditsQueryRealtimeServicemappingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingTooManyRequests creates a GetAuditsQueryRealtimeServicemappingTooManyRequests with default headers values
func NewGetAuditsQueryRealtimeServicemappingTooManyRequests() *GetAuditsQueryRealtimeServicemappingTooManyRequests {
	return &GetAuditsQueryRealtimeServicemappingTooManyRequests{}
}

/*GetAuditsQueryRealtimeServicemappingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuditsQueryRealtimeServicemappingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingInternalServerError creates a GetAuditsQueryRealtimeServicemappingInternalServerError with default headers values
func NewGetAuditsQueryRealtimeServicemappingInternalServerError() *GetAuditsQueryRealtimeServicemappingInternalServerError {
	return &GetAuditsQueryRealtimeServicemappingInternalServerError{}
}

/*GetAuditsQueryRealtimeServicemappingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuditsQueryRealtimeServicemappingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingServiceUnavailable creates a GetAuditsQueryRealtimeServicemappingServiceUnavailable with default headers values
func NewGetAuditsQueryRealtimeServicemappingServiceUnavailable() *GetAuditsQueryRealtimeServicemappingServiceUnavailable {
	return &GetAuditsQueryRealtimeServicemappingServiceUnavailable{}
}

/*GetAuditsQueryRealtimeServicemappingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuditsQueryRealtimeServicemappingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryRealtimeServicemappingGatewayTimeout creates a GetAuditsQueryRealtimeServicemappingGatewayTimeout with default headers values
func NewGetAuditsQueryRealtimeServicemappingGatewayTimeout() *GetAuditsQueryRealtimeServicemappingGatewayTimeout {
	return &GetAuditsQueryRealtimeServicemappingGatewayTimeout{}
}

/*GetAuditsQueryRealtimeServicemappingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuditsQueryRealtimeServicemappingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/realtime/servicemapping][%d] getAuditsQueryRealtimeServicemappingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryRealtimeServicemappingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
