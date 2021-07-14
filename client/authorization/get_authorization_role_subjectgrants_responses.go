// Code generated by go-swagger; DO NOT EDIT.

package authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetAuthorizationRoleSubjectgrantsReader is a Reader for the GetAuthorizationRoleSubjectgrants structure.
type GetAuthorizationRoleSubjectgrantsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAuthorizationRoleSubjectgrantsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetAuthorizationRoleSubjectgrantsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetAuthorizationRoleSubjectgrantsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetAuthorizationRoleSubjectgrantsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetAuthorizationRoleSubjectgrantsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetAuthorizationRoleSubjectgrantsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetAuthorizationRoleSubjectgrantsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetAuthorizationRoleSubjectgrantsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetAuthorizationRoleSubjectgrantsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetAuthorizationRoleSubjectgrantsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetAuthorizationRoleSubjectgrantsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetAuthorizationRoleSubjectgrantsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetAuthorizationRoleSubjectgrantsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetAuthorizationRoleSubjectgrantsOK creates a GetAuthorizationRoleSubjectgrantsOK with default headers values
func NewGetAuthorizationRoleSubjectgrantsOK() *GetAuthorizationRoleSubjectgrantsOK {
	return &GetAuthorizationRoleSubjectgrantsOK{}
}

/*GetAuthorizationRoleSubjectgrantsOK handles this case with default header values.

successful operation
*/
type GetAuthorizationRoleSubjectgrantsOK struct {
	Payload *models.SubjectDivisionGrantsEntityListing
}

func (o *GetAuthorizationRoleSubjectgrantsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsOK  %+v", 200, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsOK) GetPayload() *models.SubjectDivisionGrantsEntityListing {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SubjectDivisionGrantsEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsBadRequest creates a GetAuthorizationRoleSubjectgrantsBadRequest with default headers values
func NewGetAuthorizationRoleSubjectgrantsBadRequest() *GetAuthorizationRoleSubjectgrantsBadRequest {
	return &GetAuthorizationRoleSubjectgrantsBadRequest{}
}

/*GetAuthorizationRoleSubjectgrantsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetAuthorizationRoleSubjectgrantsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsBadRequest  %+v", 400, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsUnauthorized creates a GetAuthorizationRoleSubjectgrantsUnauthorized with default headers values
func NewGetAuthorizationRoleSubjectgrantsUnauthorized() *GetAuthorizationRoleSubjectgrantsUnauthorized {
	return &GetAuthorizationRoleSubjectgrantsUnauthorized{}
}

/*GetAuthorizationRoleSubjectgrantsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetAuthorizationRoleSubjectgrantsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsForbidden creates a GetAuthorizationRoleSubjectgrantsForbidden with default headers values
func NewGetAuthorizationRoleSubjectgrantsForbidden() *GetAuthorizationRoleSubjectgrantsForbidden {
	return &GetAuthorizationRoleSubjectgrantsForbidden{}
}

/*GetAuthorizationRoleSubjectgrantsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetAuthorizationRoleSubjectgrantsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsForbidden  %+v", 403, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsNotFound creates a GetAuthorizationRoleSubjectgrantsNotFound with default headers values
func NewGetAuthorizationRoleSubjectgrantsNotFound() *GetAuthorizationRoleSubjectgrantsNotFound {
	return &GetAuthorizationRoleSubjectgrantsNotFound{}
}

/*GetAuthorizationRoleSubjectgrantsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetAuthorizationRoleSubjectgrantsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsNotFound  %+v", 404, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsRequestTimeout creates a GetAuthorizationRoleSubjectgrantsRequestTimeout with default headers values
func NewGetAuthorizationRoleSubjectgrantsRequestTimeout() *GetAuthorizationRoleSubjectgrantsRequestTimeout {
	return &GetAuthorizationRoleSubjectgrantsRequestTimeout{}
}

/*GetAuthorizationRoleSubjectgrantsRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetAuthorizationRoleSubjectgrantsRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsRequestEntityTooLarge creates a GetAuthorizationRoleSubjectgrantsRequestEntityTooLarge with default headers values
func NewGetAuthorizationRoleSubjectgrantsRequestEntityTooLarge() *GetAuthorizationRoleSubjectgrantsRequestEntityTooLarge {
	return &GetAuthorizationRoleSubjectgrantsRequestEntityTooLarge{}
}

/*GetAuthorizationRoleSubjectgrantsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetAuthorizationRoleSubjectgrantsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsUnsupportedMediaType creates a GetAuthorizationRoleSubjectgrantsUnsupportedMediaType with default headers values
func NewGetAuthorizationRoleSubjectgrantsUnsupportedMediaType() *GetAuthorizationRoleSubjectgrantsUnsupportedMediaType {
	return &GetAuthorizationRoleSubjectgrantsUnsupportedMediaType{}
}

/*GetAuthorizationRoleSubjectgrantsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetAuthorizationRoleSubjectgrantsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsTooManyRequests creates a GetAuthorizationRoleSubjectgrantsTooManyRequests with default headers values
func NewGetAuthorizationRoleSubjectgrantsTooManyRequests() *GetAuthorizationRoleSubjectgrantsTooManyRequests {
	return &GetAuthorizationRoleSubjectgrantsTooManyRequests{}
}

/*GetAuthorizationRoleSubjectgrantsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetAuthorizationRoleSubjectgrantsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsInternalServerError creates a GetAuthorizationRoleSubjectgrantsInternalServerError with default headers values
func NewGetAuthorizationRoleSubjectgrantsInternalServerError() *GetAuthorizationRoleSubjectgrantsInternalServerError {
	return &GetAuthorizationRoleSubjectgrantsInternalServerError{}
}

/*GetAuthorizationRoleSubjectgrantsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetAuthorizationRoleSubjectgrantsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsServiceUnavailable creates a GetAuthorizationRoleSubjectgrantsServiceUnavailable with default headers values
func NewGetAuthorizationRoleSubjectgrantsServiceUnavailable() *GetAuthorizationRoleSubjectgrantsServiceUnavailable {
	return &GetAuthorizationRoleSubjectgrantsServiceUnavailable{}
}

/*GetAuthorizationRoleSubjectgrantsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetAuthorizationRoleSubjectgrantsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAuthorizationRoleSubjectgrantsGatewayTimeout creates a GetAuthorizationRoleSubjectgrantsGatewayTimeout with default headers values
func NewGetAuthorizationRoleSubjectgrantsGatewayTimeout() *GetAuthorizationRoleSubjectgrantsGatewayTimeout {
	return &GetAuthorizationRoleSubjectgrantsGatewayTimeout{}
}

/*GetAuthorizationRoleSubjectgrantsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetAuthorizationRoleSubjectgrantsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetAuthorizationRoleSubjectgrantsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/authorization/roles/{roleId}/subjectgrants][%d] getAuthorizationRoleSubjectgrantsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetAuthorizationRoleSubjectgrantsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetAuthorizationRoleSubjectgrantsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
