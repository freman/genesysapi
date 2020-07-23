// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetArchitectIvrsReader is a Reader for the GetArchitectIvrs structure.
type GetArchitectIvrsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetArchitectIvrsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetArchitectIvrsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetArchitectIvrsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetArchitectIvrsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetArchitectIvrsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetArchitectIvrsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetArchitectIvrsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetArchitectIvrsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetArchitectIvrsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetArchitectIvrsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetArchitectIvrsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetArchitectIvrsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetArchitectIvrsOK creates a GetArchitectIvrsOK with default headers values
func NewGetArchitectIvrsOK() *GetArchitectIvrsOK {
	return &GetArchitectIvrsOK{}
}

/*GetArchitectIvrsOK handles this case with default header values.

successful operation
*/
type GetArchitectIvrsOK struct {
	Payload *models.IVREntityListing
}

func (o *GetArchitectIvrsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsOK  %+v", 200, o.Payload)
}

func (o *GetArchitectIvrsOK) GetPayload() *models.IVREntityListing {
	return o.Payload
}

func (o *GetArchitectIvrsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IVREntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsBadRequest creates a GetArchitectIvrsBadRequest with default headers values
func NewGetArchitectIvrsBadRequest() *GetArchitectIvrsBadRequest {
	return &GetArchitectIvrsBadRequest{}
}

/*GetArchitectIvrsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetArchitectIvrsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsBadRequest  %+v", 400, o.Payload)
}

func (o *GetArchitectIvrsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsUnauthorized creates a GetArchitectIvrsUnauthorized with default headers values
func NewGetArchitectIvrsUnauthorized() *GetArchitectIvrsUnauthorized {
	return &GetArchitectIvrsUnauthorized{}
}

/*GetArchitectIvrsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetArchitectIvrsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetArchitectIvrsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsForbidden creates a GetArchitectIvrsForbidden with default headers values
func NewGetArchitectIvrsForbidden() *GetArchitectIvrsForbidden {
	return &GetArchitectIvrsForbidden{}
}

/*GetArchitectIvrsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetArchitectIvrsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsForbidden  %+v", 403, o.Payload)
}

func (o *GetArchitectIvrsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsNotFound creates a GetArchitectIvrsNotFound with default headers values
func NewGetArchitectIvrsNotFound() *GetArchitectIvrsNotFound {
	return &GetArchitectIvrsNotFound{}
}

/*GetArchitectIvrsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetArchitectIvrsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsNotFound  %+v", 404, o.Payload)
}

func (o *GetArchitectIvrsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsRequestEntityTooLarge creates a GetArchitectIvrsRequestEntityTooLarge with default headers values
func NewGetArchitectIvrsRequestEntityTooLarge() *GetArchitectIvrsRequestEntityTooLarge {
	return &GetArchitectIvrsRequestEntityTooLarge{}
}

/*GetArchitectIvrsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetArchitectIvrsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetArchitectIvrsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsUnsupportedMediaType creates a GetArchitectIvrsUnsupportedMediaType with default headers values
func NewGetArchitectIvrsUnsupportedMediaType() *GetArchitectIvrsUnsupportedMediaType {
	return &GetArchitectIvrsUnsupportedMediaType{}
}

/*GetArchitectIvrsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetArchitectIvrsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetArchitectIvrsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsTooManyRequests creates a GetArchitectIvrsTooManyRequests with default headers values
func NewGetArchitectIvrsTooManyRequests() *GetArchitectIvrsTooManyRequests {
	return &GetArchitectIvrsTooManyRequests{}
}

/*GetArchitectIvrsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetArchitectIvrsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetArchitectIvrsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsInternalServerError creates a GetArchitectIvrsInternalServerError with default headers values
func NewGetArchitectIvrsInternalServerError() *GetArchitectIvrsInternalServerError {
	return &GetArchitectIvrsInternalServerError{}
}

/*GetArchitectIvrsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetArchitectIvrsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetArchitectIvrsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsServiceUnavailable creates a GetArchitectIvrsServiceUnavailable with default headers values
func NewGetArchitectIvrsServiceUnavailable() *GetArchitectIvrsServiceUnavailable {
	return &GetArchitectIvrsServiceUnavailable{}
}

/*GetArchitectIvrsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetArchitectIvrsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetArchitectIvrsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetArchitectIvrsGatewayTimeout creates a GetArchitectIvrsGatewayTimeout with default headers values
func NewGetArchitectIvrsGatewayTimeout() *GetArchitectIvrsGatewayTimeout {
	return &GetArchitectIvrsGatewayTimeout{}
}

/*GetArchitectIvrsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetArchitectIvrsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetArchitectIvrsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/architect/ivrs][%d] getArchitectIvrsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetArchitectIvrsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetArchitectIvrsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
