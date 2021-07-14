// Code generated by go-swagger; DO NOT EDIT.

package outbound

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOutboundAttemptlimitsReader is a Reader for the GetOutboundAttemptlimits structure.
type GetOutboundAttemptlimitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOutboundAttemptlimitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOutboundAttemptlimitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOutboundAttemptlimitsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOutboundAttemptlimitsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOutboundAttemptlimitsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOutboundAttemptlimitsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOutboundAttemptlimitsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOutboundAttemptlimitsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOutboundAttemptlimitsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOutboundAttemptlimitsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOutboundAttemptlimitsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOutboundAttemptlimitsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOutboundAttemptlimitsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOutboundAttemptlimitsOK creates a GetOutboundAttemptlimitsOK with default headers values
func NewGetOutboundAttemptlimitsOK() *GetOutboundAttemptlimitsOK {
	return &GetOutboundAttemptlimitsOK{}
}

/*GetOutboundAttemptlimitsOK handles this case with default header values.

successful operation
*/
type GetOutboundAttemptlimitsOK struct {
	Payload *models.AttemptLimitsEntityListing
}

func (o *GetOutboundAttemptlimitsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsOK  %+v", 200, o.Payload)
}

func (o *GetOutboundAttemptlimitsOK) GetPayload() *models.AttemptLimitsEntityListing {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttemptLimitsEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsBadRequest creates a GetOutboundAttemptlimitsBadRequest with default headers values
func NewGetOutboundAttemptlimitsBadRequest() *GetOutboundAttemptlimitsBadRequest {
	return &GetOutboundAttemptlimitsBadRequest{}
}

/*GetOutboundAttemptlimitsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOutboundAttemptlimitsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsBadRequest  %+v", 400, o.Payload)
}

func (o *GetOutboundAttemptlimitsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsUnauthorized creates a GetOutboundAttemptlimitsUnauthorized with default headers values
func NewGetOutboundAttemptlimitsUnauthorized() *GetOutboundAttemptlimitsUnauthorized {
	return &GetOutboundAttemptlimitsUnauthorized{}
}

/*GetOutboundAttemptlimitsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOutboundAttemptlimitsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOutboundAttemptlimitsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsForbidden creates a GetOutboundAttemptlimitsForbidden with default headers values
func NewGetOutboundAttemptlimitsForbidden() *GetOutboundAttemptlimitsForbidden {
	return &GetOutboundAttemptlimitsForbidden{}
}

/*GetOutboundAttemptlimitsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetOutboundAttemptlimitsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsForbidden  %+v", 403, o.Payload)
}

func (o *GetOutboundAttemptlimitsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsNotFound creates a GetOutboundAttemptlimitsNotFound with default headers values
func NewGetOutboundAttemptlimitsNotFound() *GetOutboundAttemptlimitsNotFound {
	return &GetOutboundAttemptlimitsNotFound{}
}

/*GetOutboundAttemptlimitsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetOutboundAttemptlimitsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsNotFound  %+v", 404, o.Payload)
}

func (o *GetOutboundAttemptlimitsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsRequestTimeout creates a GetOutboundAttemptlimitsRequestTimeout with default headers values
func NewGetOutboundAttemptlimitsRequestTimeout() *GetOutboundAttemptlimitsRequestTimeout {
	return &GetOutboundAttemptlimitsRequestTimeout{}
}

/*GetOutboundAttemptlimitsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOutboundAttemptlimitsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOutboundAttemptlimitsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsRequestEntityTooLarge creates a GetOutboundAttemptlimitsRequestEntityTooLarge with default headers values
func NewGetOutboundAttemptlimitsRequestEntityTooLarge() *GetOutboundAttemptlimitsRequestEntityTooLarge {
	return &GetOutboundAttemptlimitsRequestEntityTooLarge{}
}

/*GetOutboundAttemptlimitsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetOutboundAttemptlimitsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsUnsupportedMediaType creates a GetOutboundAttemptlimitsUnsupportedMediaType with default headers values
func NewGetOutboundAttemptlimitsUnsupportedMediaType() *GetOutboundAttemptlimitsUnsupportedMediaType {
	return &GetOutboundAttemptlimitsUnsupportedMediaType{}
}

/*GetOutboundAttemptlimitsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOutboundAttemptlimitsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOutboundAttemptlimitsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsTooManyRequests creates a GetOutboundAttemptlimitsTooManyRequests with default headers values
func NewGetOutboundAttemptlimitsTooManyRequests() *GetOutboundAttemptlimitsTooManyRequests {
	return &GetOutboundAttemptlimitsTooManyRequests{}
}

/*GetOutboundAttemptlimitsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOutboundAttemptlimitsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOutboundAttemptlimitsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsInternalServerError creates a GetOutboundAttemptlimitsInternalServerError with default headers values
func NewGetOutboundAttemptlimitsInternalServerError() *GetOutboundAttemptlimitsInternalServerError {
	return &GetOutboundAttemptlimitsInternalServerError{}
}

/*GetOutboundAttemptlimitsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOutboundAttemptlimitsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOutboundAttemptlimitsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsServiceUnavailable creates a GetOutboundAttemptlimitsServiceUnavailable with default headers values
func NewGetOutboundAttemptlimitsServiceUnavailable() *GetOutboundAttemptlimitsServiceUnavailable {
	return &GetOutboundAttemptlimitsServiceUnavailable{}
}

/*GetOutboundAttemptlimitsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOutboundAttemptlimitsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOutboundAttemptlimitsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOutboundAttemptlimitsGatewayTimeout creates a GetOutboundAttemptlimitsGatewayTimeout with default headers values
func NewGetOutboundAttemptlimitsGatewayTimeout() *GetOutboundAttemptlimitsGatewayTimeout {
	return &GetOutboundAttemptlimitsGatewayTimeout{}
}

/*GetOutboundAttemptlimitsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetOutboundAttemptlimitsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetOutboundAttemptlimitsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/outbound/attemptlimits][%d] getOutboundAttemptlimitsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOutboundAttemptlimitsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOutboundAttemptlimitsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
