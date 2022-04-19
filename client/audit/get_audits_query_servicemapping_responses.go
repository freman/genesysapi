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

// GetAuditsQueryServicemappingReader is a Reader for the GetAuditsQueryServicemapping structure.
type GetAuditsQueryServicemappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuditsQueryServicemappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuditsQueryServicemappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuditsQueryServicemappingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuditsQueryServicemappingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuditsQueryServicemappingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuditsQueryServicemappingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuditsQueryServicemappingRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuditsQueryServicemappingRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuditsQueryServicemappingUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuditsQueryServicemappingTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuditsQueryServicemappingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuditsQueryServicemappingServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuditsQueryServicemappingGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuditsQueryServicemappingOK creates a GetAuditsQueryServicemappingOK with default headers values
func NewGetAuditsQueryServicemappingOK() *GetAuditsQueryServicemappingOK {
	return &GetAuditsQueryServicemappingOK{}
}

/*GetAuditsQueryServicemappingOK handles this case with default header values.

successful operation
*/
type GetAuditsQueryServicemappingOK struct {
	Payload *models.AuditQueryServiceMapping
}

func (o *GetAuditsQueryServicemappingOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingOK  %+v", 200, o.Payload)
}

func (o *GetAuditsQueryServicemappingOK) GetPayload() *models.AuditQueryServiceMapping {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuditQueryServiceMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingBadRequest creates a GetAuditsQueryServicemappingBadRequest with default headers values
func NewGetAuditsQueryServicemappingBadRequest() *GetAuditsQueryServicemappingBadRequest {
	return &GetAuditsQueryServicemappingBadRequest{}
}

/*GetAuditsQueryServicemappingBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuditsQueryServicemappingBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuditsQueryServicemappingBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingUnauthorized creates a GetAuditsQueryServicemappingUnauthorized with default headers values
func NewGetAuditsQueryServicemappingUnauthorized() *GetAuditsQueryServicemappingUnauthorized {
	return &GetAuditsQueryServicemappingUnauthorized{}
}

/*GetAuditsQueryServicemappingUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuditsQueryServicemappingUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuditsQueryServicemappingUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingForbidden creates a GetAuditsQueryServicemappingForbidden with default headers values
func NewGetAuditsQueryServicemappingForbidden() *GetAuditsQueryServicemappingForbidden {
	return &GetAuditsQueryServicemappingForbidden{}
}

/*GetAuditsQueryServicemappingForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuditsQueryServicemappingForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingForbidden  %+v", 403, o.Payload)
}

func (o *GetAuditsQueryServicemappingForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingNotFound creates a GetAuditsQueryServicemappingNotFound with default headers values
func NewGetAuditsQueryServicemappingNotFound() *GetAuditsQueryServicemappingNotFound {
	return &GetAuditsQueryServicemappingNotFound{}
}

/*GetAuditsQueryServicemappingNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuditsQueryServicemappingNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingNotFound  %+v", 404, o.Payload)
}

func (o *GetAuditsQueryServicemappingNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingRequestTimeout creates a GetAuditsQueryServicemappingRequestTimeout with default headers values
func NewGetAuditsQueryServicemappingRequestTimeout() *GetAuditsQueryServicemappingRequestTimeout {
	return &GetAuditsQueryServicemappingRequestTimeout{}
}

/*GetAuditsQueryServicemappingRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuditsQueryServicemappingRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuditsQueryServicemappingRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingRequestEntityTooLarge creates a GetAuditsQueryServicemappingRequestEntityTooLarge with default headers values
func NewGetAuditsQueryServicemappingRequestEntityTooLarge() *GetAuditsQueryServicemappingRequestEntityTooLarge {
	return &GetAuditsQueryServicemappingRequestEntityTooLarge{}
}

/*GetAuditsQueryServicemappingRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetAuditsQueryServicemappingRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuditsQueryServicemappingRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingUnsupportedMediaType creates a GetAuditsQueryServicemappingUnsupportedMediaType with default headers values
func NewGetAuditsQueryServicemappingUnsupportedMediaType() *GetAuditsQueryServicemappingUnsupportedMediaType {
	return &GetAuditsQueryServicemappingUnsupportedMediaType{}
}

/*GetAuditsQueryServicemappingUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuditsQueryServicemappingUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuditsQueryServicemappingUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingTooManyRequests creates a GetAuditsQueryServicemappingTooManyRequests with default headers values
func NewGetAuditsQueryServicemappingTooManyRequests() *GetAuditsQueryServicemappingTooManyRequests {
	return &GetAuditsQueryServicemappingTooManyRequests{}
}

/*GetAuditsQueryServicemappingTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuditsQueryServicemappingTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuditsQueryServicemappingTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingInternalServerError creates a GetAuditsQueryServicemappingInternalServerError with default headers values
func NewGetAuditsQueryServicemappingInternalServerError() *GetAuditsQueryServicemappingInternalServerError {
	return &GetAuditsQueryServicemappingInternalServerError{}
}

/*GetAuditsQueryServicemappingInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuditsQueryServicemappingInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuditsQueryServicemappingInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingServiceUnavailable creates a GetAuditsQueryServicemappingServiceUnavailable with default headers values
func NewGetAuditsQueryServicemappingServiceUnavailable() *GetAuditsQueryServicemappingServiceUnavailable {
	return &GetAuditsQueryServicemappingServiceUnavailable{}
}

/*GetAuditsQueryServicemappingServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuditsQueryServicemappingServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuditsQueryServicemappingServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuditsQueryServicemappingGatewayTimeout creates a GetAuditsQueryServicemappingGatewayTimeout with default headers values
func NewGetAuditsQueryServicemappingGatewayTimeout() *GetAuditsQueryServicemappingGatewayTimeout {
	return &GetAuditsQueryServicemappingGatewayTimeout{}
}

/*GetAuditsQueryServicemappingGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuditsQueryServicemappingGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuditsQueryServicemappingGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/audits/query/servicemapping][%d] getAuditsQueryServicemappingGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuditsQueryServicemappingGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuditsQueryServicemappingGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
