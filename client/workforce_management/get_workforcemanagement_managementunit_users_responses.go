// Code generated by go-swagger; DO NOT EDIT.

package workforce_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetWorkforcemanagementManagementunitUsersReader is a Reader for the GetWorkforcemanagementManagementunitUsers structure.
type GetWorkforcemanagementManagementunitUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitUsersOK creates a GetWorkforcemanagementManagementunitUsersOK with default headers values
func NewGetWorkforcemanagementManagementunitUsersOK() *GetWorkforcemanagementManagementunitUsersOK {
	return &GetWorkforcemanagementManagementunitUsersOK{}
}

/*GetWorkforcemanagementManagementunitUsersOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitUsersOK struct {
	Payload *models.WfmUserEntityListing
}

func (o *GetWorkforcemanagementManagementunitUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersOK) GetPayload() *models.WfmUserEntityListing {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmUserEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersBadRequest creates a GetWorkforcemanagementManagementunitUsersBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitUsersBadRequest() *GetWorkforcemanagementManagementunitUsersBadRequest {
	return &GetWorkforcemanagementManagementunitUsersBadRequest{}
}

/*GetWorkforcemanagementManagementunitUsersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitUsersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersUnauthorized creates a GetWorkforcemanagementManagementunitUsersUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitUsersUnauthorized() *GetWorkforcemanagementManagementunitUsersUnauthorized {
	return &GetWorkforcemanagementManagementunitUsersUnauthorized{}
}

/*GetWorkforcemanagementManagementunitUsersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitUsersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersForbidden creates a GetWorkforcemanagementManagementunitUsersForbidden with default headers values
func NewGetWorkforcemanagementManagementunitUsersForbidden() *GetWorkforcemanagementManagementunitUsersForbidden {
	return &GetWorkforcemanagementManagementunitUsersForbidden{}
}

/*GetWorkforcemanagementManagementunitUsersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitUsersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersNotFound creates a GetWorkforcemanagementManagementunitUsersNotFound with default headers values
func NewGetWorkforcemanagementManagementunitUsersNotFound() *GetWorkforcemanagementManagementunitUsersNotFound {
	return &GetWorkforcemanagementManagementunitUsersNotFound{}
}

/*GetWorkforcemanagementManagementunitUsersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitUsersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitUsersRequestEntityTooLarge() *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersUnsupportedMediaType creates a GetWorkforcemanagementManagementunitUsersUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitUsersUnsupportedMediaType() *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitUsersUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitUsersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersTooManyRequests creates a GetWorkforcemanagementManagementunitUsersTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitUsersTooManyRequests() *GetWorkforcemanagementManagementunitUsersTooManyRequests {
	return &GetWorkforcemanagementManagementunitUsersTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitUsersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementManagementunitUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersInternalServerError creates a GetWorkforcemanagementManagementunitUsersInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitUsersInternalServerError() *GetWorkforcemanagementManagementunitUsersInternalServerError {
	return &GetWorkforcemanagementManagementunitUsersInternalServerError{}
}

/*GetWorkforcemanagementManagementunitUsersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitUsersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersServiceUnavailable creates a GetWorkforcemanagementManagementunitUsersServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitUsersServiceUnavailable() *GetWorkforcemanagementManagementunitUsersServiceUnavailable {
	return &GetWorkforcemanagementManagementunitUsersServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitUsersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitUsersGatewayTimeout creates a GetWorkforcemanagementManagementunitUsersGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitUsersGatewayTimeout() *GetWorkforcemanagementManagementunitUsersGatewayTimeout {
	return &GetWorkforcemanagementManagementunitUsersGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitUsersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
