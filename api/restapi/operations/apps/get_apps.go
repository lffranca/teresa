package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
	"github.com/luizalabs/paas/api/models"
)

// GetAppsHandlerFunc turns a function with the right signature into a get apps handler
type GetAppsHandlerFunc func(GetAppsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAppsHandlerFunc) Handle(params GetAppsParams) middleware.Responder {
	return fn(params)
}

// GetAppsHandler interface for that can handle valid get apps params
type GetAppsHandler interface {
	Handle(GetAppsParams) middleware.Responder
}

// NewGetApps creates a new http.Handler for the get apps operation
func NewGetApps(ctx *middleware.Context, handler GetAppsHandler) *GetApps {
	return &GetApps{Context: ctx, Handler: handler}
}

/*GetApps swagger:route GET /teams/{team_id}/apps apps getApps

Get team apps

Get team apps

*/
type GetApps struct {
	Context *middleware.Context
	Handler GetAppsHandler
}

func (o *GetApps) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetAppsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}

/*GetAppsOKBodyBody get apps o k body body

swagger:model GetAppsOKBodyBody
*/
type GetAppsOKBodyBody struct {
	models.Pagination

	/* items

	Required: true
	*/
	Items []*models.App `json:"items"`
}

// Validate validates this get apps o k body body
func (o *GetAppsOKBodyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.Pagination.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateItems(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetAppsOKBodyBody) validateItems(formats strfmt.Registry) error {

	if err := validate.Required("getAppsOK"+"."+"items", "body", o.Items); err != nil {
		return err
	}

	return nil
}