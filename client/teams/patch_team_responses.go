// Code generated by go-swagger; DO NOT EDIT.

package teams

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchTeamReader is a Reader for the PatchTeam structure.
type PatchTeamReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchTeamReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchTeamOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchTeamBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchTeamUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchTeamForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchTeamNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchTeamRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchTeamRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchTeamUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchTeamTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchTeamInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchTeamServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchTeamGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchTeamOK creates a PatchTeamOK with default headers values
func NewPatchTeamOK() *PatchTeamOK {
	return &PatchTeamOK{}
}

/*PatchTeamOK handles this case with default header values.

successful operation
*/
type PatchTeamOK struct {
	Payload *models.Team
}

func (o *PatchTeamOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamOK  %+v", 200, o.Payload)
}

func (o *PatchTeamOK) GetPayload() *models.Team {
	return o.Payload
}

func (o *PatchTeamOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Team)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamBadRequest creates a PatchTeamBadRequest with default headers values
func NewPatchTeamBadRequest() *PatchTeamBadRequest {
	return &PatchTeamBadRequest{}
}

/*PatchTeamBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchTeamBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamBadRequest  %+v", 400, o.Payload)
}

func (o *PatchTeamBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamUnauthorized creates a PatchTeamUnauthorized with default headers values
func NewPatchTeamUnauthorized() *PatchTeamUnauthorized {
	return &PatchTeamUnauthorized{}
}

/*PatchTeamUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchTeamUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchTeamUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamForbidden creates a PatchTeamForbidden with default headers values
func NewPatchTeamForbidden() *PatchTeamForbidden {
	return &PatchTeamForbidden{}
}

/*PatchTeamForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchTeamForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamForbidden  %+v", 403, o.Payload)
}

func (o *PatchTeamForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamNotFound creates a PatchTeamNotFound with default headers values
func NewPatchTeamNotFound() *PatchTeamNotFound {
	return &PatchTeamNotFound{}
}

/*PatchTeamNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchTeamNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamNotFound  %+v", 404, o.Payload)
}

func (o *PatchTeamNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamRequestTimeout creates a PatchTeamRequestTimeout with default headers values
func NewPatchTeamRequestTimeout() *PatchTeamRequestTimeout {
	return &PatchTeamRequestTimeout{}
}

/*PatchTeamRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchTeamRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchTeamRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamRequestEntityTooLarge creates a PatchTeamRequestEntityTooLarge with default headers values
func NewPatchTeamRequestEntityTooLarge() *PatchTeamRequestEntityTooLarge {
	return &PatchTeamRequestEntityTooLarge{}
}

/*PatchTeamRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type PatchTeamRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchTeamRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamUnsupportedMediaType creates a PatchTeamUnsupportedMediaType with default headers values
func NewPatchTeamUnsupportedMediaType() *PatchTeamUnsupportedMediaType {
	return &PatchTeamUnsupportedMediaType{}
}

/*PatchTeamUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchTeamUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchTeamUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamTooManyRequests creates a PatchTeamTooManyRequests with default headers values
func NewPatchTeamTooManyRequests() *PatchTeamTooManyRequests {
	return &PatchTeamTooManyRequests{}
}

/*PatchTeamTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchTeamTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchTeamTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamInternalServerError creates a PatchTeamInternalServerError with default headers values
func NewPatchTeamInternalServerError() *PatchTeamInternalServerError {
	return &PatchTeamInternalServerError{}
}

/*PatchTeamInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchTeamInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchTeamInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamServiceUnavailable creates a PatchTeamServiceUnavailable with default headers values
func NewPatchTeamServiceUnavailable() *PatchTeamServiceUnavailable {
	return &PatchTeamServiceUnavailable{}
}

/*PatchTeamServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchTeamServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchTeamServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchTeamGatewayTimeout creates a PatchTeamGatewayTimeout with default headers values
func NewPatchTeamGatewayTimeout() *PatchTeamGatewayTimeout {
	return &PatchTeamGatewayTimeout{}
}

/*PatchTeamGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchTeamGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchTeamGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/teams/{teamId}][%d] patchTeamGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchTeamGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchTeamGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
