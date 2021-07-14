// Code generated by go-swagger; DO NOT EDIT.

package gamification

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGamificationProfilesReader is a Reader for the GetGamificationProfiles structure.
type GetGamificationProfilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGamificationProfilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGamificationProfilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGamificationProfilesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGamificationProfilesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGamificationProfilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGamificationProfilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetGamificationProfilesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGamificationProfilesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGamificationProfilesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGamificationProfilesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGamificationProfilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGamificationProfilesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGamificationProfilesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGamificationProfilesOK creates a GetGamificationProfilesOK with default headers values
func NewGetGamificationProfilesOK() *GetGamificationProfilesOK {
	return &GetGamificationProfilesOK{}
}

/*GetGamificationProfilesOK handles this case with default header values.

successful operation
*/
type GetGamificationProfilesOK struct {
	Payload *models.GetProfilesResponse
}

func (o *GetGamificationProfilesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesOK  %+v", 200, o.Payload)
}

func (o *GetGamificationProfilesOK) GetPayload() *models.GetProfilesResponse {
	return o.Payload
}

func (o *GetGamificationProfilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetProfilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesBadRequest creates a GetGamificationProfilesBadRequest with default headers values
func NewGetGamificationProfilesBadRequest() *GetGamificationProfilesBadRequest {
	return &GetGamificationProfilesBadRequest{}
}

/*GetGamificationProfilesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGamificationProfilesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesBadRequest  %+v", 400, o.Payload)
}

func (o *GetGamificationProfilesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUnauthorized creates a GetGamificationProfilesUnauthorized with default headers values
func NewGetGamificationProfilesUnauthorized() *GetGamificationProfilesUnauthorized {
	return &GetGamificationProfilesUnauthorized{}
}

/*GetGamificationProfilesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGamificationProfilesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGamificationProfilesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesForbidden creates a GetGamificationProfilesForbidden with default headers values
func NewGetGamificationProfilesForbidden() *GetGamificationProfilesForbidden {
	return &GetGamificationProfilesForbidden{}
}

/*GetGamificationProfilesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGamificationProfilesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesForbidden  %+v", 403, o.Payload)
}

func (o *GetGamificationProfilesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesNotFound creates a GetGamificationProfilesNotFound with default headers values
func NewGetGamificationProfilesNotFound() *GetGamificationProfilesNotFound {
	return &GetGamificationProfilesNotFound{}
}

/*GetGamificationProfilesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGamificationProfilesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesNotFound  %+v", 404, o.Payload)
}

func (o *GetGamificationProfilesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesRequestTimeout creates a GetGamificationProfilesRequestTimeout with default headers values
func NewGetGamificationProfilesRequestTimeout() *GetGamificationProfilesRequestTimeout {
	return &GetGamificationProfilesRequestTimeout{}
}

/*GetGamificationProfilesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetGamificationProfilesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetGamificationProfilesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesRequestEntityTooLarge creates a GetGamificationProfilesRequestEntityTooLarge with default headers values
func NewGetGamificationProfilesRequestEntityTooLarge() *GetGamificationProfilesRequestEntityTooLarge {
	return &GetGamificationProfilesRequestEntityTooLarge{}
}

/*GetGamificationProfilesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGamificationProfilesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGamificationProfilesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesUnsupportedMediaType creates a GetGamificationProfilesUnsupportedMediaType with default headers values
func NewGetGamificationProfilesUnsupportedMediaType() *GetGamificationProfilesUnsupportedMediaType {
	return &GetGamificationProfilesUnsupportedMediaType{}
}

/*GetGamificationProfilesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGamificationProfilesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGamificationProfilesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesTooManyRequests creates a GetGamificationProfilesTooManyRequests with default headers values
func NewGetGamificationProfilesTooManyRequests() *GetGamificationProfilesTooManyRequests {
	return &GetGamificationProfilesTooManyRequests{}
}

/*GetGamificationProfilesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetGamificationProfilesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGamificationProfilesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesInternalServerError creates a GetGamificationProfilesInternalServerError with default headers values
func NewGetGamificationProfilesInternalServerError() *GetGamificationProfilesInternalServerError {
	return &GetGamificationProfilesInternalServerError{}
}

/*GetGamificationProfilesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGamificationProfilesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGamificationProfilesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesServiceUnavailable creates a GetGamificationProfilesServiceUnavailable with default headers values
func NewGetGamificationProfilesServiceUnavailable() *GetGamificationProfilesServiceUnavailable {
	return &GetGamificationProfilesServiceUnavailable{}
}

/*GetGamificationProfilesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGamificationProfilesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGamificationProfilesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGamificationProfilesGatewayTimeout creates a GetGamificationProfilesGatewayTimeout with default headers values
func NewGetGamificationProfilesGatewayTimeout() *GetGamificationProfilesGatewayTimeout {
	return &GetGamificationProfilesGatewayTimeout{}
}

/*GetGamificationProfilesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGamificationProfilesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGamificationProfilesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/gamification/profiles][%d] getGamificationProfilesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGamificationProfilesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGamificationProfilesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
