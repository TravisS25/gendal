// Package ischema contains the types for schema 'information_schema'.
package ischema

import "github.com/turnkey-commerce/gendal/examples/pgcatalog/pgtypes"

// Code generated by xo. DO NOT EDIT.

// ElementType represents a row from 'information_schema.element_types'.
type ElementType struct {
	ObjectCatalog            pgtypes.SQLIdentifier  `json:"object_catalog"`             // object_catalog
	ObjectSchema             pgtypes.SQLIdentifier  `json:"object_schema"`              // object_schema
	ObjectName               pgtypes.SQLIdentifier  `json:"object_name"`                // object_name
	ObjectType               pgtypes.CharacterData  `json:"object_type"`                // object_type
	CollectionTypeIdentifier pgtypes.SQLIdentifier  `json:"collection_type_identifier"` // collection_type_identifier
	DataType                 pgtypes.CharacterData  `json:"data_type"`                  // data_type
	CharacterMaximumLength   pgtypes.CardinalNumber `json:"character_maximum_length"`   // character_maximum_length
	CharacterOctetLength     pgtypes.CardinalNumber `json:"character_octet_length"`     // character_octet_length
	CharacterSetCatalog      pgtypes.SQLIdentifier  `json:"character_set_catalog"`      // character_set_catalog
	CharacterSetSchema       pgtypes.SQLIdentifier  `json:"character_set_schema"`       // character_set_schema
	CharacterSetName         pgtypes.SQLIdentifier  `json:"character_set_name"`         // character_set_name
	CollationCatalog         pgtypes.SQLIdentifier  `json:"collation_catalog"`          // collation_catalog
	CollationSchema          pgtypes.SQLIdentifier  `json:"collation_schema"`           // collation_schema
	CollationName            pgtypes.SQLIdentifier  `json:"collation_name"`             // collation_name
	NumericPrecision         pgtypes.CardinalNumber `json:"numeric_precision"`          // numeric_precision
	NumericPrecisionRadix    pgtypes.CardinalNumber `json:"numeric_precision_radix"`    // numeric_precision_radix
	NumericScale             pgtypes.CardinalNumber `json:"numeric_scale"`              // numeric_scale
	DatetimePrecision        pgtypes.CardinalNumber `json:"datetime_precision"`         // datetime_precision
	IntervalType             pgtypes.CharacterData  `json:"interval_type"`              // interval_type
	IntervalPrecision        pgtypes.CardinalNumber `json:"interval_precision"`         // interval_precision
	DomainDefault            pgtypes.CharacterData  `json:"domain_default"`             // domain_default
	UdtCatalog               pgtypes.SQLIdentifier  `json:"udt_catalog"`                // udt_catalog
	UdtSchema                pgtypes.SQLIdentifier  `json:"udt_schema"`                 // udt_schema
	UdtName                  pgtypes.SQLIdentifier  `json:"udt_name"`                   // udt_name
	ScopeCatalog             pgtypes.SQLIdentifier  `json:"scope_catalog"`              // scope_catalog
	ScopeSchema              pgtypes.SQLIdentifier  `json:"scope_schema"`               // scope_schema
	ScopeName                pgtypes.SQLIdentifier  `json:"scope_name"`                 // scope_name
	MaximumCardinality       pgtypes.CardinalNumber `json:"maximum_cardinality"`        // maximum_cardinality
	DtdIdentifier            pgtypes.SQLIdentifier  `json:"dtd_identifier"`             // dtd_identifier
}
