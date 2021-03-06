// Code generated by go-swagger; DO NOT EDIT.

package groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetGroupIndividualsReader is a Reader for the GetGroupIndividuals structure.
type GetGroupIndividualsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGroupIndividualsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGroupIndividualsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetGroupIndividualsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetGroupIndividualsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGroupIndividualsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetGroupIndividualsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetGroupIndividualsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetGroupIndividualsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetGroupIndividualsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGroupIndividualsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetGroupIndividualsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetGroupIndividualsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGroupIndividualsOK creates a GetGroupIndividualsOK with default headers values
func NewGetGroupIndividualsOK() *GetGroupIndividualsOK {
	return &GetGroupIndividualsOK{}
}

/*GetGroupIndividualsOK handles this case with default header values.

successful operation
*/
type GetGroupIndividualsOK struct {
	Payload *models.UserEntityListing
}

func (o *GetGroupIndividualsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsOK  %+v", 200, o.Payload)
}

func (o *GetGroupIndividualsOK) GetPayload() *models.UserEntityListing {
	return o.Payload
}

func (o *GetGroupIndividualsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsBadRequest creates a GetGroupIndividualsBadRequest with default headers values
func NewGetGroupIndividualsBadRequest() *GetGroupIndividualsBadRequest {
	return &GetGroupIndividualsBadRequest{}
}

/*GetGroupIndividualsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetGroupIndividualsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsBadRequest  %+v", 400, o.Payload)
}

func (o *GetGroupIndividualsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsUnauthorized creates a GetGroupIndividualsUnauthorized with default headers values
func NewGetGroupIndividualsUnauthorized() *GetGroupIndividualsUnauthorized {
	return &GetGroupIndividualsUnauthorized{}
}

/*GetGroupIndividualsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetGroupIndividualsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGroupIndividualsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsForbidden creates a GetGroupIndividualsForbidden with default headers values
func NewGetGroupIndividualsForbidden() *GetGroupIndividualsForbidden {
	return &GetGroupIndividualsForbidden{}
}

/*GetGroupIndividualsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetGroupIndividualsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsForbidden  %+v", 403, o.Payload)
}

func (o *GetGroupIndividualsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsNotFound creates a GetGroupIndividualsNotFound with default headers values
func NewGetGroupIndividualsNotFound() *GetGroupIndividualsNotFound {
	return &GetGroupIndividualsNotFound{}
}

/*GetGroupIndividualsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetGroupIndividualsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsNotFound  %+v", 404, o.Payload)
}

func (o *GetGroupIndividualsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsRequestEntityTooLarge creates a GetGroupIndividualsRequestEntityTooLarge with default headers values
func NewGetGroupIndividualsRequestEntityTooLarge() *GetGroupIndividualsRequestEntityTooLarge {
	return &GetGroupIndividualsRequestEntityTooLarge{}
}

/*GetGroupIndividualsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetGroupIndividualsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetGroupIndividualsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsUnsupportedMediaType creates a GetGroupIndividualsUnsupportedMediaType with default headers values
func NewGetGroupIndividualsUnsupportedMediaType() *GetGroupIndividualsUnsupportedMediaType {
	return &GetGroupIndividualsUnsupportedMediaType{}
}

/*GetGroupIndividualsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetGroupIndividualsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetGroupIndividualsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsTooManyRequests creates a GetGroupIndividualsTooManyRequests with default headers values
func NewGetGroupIndividualsTooManyRequests() *GetGroupIndividualsTooManyRequests {
	return &GetGroupIndividualsTooManyRequests{}
}

/*GetGroupIndividualsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetGroupIndividualsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetGroupIndividualsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsInternalServerError creates a GetGroupIndividualsInternalServerError with default headers values
func NewGetGroupIndividualsInternalServerError() *GetGroupIndividualsInternalServerError {
	return &GetGroupIndividualsInternalServerError{}
}

/*GetGroupIndividualsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetGroupIndividualsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGroupIndividualsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsServiceUnavailable creates a GetGroupIndividualsServiceUnavailable with default headers values
func NewGetGroupIndividualsServiceUnavailable() *GetGroupIndividualsServiceUnavailable {
	return &GetGroupIndividualsServiceUnavailable{}
}

/*GetGroupIndividualsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetGroupIndividualsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetGroupIndividualsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGroupIndividualsGatewayTimeout creates a GetGroupIndividualsGatewayTimeout with default headers values
func NewGetGroupIndividualsGatewayTimeout() *GetGroupIndividualsGatewayTimeout {
	return &GetGroupIndividualsGatewayTimeout{}
}

/*GetGroupIndividualsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetGroupIndividualsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetGroupIndividualsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/groups/{groupId}/individuals][%d] getGroupIndividualsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetGroupIndividualsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetGroupIndividualsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
