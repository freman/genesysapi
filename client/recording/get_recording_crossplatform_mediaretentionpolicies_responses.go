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

// GetRecordingCrossplatformMediaretentionpoliciesReader is a Reader for the GetRecordingCrossplatformMediaretentionpolicies structure.
type GetRecordingCrossplatformMediaretentionpoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecordingCrossplatformMediaretentionpoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecordingCrossplatformMediaretentionpoliciesOK creates a GetRecordingCrossplatformMediaretentionpoliciesOK with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesOK() *GetRecordingCrossplatformMediaretentionpoliciesOK {
	return &GetRecordingCrossplatformMediaretentionpoliciesOK{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesOK handles this case with default header values.

successful operation
*/
type GetRecordingCrossplatformMediaretentionpoliciesOK struct {
	Payload *models.PolicyEntityListing
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesOK  %+v", 200, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesOK) GetPayload() *models.PolicyEntityListing {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesBadRequest creates a GetRecordingCrossplatformMediaretentionpoliciesBadRequest with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesBadRequest() *GetRecordingCrossplatformMediaretentionpoliciesBadRequest {
	return &GetRecordingCrossplatformMediaretentionpoliciesBadRequest{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingCrossplatformMediaretentionpoliciesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesUnauthorized creates a GetRecordingCrossplatformMediaretentionpoliciesUnauthorized with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesUnauthorized() *GetRecordingCrossplatformMediaretentionpoliciesUnauthorized {
	return &GetRecordingCrossplatformMediaretentionpoliciesUnauthorized{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingCrossplatformMediaretentionpoliciesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesForbidden creates a GetRecordingCrossplatformMediaretentionpoliciesForbidden with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesForbidden() *GetRecordingCrossplatformMediaretentionpoliciesForbidden {
	return &GetRecordingCrossplatformMediaretentionpoliciesForbidden{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingCrossplatformMediaretentionpoliciesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesNotFound creates a GetRecordingCrossplatformMediaretentionpoliciesNotFound with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesNotFound() *GetRecordingCrossplatformMediaretentionpoliciesNotFound {
	return &GetRecordingCrossplatformMediaretentionpoliciesNotFound{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRecordingCrossplatformMediaretentionpoliciesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge creates a GetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge() *GetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge {
	return &GetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s
*/
type GetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType creates a GetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType() *GetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType {
	return &GetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesTooManyRequests creates a GetRecordingCrossplatformMediaretentionpoliciesTooManyRequests with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesTooManyRequests() *GetRecordingCrossplatformMediaretentionpoliciesTooManyRequests {
	return &GetRecordingCrossplatformMediaretentionpoliciesTooManyRequests{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum [%s] requests within [%s] seconds
*/
type GetRecordingCrossplatformMediaretentionpoliciesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesInternalServerError creates a GetRecordingCrossplatformMediaretentionpoliciesInternalServerError with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesInternalServerError() *GetRecordingCrossplatformMediaretentionpoliciesInternalServerError {
	return &GetRecordingCrossplatformMediaretentionpoliciesInternalServerError{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingCrossplatformMediaretentionpoliciesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable creates a GetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable() *GetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable {
	return &GetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout creates a GetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout with default headers values
func NewGetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout() *GetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout {
	return &GetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout{}
}

/*GetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/crossplatform/mediaretentionpolicies][%d] getRecordingCrossplatformMediaretentionpoliciesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingCrossplatformMediaretentionpoliciesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
