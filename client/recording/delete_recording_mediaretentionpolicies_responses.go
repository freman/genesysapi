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

// DeleteRecordingMediaretentionpoliciesReader is a Reader for the DeleteRecordingMediaretentionpolicies structure.
type DeleteRecordingMediaretentionpoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRecordingMediaretentionpoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRecordingMediaretentionpoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRecordingMediaretentionpoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRecordingMediaretentionpoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRecordingMediaretentionpoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRecordingMediaretentionpoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewDeleteRecordingMediaretentionpoliciesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewDeleteRecordingMediaretentionpoliciesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewDeleteRecordingMediaretentionpoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRecordingMediaretentionpoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewDeleteRecordingMediaretentionpoliciesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewDeleteRecordingMediaretentionpoliciesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRecordingMediaretentionpoliciesOK creates a DeleteRecordingMediaretentionpoliciesOK with default headers values
func NewDeleteRecordingMediaretentionpoliciesOK() *DeleteRecordingMediaretentionpoliciesOK {
	return &DeleteRecordingMediaretentionpoliciesOK{}
}

/*DeleteRecordingMediaretentionpoliciesOK handles this case with default header values.

Operation was successful.
*/
type DeleteRecordingMediaretentionpoliciesOK struct {
}

func (o *DeleteRecordingMediaretentionpoliciesOK) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesOK ", 200)
}

func (o *DeleteRecordingMediaretentionpoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesBadRequest creates a DeleteRecordingMediaretentionpoliciesBadRequest with default headers values
func NewDeleteRecordingMediaretentionpoliciesBadRequest() *DeleteRecordingMediaretentionpoliciesBadRequest {
	return &DeleteRecordingMediaretentionpoliciesBadRequest{}
}

/*DeleteRecordingMediaretentionpoliciesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type DeleteRecordingMediaretentionpoliciesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesUnauthorized creates a DeleteRecordingMediaretentionpoliciesUnauthorized with default headers values
func NewDeleteRecordingMediaretentionpoliciesUnauthorized() *DeleteRecordingMediaretentionpoliciesUnauthorized {
	return &DeleteRecordingMediaretentionpoliciesUnauthorized{}
}

/*DeleteRecordingMediaretentionpoliciesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type DeleteRecordingMediaretentionpoliciesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesForbidden creates a DeleteRecordingMediaretentionpoliciesForbidden with default headers values
func NewDeleteRecordingMediaretentionpoliciesForbidden() *DeleteRecordingMediaretentionpoliciesForbidden {
	return &DeleteRecordingMediaretentionpoliciesForbidden{}
}

/*DeleteRecordingMediaretentionpoliciesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type DeleteRecordingMediaretentionpoliciesForbidden struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesForbidden) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesForbidden  %+v", 403, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesNotFound creates a DeleteRecordingMediaretentionpoliciesNotFound with default headers values
func NewDeleteRecordingMediaretentionpoliciesNotFound() *DeleteRecordingMediaretentionpoliciesNotFound {
	return &DeleteRecordingMediaretentionpoliciesNotFound{}
}

/*DeleteRecordingMediaretentionpoliciesNotFound handles this case with default header values.

The requested resource was not found.
*/
type DeleteRecordingMediaretentionpoliciesNotFound struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesNotFound) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesNotFound  %+v", 404, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesRequestEntityTooLarge creates a DeleteRecordingMediaretentionpoliciesRequestEntityTooLarge with default headers values
func NewDeleteRecordingMediaretentionpoliciesRequestEntityTooLarge() *DeleteRecordingMediaretentionpoliciesRequestEntityTooLarge {
	return &DeleteRecordingMediaretentionpoliciesRequestEntityTooLarge{}
}

/*DeleteRecordingMediaretentionpoliciesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type DeleteRecordingMediaretentionpoliciesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesUnsupportedMediaType creates a DeleteRecordingMediaretentionpoliciesUnsupportedMediaType with default headers values
func NewDeleteRecordingMediaretentionpoliciesUnsupportedMediaType() *DeleteRecordingMediaretentionpoliciesUnsupportedMediaType {
	return &DeleteRecordingMediaretentionpoliciesUnsupportedMediaType{}
}

/*DeleteRecordingMediaretentionpoliciesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type DeleteRecordingMediaretentionpoliciesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesTooManyRequests creates a DeleteRecordingMediaretentionpoliciesTooManyRequests with default headers values
func NewDeleteRecordingMediaretentionpoliciesTooManyRequests() *DeleteRecordingMediaretentionpoliciesTooManyRequests {
	return &DeleteRecordingMediaretentionpoliciesTooManyRequests{}
}

/*DeleteRecordingMediaretentionpoliciesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type DeleteRecordingMediaretentionpoliciesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesInternalServerError creates a DeleteRecordingMediaretentionpoliciesInternalServerError with default headers values
func NewDeleteRecordingMediaretentionpoliciesInternalServerError() *DeleteRecordingMediaretentionpoliciesInternalServerError {
	return &DeleteRecordingMediaretentionpoliciesInternalServerError{}
}

/*DeleteRecordingMediaretentionpoliciesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type DeleteRecordingMediaretentionpoliciesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesServiceUnavailable creates a DeleteRecordingMediaretentionpoliciesServiceUnavailable with default headers values
func NewDeleteRecordingMediaretentionpoliciesServiceUnavailable() *DeleteRecordingMediaretentionpoliciesServiceUnavailable {
	return &DeleteRecordingMediaretentionpoliciesServiceUnavailable{}
}

/*DeleteRecordingMediaretentionpoliciesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type DeleteRecordingMediaretentionpoliciesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesServiceUnavailable) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRecordingMediaretentionpoliciesGatewayTimeout creates a DeleteRecordingMediaretentionpoliciesGatewayTimeout with default headers values
func NewDeleteRecordingMediaretentionpoliciesGatewayTimeout() *DeleteRecordingMediaretentionpoliciesGatewayTimeout {
	return &DeleteRecordingMediaretentionpoliciesGatewayTimeout{}
}

/*DeleteRecordingMediaretentionpoliciesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type DeleteRecordingMediaretentionpoliciesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *DeleteRecordingMediaretentionpoliciesGatewayTimeout) Error() string {
	return fmt.Sprintf("[DELETE /api/v2/recording/mediaretentionpolicies][%d] deleteRecordingMediaretentionpoliciesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *DeleteRecordingMediaretentionpoliciesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *DeleteRecordingMediaretentionpoliciesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
