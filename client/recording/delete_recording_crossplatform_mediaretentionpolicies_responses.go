// Code generated by go-swagger; DO NOT EDIT.

package recording

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// DeleteRecordingCrossplatformMediaretentionpoliciesReader is a Reader for the DeleteRecordingCrossplatformMediaretentionpolicies structure.
type DeleteRecordingCrossplatformMediaretentionpoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecordingCrossplatformMediaretentionpoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesOK creates a DeleteRecordingCrossplatformMediaretentionpoliciesOK with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesOK() *DeleteRecordingCrossplatformMediaretentionpoliciesOK {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesOK{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesOK handles this case with default header values.

Operation was successful.
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesOK struct {
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesOK ", 200)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesBadRequest creates a DeleteRecordingCrossplatformMediaretentionpoliciesBadRequest with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesBadRequest() *DeleteRecordingCrossplatformMediaretentionpoliciesBadRequest {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesBadRequest{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized creates a DeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized() *DeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesForbidden creates a DeleteRecordingCrossplatformMediaretentionpoliciesForbidden with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesForbidden() *DeleteRecordingCrossplatformMediaretentionpoliciesForbidden {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesForbidden{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesNotFound creates a DeleteRecordingCrossplatformMediaretentionpoliciesNotFound with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesNotFound() *DeleteRecordingCrossplatformMediaretentionpoliciesNotFound {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesNotFound{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge creates a DeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge() *DeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType creates a DeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType() *DeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests creates a DeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests() *DeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError creates a DeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError() *DeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable creates a DeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable() *DeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout creates a DeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout with default headers values
func NewDeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout() *DeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout {
	return &DeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout{}
}

/*DeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/crossplatform/mediaretentionpolicies][%d] deleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingCrossplatformMediaretentionpoliciesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}