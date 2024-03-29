// Code generated by go-swagger; DO NOT EDIT.

package organization_authorization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetOrgauthorizationTrusteeUserRolesReader is a Reader for the GetOrgauthorizationTrusteeUserRoles structure.
type GetOrgauthorizationTrusteeUserRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgauthorizationTrusteeUserRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgauthorizationTrusteeUserRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgauthorizationTrusteeUserRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrgauthorizationTrusteeUserRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgauthorizationTrusteeUserRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgauthorizationTrusteeUserRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrgauthorizationTrusteeUserRolesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrgauthorizationTrusteeUserRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrgauthorizationTrusteeUserRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgauthorizationTrusteeUserRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrgauthorizationTrusteeUserRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrgauthorizationTrusteeUserRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrgauthorizationTrusteeUserRolesOK creates a GetOrgauthorizationTrusteeUserRolesOK with default headers values
func NewGetOrgauthorizationTrusteeUserRolesOK() *GetOrgauthorizationTrusteeUserRolesOK {
	return &GetOrgauthorizationTrusteeUserRolesOK{}
}

/*
GetOrgauthorizationTrusteeUserRolesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOrgauthorizationTrusteeUserRolesOK struct {
	Payload *models.UserAuthorization
}

// IsSuccess returns true when this get orgauthorization trustee user roles o k response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get orgauthorization trustee user roles o k response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles o k response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee user roles o k response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee user roles o k response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrgauthorizationTrusteeUserRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesOK) GetPayload() *models.UserAuthorization {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAuthorization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesBadRequest creates a GetOrgauthorizationTrusteeUserRolesBadRequest with default headers values
func NewGetOrgauthorizationTrusteeUserRolesBadRequest() *GetOrgauthorizationTrusteeUserRolesBadRequest {
	return &GetOrgauthorizationTrusteeUserRolesBadRequest{}
}

/*
GetOrgauthorizationTrusteeUserRolesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrgauthorizationTrusteeUserRolesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles bad request response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles bad request response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles bad request response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee user roles bad request response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee user roles bad request response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrgauthorizationTrusteeUserRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesUnauthorized creates a GetOrgauthorizationTrusteeUserRolesUnauthorized with default headers values
func NewGetOrgauthorizationTrusteeUserRolesUnauthorized() *GetOrgauthorizationTrusteeUserRolesUnauthorized {
	return &GetOrgauthorizationTrusteeUserRolesUnauthorized{}
}

/*
GetOrgauthorizationTrusteeUserRolesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrgauthorizationTrusteeUserRolesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles unauthorized response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles unauthorized response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles unauthorized response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee user roles unauthorized response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee user roles unauthorized response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOrgauthorizationTrusteeUserRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesForbidden creates a GetOrgauthorizationTrusteeUserRolesForbidden with default headers values
func NewGetOrgauthorizationTrusteeUserRolesForbidden() *GetOrgauthorizationTrusteeUserRolesForbidden {
	return &GetOrgauthorizationTrusteeUserRolesForbidden{}
}

/*
GetOrgauthorizationTrusteeUserRolesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOrgauthorizationTrusteeUserRolesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles forbidden response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles forbidden response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles forbidden response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee user roles forbidden response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee user roles forbidden response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrgauthorizationTrusteeUserRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesNotFound creates a GetOrgauthorizationTrusteeUserRolesNotFound with default headers values
func NewGetOrgauthorizationTrusteeUserRolesNotFound() *GetOrgauthorizationTrusteeUserRolesNotFound {
	return &GetOrgauthorizationTrusteeUserRolesNotFound{}
}

/*
GetOrgauthorizationTrusteeUserRolesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOrgauthorizationTrusteeUserRolesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles not found response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles not found response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles not found response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee user roles not found response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee user roles not found response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrgauthorizationTrusteeUserRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesRequestTimeout creates a GetOrgauthorizationTrusteeUserRolesRequestTimeout with default headers values
func NewGetOrgauthorizationTrusteeUserRolesRequestTimeout() *GetOrgauthorizationTrusteeUserRolesRequestTimeout {
	return &GetOrgauthorizationTrusteeUserRolesRequestTimeout{}
}

/*
GetOrgauthorizationTrusteeUserRolesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOrgauthorizationTrusteeUserRolesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles request timeout response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles request timeout response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles request timeout response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee user roles request timeout response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee user roles request timeout response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOrgauthorizationTrusteeUserRolesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge creates a GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge with default headers values
func NewGetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge() *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge {
	return &GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge{}
}

/*
GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles request entity too large response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles request entity too large response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles request entity too large response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee user roles request entity too large response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee user roles request entity too large response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesUnsupportedMediaType creates a GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType with default headers values
func NewGetOrgauthorizationTrusteeUserRolesUnsupportedMediaType() *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType {
	return &GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType{}
}

/*
GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles unsupported media type response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles unsupported media type response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles unsupported media type response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee user roles unsupported media type response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee user roles unsupported media type response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesTooManyRequests creates a GetOrgauthorizationTrusteeUserRolesTooManyRequests with default headers values
func NewGetOrgauthorizationTrusteeUserRolesTooManyRequests() *GetOrgauthorizationTrusteeUserRolesTooManyRequests {
	return &GetOrgauthorizationTrusteeUserRolesTooManyRequests{}
}

/*
GetOrgauthorizationTrusteeUserRolesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOrgauthorizationTrusteeUserRolesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles too many requests response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles too many requests response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles too many requests response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee user roles too many requests response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee user roles too many requests response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrgauthorizationTrusteeUserRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesInternalServerError creates a GetOrgauthorizationTrusteeUserRolesInternalServerError with default headers values
func NewGetOrgauthorizationTrusteeUserRolesInternalServerError() *GetOrgauthorizationTrusteeUserRolesInternalServerError {
	return &GetOrgauthorizationTrusteeUserRolesInternalServerError{}
}

/*
GetOrgauthorizationTrusteeUserRolesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrgauthorizationTrusteeUserRolesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles internal server error response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles internal server error response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles internal server error response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee user roles internal server error response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustee user roles internal server error response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrgauthorizationTrusteeUserRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesServiceUnavailable creates a GetOrgauthorizationTrusteeUserRolesServiceUnavailable with default headers values
func NewGetOrgauthorizationTrusteeUserRolesServiceUnavailable() *GetOrgauthorizationTrusteeUserRolesServiceUnavailable {
	return &GetOrgauthorizationTrusteeUserRolesServiceUnavailable{}
}

/*
GetOrgauthorizationTrusteeUserRolesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrgauthorizationTrusteeUserRolesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles service unavailable response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles service unavailable response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles service unavailable response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee user roles service unavailable response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustee user roles service unavailable response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrgauthorizationTrusteeUserRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUserRolesGatewayTimeout creates a GetOrgauthorizationTrusteeUserRolesGatewayTimeout with default headers values
func NewGetOrgauthorizationTrusteeUserRolesGatewayTimeout() *GetOrgauthorizationTrusteeUserRolesGatewayTimeout {
	return &GetOrgauthorizationTrusteeUserRolesGatewayTimeout{}
}

/*
GetOrgauthorizationTrusteeUserRolesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOrgauthorizationTrusteeUserRolesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee user roles gateway timeout response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUserRolesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee user roles gateway timeout response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUserRolesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee user roles gateway timeout response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUserRolesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee user roles gateway timeout response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUserRolesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustee user roles gateway timeout response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUserRolesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOrgauthorizationTrusteeUserRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users/{trusteeUserId}/roles][%d] getOrgauthorizationTrusteeUserRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUserRolesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUserRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
