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

// GetUsersDevelopmentActivitiesReader is a Reader for the GetUsersDevelopmentActivities structure.
type GetUsersDevelopmentActivitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersDevelopmentActivitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersDevelopmentActivitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersDevelopmentActivitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsersDevelopmentActivitiesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersDevelopmentActivitiesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersDevelopmentActivitiesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetUsersDevelopmentActivitiesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetUsersDevelopmentActivitiesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetUsersDevelopmentActivitiesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersDevelopmentActivitiesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetUsersDevelopmentActivitiesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetUsersDevelopmentActivitiesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersDevelopmentActivitiesOK creates a GetUsersDevelopmentActivitiesOK with default headers values
func NewGetUsersDevelopmentActivitiesOK() *GetUsersDevelopmentActivitiesOK {
	return &GetUsersDevelopmentActivitiesOK{}
}

/*GetUsersDevelopmentActivitiesOK handles this case with default header values.

successful operation
*/
type GetUsersDevelopmentActivitiesOK struct {
	Payload *models.DevelopmentActivityListing
}

func (o *GetUsersDevelopmentActivitiesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesOK  %+v", 200, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesOK) GetPayload() *models.DevelopmentActivityListing {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DevelopmentActivityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesBadRequest creates a GetUsersDevelopmentActivitiesBadRequest with default headers values
func NewGetUsersDevelopmentActivitiesBadRequest() *GetUsersDevelopmentActivitiesBadRequest {
	return &GetUsersDevelopmentActivitiesBadRequest{}
}

/*GetUsersDevelopmentActivitiesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetUsersDevelopmentActivitiesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesBadRequest  %+v", 400, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesUnauthorized creates a GetUsersDevelopmentActivitiesUnauthorized with default headers values
func NewGetUsersDevelopmentActivitiesUnauthorized() *GetUsersDevelopmentActivitiesUnauthorized {
	return &GetUsersDevelopmentActivitiesUnauthorized{}
}

/*GetUsersDevelopmentActivitiesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetUsersDevelopmentActivitiesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesForbidden creates a GetUsersDevelopmentActivitiesForbidden with default headers values
func NewGetUsersDevelopmentActivitiesForbidden() *GetUsersDevelopmentActivitiesForbidden {
	return &GetUsersDevelopmentActivitiesForbidden{}
}

/*GetUsersDevelopmentActivitiesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetUsersDevelopmentActivitiesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesForbidden  %+v", 403, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesNotFound creates a GetUsersDevelopmentActivitiesNotFound with default headers values
func NewGetUsersDevelopmentActivitiesNotFound() *GetUsersDevelopmentActivitiesNotFound {
	return &GetUsersDevelopmentActivitiesNotFound{}
}

/*GetUsersDevelopmentActivitiesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetUsersDevelopmentActivitiesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesNotFound  %+v", 404, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesRequestEntityTooLarge creates a GetUsersDevelopmentActivitiesRequestEntityTooLarge with default headers values
func NewGetUsersDevelopmentActivitiesRequestEntityTooLarge() *GetUsersDevelopmentActivitiesRequestEntityTooLarge {
	return &GetUsersDevelopmentActivitiesRequestEntityTooLarge{}
}

/*GetUsersDevelopmentActivitiesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetUsersDevelopmentActivitiesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesUnsupportedMediaType creates a GetUsersDevelopmentActivitiesUnsupportedMediaType with default headers values
func NewGetUsersDevelopmentActivitiesUnsupportedMediaType() *GetUsersDevelopmentActivitiesUnsupportedMediaType {
	return &GetUsersDevelopmentActivitiesUnsupportedMediaType{}
}

/*GetUsersDevelopmentActivitiesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetUsersDevelopmentActivitiesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesTooManyRequests creates a GetUsersDevelopmentActivitiesTooManyRequests with default headers values
func NewGetUsersDevelopmentActivitiesTooManyRequests() *GetUsersDevelopmentActivitiesTooManyRequests {
	return &GetUsersDevelopmentActivitiesTooManyRequests{}
}

/*GetUsersDevelopmentActivitiesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetUsersDevelopmentActivitiesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesInternalServerError creates a GetUsersDevelopmentActivitiesInternalServerError with default headers values
func NewGetUsersDevelopmentActivitiesInternalServerError() *GetUsersDevelopmentActivitiesInternalServerError {
	return &GetUsersDevelopmentActivitiesInternalServerError{}
}

/*GetUsersDevelopmentActivitiesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetUsersDevelopmentActivitiesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesServiceUnavailable creates a GetUsersDevelopmentActivitiesServiceUnavailable with default headers values
func NewGetUsersDevelopmentActivitiesServiceUnavailable() *GetUsersDevelopmentActivitiesServiceUnavailable {
	return &GetUsersDevelopmentActivitiesServiceUnavailable{}
}

/*GetUsersDevelopmentActivitiesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetUsersDevelopmentActivitiesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDevelopmentActivitiesGatewayTimeout creates a GetUsersDevelopmentActivitiesGatewayTimeout with default headers values
func NewGetUsersDevelopmentActivitiesGatewayTimeout() *GetUsersDevelopmentActivitiesGatewayTimeout {
	return &GetUsersDevelopmentActivitiesGatewayTimeout{}
}

/*GetUsersDevelopmentActivitiesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetUsersDevelopmentActivitiesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetUsersDevelopmentActivitiesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/development/activities][%d] getUsersDevelopmentActivitiesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetUsersDevelopmentActivitiesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetUsersDevelopmentActivitiesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
