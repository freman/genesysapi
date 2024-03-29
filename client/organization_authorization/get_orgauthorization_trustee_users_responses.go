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

// GetOrgauthorizationTrusteeUsersReader is a Reader for the GetOrgauthorizationTrusteeUsers structure.
type GetOrgauthorizationTrusteeUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgauthorizationTrusteeUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgauthorizationTrusteeUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgauthorizationTrusteeUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrgauthorizationTrusteeUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgauthorizationTrusteeUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgauthorizationTrusteeUsersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrgauthorizationTrusteeUsersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrgauthorizationTrusteeUsersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrgauthorizationTrusteeUsersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrgauthorizationTrusteeUsersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgauthorizationTrusteeUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrgauthorizationTrusteeUsersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrgauthorizationTrusteeUsersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrgauthorizationTrusteeUsersOK creates a GetOrgauthorizationTrusteeUsersOK with default headers values
func NewGetOrgauthorizationTrusteeUsersOK() *GetOrgauthorizationTrusteeUsersOK {
	return &GetOrgauthorizationTrusteeUsersOK{}
}

/*
GetOrgauthorizationTrusteeUsersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOrgauthorizationTrusteeUsersOK struct {
	Payload *models.TrustUserEntityListing
}

// IsSuccess returns true when this get orgauthorization trustee users o k response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get orgauthorization trustee users o k response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users o k response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee users o k response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee users o k response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrgauthorizationTrusteeUsersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersOK) GetPayload() *models.TrustUserEntityListing {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrustUserEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersBadRequest creates a GetOrgauthorizationTrusteeUsersBadRequest with default headers values
func NewGetOrgauthorizationTrusteeUsersBadRequest() *GetOrgauthorizationTrusteeUsersBadRequest {
	return &GetOrgauthorizationTrusteeUsersBadRequest{}
}

/*
GetOrgauthorizationTrusteeUsersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrgauthorizationTrusteeUsersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users bad request response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users bad request response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users bad request response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee users bad request response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee users bad request response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrgauthorizationTrusteeUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersUnauthorized creates a GetOrgauthorizationTrusteeUsersUnauthorized with default headers values
func NewGetOrgauthorizationTrusteeUsersUnauthorized() *GetOrgauthorizationTrusteeUsersUnauthorized {
	return &GetOrgauthorizationTrusteeUsersUnauthorized{}
}

/*
GetOrgauthorizationTrusteeUsersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrgauthorizationTrusteeUsersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users unauthorized response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users unauthorized response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users unauthorized response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee users unauthorized response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee users unauthorized response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOrgauthorizationTrusteeUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersForbidden creates a GetOrgauthorizationTrusteeUsersForbidden with default headers values
func NewGetOrgauthorizationTrusteeUsersForbidden() *GetOrgauthorizationTrusteeUsersForbidden {
	return &GetOrgauthorizationTrusteeUsersForbidden{}
}

/*
GetOrgauthorizationTrusteeUsersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOrgauthorizationTrusteeUsersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users forbidden response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users forbidden response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users forbidden response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee users forbidden response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee users forbidden response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrgauthorizationTrusteeUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersNotFound creates a GetOrgauthorizationTrusteeUsersNotFound with default headers values
func NewGetOrgauthorizationTrusteeUsersNotFound() *GetOrgauthorizationTrusteeUsersNotFound {
	return &GetOrgauthorizationTrusteeUsersNotFound{}
}

/*
GetOrgauthorizationTrusteeUsersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOrgauthorizationTrusteeUsersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users not found response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users not found response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users not found response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee users not found response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee users not found response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrgauthorizationTrusteeUsersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersRequestTimeout creates a GetOrgauthorizationTrusteeUsersRequestTimeout with default headers values
func NewGetOrgauthorizationTrusteeUsersRequestTimeout() *GetOrgauthorizationTrusteeUsersRequestTimeout {
	return &GetOrgauthorizationTrusteeUsersRequestTimeout{}
}

/*
GetOrgauthorizationTrusteeUsersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOrgauthorizationTrusteeUsersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users request timeout response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users request timeout response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users request timeout response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee users request timeout response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee users request timeout response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOrgauthorizationTrusteeUsersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersRequestEntityTooLarge creates a GetOrgauthorizationTrusteeUsersRequestEntityTooLarge with default headers values
func NewGetOrgauthorizationTrusteeUsersRequestEntityTooLarge() *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge {
	return &GetOrgauthorizationTrusteeUsersRequestEntityTooLarge{}
}

/*
GetOrgauthorizationTrusteeUsersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetOrgauthorizationTrusteeUsersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users request entity too large response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users request entity too large response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users request entity too large response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee users request entity too large response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee users request entity too large response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersUnsupportedMediaType creates a GetOrgauthorizationTrusteeUsersUnsupportedMediaType with default headers values
func NewGetOrgauthorizationTrusteeUsersUnsupportedMediaType() *GetOrgauthorizationTrusteeUsersUnsupportedMediaType {
	return &GetOrgauthorizationTrusteeUsersUnsupportedMediaType{}
}

/*
GetOrgauthorizationTrusteeUsersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrgauthorizationTrusteeUsersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users unsupported media type response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users unsupported media type response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users unsupported media type response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee users unsupported media type response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee users unsupported media type response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOrgauthorizationTrusteeUsersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersTooManyRequests creates a GetOrgauthorizationTrusteeUsersTooManyRequests with default headers values
func NewGetOrgauthorizationTrusteeUsersTooManyRequests() *GetOrgauthorizationTrusteeUsersTooManyRequests {
	return &GetOrgauthorizationTrusteeUsersTooManyRequests{}
}

/*
GetOrgauthorizationTrusteeUsersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOrgauthorizationTrusteeUsersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users too many requests response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users too many requests response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users too many requests response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustee users too many requests response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustee users too many requests response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrgauthorizationTrusteeUsersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersInternalServerError creates a GetOrgauthorizationTrusteeUsersInternalServerError with default headers values
func NewGetOrgauthorizationTrusteeUsersInternalServerError() *GetOrgauthorizationTrusteeUsersInternalServerError {
	return &GetOrgauthorizationTrusteeUsersInternalServerError{}
}

/*
GetOrgauthorizationTrusteeUsersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrgauthorizationTrusteeUsersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users internal server error response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users internal server error response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users internal server error response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee users internal server error response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustee users internal server error response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrgauthorizationTrusteeUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersServiceUnavailable creates a GetOrgauthorizationTrusteeUsersServiceUnavailable with default headers values
func NewGetOrgauthorizationTrusteeUsersServiceUnavailable() *GetOrgauthorizationTrusteeUsersServiceUnavailable {
	return &GetOrgauthorizationTrusteeUsersServiceUnavailable{}
}

/*
GetOrgauthorizationTrusteeUsersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrgauthorizationTrusteeUsersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users service unavailable response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users service unavailable response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users service unavailable response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee users service unavailable response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustee users service unavailable response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrgauthorizationTrusteeUsersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrusteeUsersGatewayTimeout creates a GetOrgauthorizationTrusteeUsersGatewayTimeout with default headers values
func NewGetOrgauthorizationTrusteeUsersGatewayTimeout() *GetOrgauthorizationTrusteeUsersGatewayTimeout {
	return &GetOrgauthorizationTrusteeUsersGatewayTimeout{}
}

/*
GetOrgauthorizationTrusteeUsersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOrgauthorizationTrusteeUsersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustee users gateway timeout response has a 2xx status code
func (o *GetOrgauthorizationTrusteeUsersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustee users gateway timeout response has a 3xx status code
func (o *GetOrgauthorizationTrusteeUsersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustee users gateway timeout response has a 4xx status code
func (o *GetOrgauthorizationTrusteeUsersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustee users gateway timeout response has a 5xx status code
func (o *GetOrgauthorizationTrusteeUsersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustee users gateway timeout response a status code equal to that given
func (o *GetOrgauthorizationTrusteeUsersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOrgauthorizationTrusteeUsersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustees/{trusteeOrgId}/users][%d] getOrgauthorizationTrusteeUsersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrusteeUsersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrusteeUsersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
