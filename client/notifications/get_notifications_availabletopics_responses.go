// Code generated by go-swagger; DO NOT EDIT.

package notifications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetNotificationsAvailabletopicsReader is a Reader for the GetNotificationsAvailabletopics structure.
type GetNotificationsAvailabletopicsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNotificationsAvailabletopicsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNotificationsAvailabletopicsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetNotificationsAvailabletopicsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetNotificationsAvailabletopicsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetNotificationsAvailabletopicsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNotificationsAvailabletopicsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetNotificationsAvailabletopicsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetNotificationsAvailabletopicsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetNotificationsAvailabletopicsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetNotificationsAvailabletopicsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetNotificationsAvailabletopicsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetNotificationsAvailabletopicsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNotificationsAvailabletopicsOK creates a GetNotificationsAvailabletopicsOK with default headers values
func NewGetNotificationsAvailabletopicsOK() *GetNotificationsAvailabletopicsOK {
	return &GetNotificationsAvailabletopicsOK{}
}

/*GetNotificationsAvailabletopicsOK handles this case with default header values.

successful operation
*/
type GetNotificationsAvailabletopicsOK struct {
	Payload *models.AvailableTopicEntityListing
}

func (o *GetNotificationsAvailabletopicsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsOK  %+v", 200, o.Payload)
}

func (o *GetNotificationsAvailabletopicsOK) GetPayload() *models.AvailableTopicEntityListing {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AvailableTopicEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsBadRequest creates a GetNotificationsAvailabletopicsBadRequest with default headers values
func NewGetNotificationsAvailabletopicsBadRequest() *GetNotificationsAvailabletopicsBadRequest {
	return &GetNotificationsAvailabletopicsBadRequest{}
}

/*GetNotificationsAvailabletopicsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetNotificationsAvailabletopicsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsBadRequest  %+v", 400, o.Payload)
}

func (o *GetNotificationsAvailabletopicsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsUnauthorized creates a GetNotificationsAvailabletopicsUnauthorized with default headers values
func NewGetNotificationsAvailabletopicsUnauthorized() *GetNotificationsAvailabletopicsUnauthorized {
	return &GetNotificationsAvailabletopicsUnauthorized{}
}

/*GetNotificationsAvailabletopicsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetNotificationsAvailabletopicsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNotificationsAvailabletopicsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsForbidden creates a GetNotificationsAvailabletopicsForbidden with default headers values
func NewGetNotificationsAvailabletopicsForbidden() *GetNotificationsAvailabletopicsForbidden {
	return &GetNotificationsAvailabletopicsForbidden{}
}

/*GetNotificationsAvailabletopicsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetNotificationsAvailabletopicsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsForbidden  %+v", 403, o.Payload)
}

func (o *GetNotificationsAvailabletopicsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsNotFound creates a GetNotificationsAvailabletopicsNotFound with default headers values
func NewGetNotificationsAvailabletopicsNotFound() *GetNotificationsAvailabletopicsNotFound {
	return &GetNotificationsAvailabletopicsNotFound{}
}

/*GetNotificationsAvailabletopicsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetNotificationsAvailabletopicsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsNotFound  %+v", 404, o.Payload)
}

func (o *GetNotificationsAvailabletopicsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsRequestEntityTooLarge creates a GetNotificationsAvailabletopicsRequestEntityTooLarge with default headers values
func NewGetNotificationsAvailabletopicsRequestEntityTooLarge() *GetNotificationsAvailabletopicsRequestEntityTooLarge {
	return &GetNotificationsAvailabletopicsRequestEntityTooLarge{}
}

/*GetNotificationsAvailabletopicsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetNotificationsAvailabletopicsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsUnsupportedMediaType creates a GetNotificationsAvailabletopicsUnsupportedMediaType with default headers values
func NewGetNotificationsAvailabletopicsUnsupportedMediaType() *GetNotificationsAvailabletopicsUnsupportedMediaType {
	return &GetNotificationsAvailabletopicsUnsupportedMediaType{}
}

/*GetNotificationsAvailabletopicsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetNotificationsAvailabletopicsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsTooManyRequests creates a GetNotificationsAvailabletopicsTooManyRequests with default headers values
func NewGetNotificationsAvailabletopicsTooManyRequests() *GetNotificationsAvailabletopicsTooManyRequests {
	return &GetNotificationsAvailabletopicsTooManyRequests{}
}

/*GetNotificationsAvailabletopicsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetNotificationsAvailabletopicsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetNotificationsAvailabletopicsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsInternalServerError creates a GetNotificationsAvailabletopicsInternalServerError with default headers values
func NewGetNotificationsAvailabletopicsInternalServerError() *GetNotificationsAvailabletopicsInternalServerError {
	return &GetNotificationsAvailabletopicsInternalServerError{}
}

/*GetNotificationsAvailabletopicsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetNotificationsAvailabletopicsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetNotificationsAvailabletopicsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsServiceUnavailable creates a GetNotificationsAvailabletopicsServiceUnavailable with default headers values
func NewGetNotificationsAvailabletopicsServiceUnavailable() *GetNotificationsAvailabletopicsServiceUnavailable {
	return &GetNotificationsAvailabletopicsServiceUnavailable{}
}

/*GetNotificationsAvailabletopicsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetNotificationsAvailabletopicsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetNotificationsAvailabletopicsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNotificationsAvailabletopicsGatewayTimeout creates a GetNotificationsAvailabletopicsGatewayTimeout with default headers values
func NewGetNotificationsAvailabletopicsGatewayTimeout() *GetNotificationsAvailabletopicsGatewayTimeout {
	return &GetNotificationsAvailabletopicsGatewayTimeout{}
}

/*GetNotificationsAvailabletopicsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetNotificationsAvailabletopicsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetNotificationsAvailabletopicsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/notifications/availabletopics][%d] getNotificationsAvailabletopicsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetNotificationsAvailabletopicsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetNotificationsAvailabletopicsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}