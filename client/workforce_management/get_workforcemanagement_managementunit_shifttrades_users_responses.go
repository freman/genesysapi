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

// GetWorkforcemanagementManagementunitShifttradesUsersReader is a Reader for the GetWorkforcemanagementManagementunitShifttradesUsers structure.
type GetWorkforcemanagementManagementunitShifttradesUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWorkforcemanagementManagementunitShifttradesUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersOK creates a GetWorkforcemanagementManagementunitShifttradesUsersOK with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersOK() *GetWorkforcemanagementManagementunitShifttradesUsersOK {
	return &GetWorkforcemanagementManagementunitShifttradesUsersOK{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersOK handles this case with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitShifttradesUsersOK struct {
	Payload *models.WfmUserEntityListing
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersOK) GetPayload() *models.WfmUserEntityListing {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WfmUserEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersBadRequest creates a GetWorkforcemanagementManagementunitShifttradesUsersBadRequest with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersBadRequest() *GetWorkforcemanagementManagementunitShifttradesUsersBadRequest {
	return &GetWorkforcemanagementManagementunitShifttradesUsersBadRequest{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitShifttradesUsersBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersUnauthorized creates a GetWorkforcemanagementManagementunitShifttradesUsersUnauthorized with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersUnauthorized() *GetWorkforcemanagementManagementunitShifttradesUsersUnauthorized {
	return &GetWorkforcemanagementManagementunitShifttradesUsersUnauthorized{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitShifttradesUsersUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersForbidden creates a GetWorkforcemanagementManagementunitShifttradesUsersForbidden with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersForbidden() *GetWorkforcemanagementManagementunitShifttradesUsersForbidden {
	return &GetWorkforcemanagementManagementunitShifttradesUsersForbidden{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitShifttradesUsersForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersNotFound creates a GetWorkforcemanagementManagementunitShifttradesUsersNotFound with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersNotFound() *GetWorkforcemanagementManagementunitShifttradesUsersNotFound {
	return &GetWorkforcemanagementManagementunitShifttradesUsersNotFound{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitShifttradesUsersNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge creates a GetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge() *GetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge {
	return &GetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType creates a GetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType() *GetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType {
	return &GetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests creates a GetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests() *GetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests {
	return &GetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersInternalServerError creates a GetWorkforcemanagementManagementunitShifttradesUsersInternalServerError with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersInternalServerError() *GetWorkforcemanagementManagementunitShifttradesUsersInternalServerError {
	return &GetWorkforcemanagementManagementunitShifttradesUsersInternalServerError{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitShifttradesUsersInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable creates a GetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable() *GetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable {
	return &GetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout creates a GetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout with default headers values
func NewGetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout() *GetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout {
	return &GetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout{}
}

/*GetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{muId}/shifttrades/users][%d] getWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitShifttradesUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
