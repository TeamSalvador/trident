// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MetroclusterDiagConnectionDetails metrocluster diag connection details
//
// swagger:model metrocluster_diag_connection_details
type MetroclusterDiagConnectionDetails struct {

	// cluster
	Cluster *MetroclusterDiagConnectionDetailsInlineCluster `json:"cluster,omitempty"`

	// metrocluster diag connection details inline connections
	// Read Only: true
	MetroclusterDiagConnectionDetailsInlineConnections []*MetroclusterDiagConnection `json:"connections,omitempty"`

	// node
	Node *MetroclusterDiagConnectionDetailsInlineNode `json:"node,omitempty"`
}

// Validate validates this metrocluster diag connection details
func (m *MetroclusterDiagConnectionDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCluster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetroclusterDiagConnectionDetailsInlineConnections(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetails) validateCluster(formats strfmt.Registry) error {
	if swag.IsZero(m.Cluster) { // not required
		return nil
	}

	if m.Cluster != nil {
		if err := m.Cluster.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterDiagConnectionDetails) validateMetroclusterDiagConnectionDetailsInlineConnections(formats strfmt.Registry) error {
	if swag.IsZero(m.MetroclusterDiagConnectionDetailsInlineConnections) { // not required
		return nil
	}

	for i := 0; i < len(m.MetroclusterDiagConnectionDetailsInlineConnections); i++ {
		if swag.IsZero(m.MetroclusterDiagConnectionDetailsInlineConnections[i]) { // not required
			continue
		}

		if m.MetroclusterDiagConnectionDetailsInlineConnections[i] != nil {
			if err := m.MetroclusterDiagConnectionDetailsInlineConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MetroclusterDiagConnectionDetails) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster diag connection details based on the context it is used
func (m *MetroclusterDiagConnectionDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCluster(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateMetroclusterDiagConnectionDetailsInlineConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetails) contextValidateCluster(ctx context.Context, formats strfmt.Registry) error {

	if m.Cluster != nil {
		if err := m.Cluster.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterDiagConnectionDetails) contextValidateMetroclusterDiagConnectionDetailsInlineConnections(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connections", "body", []*MetroclusterDiagConnection(m.MetroclusterDiagConnectionDetailsInlineConnections)); err != nil {
		return err
	}

	for i := 0; i < len(m.MetroclusterDiagConnectionDetailsInlineConnections); i++ {

		if m.MetroclusterDiagConnectionDetailsInlineConnections[i] != nil {
			if err := m.MetroclusterDiagConnectionDetailsInlineConnections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("connections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *MetroclusterDiagConnectionDetails) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetails) UnmarshalBinary(b []byte) error {
	var res MetroclusterDiagConnectionDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterDiagConnectionDetailsInlineCluster metrocluster diag connection details inline cluster
//
// swagger:model metrocluster_diag_connection_details_inline_cluster
type MetroclusterDiagConnectionDetailsInlineCluster struct {

	// links
	Links *MetroclusterDiagConnectionDetailsInlineClusterInlineLinks `json:"_links,omitempty"`

	// name
	// Example: cluster1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Format: uuid
	UUID *strfmt.UUID `json:"uuid,omitempty"`
}

// Validate validates this metrocluster diag connection details inline cluster
func (m *MetroclusterDiagConnectionDetailsInlineCluster) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUUID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetailsInlineCluster) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *MetroclusterDiagConnectionDetailsInlineCluster) validateUUID(formats strfmt.Registry) error {
	if swag.IsZero(m.UUID) { // not required
		return nil
	}

	if err := validate.FormatOf("cluster"+"."+"uuid", "body", "uuid", m.UUID.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this metrocluster diag connection details inline cluster based on the context it is used
func (m *MetroclusterDiagConnectionDetailsInlineCluster) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetailsInlineCluster) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetailsInlineCluster) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetailsInlineCluster) UnmarshalBinary(b []byte) error {
	var res MetroclusterDiagConnectionDetailsInlineCluster
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterDiagConnectionDetailsInlineClusterInlineLinks metrocluster diag connection details inline cluster inline links
//
// swagger:model metrocluster_diag_connection_details_inline_cluster_inline__links
type MetroclusterDiagConnectionDetailsInlineClusterInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this metrocluster diag connection details inline cluster inline links
func (m *MetroclusterDiagConnectionDetailsInlineClusterInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetailsInlineClusterInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster diag connection details inline cluster inline links based on the context it is used
func (m *MetroclusterDiagConnectionDetailsInlineClusterInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetailsInlineClusterInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cluster" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetailsInlineClusterInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetailsInlineClusterInlineLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterDiagConnectionDetailsInlineClusterInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterDiagConnectionDetailsInlineNode metrocluster diag connection details inline node
//
// swagger:model metrocluster_diag_connection_details_inline_node
type MetroclusterDiagConnectionDetailsInlineNode struct {

	// links
	Links *MetroclusterDiagConnectionDetailsInlineNodeInlineLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name *string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID *string `json:"uuid,omitempty"`
}

// Validate validates this metrocluster diag connection details inline node
func (m *MetroclusterDiagConnectionDetailsInlineNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetailsInlineNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster diag connection details inline node based on the context it is used
func (m *MetroclusterDiagConnectionDetailsInlineNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetailsInlineNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetailsInlineNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetailsInlineNode) UnmarshalBinary(b []byte) error {
	var res MetroclusterDiagConnectionDetailsInlineNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MetroclusterDiagConnectionDetailsInlineNodeInlineLinks metrocluster diag connection details inline node inline links
//
// swagger:model metrocluster_diag_connection_details_inline_node_inline__links
type MetroclusterDiagConnectionDetailsInlineNodeInlineLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this metrocluster diag connection details inline node inline links
func (m *MetroclusterDiagConnectionDetailsInlineNodeInlineLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetailsInlineNodeInlineLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this metrocluster diag connection details inline node inline links based on the context it is used
func (m *MetroclusterDiagConnectionDetailsInlineNodeInlineLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MetroclusterDiagConnectionDetailsInlineNodeInlineLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetailsInlineNodeInlineLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MetroclusterDiagConnectionDetailsInlineNodeInlineLinks) UnmarshalBinary(b []byte) error {
	var res MetroclusterDiagConnectionDetailsInlineNodeInlineLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}