// Code generated by go-swagger; DO NOT EDIT.

package quality

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/freman/genesysapi/models"
)

// GetQualityFormsEvaluationsReader is a Reader for the GetQualityFormsEvaluations structure.
type GetQualityFormsEvaluationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQualityFormsEvaluationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQualityFormsEvaluationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetQualityFormsEvaluationsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetQualityFormsEvaluationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQualityFormsEvaluationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQualityFormsEvaluationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetQualityFormsEvaluationsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetQualityFormsEvaluationsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetQualityFormsEvaluationsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQualityFormsEvaluationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetQualityFormsEvaluationsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetQualityFormsEvaluationsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQualityFormsEvaluationsOK creates a GetQualityFormsEvaluationsOK with default headers values
func NewGetQualityFormsEvaluationsOK() *GetQualityFormsEvaluationsOK {
	return &GetQualityFormsEvaluationsOK{}
}

/*GetQualityFormsEvaluationsOK handles this case with default header values.

successful operation
*/
type GetQualityFormsEvaluationsOK struct {
	Payload *models.EvaluationFormEntityListing
}

func (o *GetQualityFormsEvaluationsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsOK  %+v", 200, o.Payload)
}

func (o *GetQualityFormsEvaluationsOK) GetPayload() *models.EvaluationFormEntityListing {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EvaluationFormEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsBadRequest creates a GetQualityFormsEvaluationsBadRequest with default headers values
func NewGetQualityFormsEvaluationsBadRequest() *GetQualityFormsEvaluationsBadRequest {
	return &GetQualityFormsEvaluationsBadRequest{}
}

/*GetQualityFormsEvaluationsBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetQualityFormsEvaluationsBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsBadRequest  %+v", 400, o.Payload)
}

func (o *GetQualityFormsEvaluationsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsUnauthorized creates a GetQualityFormsEvaluationsUnauthorized with default headers values
func NewGetQualityFormsEvaluationsUnauthorized() *GetQualityFormsEvaluationsUnauthorized {
	return &GetQualityFormsEvaluationsUnauthorized{}
}

/*GetQualityFormsEvaluationsUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetQualityFormsEvaluationsUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetQualityFormsEvaluationsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsForbidden creates a GetQualityFormsEvaluationsForbidden with default headers values
func NewGetQualityFormsEvaluationsForbidden() *GetQualityFormsEvaluationsForbidden {
	return &GetQualityFormsEvaluationsForbidden{}
}

/*GetQualityFormsEvaluationsForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetQualityFormsEvaluationsForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsForbidden  %+v", 403, o.Payload)
}

func (o *GetQualityFormsEvaluationsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsNotFound creates a GetQualityFormsEvaluationsNotFound with default headers values
func NewGetQualityFormsEvaluationsNotFound() *GetQualityFormsEvaluationsNotFound {
	return &GetQualityFormsEvaluationsNotFound{}
}

/*GetQualityFormsEvaluationsNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetQualityFormsEvaluationsNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsNotFound  %+v", 404, o.Payload)
}

func (o *GetQualityFormsEvaluationsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsRequestEntityTooLarge creates a GetQualityFormsEvaluationsRequestEntityTooLarge with default headers values
func NewGetQualityFormsEvaluationsRequestEntityTooLarge() *GetQualityFormsEvaluationsRequestEntityTooLarge {
	return &GetQualityFormsEvaluationsRequestEntityTooLarge{}
}

/*GetQualityFormsEvaluationsRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetQualityFormsEvaluationsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetQualityFormsEvaluationsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsUnsupportedMediaType creates a GetQualityFormsEvaluationsUnsupportedMediaType with default headers values
func NewGetQualityFormsEvaluationsUnsupportedMediaType() *GetQualityFormsEvaluationsUnsupportedMediaType {
	return &GetQualityFormsEvaluationsUnsupportedMediaType{}
}

/*GetQualityFormsEvaluationsUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetQualityFormsEvaluationsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetQualityFormsEvaluationsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsTooManyRequests creates a GetQualityFormsEvaluationsTooManyRequests with default headers values
func NewGetQualityFormsEvaluationsTooManyRequests() *GetQualityFormsEvaluationsTooManyRequests {
	return &GetQualityFormsEvaluationsTooManyRequests{}
}

/*GetQualityFormsEvaluationsTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetQualityFormsEvaluationsTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetQualityFormsEvaluationsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsInternalServerError creates a GetQualityFormsEvaluationsInternalServerError with default headers values
func NewGetQualityFormsEvaluationsInternalServerError() *GetQualityFormsEvaluationsInternalServerError {
	return &GetQualityFormsEvaluationsInternalServerError{}
}

/*GetQualityFormsEvaluationsInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetQualityFormsEvaluationsInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetQualityFormsEvaluationsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsServiceUnavailable creates a GetQualityFormsEvaluationsServiceUnavailable with default headers values
func NewGetQualityFormsEvaluationsServiceUnavailable() *GetQualityFormsEvaluationsServiceUnavailable {
	return &GetQualityFormsEvaluationsServiceUnavailable{}
}

/*GetQualityFormsEvaluationsServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetQualityFormsEvaluationsServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetQualityFormsEvaluationsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQualityFormsEvaluationsGatewayTimeout creates a GetQualityFormsEvaluationsGatewayTimeout with default headers values
func NewGetQualityFormsEvaluationsGatewayTimeout() *GetQualityFormsEvaluationsGatewayTimeout {
	return &GetQualityFormsEvaluationsGatewayTimeout{}
}

/*GetQualityFormsEvaluationsGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetQualityFormsEvaluationsGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetQualityFormsEvaluationsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/quality/forms/evaluations][%d] getQualityFormsEvaluationsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetQualityFormsEvaluationsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetQualityFormsEvaluationsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
