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

// GetOrgauthorizationTrustorClonedusersReader is a Reader for the GetOrgauthorizationTrustorClonedusers structure.
type GetOrgauthorizationTrustorClonedusersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrgauthorizationTrustorClonedusersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrgauthorizationTrustorClonedusersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetOrgauthorizationTrustorClonedusersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetOrgauthorizationTrustorClonedusersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrgauthorizationTrustorClonedusersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrgauthorizationTrustorClonedusersNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetOrgauthorizationTrustorClonedusersRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetOrgauthorizationTrustorClonedusersRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetOrgauthorizationTrustorClonedusersUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetOrgauthorizationTrustorClonedusersTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrgauthorizationTrustorClonedusersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetOrgauthorizationTrustorClonedusersServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetOrgauthorizationTrustorClonedusersGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrgauthorizationTrustorClonedusersOK creates a GetOrgauthorizationTrustorClonedusersOK with default headers values
func NewGetOrgauthorizationTrustorClonedusersOK() *GetOrgauthorizationTrustorClonedusersOK {
	return &GetOrgauthorizationTrustorClonedusersOK{}
}

/*
GetOrgauthorizationTrustorClonedusersOK describes a response with status code 200, with default header values.

successful operation
*/
type GetOrgauthorizationTrustorClonedusersOK struct {
	Payload *models.ClonedUserEntityListing
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers o k response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers o k response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers o k response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustor clonedusers o k response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustor clonedusers o k response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrgauthorizationTrustorClonedusersOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersOK) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersOK  %+v", 200, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersOK) GetPayload() *models.ClonedUserEntityListing {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClonedUserEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersBadRequest creates a GetOrgauthorizationTrustorClonedusersBadRequest with default headers values
func NewGetOrgauthorizationTrustorClonedusersBadRequest() *GetOrgauthorizationTrustorClonedusersBadRequest {
	return &GetOrgauthorizationTrustorClonedusersBadRequest{}
}

/*
GetOrgauthorizationTrustorClonedusersBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetOrgauthorizationTrustorClonedusersBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers bad request response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers bad request response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers bad request response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustor clonedusers bad request response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustor clonedusers bad request response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetOrgauthorizationTrustorClonedusersBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersBadRequest  %+v", 400, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersUnauthorized creates a GetOrgauthorizationTrustorClonedusersUnauthorized with default headers values
func NewGetOrgauthorizationTrustorClonedusersUnauthorized() *GetOrgauthorizationTrustorClonedusersUnauthorized {
	return &GetOrgauthorizationTrustorClonedusersUnauthorized{}
}

/*
GetOrgauthorizationTrustorClonedusersUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetOrgauthorizationTrustorClonedusersUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers unauthorized response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers unauthorized response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers unauthorized response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustor clonedusers unauthorized response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustor clonedusers unauthorized response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetOrgauthorizationTrustorClonedusersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersForbidden creates a GetOrgauthorizationTrustorClonedusersForbidden with default headers values
func NewGetOrgauthorizationTrustorClonedusersForbidden() *GetOrgauthorizationTrustorClonedusersForbidden {
	return &GetOrgauthorizationTrustorClonedusersForbidden{}
}

/*
GetOrgauthorizationTrustorClonedusersForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetOrgauthorizationTrustorClonedusersForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers forbidden response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers forbidden response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers forbidden response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustor clonedusers forbidden response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustor clonedusers forbidden response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetOrgauthorizationTrustorClonedusersForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersForbidden  %+v", 403, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersNotFound creates a GetOrgauthorizationTrustorClonedusersNotFound with default headers values
func NewGetOrgauthorizationTrustorClonedusersNotFound() *GetOrgauthorizationTrustorClonedusersNotFound {
	return &GetOrgauthorizationTrustorClonedusersNotFound{}
}

/*
GetOrgauthorizationTrustorClonedusersNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetOrgauthorizationTrustorClonedusersNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers not found response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers not found response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers not found response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustor clonedusers not found response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustor clonedusers not found response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetOrgauthorizationTrustorClonedusersNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersNotFound  %+v", 404, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersRequestTimeout creates a GetOrgauthorizationTrustorClonedusersRequestTimeout with default headers values
func NewGetOrgauthorizationTrustorClonedusersRequestTimeout() *GetOrgauthorizationTrustorClonedusersRequestTimeout {
	return &GetOrgauthorizationTrustorClonedusersRequestTimeout{}
}

/*
GetOrgauthorizationTrustorClonedusersRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetOrgauthorizationTrustorClonedusersRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers request timeout response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers request timeout response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers request timeout response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustor clonedusers request timeout response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustor clonedusers request timeout response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetOrgauthorizationTrustorClonedusersRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersRequestEntityTooLarge creates a GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge with default headers values
func NewGetOrgauthorizationTrustorClonedusersRequestEntityTooLarge() *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge {
	return &GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge{}
}

/*
GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers request entity too large response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers request entity too large response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers request entity too large response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustor clonedusers request entity too large response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustor clonedusers request entity too large response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersUnsupportedMediaType creates a GetOrgauthorizationTrustorClonedusersUnsupportedMediaType with default headers values
func NewGetOrgauthorizationTrustorClonedusersUnsupportedMediaType() *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType {
	return &GetOrgauthorizationTrustorClonedusersUnsupportedMediaType{}
}

/*
GetOrgauthorizationTrustorClonedusersUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetOrgauthorizationTrustorClonedusersUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers unsupported media type response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers unsupported media type response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers unsupported media type response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustor clonedusers unsupported media type response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustor clonedusers unsupported media type response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersTooManyRequests creates a GetOrgauthorizationTrustorClonedusersTooManyRequests with default headers values
func NewGetOrgauthorizationTrustorClonedusersTooManyRequests() *GetOrgauthorizationTrustorClonedusersTooManyRequests {
	return &GetOrgauthorizationTrustorClonedusersTooManyRequests{}
}

/*
GetOrgauthorizationTrustorClonedusersTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetOrgauthorizationTrustorClonedusersTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers too many requests response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers too many requests response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers too many requests response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get orgauthorization trustor clonedusers too many requests response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get orgauthorization trustor clonedusers too many requests response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetOrgauthorizationTrustorClonedusersTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersInternalServerError creates a GetOrgauthorizationTrustorClonedusersInternalServerError with default headers values
func NewGetOrgauthorizationTrustorClonedusersInternalServerError() *GetOrgauthorizationTrustorClonedusersInternalServerError {
	return &GetOrgauthorizationTrustorClonedusersInternalServerError{}
}

/*
GetOrgauthorizationTrustorClonedusersInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetOrgauthorizationTrustorClonedusersInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers internal server error response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers internal server error response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers internal server error response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustor clonedusers internal server error response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustor clonedusers internal server error response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetOrgauthorizationTrustorClonedusersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersInternalServerError  %+v", 500, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersServiceUnavailable creates a GetOrgauthorizationTrustorClonedusersServiceUnavailable with default headers values
func NewGetOrgauthorizationTrustorClonedusersServiceUnavailable() *GetOrgauthorizationTrustorClonedusersServiceUnavailable {
	return &GetOrgauthorizationTrustorClonedusersServiceUnavailable{}
}

/*
GetOrgauthorizationTrustorClonedusersServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetOrgauthorizationTrustorClonedusersServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers service unavailable response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers service unavailable response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers service unavailable response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustor clonedusers service unavailable response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustor clonedusers service unavailable response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetOrgauthorizationTrustorClonedusersServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrgauthorizationTrustorClonedusersGatewayTimeout creates a GetOrgauthorizationTrustorClonedusersGatewayTimeout with default headers values
func NewGetOrgauthorizationTrustorClonedusersGatewayTimeout() *GetOrgauthorizationTrustorClonedusersGatewayTimeout {
	return &GetOrgauthorizationTrustorClonedusersGatewayTimeout{}
}

/*
GetOrgauthorizationTrustorClonedusersGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetOrgauthorizationTrustorClonedusersGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get orgauthorization trustor clonedusers gateway timeout response has a 2xx status code
func (o *GetOrgauthorizationTrustorClonedusersGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get orgauthorization trustor clonedusers gateway timeout response has a 3xx status code
func (o *GetOrgauthorizationTrustorClonedusersGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get orgauthorization trustor clonedusers gateway timeout response has a 4xx status code
func (o *GetOrgauthorizationTrustorClonedusersGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get orgauthorization trustor clonedusers gateway timeout response has a 5xx status code
func (o *GetOrgauthorizationTrustorClonedusersGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get orgauthorization trustor clonedusers gateway timeout response a status code equal to that given
func (o *GetOrgauthorizationTrustorClonedusersGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetOrgauthorizationTrustorClonedusersGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/orgauthorization/trustors/{trustorOrgId}/clonedusers][%d] getOrgauthorizationTrustorClonedusersGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetOrgauthorizationTrustorClonedusersGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetOrgauthorizationTrustorClonedusersGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
