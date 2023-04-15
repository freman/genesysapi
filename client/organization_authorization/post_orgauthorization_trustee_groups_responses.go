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

// PostOrgauthorizationTrusteeGroupsReader is a Reader for the PostOrgauthorizationTrusteeGroups structure.
type PostOrgauthorizationTrusteeGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostOrgauthorizationTrusteeGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostOrgauthorizationTrusteeGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostOrgauthorizationTrusteeGroupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostOrgauthorizationTrusteeGroupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostOrgauthorizationTrusteeGroupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostOrgauthorizationTrusteeGroupsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPostOrgauthorizationTrusteeGroupsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostOrgauthorizationTrusteeGroupsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostOrgauthorizationTrusteeGroupsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostOrgauthorizationTrusteeGroupsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostOrgauthorizationTrusteeGroupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostOrgauthorizationTrusteeGroupsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostOrgauthorizationTrusteeGroupsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostOrgauthorizationTrusteeGroupsOK creates a PostOrgauthorizationTrusteeGroupsOK with default headers values
func NewPostOrgauthorizationTrusteeGroupsOK() *PostOrgauthorizationTrusteeGroupsOK {
	return &PostOrgauthorizationTrusteeGroupsOK{}
}

/*
PostOrgauthorizationTrusteeGroupsOK describes a response with status code 200, with default header values.

successful operation
*/
type PostOrgauthorizationTrusteeGroupsOK struct {
	Payload *models.TrustGroup
}

// IsSuccess returns true when this post orgauthorization trustee groups o k response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this post orgauthorization trustee groups o k response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups o k response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this post orgauthorization trustee groups o k response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee groups o k response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsOK) IsCode(code int) bool {
	return code == 200
}

func (o *PostOrgauthorizationTrusteeGroupsOK) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsOK  %+v", 200, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsOK) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsOK  %+v", 200, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsOK) GetPayload() *models.TrustGroup {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TrustGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsBadRequest creates a PostOrgauthorizationTrusteeGroupsBadRequest with default headers values
func NewPostOrgauthorizationTrusteeGroupsBadRequest() *PostOrgauthorizationTrusteeGroupsBadRequest {
	return &PostOrgauthorizationTrusteeGroupsBadRequest{}
}

/*
PostOrgauthorizationTrusteeGroupsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostOrgauthorizationTrusteeGroupsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups bad request response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups bad request response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups bad request response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee groups bad request response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee groups bad request response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PostOrgauthorizationTrusteeGroupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsBadRequest) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsUnauthorized creates a PostOrgauthorizationTrusteeGroupsUnauthorized with default headers values
func NewPostOrgauthorizationTrusteeGroupsUnauthorized() *PostOrgauthorizationTrusteeGroupsUnauthorized {
	return &PostOrgauthorizationTrusteeGroupsUnauthorized{}
}

/*
PostOrgauthorizationTrusteeGroupsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type PostOrgauthorizationTrusteeGroupsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups unauthorized response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups unauthorized response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups unauthorized response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee groups unauthorized response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee groups unauthorized response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PostOrgauthorizationTrusteeGroupsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsUnauthorized) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsForbidden creates a PostOrgauthorizationTrusteeGroupsForbidden with default headers values
func NewPostOrgauthorizationTrusteeGroupsForbidden() *PostOrgauthorizationTrusteeGroupsForbidden {
	return &PostOrgauthorizationTrusteeGroupsForbidden{}
}

/*
PostOrgauthorizationTrusteeGroupsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type PostOrgauthorizationTrusteeGroupsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups forbidden response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups forbidden response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups forbidden response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee groups forbidden response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee groups forbidden response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PostOrgauthorizationTrusteeGroupsForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsForbidden  %+v", 403, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsForbidden) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsForbidden  %+v", 403, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsNotFound creates a PostOrgauthorizationTrusteeGroupsNotFound with default headers values
func NewPostOrgauthorizationTrusteeGroupsNotFound() *PostOrgauthorizationTrusteeGroupsNotFound {
	return &PostOrgauthorizationTrusteeGroupsNotFound{}
}

/*
PostOrgauthorizationTrusteeGroupsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type PostOrgauthorizationTrusteeGroupsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups not found response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups not found response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups not found response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee groups not found response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee groups not found response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PostOrgauthorizationTrusteeGroupsNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsNotFound  %+v", 404, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsNotFound) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsNotFound  %+v", 404, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsRequestTimeout creates a PostOrgauthorizationTrusteeGroupsRequestTimeout with default headers values
func NewPostOrgauthorizationTrusteeGroupsRequestTimeout() *PostOrgauthorizationTrusteeGroupsRequestTimeout {
	return &PostOrgauthorizationTrusteeGroupsRequestTimeout{}
}

/*
PostOrgauthorizationTrusteeGroupsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PostOrgauthorizationTrusteeGroupsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups request timeout response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups request timeout response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups request timeout response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee groups request timeout response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee groups request timeout response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *PostOrgauthorizationTrusteeGroupsRequestTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsRequestTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsRequestTimeout  %+v", 408, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsRequestEntityTooLarge creates a PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge with default headers values
func NewPostOrgauthorizationTrusteeGroupsRequestEntityTooLarge() *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge {
	return &PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge{}
}

/*
PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups request entity too large response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups request entity too large response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups request entity too large response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee groups request entity too large response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee groups request entity too large response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsUnsupportedMediaType creates a PostOrgauthorizationTrusteeGroupsUnsupportedMediaType with default headers values
func NewPostOrgauthorizationTrusteeGroupsUnsupportedMediaType() *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType {
	return &PostOrgauthorizationTrusteeGroupsUnsupportedMediaType{}
}

/*
PostOrgauthorizationTrusteeGroupsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostOrgauthorizationTrusteeGroupsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups unsupported media type response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups unsupported media type response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups unsupported media type response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee groups unsupported media type response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee groups unsupported media type response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsTooManyRequests creates a PostOrgauthorizationTrusteeGroupsTooManyRequests with default headers values
func NewPostOrgauthorizationTrusteeGroupsTooManyRequests() *PostOrgauthorizationTrusteeGroupsTooManyRequests {
	return &PostOrgauthorizationTrusteeGroupsTooManyRequests{}
}

/*
PostOrgauthorizationTrusteeGroupsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PostOrgauthorizationTrusteeGroupsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups too many requests response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups too many requests response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups too many requests response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this post orgauthorization trustee groups too many requests response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this post orgauthorization trustee groups too many requests response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *PostOrgauthorizationTrusteeGroupsTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsTooManyRequests) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsInternalServerError creates a PostOrgauthorizationTrusteeGroupsInternalServerError with default headers values
func NewPostOrgauthorizationTrusteeGroupsInternalServerError() *PostOrgauthorizationTrusteeGroupsInternalServerError {
	return &PostOrgauthorizationTrusteeGroupsInternalServerError{}
}

/*
PostOrgauthorizationTrusteeGroupsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostOrgauthorizationTrusteeGroupsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups internal server error response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups internal server error response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups internal server error response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this post orgauthorization trustee groups internal server error response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this post orgauthorization trustee groups internal server error response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PostOrgauthorizationTrusteeGroupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsInternalServerError) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsServiceUnavailable creates a PostOrgauthorizationTrusteeGroupsServiceUnavailable with default headers values
func NewPostOrgauthorizationTrusteeGroupsServiceUnavailable() *PostOrgauthorizationTrusteeGroupsServiceUnavailable {
	return &PostOrgauthorizationTrusteeGroupsServiceUnavailable{}
}

/*
PostOrgauthorizationTrusteeGroupsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostOrgauthorizationTrusteeGroupsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups service unavailable response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups service unavailable response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups service unavailable response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this post orgauthorization trustee groups service unavailable response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this post orgauthorization trustee groups service unavailable response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *PostOrgauthorizationTrusteeGroupsServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsServiceUnavailable) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostOrgauthorizationTrusteeGroupsGatewayTimeout creates a PostOrgauthorizationTrusteeGroupsGatewayTimeout with default headers values
func NewPostOrgauthorizationTrusteeGroupsGatewayTimeout() *PostOrgauthorizationTrusteeGroupsGatewayTimeout {
	return &PostOrgauthorizationTrusteeGroupsGatewayTimeout{}
}

/*
PostOrgauthorizationTrusteeGroupsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type PostOrgauthorizationTrusteeGroupsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this post orgauthorization trustee groups gateway timeout response has a 2xx status code
func (o *PostOrgauthorizationTrusteeGroupsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this post orgauthorization trustee groups gateway timeout response has a 3xx status code
func (o *PostOrgauthorizationTrusteeGroupsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this post orgauthorization trustee groups gateway timeout response has a 4xx status code
func (o *PostOrgauthorizationTrusteeGroupsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this post orgauthorization trustee groups gateway timeout response has a 5xx status code
func (o *PostOrgauthorizationTrusteeGroupsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this post orgauthorization trustee groups gateway timeout response a status code equal to that given
func (o *PostOrgauthorizationTrusteeGroupsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *PostOrgauthorizationTrusteeGroupsGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsGatewayTimeout) String() string {
	return fmt.Sprintf("[POST /api/v2/orgauthorization/trustees/{trusteeOrgId}/groups][%d] postOrgauthorizationTrusteeGroupsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostOrgauthorizationTrusteeGroupsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostOrgauthorizationTrusteeGroupsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}