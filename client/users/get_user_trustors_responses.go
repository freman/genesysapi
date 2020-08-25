// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetUserTrustorsReader is a Reader for the GetUserTrustors structure.
type GetUserTrustorsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserTrustorsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserTrustorsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserTrustorsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserTrustorsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserTrustorsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserTrustorsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUserTrustorsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUserTrustorsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUserTrustorsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserTrustorsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUserTrustorsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUserTrustorsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUserTrustorsOK creates a GetUserTrustorsOK with default headers values
func NewGetUserTrustorsOK() *GetUserTrustorsOK {
	return &GetUserTrustorsOK{}
}

/*GetUserTrustorsOK handles this case with default header values.

successful operation
*/
type GetUserTrustorsOK struct {
	Payload *models.TrustorEntityListing
}

func (o *GetUserTrustorsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsOK  %+v", 200, o.Payload)
}

func (o *GetUserTrustorsOK) GetPayload() *models.TrustorEntityListing {
	return o.Payload
}

func (o *GetUserTrustorsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrustorEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsBadRequest creates a GetUserTrustorsBadRequest with default headers values
func NewGetUserTrustorsBadRequest() *GetUserTrustorsBadRequest {
	return &GetUserTrustorsBadRequest{}
}

/*GetUserTrustorsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUserTrustorsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserTrustorsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsUnauthorized creates a GetUserTrustorsUnauthorized with default headers values
func NewGetUserTrustorsUnauthorized() *GetUserTrustorsUnauthorized {
	return &GetUserTrustorsUnauthorized{}
}

/*GetUserTrustorsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUserTrustorsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserTrustorsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsForbidden creates a GetUserTrustorsForbidden with default headers values
func NewGetUserTrustorsForbidden() *GetUserTrustorsForbidden {
	return &GetUserTrustorsForbidden{}
}

/*GetUserTrustorsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUserTrustorsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserTrustorsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsNotFound creates a GetUserTrustorsNotFound with default headers values
func NewGetUserTrustorsNotFound() *GetUserTrustorsNotFound {
	return &GetUserTrustorsNotFound{}
}

/*GetUserTrustorsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUserTrustorsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserTrustorsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsRequestEntityTooLarge creates a GetUserTrustorsRequestEntityTooLarge with default headers values
func NewGetUserTrustorsRequestEntityTooLarge() *GetUserTrustorsRequestEntityTooLarge {
	return &GetUserTrustorsRequestEntityTooLarge{}
}

/*GetUserTrustorsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetUserTrustorsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUserTrustorsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsUnsupportedMediaType creates a GetUserTrustorsUnsupportedMediaType with default headers values
func NewGetUserTrustorsUnsupportedMediaType() *GetUserTrustorsUnsupportedMediaType {
	return &GetUserTrustorsUnsupportedMediaType{}
}

/*GetUserTrustorsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUserTrustorsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUserTrustorsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsTooManyRequests creates a GetUserTrustorsTooManyRequests with default headers values
func NewGetUserTrustorsTooManyRequests() *GetUserTrustorsTooManyRequests {
	return &GetUserTrustorsTooManyRequests{}
}

/*GetUserTrustorsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetUserTrustorsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUserTrustorsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsInternalServerError creates a GetUserTrustorsInternalServerError with default headers values
func NewGetUserTrustorsInternalServerError() *GetUserTrustorsInternalServerError {
	return &GetUserTrustorsInternalServerError{}
}

/*GetUserTrustorsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUserTrustorsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserTrustorsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsServiceUnavailable creates a GetUserTrustorsServiceUnavailable with default headers values
func NewGetUserTrustorsServiceUnavailable() *GetUserTrustorsServiceUnavailable {
	return &GetUserTrustorsServiceUnavailable{}
}

/*GetUserTrustorsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUserTrustorsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUserTrustorsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserTrustorsGatewayTimeout creates a GetUserTrustorsGatewayTimeout with default headers values
func NewGetUserTrustorsGatewayTimeout() *GetUserTrustorsGatewayTimeout {
	return &GetUserTrustorsGatewayTimeout{}
}

/*GetUserTrustorsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUserTrustorsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUserTrustorsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{userId}/trustors][%d] getUserTrustorsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUserTrustorsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUserTrustorsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}