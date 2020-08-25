// Code generated by go-swagger; DO NOT EDIT.

package architect

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PostArchitectDependencytrackingBuildReader is a Reader for the PostArchitectDependencytrackingBuild structure.
type PostArchitectDependencytrackingBuildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostArchitectDependencytrackingBuildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewPostArchitectDependencytrackingBuildAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostArchitectDependencytrackingBuildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostArchitectDependencytrackingBuildUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostArchitectDependencytrackingBuildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostArchitectDependencytrackingBuildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostArchitectDependencytrackingBuildConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPostArchitectDependencytrackingBuildRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPostArchitectDependencytrackingBuildUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPostArchitectDependencytrackingBuildTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostArchitectDependencytrackingBuildInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPostArchitectDependencytrackingBuildServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPostArchitectDependencytrackingBuildGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostArchitectDependencytrackingBuildAccepted creates a PostArchitectDependencytrackingBuildAccepted with default headers values
func NewPostArchitectDependencytrackingBuildAccepted() *PostArchitectDependencytrackingBuildAccepted {
	return &PostArchitectDependencytrackingBuildAccepted{}
}

/*PostArchitectDependencytrackingBuildAccepted handles this case with default header values.

Accepted - the rebuild has begun.
*/
type PostArchitectDependencytrackingBuildAccepted struct {
}

func (o *PostArchitectDependencytrackingBuildAccepted) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildAccepted ", 202)
}

func (o *PostArchitectDependencytrackingBuildAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostArchitectDependencytrackingBuildBadRequest creates a PostArchitectDependencytrackingBuildBadRequest with default headers values
func NewPostArchitectDependencytrackingBuildBadRequest() *PostArchitectDependencytrackingBuildBadRequest {
	return &PostArchitectDependencytrackingBuildBadRequest{}
}

/*PostArchitectDependencytrackingBuildBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PostArchitectDependencytrackingBuildBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildBadRequest  %+v", 400, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildUnauthorized creates a PostArchitectDependencytrackingBuildUnauthorized with default headers values
func NewPostArchitectDependencytrackingBuildUnauthorized() *PostArchitectDependencytrackingBuildUnauthorized {
	return &PostArchitectDependencytrackingBuildUnauthorized{}
}

/*PostArchitectDependencytrackingBuildUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PostArchitectDependencytrackingBuildUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildUnauthorized) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildUnauthorized  %+v", 401, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildForbidden creates a PostArchitectDependencytrackingBuildForbidden with default headers values
func NewPostArchitectDependencytrackingBuildForbidden() *PostArchitectDependencytrackingBuildForbidden {
	return &PostArchitectDependencytrackingBuildForbidden{}
}

/*PostArchitectDependencytrackingBuildForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PostArchitectDependencytrackingBuildForbidden struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildForbidden  %+v", 403, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildNotFound creates a PostArchitectDependencytrackingBuildNotFound with default headers values
func NewPostArchitectDependencytrackingBuildNotFound() *PostArchitectDependencytrackingBuildNotFound {
	return &PostArchitectDependencytrackingBuildNotFound{}
}

/*PostArchitectDependencytrackingBuildNotFound handles this case with default header values.

The requested resource was not found.
*/
type PostArchitectDependencytrackingBuildNotFound struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildNotFound  %+v", 404, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildConflict creates a PostArchitectDependencytrackingBuildConflict with default headers values
func NewPostArchitectDependencytrackingBuildConflict() *PostArchitectDependencytrackingBuildConflict {
	return &PostArchitectDependencytrackingBuildConflict{}
}

/*PostArchitectDependencytrackingBuildConflict handles this case with default header values.

Conflict
*/
type PostArchitectDependencytrackingBuildConflict struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildConflict) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildConflict  %+v", 409, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildConflict) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildRequestEntityTooLarge creates a PostArchitectDependencytrackingBuildRequestEntityTooLarge with default headers values
func NewPostArchitectDependencytrackingBuildRequestEntityTooLarge() *PostArchitectDependencytrackingBuildRequestEntityTooLarge {
	return &PostArchitectDependencytrackingBuildRequestEntityTooLarge{}
}

/*PostArchitectDependencytrackingBuildRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PostArchitectDependencytrackingBuildRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildUnsupportedMediaType creates a PostArchitectDependencytrackingBuildUnsupportedMediaType with default headers values
func NewPostArchitectDependencytrackingBuildUnsupportedMediaType() *PostArchitectDependencytrackingBuildUnsupportedMediaType {
	return &PostArchitectDependencytrackingBuildUnsupportedMediaType{}
}

/*PostArchitectDependencytrackingBuildUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PostArchitectDependencytrackingBuildUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildTooManyRequests creates a PostArchitectDependencytrackingBuildTooManyRequests with default headers values
func NewPostArchitectDependencytrackingBuildTooManyRequests() *PostArchitectDependencytrackingBuildTooManyRequests {
	return &PostArchitectDependencytrackingBuildTooManyRequests{}
}

/*PostArchitectDependencytrackingBuildTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type PostArchitectDependencytrackingBuildTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildTooManyRequests) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildTooManyRequests  %+v", 429, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildInternalServerError creates a PostArchitectDependencytrackingBuildInternalServerError with default headers values
func NewPostArchitectDependencytrackingBuildInternalServerError() *PostArchitectDependencytrackingBuildInternalServerError {
	return &PostArchitectDependencytrackingBuildInternalServerError{}
}

/*PostArchitectDependencytrackingBuildInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PostArchitectDependencytrackingBuildInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildInternalServerError) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildInternalServerError  %+v", 500, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildServiceUnavailable creates a PostArchitectDependencytrackingBuildServiceUnavailable with default headers values
func NewPostArchitectDependencytrackingBuildServiceUnavailable() *PostArchitectDependencytrackingBuildServiceUnavailable {
	return &PostArchitectDependencytrackingBuildServiceUnavailable{}
}

/*PostArchitectDependencytrackingBuildServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PostArchitectDependencytrackingBuildServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildServiceUnavailable) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostArchitectDependencytrackingBuildGatewayTimeout creates a PostArchitectDependencytrackingBuildGatewayTimeout with default headers values
func NewPostArchitectDependencytrackingBuildGatewayTimeout() *PostArchitectDependencytrackingBuildGatewayTimeout {
	return &PostArchitectDependencytrackingBuildGatewayTimeout{}
}

/*PostArchitectDependencytrackingBuildGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PostArchitectDependencytrackingBuildGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PostArchitectDependencytrackingBuildGatewayTimeout) Error() string {
	return fmt.Sprintf("[POST /api/v2/architect/dependencytracking/build][%d] postArchitectDependencytrackingBuildGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PostArchitectDependencytrackingBuildGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PostArchitectDependencytrackingBuildGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}