// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/type/postal_address.proto

package google_type

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Represents a postal address, e.g. for postal delivery or payments addresses.
// Given a postal address, a postal service can deliver items to a premise, P.O.
// Box or similar.
// It is not intended to model geographical locations (roads, towns,
// mountains).
//
// In typical usage an address would be created via user input or from importing
// existing data, depending on the type of process.
//
// Advice on address input / editing:
//  - Use an i18n-ready address widget such as
//    https://github.com/googlei18n/libaddressinput)
// - Users should not be presented with UI elements for input or editing of
//   fields outside countries where that field is used.
//
// For more guidance on how to use this schema, please see:
// https://support.google.com/business/answer/6397478
type PostalAddress struct {
	// The schema revision of the `PostalAddress`.
	// All new revisions **must** be backward compatible with old revisions.
	Revision int32 `protobuf:"varint,1,opt,name=revision,proto3" json:"revision,omitempty"`
	// Required. CLDR region code of the country/region of the address. This
	// is never inferred and it is up to the user to ensure the value is
	// correct. See http://cldr.unicode.org/ and
	// http://www.unicode.org/cldr/charts/30/supplemental/territory_information.html
	// for details. Example: "CH" for Switzerland.
	RegionCode string `protobuf:"bytes,2,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// Optional. BCP-47 language code of the contents of this address (if
	// known). This is often the UI language of the input form or is expected
	// to match one of the languages used in the address' country/region, or their
	// transliterated equivalents.
	// This can affect formatting in certain countries, but is not critical
	// to the correctness of the data and will never affect any validation or
	// other non-formatting related operations.
	//
	// If this value is not known, it should be omitted (rather than specifying a
	// possibly incorrect default).
	//
	// Examples: "zh-Hant", "ja", "ja-Latn", "en".
	LanguageCode string `protobuf:"bytes,3,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional. Postal code of the address. Not all countries use or require
	// postal codes to be present, but where they are used, they may trigger
	// additional validation with other parts of the address (e.g. state/zip
	// validation in the U.S.A.).
	PostalCode string `protobuf:"bytes,4,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	// Optional. Additional, country-specific, sorting code. This is not used
	// in most regions. Where it is used, the value is either a string like
	// "CEDEX", optionally followed by a number (e.g. "CEDEX 7"), or just a number
	// alone, representing the "sector code" (Jamaica), "delivery area indicator"
	// (Malawi) or "post office indicator" (e.g. Côte d'Ivoire).
	SortingCode string `protobuf:"bytes,5,opt,name=sorting_code,json=sortingCode,proto3" json:"sorting_code,omitempty"`
	// Optional. Highest administrative subdivision which is used for postal
	// addresses of a country or region.
	// For example, this can be a state, a province, an oblast, or a prefecture.
	// Specifically, for Spain this is the province and not the autonomous
	// community (e.g. "Barcelona" and not "Catalonia").
	// Many countries don't use an administrative area in postal addresses. E.g.
	// in Switzerland this should be left unpopulated.
	AdministrativeArea string `protobuf:"bytes,6,opt,name=administrative_area,json=administrativeArea,proto3" json:"administrative_area,omitempty"`
	// Optional. Generally refers to the city/town portion of the address.
	// Examples: US city, IT comune, UK post town.
	// In regions of the world where localities are not well defined or do not fit
	// into this structure well, leave locality empty and use address_lines.
	Locality string `protobuf:"bytes,7,opt,name=locality,proto3" json:"locality,omitempty"`
	// Optional. Sublocality of the address.
	// For example, this can be neighborhoods, boroughs, districts.
	Sublocality string `protobuf:"bytes,8,opt,name=sublocality,proto3" json:"sublocality,omitempty"`
	// Unstructured address lines describing the lower levels of an address.
	//
	// Because values in address_lines do not have type information and may
	// sometimes contain multiple values in a single field (e.g.
	// "Austin, TX"), it is important that the line order is clear. The order of
	// address lines should be "envelope order" for the country/region of the
	// address. In places where this can vary (e.g. Japan), address_language is
	// used to make it explicit (e.g. "ja" for large-to-small ordering and
	// "ja-Latn" or "en" for small-to-large). This way, the most specific line of
	// an address can be selected based on the language.
	//
	// The minimum permitted structural representation of an address consists
	// of a region_code with all remaining information placed in the
	// address_lines. It would be possible to format such an address very
	// approximately without geocoding, but no semantic reasoning could be
	// made about any of the address components until it was at least
	// partially resolved.
	//
	// Creating an address only containing a region_code and address_lines, and
	// then geocoding is the recommended way to handle completely unstructured
	// addresses (as opposed to guessing which parts of the address should be
	// localities or administrative areas).
	AddressLines []string `protobuf:"bytes,9,rep,name=address_lines,json=addressLines" json:"address_lines,omitempty"`
	// Optional. The recipient at the address.
	// This field may, under certain circumstances, contain multiline information.
	// For example, it might contain "care of" information.
	Recipients []string `protobuf:"bytes,10,rep,name=recipients" json:"recipients,omitempty"`
	// Optional. The name of the organization at the address.
	Organization string `protobuf:"bytes,11,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (m *PostalAddress) Reset()                    { *m = PostalAddress{} }
func (m *PostalAddress) String() string            { return proto.CompactTextString(m) }
func (*PostalAddress) ProtoMessage()               {}
func (*PostalAddress) Descriptor() ([]byte, []int) { return fileDescriptorPostalAddress, []int{0} }

func (m *PostalAddress) GetRevision() int32 {
	if m != nil {
		return m.Revision
	}
	return 0
}

func (m *PostalAddress) GetRegionCode() string {
	if m != nil {
		return m.RegionCode
	}
	return ""
}

func (m *PostalAddress) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *PostalAddress) GetPostalCode() string {
	if m != nil {
		return m.PostalCode
	}
	return ""
}

func (m *PostalAddress) GetSortingCode() string {
	if m != nil {
		return m.SortingCode
	}
	return ""
}

func (m *PostalAddress) GetAdministrativeArea() string {
	if m != nil {
		return m.AdministrativeArea
	}
	return ""
}

func (m *PostalAddress) GetLocality() string {
	if m != nil {
		return m.Locality
	}
	return ""
}

func (m *PostalAddress) GetSublocality() string {
	if m != nil {
		return m.Sublocality
	}
	return ""
}

func (m *PostalAddress) GetAddressLines() []string {
	if m != nil {
		return m.AddressLines
	}
	return nil
}

func (m *PostalAddress) GetRecipients() []string {
	if m != nil {
		return m.Recipients
	}
	return nil
}

func (m *PostalAddress) GetOrganization() string {
	if m != nil {
		return m.Organization
	}
	return ""
}

func init() {
	proto.RegisterType((*PostalAddress)(nil), "google.type.PostalAddress")
}

func init() { proto.RegisterFile("google/type/postal_address.proto", fileDescriptorPostalAddress) }

var fileDescriptorPostalAddress = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x54, 0x91, 0xb1, 0x6e, 0xc2, 0x30,
	0x10, 0x86, 0x15, 0x52, 0x28, 0x5c, 0x40, 0x95, 0xdc, 0x25, 0xea, 0x50, 0x52, 0xba, 0xb0, 0x14,
	0x86, 0x3e, 0x01, 0x74, 0xe8, 0xd2, 0x21, 0x42, 0xdd, 0x23, 0x93, 0x9c, 0x2c, 0x4b, 0xc6, 0x17,
	0xd9, 0x06, 0x89, 0xbe, 0x43, 0x5f, 0xa2, 0x4f, 0x5a, 0xd9, 0x4e, 0x69, 0x18, 0xef, 0xbb, 0x4f,
	0x89, 0xef, 0xff, 0xa1, 0x10, 0x44, 0x42, 0xe1, 0xda, 0x9d, 0x5b, 0x5c, 0xb7, 0x64, 0x1d, 0x57,
	0x15, 0x6f, 0x1a, 0x83, 0xd6, 0xae, 0x5a, 0x43, 0x8e, 0x58, 0x16, 0x8d, 0x95, 0x37, 0x16, 0xdf,
	0x29, 0xcc, 0xca, 0x60, 0x6d, 0xa2, 0xc4, 0x1e, 0x60, 0x6c, 0xf0, 0x24, 0xad, 0x24, 0x9d, 0x27,
	0x45, 0xb2, 0x1c, 0xee, 0x2e, 0x33, 0x9b, 0x43, 0x66, 0x50, 0x48, 0xd2, 0x55, 0x4d, 0x0d, 0xe6,
	0x83, 0x22, 0x59, 0x4e, 0x76, 0x10, 0xd1, 0x1b, 0x35, 0xc8, 0x9e, 0x61, 0xa6, 0xb8, 0x16, 0x47,
	0x2e, 0x30, 0x2a, 0x69, 0x50, 0xa6, 0x7f, 0x30, 0x48, 0x73, 0xc8, 0xba, 0x87, 0x05, 0xe5, 0x26,
	0x7e, 0x25, 0xa2, 0x20, 0x3c, 0xc1, 0xd4, 0x92, 0x71, 0x52, 0x8b, 0x68, 0x0c, 0x83, 0x91, 0x75,
	0x2c, 0x28, 0x6b, 0xb8, 0xe7, 0xcd, 0x41, 0x6a, 0x69, 0x9d, 0xe1, 0x4e, 0x9e, 0xb0, 0xe2, 0x06,
	0x79, 0x3e, 0x0a, 0x26, 0xbb, 0x5e, 0x6d, 0x0c, 0x72, 0x7f, 0x96, 0xa2, 0x9a, 0x2b, 0xe9, 0xce,
	0xf9, 0x6d, 0xb0, 0x2e, 0x33, 0x2b, 0x20, 0xb3, 0xc7, 0xfd, 0x65, 0x3d, 0xee, 0x7e, 0xf7, 0x8f,
	0xfc, 0x5d, 0x5d, 0x88, 0x95, 0x92, 0x1a, 0x6d, 0x3e, 0x29, 0x52, 0x7f, 0x57, 0x07, 0x3f, 0x3c,
	0x63, 0x8f, 0x00, 0x06, 0x6b, 0xd9, 0x4a, 0xd4, 0xce, 0xe6, 0x10, 0x8c, 0x1e, 0x61, 0x0b, 0x98,
	0x92, 0x11, 0x5c, 0xcb, 0x2f, 0xee, 0x7c, 0xba, 0x59, 0xcc, 0xa6, 0xcf, 0xb6, 0x2f, 0x70, 0x57,
	0xd3, 0x61, 0xd5, 0xab, 0x68, 0xcb, 0xae, 0xfa, 0x29, 0x7d, 0x87, 0x65, 0xf2, 0x33, 0x48, 0xdf,
	0x3f, 0xcb, 0xfd, 0x28, 0x54, 0xfa, 0xfa, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xfb, 0x0d, 0x4a, 0xc4,
	0xf6, 0x01, 0x00, 0x00,
}
