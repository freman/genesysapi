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
	case 408:
		result := NewGetWorkforcemanagementManagementunitUsersRequestTimeout()
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

/*
GetWorkforcemanagementManagementunitUsersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetWorkforcemanagementManagementunitUsersOK struct {
	Payload *models.WfmUserEntityListing
}

// IsSuccess returns true when this get workforcemanagement managementunit users o k response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get workforcemanagement managementunit users o k response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users o k response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit users o k response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit users o k response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetWorkforcemanagementManagementunitUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersOK  %+v", 200, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersOK) String() string {
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

/*
GetWorkforcemanagementManagementunitUsersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetWorkforcemanagementManagementunitUsersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users bad request response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users bad request response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users bad request response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit users bad request response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit users bad request response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetWorkforcemanagementManagementunitUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersBadRequest) String() string {
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

/*
GetWorkforcemanagementManagementunitUsersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetWorkforcemanagementManagementunitUsersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users unauthorized response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users unauthorized response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users unauthorized response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit users unauthorized response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit users unauthorized response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersUnauthorized) String() string {
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

/*
GetWorkforcemanagementManagementunitUsersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetWorkforcemanagementManagementunitUsersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users forbidden response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users forbidden response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users forbidden response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit users forbidden response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit users forbidden response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetWorkforcemanagementManagementunitUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersForbidden) String() string {
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

/*
GetWorkforcemanagementManagementunitUsersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetWorkforcemanagementManagementunitUsersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users not found response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users not found response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users not found response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit users not found response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit users not found response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetWorkforcemanagementManagementunitUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersNotFound) String() string {
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

// NewGetWorkforcemanagementManagementunitUsersRequestTimeout creates a GetWorkforcemanagementManagementunitUsersRequestTimeout with default headers values
func NewGetWorkforcemanagementManagementunitUsersRequestTimeout() *GetWorkforcemanagementManagementunitUsersRequestTimeout {
	return &GetWorkforcemanagementManagementunitUsersRequestTimeout{}
}

/*
GetWorkforcemanagementManagementunitUsersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetWorkforcemanagementManagementunitUsersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users request timeout response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users request timeout response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users request timeout response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit users request timeout response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit users request timeout response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetWorkforcemanagementManagementunitUsersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetWorkforcemanagementManagementunitUsersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users request entity too large response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users request entity too large response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users request entity too large response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit users request entity too large response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit users request entity too large response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersRequestEntityTooLarge) String() string {
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

/*
GetWorkforcemanagementManagementunitUsersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetWorkforcemanagementManagementunitUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users unsupported media type response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users unsupported media type response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users unsupported media type response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit users unsupported media type response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit users unsupported media type response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersUnsupportedMediaType) String() string {
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

/*
GetWorkforcemanagementManagementunitUsersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetWorkforcemanagementManagementunitUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users too many requests response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users too many requests response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users too many requests response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get workforcemanagement managementunit users too many requests response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get workforcemanagement managementunit users too many requests response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersTooManyRequests) String() string {
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

/*
GetWorkforcemanagementManagementunitUsersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetWorkforcemanagementManagementunitUsersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users internal server error response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users internal server error response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users internal server error response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit users internal server error response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement managementunit users internal server error response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersInternalServerError) String() string {
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

/*
GetWorkforcemanagementManagementunitUsersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetWorkforcemanagementManagementunitUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users service unavailable response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users service unavailable response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users service unavailable response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit users service unavailable response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement managementunit users service unavailable response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersServiceUnavailable) String() string {
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

/*
GetWorkforcemanagementManagementunitUsersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetWorkforcemanagementManagementunitUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get workforcemanagement managementunit users gateway timeout response has a 2xx status code
func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get workforcemanagement managementunit users gateway timeout response has a 3xx status code
func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get workforcemanagement managementunit users gateway timeout response has a 4xx status code
func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get workforcemanagement managementunit users gateway timeout response has a 5xx status code
func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get workforcemanagement managementunit users gateway timeout response a status code equal to that given
func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/workforcemanagement/managementunits/{managementUnitId}/users][%d] getWorkforcemanagementManagementunitUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetWorkforcemanagementManagementunitUsersGatewayTimeout) String() string {
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
