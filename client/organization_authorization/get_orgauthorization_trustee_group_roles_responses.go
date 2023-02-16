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

// GetOrgauthorizationTrusteeGroupRolesReader is a Reader for the GetOrgauthorizationTrusteeGroupRoles structure.
type GetOrgauthorizationTrusteeGroupRolesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgauthorizationTrusteeGroupRolesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgauthorizationTrusteeGroupRolesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgauthorizationTrusteeGroupRolesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrgauthorizationTrusteeGroupRolesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgauthorizationTrusteeGroupRolesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgauthorizationTrusteeGroupRolesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrgauthorizationTrusteeGroupRolesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrgauthorizationTrusteeGroupRolesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgauthorizationTrusteeGroupRolesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrgauthorizationTrusteeGroupRolesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrgauthorizationTrusteeGroupRolesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrgauthorizationTrusteeGroupRolesOK creates a GetOrgauthorizationTrusteeGroupRolesOK with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesOK() *GetOrgauthorizationTrusteeGroupRolesOK {
	return &GetOrgauthorizationTrusteeGroupRolesOK{}
}

/*
GetOrgauthorizationTrusteeGroupRolesOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOrgauthorizationTrusteeGroupRolesOK struct {
	Payload *models.UserAuthorization
}

// IsSuccess returns true when this get orgauthorization trustee group roles o k response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get orgauthorization trustee group roles o k response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles o k response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee group roles o k response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee group roles o k response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrgauthorizationTrusteeGroupRolesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesOK) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesOK) GetPayload() *models.UserAuthorization {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserAuthorization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesBadRequest creates a GetOrgauthorizationTrusteeGroupRolesBadRequest with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesBadRequest() *GetOrgauthorizationTrusteeGroupRolesBadRequest {
	return &GetOrgauthorizationTrusteeGroupRolesBadRequest{}
}

/*
GetOrgauthorizationTrusteeGroupRolesBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrgauthorizationTrusteeGroupRolesBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles bad request response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles bad request response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles bad request response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee group roles bad request response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee group roles bad request response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrgauthorizationTrusteeGroupRolesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesUnauthorized creates a GetOrgauthorizationTrusteeGroupRolesUnauthorized with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesUnauthorized() *GetOrgauthorizationTrusteeGroupRolesUnauthorized {
	return &GetOrgauthorizationTrusteeGroupRolesUnauthorized{}
}

/*
GetOrgauthorizationTrusteeGroupRolesUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrgauthorizationTrusteeGroupRolesUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles unauthorized response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles unauthorized response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles unauthorized response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee group roles unauthorized response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee group roles unauthorized response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOrgauthorizationTrusteeGroupRolesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesForbidden creates a GetOrgauthorizationTrusteeGroupRolesForbidden with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesForbidden() *GetOrgauthorizationTrusteeGroupRolesForbidden {
	return &GetOrgauthorizationTrusteeGroupRolesForbidden{}
}

/*
GetOrgauthorizationTrusteeGroupRolesForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOrgauthorizationTrusteeGroupRolesForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles forbidden response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles forbidden response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles forbidden response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee group roles forbidden response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee group roles forbidden response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrgauthorizationTrusteeGroupRolesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesNotFound creates a GetOrgauthorizationTrusteeGroupRolesNotFound with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesNotFound() *GetOrgauthorizationTrusteeGroupRolesNotFound {
	return &GetOrgauthorizationTrusteeGroupRolesNotFound{}
}

/*
GetOrgauthorizationTrusteeGroupRolesNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOrgauthorizationTrusteeGroupRolesNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles not found response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles not found response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles not found response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee group roles not found response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee group roles not found response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrgauthorizationTrusteeGroupRolesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesRequestTimeout creates a GetOrgauthorizationTrusteeGroupRolesRequestTimeout with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesRequestTimeout() *GetOrgauthorizationTrusteeGroupRolesRequestTimeout {
	return &GetOrgauthorizationTrusteeGroupRolesRequestTimeout{}
}

/*
GetOrgauthorizationTrusteeGroupRolesRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOrgauthorizationTrusteeGroupRolesRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles request timeout response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles request timeout response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles request timeout response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee group roles request timeout response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee group roles request timeout response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOrgauthorizationTrusteeGroupRolesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge creates a GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge() *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge {
	return &GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge{}
}

/*
GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles request entity too large response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles request entity too large response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles request entity too large response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee group roles request entity too large response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee group roles request entity too large response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType creates a GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType() *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType {
	return &GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType{}
}

/*
GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles unsupported media type response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles unsupported media type response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles unsupported media type response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee group roles unsupported media type response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee group roles unsupported media type response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesTooManyRequests creates a GetOrgauthorizationTrusteeGroupRolesTooManyRequests with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesTooManyRequests() *GetOrgauthorizationTrusteeGroupRolesTooManyRequests {
	return &GetOrgauthorizationTrusteeGroupRolesTooManyRequests{}
}

/*
GetOrgauthorizationTrusteeGroupRolesTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOrgauthorizationTrusteeGroupRolesTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles too many requests response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles too many requests response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles too many requests response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee group roles too many requests response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee group roles too many requests response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrgauthorizationTrusteeGroupRolesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesInternalServerError creates a GetOrgauthorizationTrusteeGroupRolesInternalServerError with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesInternalServerError() *GetOrgauthorizationTrusteeGroupRolesInternalServerError {
	return &GetOrgauthorizationTrusteeGroupRolesInternalServerError{}
}

/*
GetOrgauthorizationTrusteeGroupRolesInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrgauthorizationTrusteeGroupRolesInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles internal server error response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles internal server error response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles internal server error response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee group roles internal server error response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustee group roles internal server error response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrgauthorizationTrusteeGroupRolesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesServiceUnavailable creates a GetOrgauthorizationTrusteeGroupRolesServiceUnavailable with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesServiceUnavailable() *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable {
	return &GetOrgauthorizationTrusteeGroupRolesServiceUnavailable{}
}

/*
GetOrgauthorizationTrusteeGroupRolesServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrgauthorizationTrusteeGroupRolesServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles service unavailable response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles service unavailable response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles service unavailable response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee group roles service unavailable response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustee group roles service unavailable response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeGroupRolesGatewayTimeout creates a GetOrgauthorizationTrusteeGroupRolesGatewayTimeout with default headers values
func NewGetOrgauthorizationTrusteeGroupRolesGatewayTimeout() *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout {
	return &GetOrgauthorizationTrusteeGroupRolesGatewayTimeout{}
}

/*
GetOrgauthorizationTrusteeGroupRolesGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOrgauthorizationTrusteeGroupRolesGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee group roles gateway timeout response has a 2xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee group roles gateway timeout response has a 3xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee group roles gateway timeout response has a 4xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee group roles gateway timeout response has a 5xx status code
func (o *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustee group roles gateway timeout response a status code equal to that given
func (o *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups/{trusteeGroupId}/roles][%d] getOrgauthorizationTrusteeGroupRolesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeGroupRolesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
