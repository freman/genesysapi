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

// GetRecordingMediaretentionpoliciesReader is a Reader for the GetRecordingMediaretentionpolicies structure.
type GetRecordingMediaretentionpoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRecordingMediaretentionpoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRecordingMediaretentionpoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRecordingMediaretentionpoliciesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRecordingMediaretentionpoliciesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRecordingMediaretentionpoliciesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRecordingMediaretentionpoliciesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRecordingMediaretentionpoliciesRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRecordingMediaretentionpoliciesRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRecordingMediaretentionpoliciesUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRecordingMediaretentionpoliciesTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRecordingMediaretentionpoliciesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRecordingMediaretentionpoliciesServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRecordingMediaretentionpoliciesGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRecordingMediaretentionpoliciesOK creates a GetRecordingMediaretentionpoliciesOK with default headers values
func NewGetRecordingMediaretentionpoliciesOK() *GetRecordingMediaretentionpoliciesOK {
	return &GetRecordingMediaretentionpoliciesOK{}
}

/*GetRecordingMediaretentionpoliciesOK handles this case with default header values.

successful operation
*/
type GetRecordingMediaretentionpoliciesOK struct {
	Payload *models.PolicyEntityListing
}

func (o *GetRecordingMediaretentionpoliciesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesOK  %+v", 200, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesOK) GetPayload() *models.PolicyEntityListing {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PolicyEntityListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesBadRequest creates a GetRecordingMediaretentionpoliciesBadRequest with default headers values
func NewGetRecordingMediaretentionpoliciesBadRequest() *GetRecordingMediaretentionpoliciesBadRequest {
	return &GetRecordingMediaretentionpoliciesBadRequest{}
}

/*GetRecordingMediaretentionpoliciesBadRequest handles this case with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRecordingMediaretentionpoliciesBadRequest struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesUnauthorized creates a GetRecordingMediaretentionpoliciesUnauthorized with default headers values
func NewGetRecordingMediaretentionpoliciesUnauthorized() *GetRecordingMediaretentionpoliciesUnauthorized {
	return &GetRecordingMediaretentionpoliciesUnauthorized{}
}

/*GetRecordingMediaretentionpoliciesUnauthorized handles this case with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRecordingMediaretentionpoliciesUnauthorized struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesForbidden creates a GetRecordingMediaretentionpoliciesForbidden with default headers values
func NewGetRecordingMediaretentionpoliciesForbidden() *GetRecordingMediaretentionpoliciesForbidden {
	return &GetRecordingMediaretentionpoliciesForbidden{}
}

/*GetRecordingMediaretentionpoliciesForbidden handles this case with default header values.

You are not authorized to perform the requested action.
*/
type GetRecordingMediaretentionpoliciesForbidden struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesForbidden  %+v", 403, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesNotFound creates a GetRecordingMediaretentionpoliciesNotFound with default headers values
func NewGetRecordingMediaretentionpoliciesNotFound() *GetRecordingMediaretentionpoliciesNotFound {
	return &GetRecordingMediaretentionpoliciesNotFound{}
}

/*GetRecordingMediaretentionpoliciesNotFound handles this case with default header values.

The requested resource was not found.
*/
type GetRecordingMediaretentionpoliciesNotFound struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesNotFound  %+v", 404, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesRequestTimeout creates a GetRecordingMediaretentionpoliciesRequestTimeout with default headers values
func NewGetRecordingMediaretentionpoliciesRequestTimeout() *GetRecordingMediaretentionpoliciesRequestTimeout {
	return &GetRecordingMediaretentionpoliciesRequestTimeout{}
}

/*GetRecordingMediaretentionpoliciesRequestTimeout handles this case with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRecordingMediaretentionpoliciesRequestTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesRequestEntityTooLarge creates a GetRecordingMediaretentionpoliciesRequestEntityTooLarge with default headers values
func NewGetRecordingMediaretentionpoliciesRequestEntityTooLarge() *GetRecordingMediaretentionpoliciesRequestEntityTooLarge {
	return &GetRecordingMediaretentionpoliciesRequestEntityTooLarge{}
}

/*GetRecordingMediaretentionpoliciesRequestEntityTooLarge handles this case with default header values.

The request is over the size limit. Content-Length: %s, Maximum bytes: %s
*/
type GetRecordingMediaretentionpoliciesRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesUnsupportedMediaType creates a GetRecordingMediaretentionpoliciesUnsupportedMediaType with default headers values
func NewGetRecordingMediaretentionpoliciesUnsupportedMediaType() *GetRecordingMediaretentionpoliciesUnsupportedMediaType {
	return &GetRecordingMediaretentionpoliciesUnsupportedMediaType{}
}

/*GetRecordingMediaretentionpoliciesUnsupportedMediaType handles this case with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRecordingMediaretentionpoliciesUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesTooManyRequests creates a GetRecordingMediaretentionpoliciesTooManyRequests with default headers values
func NewGetRecordingMediaretentionpoliciesTooManyRequests() *GetRecordingMediaretentionpoliciesTooManyRequests {
	return &GetRecordingMediaretentionpoliciesTooManyRequests{}
}

/*GetRecordingMediaretentionpoliciesTooManyRequests handles this case with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRecordingMediaretentionpoliciesTooManyRequests struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesInternalServerError creates a GetRecordingMediaretentionpoliciesInternalServerError with default headers values
func NewGetRecordingMediaretentionpoliciesInternalServerError() *GetRecordingMediaretentionpoliciesInternalServerError {
	return &GetRecordingMediaretentionpoliciesInternalServerError{}
}

/*GetRecordingMediaretentionpoliciesInternalServerError handles this case with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRecordingMediaretentionpoliciesInternalServerError struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesServiceUnavailable creates a GetRecordingMediaretentionpoliciesServiceUnavailable with default headers values
func NewGetRecordingMediaretentionpoliciesServiceUnavailable() *GetRecordingMediaretentionpoliciesServiceUnavailable {
	return &GetRecordingMediaretentionpoliciesServiceUnavailable{}
}

/*GetRecordingMediaretentionpoliciesServiceUnavailable handles this case with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRecordingMediaretentionpoliciesServiceUnavailable struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRecordingMediaretentionpoliciesGatewayTimeout creates a GetRecordingMediaretentionpoliciesGatewayTimeout with default headers values
func NewGetRecordingMediaretentionpoliciesGatewayTimeout() *GetRecordingMediaretentionpoliciesGatewayTimeout {
	return &GetRecordingMediaretentionpoliciesGatewayTimeout{}
}

/*GetRecordingMediaretentionpoliciesGatewayTimeout handles this case with default header values.

The request timed out.
*/
type GetRecordingMediaretentionpoliciesGatewayTimeout struct {
	Payload *models.ErrorBody
}

func (o *GetRecordingMediaretentionpoliciesGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/recording/mediaretentionpolicies][%d] getRecordingMediaretentionpoliciesGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRecordingMediaretentionpoliciesGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRecordingMediaretentionpoliciesGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
