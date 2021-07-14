// Code generated by go-swagger; DO NOT EDIT.

package routing

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// PatchRoutingPredictorReader is a Reader for the PatchRoutingPredictor structure.
type PatchRoutingPredictorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchRoutingPredictorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchRoutingPredictorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPatchRoutingPredictorBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPatchRoutingPredictorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchRoutingPredictorForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchRoutingPredictorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPatchRoutingPredictorRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewPatchRoutingPredictorRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewPatchRoutingPredictorUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewPatchRoutingPredictorTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchRoutingPredictorInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewPatchRoutingPredictorServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewPatchRoutingPredictorGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPatchRoutingPredictorOK creates a PatchRoutingPredictorOK with default headers values
func NewPatchRoutingPredictorOK() *PatchRoutingPredictorOK {
	return &PatchRoutingPredictorOK{}
}

/*PatchRoutingPredictorOK handles this case with default header values.

successful operation
*/
type PatchRoutingPredictorOK struct {
	Payload *models.Predictor
}

func (o *PatchRoutingPredictorOK) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorOK  %+v", 200, o.Payload)
}

func (o *PatchRoutingPredictorOK) GetPayload() *models.Predictor {
	return o.Payload
}

func (o *PatchRoutingPredictorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Predictor)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorBadRequest creates a PatchRoutingPredictorBadRequest with default headers values
func NewPatchRoutingPredictorBadRequest() *PatchRoutingPredictorBadRequest {
	return &PatchRoutingPredictorBadRequest{}
}

/*PatchRoutingPredictorBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type PatchRoutingPredictorBadRequest struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorBadRequest  %+v", 400, o.Payload)
}

func (o *PatchRoutingPredictorBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorUnauthorized creates a PatchRoutingPredictorUnauthorized with default headers values
func NewPatchRoutingPredictorUnauthorized() *PatchRoutingPredictorUnauthorized {
	return &PatchRoutingPredictorUnauthorized{}
}

/*PatchRoutingPredictorUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type PatchRoutingPredictorUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorUnauthorized  %+v", 401, o.Payload)
}

func (o *PatchRoutingPredictorUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorForbidden creates a PatchRoutingPredictorForbidden with default headers values
func NewPatchRoutingPredictorForbidden() *PatchRoutingPredictorForbidden {
	return &PatchRoutingPredictorForbidden{}
}

/*PatchRoutingPredictorForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type PatchRoutingPredictorForbidden struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorForbidden) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorForbidden  %+v", 403, o.Payload)
}

func (o *PatchRoutingPredictorForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorNotFound creates a PatchRoutingPredictorNotFound with default headers values
func NewPatchRoutingPredictorNotFound() *PatchRoutingPredictorNotFound {
	return &PatchRoutingPredictorNotFound{}
}

/*PatchRoutingPredictorNotFound handles this case with default header values.

The requested resource was not found.
*/
type PatchRoutingPredictorNotFound struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorNotFound) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorNotFound  %+v", 404, o.Payload)
}

func (o *PatchRoutingPredictorNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorRequestTimeout creates a PatchRoutingPredictorRequestTimeout with default headers values
func NewPatchRoutingPredictorRequestTimeout() *PatchRoutingPredictorRequestTimeout {
	return &PatchRoutingPredictorRequestTimeout{}
}

/*PatchRoutingPredictorRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type PatchRoutingPredictorRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorRequestTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorRequestTimeout  %+v", 408, o.Payload)
}

func (o *PatchRoutingPredictorRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorRequestEntityTooLarge creates a PatchRoutingPredictorRequestEntityTooLarge with default headers values
func NewPatchRoutingPredictorRequestEntityTooLarge() *PatchRoutingPredictorRequestEntityTooLarge {
	return &PatchRoutingPredictorRequestEntityTooLarge{}
}

/*PatchRoutingPredictorRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type PatchRoutingPredictorRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *PatchRoutingPredictorRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorUnsupportedMediaType creates a PatchRoutingPredictorUnsupportedMediaType with default headers values
func NewPatchRoutingPredictorUnsupportedMediaType() *PatchRoutingPredictorUnsupportedMediaType {
	return &PatchRoutingPredictorUnsupportedMediaType{}
}

/*PatchRoutingPredictorUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type PatchRoutingPredictorUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *PatchRoutingPredictorUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorTooManyRequests creates a PatchRoutingPredictorTooManyRequests with default headers values
func NewPatchRoutingPredictorTooManyRequests() *PatchRoutingPredictorTooManyRequests {
	return &PatchRoutingPredictorTooManyRequests{}
}

/*PatchRoutingPredictorTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type PatchRoutingPredictorTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorTooManyRequests) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorTooManyRequests  %+v", 429, o.Payload)
}

func (o *PatchRoutingPredictorTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorInternalServerError creates a PatchRoutingPredictorInternalServerError with default headers values
func NewPatchRoutingPredictorInternalServerError() *PatchRoutingPredictorInternalServerError {
	return &PatchRoutingPredictorInternalServerError{}
}

/*PatchRoutingPredictorInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type PatchRoutingPredictorInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorInternalServerError  %+v", 500, o.Payload)
}

func (o *PatchRoutingPredictorInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorServiceUnavailable creates a PatchRoutingPredictorServiceUnavailable with default headers values
func NewPatchRoutingPredictorServiceUnavailable() *PatchRoutingPredictorServiceUnavailable {
	return &PatchRoutingPredictorServiceUnavailable{}
}

/*PatchRoutingPredictorServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type PatchRoutingPredictorServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorServiceUnavailable) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorServiceUnavailable  %+v", 503, o.Payload)
}

func (o *PatchRoutingPredictorServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchRoutingPredictorGatewayTimeout creates a PatchRoutingPredictorGatewayTimeout with default headers values
func NewPatchRoutingPredictorGatewayTimeout() *PatchRoutingPredictorGatewayTimeout {
	return &PatchRoutingPredictorGatewayTimeout{}
}

/*PatchRoutingPredictorGatewayTimeout handles this case with default header values.

The request timed out.
*/
type PatchRoutingPredictorGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *PatchRoutingPredictorGatewayTimeout) Error() string {
	return fmt.Sprintf("[PATCH /api/v2/routing/predictors/{predictorId}][%d] patchRoutingPredictorGatewayTimeout  %+v", 504, o.Payload)
}

func (o *PatchRoutingPredictorGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *PatchRoutingPredictorGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
