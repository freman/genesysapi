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

// GetRoutingPredictorModelsReader is a Reader for the GetRoutingPredictorModels structure.
type GetRoutingPredictorModelsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoutingPredictorModelsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRoutingPredictorModelsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetRoutingPredictorModelsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetRoutingPredictorModelsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRoutingPredictorModelsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRoutingPredictorModelsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewGetRoutingPredictorModelsRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 413:
		result := NewGetRoutingPredictorModelsRequestEntityTooLarge()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 415:
		result := NewGetRoutingPredictorModelsUnsupportedMediaType()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetRoutingPredictorModelsTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRoutingPredictorModelsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 503:
		result := NewGetRoutingPredictorModelsServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewGetRoutingPredictorModelsGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRoutingPredictorModelsOK creates a GetRoutingPredictorModelsOK with default headers values
func NewGetRoutingPredictorModelsOK() *GetRoutingPredictorModelsOK {
	return &GetRoutingPredictorModelsOK{}
}

/*
GetRoutingPredictorModelsOK describes a response with status code 200, with default header values.

successful operation
*/
type GetRoutingPredictorModelsOK struct {
	Payload *models.PredictorModels
}

// IsSuccess returns true when this get routing predictor models o k response has a 2xx status code
func (o *GetRoutingPredictorModelsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get routing predictor models o k response has a 3xx status code
func (o *GetRoutingPredictorModelsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models o k response has a 4xx status code
func (o *GetRoutingPredictorModelsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing predictor models o k response has a 5xx status code
func (o *GetRoutingPredictorModelsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor models o k response a status code equal to that given
func (o *GetRoutingPredictorModelsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetRoutingPredictorModelsOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingPredictorModelsOK) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsOK  %+v", 200, o.Payload)
}

func (o *GetRoutingPredictorModelsOK) GetPayload() *models.PredictorModels {
	return o.Payload
}

func (o *GetRoutingPredictorModelsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PredictorModels)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsBadRequest creates a GetRoutingPredictorModelsBadRequest with default headers values
func NewGetRoutingPredictorModelsBadRequest() *GetRoutingPredictorModelsBadRequest {
	return &GetRoutingPredictorModelsBadRequest{}
}

/*
GetRoutingPredictorModelsBadRequest describes a response with status code 400, with default header values.

The request could not be understood by the server due to malformed syntax.
*/
type GetRoutingPredictorModelsBadRequest struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models bad request response has a 2xx status code
func (o *GetRoutingPredictorModelsBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models bad request response has a 3xx status code
func (o *GetRoutingPredictorModelsBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models bad request response has a 4xx status code
func (o *GetRoutingPredictorModelsBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor models bad request response has a 5xx status code
func (o *GetRoutingPredictorModelsBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor models bad request response a status code equal to that given
func (o *GetRoutingPredictorModelsBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *GetRoutingPredictorModelsBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingPredictorModelsBadRequest) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsBadRequest  %+v", 400, o.Payload)
}

func (o *GetRoutingPredictorModelsBadRequest) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsUnauthorized creates a GetRoutingPredictorModelsUnauthorized with default headers values
func NewGetRoutingPredictorModelsUnauthorized() *GetRoutingPredictorModelsUnauthorized {
	return &GetRoutingPredictorModelsUnauthorized{}
}

/*
GetRoutingPredictorModelsUnauthorized describes a response with status code 401, with default header values.

No authentication bearer token specified in authorization header.
*/
type GetRoutingPredictorModelsUnauthorized struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models unauthorized response has a 2xx status code
func (o *GetRoutingPredictorModelsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models unauthorized response has a 3xx status code
func (o *GetRoutingPredictorModelsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models unauthorized response has a 4xx status code
func (o *GetRoutingPredictorModelsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor models unauthorized response has a 5xx status code
func (o *GetRoutingPredictorModelsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor models unauthorized response a status code equal to that given
func (o *GetRoutingPredictorModelsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetRoutingPredictorModelsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingPredictorModelsUnauthorized) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRoutingPredictorModelsUnauthorized) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsForbidden creates a GetRoutingPredictorModelsForbidden with default headers values
func NewGetRoutingPredictorModelsForbidden() *GetRoutingPredictorModelsForbidden {
	return &GetRoutingPredictorModelsForbidden{}
}

/*
GetRoutingPredictorModelsForbidden describes a response with status code 403, with default header values.

You are not authorized to perform the requested action.
*/
type GetRoutingPredictorModelsForbidden struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models forbidden response has a 2xx status code
func (o *GetRoutingPredictorModelsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models forbidden response has a 3xx status code
func (o *GetRoutingPredictorModelsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models forbidden response has a 4xx status code
func (o *GetRoutingPredictorModelsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor models forbidden response has a 5xx status code
func (o *GetRoutingPredictorModelsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor models forbidden response a status code equal to that given
func (o *GetRoutingPredictorModelsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetRoutingPredictorModelsForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingPredictorModelsForbidden) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsForbidden  %+v", 403, o.Payload)
}

func (o *GetRoutingPredictorModelsForbidden) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsNotFound creates a GetRoutingPredictorModelsNotFound with default headers values
func NewGetRoutingPredictorModelsNotFound() *GetRoutingPredictorModelsNotFound {
	return &GetRoutingPredictorModelsNotFound{}
}

/*
GetRoutingPredictorModelsNotFound describes a response with status code 404, with default header values.

The requested resource was not found.
*/
type GetRoutingPredictorModelsNotFound struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models not found response has a 2xx status code
func (o *GetRoutingPredictorModelsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models not found response has a 3xx status code
func (o *GetRoutingPredictorModelsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models not found response has a 4xx status code
func (o *GetRoutingPredictorModelsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor models not found response has a 5xx status code
func (o *GetRoutingPredictorModelsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor models not found response a status code equal to that given
func (o *GetRoutingPredictorModelsNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetRoutingPredictorModelsNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingPredictorModelsNotFound) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsNotFound  %+v", 404, o.Payload)
}

func (o *GetRoutingPredictorModelsNotFound) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsRequestTimeout creates a GetRoutingPredictorModelsRequestTimeout with default headers values
func NewGetRoutingPredictorModelsRequestTimeout() *GetRoutingPredictorModelsRequestTimeout {
	return &GetRoutingPredictorModelsRequestTimeout{}
}

/*
GetRoutingPredictorModelsRequestTimeout describes a response with status code 408, with default header values.

The client did not produce a request within the server timeout limit. This can be caused by a slow network connection and/or large payloads.
*/
type GetRoutingPredictorModelsRequestTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models request timeout response has a 2xx status code
func (o *GetRoutingPredictorModelsRequestTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models request timeout response has a 3xx status code
func (o *GetRoutingPredictorModelsRequestTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models request timeout response has a 4xx status code
func (o *GetRoutingPredictorModelsRequestTimeout) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor models request timeout response has a 5xx status code
func (o *GetRoutingPredictorModelsRequestTimeout) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor models request timeout response a status code equal to that given
func (o *GetRoutingPredictorModelsRequestTimeout) IsCode(code int) bool {
	return code == 408
}

func (o *GetRoutingPredictorModelsRequestTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingPredictorModelsRequestTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsRequestTimeout  %+v", 408, o.Payload)
}

func (o *GetRoutingPredictorModelsRequestTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsRequestEntityTooLarge creates a GetRoutingPredictorModelsRequestEntityTooLarge with default headers values
func NewGetRoutingPredictorModelsRequestEntityTooLarge() *GetRoutingPredictorModelsRequestEntityTooLarge {
	return &GetRoutingPredictorModelsRequestEntityTooLarge{}
}

/*
GetRoutingPredictorModelsRequestEntityTooLarge describes a response with status code 413, with default header values.

The request is over the size limit. Maximum bytes: %s
*/
type GetRoutingPredictorModelsRequestEntityTooLarge struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models request entity too large response has a 2xx status code
func (o *GetRoutingPredictorModelsRequestEntityTooLarge) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models request entity too large response has a 3xx status code
func (o *GetRoutingPredictorModelsRequestEntityTooLarge) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models request entity too large response has a 4xx status code
func (o *GetRoutingPredictorModelsRequestEntityTooLarge) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor models request entity too large response has a 5xx status code
func (o *GetRoutingPredictorModelsRequestEntityTooLarge) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor models request entity too large response a status code equal to that given
func (o *GetRoutingPredictorModelsRequestEntityTooLarge) IsCode(code int) bool {
	return code == 413
}

func (o *GetRoutingPredictorModelsRequestEntityTooLarge) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingPredictorModelsRequestEntityTooLarge) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsRequestEntityTooLarge  %+v", 413, o.Payload)
}

func (o *GetRoutingPredictorModelsRequestEntityTooLarge) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsRequestEntityTooLarge) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsUnsupportedMediaType creates a GetRoutingPredictorModelsUnsupportedMediaType with default headers values
func NewGetRoutingPredictorModelsUnsupportedMediaType() *GetRoutingPredictorModelsUnsupportedMediaType {
	return &GetRoutingPredictorModelsUnsupportedMediaType{}
}

/*
GetRoutingPredictorModelsUnsupportedMediaType describes a response with status code 415, with default header values.

Unsupported Media Type - Unsupported or incorrect media type, such as an incorrect Content-Type value in the header.
*/
type GetRoutingPredictorModelsUnsupportedMediaType struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models unsupported media type response has a 2xx status code
func (o *GetRoutingPredictorModelsUnsupportedMediaType) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models unsupported media type response has a 3xx status code
func (o *GetRoutingPredictorModelsUnsupportedMediaType) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models unsupported media type response has a 4xx status code
func (o *GetRoutingPredictorModelsUnsupportedMediaType) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor models unsupported media type response has a 5xx status code
func (o *GetRoutingPredictorModelsUnsupportedMediaType) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor models unsupported media type response a status code equal to that given
func (o *GetRoutingPredictorModelsUnsupportedMediaType) IsCode(code int) bool {
	return code == 415
}

func (o *GetRoutingPredictorModelsUnsupportedMediaType) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingPredictorModelsUnsupportedMediaType) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsUnsupportedMediaType  %+v", 415, o.Payload)
}

func (o *GetRoutingPredictorModelsUnsupportedMediaType) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsUnsupportedMediaType) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsTooManyRequests creates a GetRoutingPredictorModelsTooManyRequests with default headers values
func NewGetRoutingPredictorModelsTooManyRequests() *GetRoutingPredictorModelsTooManyRequests {
	return &GetRoutingPredictorModelsTooManyRequests{}
}

/*
GetRoutingPredictorModelsTooManyRequests describes a response with status code 429, with default header values.

Rate limit exceeded the maximum. Retry the request in [%s] seconds
*/
type GetRoutingPredictorModelsTooManyRequests struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models too many requests response has a 2xx status code
func (o *GetRoutingPredictorModelsTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models too many requests response has a 3xx status code
func (o *GetRoutingPredictorModelsTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models too many requests response has a 4xx status code
func (o *GetRoutingPredictorModelsTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get routing predictor models too many requests response has a 5xx status code
func (o *GetRoutingPredictorModelsTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get routing predictor models too many requests response a status code equal to that given
func (o *GetRoutingPredictorModelsTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetRoutingPredictorModelsTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingPredictorModelsTooManyRequests) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsTooManyRequests  %+v", 429, o.Payload)
}

func (o *GetRoutingPredictorModelsTooManyRequests) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsInternalServerError creates a GetRoutingPredictorModelsInternalServerError with default headers values
func NewGetRoutingPredictorModelsInternalServerError() *GetRoutingPredictorModelsInternalServerError {
	return &GetRoutingPredictorModelsInternalServerError{}
}

/*
GetRoutingPredictorModelsInternalServerError describes a response with status code 500, with default header values.

The server encountered an unexpected condition which prevented it from fulfilling the request.
*/
type GetRoutingPredictorModelsInternalServerError struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models internal server error response has a 2xx status code
func (o *GetRoutingPredictorModelsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models internal server error response has a 3xx status code
func (o *GetRoutingPredictorModelsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models internal server error response has a 4xx status code
func (o *GetRoutingPredictorModelsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing predictor models internal server error response has a 5xx status code
func (o *GetRoutingPredictorModelsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing predictor models internal server error response a status code equal to that given
func (o *GetRoutingPredictorModelsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetRoutingPredictorModelsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingPredictorModelsInternalServerError) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRoutingPredictorModelsInternalServerError) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsServiceUnavailable creates a GetRoutingPredictorModelsServiceUnavailable with default headers values
func NewGetRoutingPredictorModelsServiceUnavailable() *GetRoutingPredictorModelsServiceUnavailable {
	return &GetRoutingPredictorModelsServiceUnavailable{}
}

/*
GetRoutingPredictorModelsServiceUnavailable describes a response with status code 503, with default header values.

Service Unavailable - The server is currently unavailable (because it is overloaded or down for maintenance).
*/
type GetRoutingPredictorModelsServiceUnavailable struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models service unavailable response has a 2xx status code
func (o *GetRoutingPredictorModelsServiceUnavailable) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models service unavailable response has a 3xx status code
func (o *GetRoutingPredictorModelsServiceUnavailable) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models service unavailable response has a 4xx status code
func (o *GetRoutingPredictorModelsServiceUnavailable) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing predictor models service unavailable response has a 5xx status code
func (o *GetRoutingPredictorModelsServiceUnavailable) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing predictor models service unavailable response a status code equal to that given
func (o *GetRoutingPredictorModelsServiceUnavailable) IsCode(code int) bool {
	return code == 503
}

func (o *GetRoutingPredictorModelsServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingPredictorModelsServiceUnavailable) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsServiceUnavailable  %+v", 503, o.Payload)
}

func (o *GetRoutingPredictorModelsServiceUnavailable) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoutingPredictorModelsGatewayTimeout creates a GetRoutingPredictorModelsGatewayTimeout with default headers values
func NewGetRoutingPredictorModelsGatewayTimeout() *GetRoutingPredictorModelsGatewayTimeout {
	return &GetRoutingPredictorModelsGatewayTimeout{}
}

/*
GetRoutingPredictorModelsGatewayTimeout describes a response with status code 504, with default header values.

The request timed out.
*/
type GetRoutingPredictorModelsGatewayTimeout struct {
	Payload *models.ErrorBody
}

// IsSuccess returns true when this get routing predictor models gateway timeout response has a 2xx status code
func (o *GetRoutingPredictorModelsGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get routing predictor models gateway timeout response has a 3xx status code
func (o *GetRoutingPredictorModelsGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get routing predictor models gateway timeout response has a 4xx status code
func (o *GetRoutingPredictorModelsGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this get routing predictor models gateway timeout response has a 5xx status code
func (o *GetRoutingPredictorModelsGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this get routing predictor models gateway timeout response a status code equal to that given
func (o *GetRoutingPredictorModelsGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

func (o *GetRoutingPredictorModelsGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingPredictorModelsGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /api/v2/routing/predictors/{predictorId}/models][%d] getRoutingPredictorModelsGatewayTimeout  %+v", 504, o.Payload)
}

func (o *GetRoutingPredictorModelsGatewayTimeout) GetPayload() *models.ErrorBody {
	return o.Payload
}

func (o *GetRoutingPredictorModelsGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
