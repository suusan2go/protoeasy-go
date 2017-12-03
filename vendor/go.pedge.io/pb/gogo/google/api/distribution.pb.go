// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: google/api/distribution.proto

package google_api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.pedge.io/pb/gogo/google/protobuf"
import _ "go.pedge.io/pb/gogo/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Distribution contains summary statistics for a population of values and,
// optionally, a histogram representing the distribution of those values across
// a specified set of histogram buckets.
//
// The summary statistics are the count, mean, sum of the squared deviation from
// the mean, the minimum, and the maximum of the set of population of values.
//
// The histogram is based on a sequence of buckets and gives a count of values
// that fall into each bucket.  The boundaries of the buckets are given either
// explicitly or by specifying parameters for a method of computing them
// (buckets of fixed width or buckets of exponentially increasing width).
//
// Although it is not forbidden, it is generally a bad idea to include
// non-finite values (infinities or NaNs) in the population of values, as this
// will render the `mean` and `sum_of_squared_deviation` fields meaningless.
type Distribution struct {
	// The number of values in the population. Must be non-negative.
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// The arithmetic mean of the values in the population. If `count` is zero
	// then this field must be zero.
	Mean float64 `protobuf:"fixed64,2,opt,name=mean,proto3" json:"mean,omitempty"`
	// The sum of squared deviations from the mean of the values in the
	// population.  For values x_i this is:
	//
	//     Sum[i=1..n]((x_i - mean)^2)
	//
	// Knuth, "The Art of Computer Programming", Vol. 2, page 323, 3rd edition
	// describes Welford's method for accumulating this sum in one pass.
	//
	// If `count` is zero then this field must be zero.
	SumOfSquaredDeviation float64 `protobuf:"fixed64,3,opt,name=sum_of_squared_deviation,json=sumOfSquaredDeviation,proto3" json:"sum_of_squared_deviation,omitempty"`
	// If specified, contains the range of the population values. The field
	// must not be present if the `count` is zero.
	Range *Distribution_Range `protobuf:"bytes,4,opt,name=range" json:"range,omitempty"`
	// Defines the histogram bucket boundaries.
	BucketOptions *Distribution_BucketOptions `protobuf:"bytes,6,opt,name=bucket_options,json=bucketOptions" json:"bucket_options,omitempty"`
	// If `bucket_options` is given, then the sum of the values in `bucket_counts`
	// must equal the value in `count`.  If `bucket_options` is not given, no
	// `bucket_counts` fields may be given.
	//
	// Bucket counts are given in order under the numbering scheme described
	// above (the underflow bucket has number 0; the finite buckets, if any,
	// have numbers 1 through N-2; the overflow bucket has number N-1).
	//
	// The size of `bucket_counts` must be no greater than N as defined in
	// `bucket_options`.
	//
	// Any suffix of trailing zero bucket_count fields may be omitted.
	BucketCounts []int64 `protobuf:"varint,7,rep,packed,name=bucket_counts,json=bucketCounts" json:"bucket_counts,omitempty"`
}

func (m *Distribution) Reset()                    { *m = Distribution{} }
func (m *Distribution) String() string            { return proto.CompactTextString(m) }
func (*Distribution) ProtoMessage()               {}
func (*Distribution) Descriptor() ([]byte, []int) { return fileDescriptorDistribution, []int{0} }

func (m *Distribution) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Distribution) GetMean() float64 {
	if m != nil {
		return m.Mean
	}
	return 0
}

func (m *Distribution) GetSumOfSquaredDeviation() float64 {
	if m != nil {
		return m.SumOfSquaredDeviation
	}
	return 0
}

func (m *Distribution) GetRange() *Distribution_Range {
	if m != nil {
		return m.Range
	}
	return nil
}

func (m *Distribution) GetBucketOptions() *Distribution_BucketOptions {
	if m != nil {
		return m.BucketOptions
	}
	return nil
}

func (m *Distribution) GetBucketCounts() []int64 {
	if m != nil {
		return m.BucketCounts
	}
	return nil
}

// The range of the population values.
type Distribution_Range struct {
	// The minimum of the population values.
	Min float64 `protobuf:"fixed64,1,opt,name=min,proto3" json:"min,omitempty"`
	// The maximum of the population values.
	Max float64 `protobuf:"fixed64,2,opt,name=max,proto3" json:"max,omitempty"`
}

func (m *Distribution_Range) Reset()         { *m = Distribution_Range{} }
func (m *Distribution_Range) String() string { return proto.CompactTextString(m) }
func (*Distribution_Range) ProtoMessage()    {}
func (*Distribution_Range) Descriptor() ([]byte, []int) {
	return fileDescriptorDistribution, []int{0, 0}
}

func (m *Distribution_Range) GetMin() float64 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *Distribution_Range) GetMax() float64 {
	if m != nil {
		return m.Max
	}
	return 0
}

// A Distribution may optionally contain a histogram of the values in the
// population.  The histogram is given in `bucket_counts` as counts of values
// that fall into one of a sequence of non-overlapping buckets.  The sequence
// of buckets is described by `bucket_options`.
//
// A bucket specifies an inclusive lower bound and exclusive upper bound for
// the values that are counted for that bucket.  The upper bound of a bucket
// is strictly greater than the lower bound.
//
// The sequence of N buckets for a Distribution consists of an underflow
// bucket (number 0), zero or more finite buckets (number 1 through N - 2) and
// an overflow bucket (number N - 1).  The buckets are contiguous:  the lower
// bound of bucket i (i > 0) is the same as the upper bound of bucket i - 1.
// The buckets span the whole range of finite values: lower bound of the
// underflow bucket is -infinity and the upper bound of the overflow bucket is
// +infinity.  The finite buckets are so-called because both bounds are
// finite.
//
// `BucketOptions` describes bucket boundaries in one of three ways.  Two
// describe the boundaries by giving parameters for a formula to generate
// boundaries and one gives the bucket boundaries explicitly.
//
// If `bucket_boundaries` is not given, then no `bucket_counts` may be given.
type Distribution_BucketOptions struct {
	// Exactly one of these three fields must be set.
	//
	// Types that are valid to be assigned to Options:
	//	*Distribution_BucketOptions_LinearBuckets
	//	*Distribution_BucketOptions_ExponentialBuckets
	//	*Distribution_BucketOptions_ExplicitBuckets
	Options isDistribution_BucketOptions_Options `protobuf_oneof:"options"`
}

func (m *Distribution_BucketOptions) Reset()         { *m = Distribution_BucketOptions{} }
func (m *Distribution_BucketOptions) String() string { return proto.CompactTextString(m) }
func (*Distribution_BucketOptions) ProtoMessage()    {}
func (*Distribution_BucketOptions) Descriptor() ([]byte, []int) {
	return fileDescriptorDistribution, []int{0, 1}
}

type isDistribution_BucketOptions_Options interface {
	isDistribution_BucketOptions_Options()
}

type Distribution_BucketOptions_LinearBuckets struct {
	LinearBuckets *Distribution_BucketOptions_Linear `protobuf:"bytes,1,opt,name=linear_buckets,json=linearBuckets,oneof"`
}
type Distribution_BucketOptions_ExponentialBuckets struct {
	ExponentialBuckets *Distribution_BucketOptions_Exponential `protobuf:"bytes,2,opt,name=exponential_buckets,json=exponentialBuckets,oneof"`
}
type Distribution_BucketOptions_ExplicitBuckets struct {
	ExplicitBuckets *Distribution_BucketOptions_Explicit `protobuf:"bytes,3,opt,name=explicit_buckets,json=explicitBuckets,oneof"`
}

func (*Distribution_BucketOptions_LinearBuckets) isDistribution_BucketOptions_Options()      {}
func (*Distribution_BucketOptions_ExponentialBuckets) isDistribution_BucketOptions_Options() {}
func (*Distribution_BucketOptions_ExplicitBuckets) isDistribution_BucketOptions_Options()    {}

func (m *Distribution_BucketOptions) GetOptions() isDistribution_BucketOptions_Options {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *Distribution_BucketOptions) GetLinearBuckets() *Distribution_BucketOptions_Linear {
	if x, ok := m.GetOptions().(*Distribution_BucketOptions_LinearBuckets); ok {
		return x.LinearBuckets
	}
	return nil
}

func (m *Distribution_BucketOptions) GetExponentialBuckets() *Distribution_BucketOptions_Exponential {
	if x, ok := m.GetOptions().(*Distribution_BucketOptions_ExponentialBuckets); ok {
		return x.ExponentialBuckets
	}
	return nil
}

func (m *Distribution_BucketOptions) GetExplicitBuckets() *Distribution_BucketOptions_Explicit {
	if x, ok := m.GetOptions().(*Distribution_BucketOptions_ExplicitBuckets); ok {
		return x.ExplicitBuckets
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Distribution_BucketOptions) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Distribution_BucketOptions_OneofMarshaler, _Distribution_BucketOptions_OneofUnmarshaler, _Distribution_BucketOptions_OneofSizer, []interface{}{
		(*Distribution_BucketOptions_LinearBuckets)(nil),
		(*Distribution_BucketOptions_ExponentialBuckets)(nil),
		(*Distribution_BucketOptions_ExplicitBuckets)(nil),
	}
}

func _Distribution_BucketOptions_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Distribution_BucketOptions)
	// options
	switch x := m.Options.(type) {
	case *Distribution_BucketOptions_LinearBuckets:
		_ = b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.LinearBuckets); err != nil {
			return err
		}
	case *Distribution_BucketOptions_ExponentialBuckets:
		_ = b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExponentialBuckets); err != nil {
			return err
		}
	case *Distribution_BucketOptions_ExplicitBuckets:
		_ = b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.ExplicitBuckets); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Distribution_BucketOptions.Options has unexpected type %T", x)
	}
	return nil
}

func _Distribution_BucketOptions_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Distribution_BucketOptions)
	switch tag {
	case 1: // options.linear_buckets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution_BucketOptions_Linear)
		err := b.DecodeMessage(msg)
		m.Options = &Distribution_BucketOptions_LinearBuckets{msg}
		return true, err
	case 2: // options.exponential_buckets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution_BucketOptions_Exponential)
		err := b.DecodeMessage(msg)
		m.Options = &Distribution_BucketOptions_ExponentialBuckets{msg}
		return true, err
	case 3: // options.explicit_buckets
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Distribution_BucketOptions_Explicit)
		err := b.DecodeMessage(msg)
		m.Options = &Distribution_BucketOptions_ExplicitBuckets{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Distribution_BucketOptions_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Distribution_BucketOptions)
	// options
	switch x := m.Options.(type) {
	case *Distribution_BucketOptions_LinearBuckets:
		s := proto.Size(x.LinearBuckets)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Distribution_BucketOptions_ExponentialBuckets:
		s := proto.Size(x.ExponentialBuckets)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Distribution_BucketOptions_ExplicitBuckets:
		s := proto.Size(x.ExplicitBuckets)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Specify a sequence of buckets that all have the same width (except
// overflow and underflow).  Each bucket represents a constant absolute
// uncertainty on the specific value in the bucket.
//
// Defines `num_finite_buckets + 2` (= N) buckets with these boundaries for
// bucket `i`:
//
//    Upper bound (0 <= i < N-1):     offset + (width * i).
//    Lower bound (1 <= i < N):       offset + (width * (i - 1)).
type Distribution_BucketOptions_Linear struct {
	// Must be greater than 0.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets,proto3" json:"num_finite_buckets,omitempty"`
	// Must be greater than 0.
	Width float64 `protobuf:"fixed64,2,opt,name=width,proto3" json:"width,omitempty"`
	// Lower bound of the first bucket.
	Offset float64 `protobuf:"fixed64,3,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (m *Distribution_BucketOptions_Linear) Reset()         { *m = Distribution_BucketOptions_Linear{} }
func (m *Distribution_BucketOptions_Linear) String() string { return proto.CompactTextString(m) }
func (*Distribution_BucketOptions_Linear) ProtoMessage()    {}
func (*Distribution_BucketOptions_Linear) Descriptor() ([]byte, []int) {
	return fileDescriptorDistribution, []int{0, 1, 0}
}

func (m *Distribution_BucketOptions_Linear) GetNumFiniteBuckets() int32 {
	if m != nil {
		return m.NumFiniteBuckets
	}
	return 0
}

func (m *Distribution_BucketOptions_Linear) GetWidth() float64 {
	if m != nil {
		return m.Width
	}
	return 0
}

func (m *Distribution_BucketOptions_Linear) GetOffset() float64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

// Specify a sequence of buckets that have a width that is proportional to
// the value of the lower bound.  Each bucket represents a constant relative
// uncertainty on a specific value in the bucket.
//
// Defines `num_finite_buckets + 2` (= N) buckets with these boundaries for
// bucket i:
//
//    Upper bound (0 <= i < N-1):     scale * (growth_factor ^ i).
//    Lower bound (1 <= i < N):       scale * (growth_factor ^ (i - 1)).
type Distribution_BucketOptions_Exponential struct {
	// Must be greater than 0.
	NumFiniteBuckets int32 `protobuf:"varint,1,opt,name=num_finite_buckets,json=numFiniteBuckets,proto3" json:"num_finite_buckets,omitempty"`
	// Must be greater than 1.
	GrowthFactor float64 `protobuf:"fixed64,2,opt,name=growth_factor,json=growthFactor,proto3" json:"growth_factor,omitempty"`
	// Must be greater than 0.
	Scale float64 `protobuf:"fixed64,3,opt,name=scale,proto3" json:"scale,omitempty"`
}

func (m *Distribution_BucketOptions_Exponential) Reset() {
	*m = Distribution_BucketOptions_Exponential{}
}
func (m *Distribution_BucketOptions_Exponential) String() string { return proto.CompactTextString(m) }
func (*Distribution_BucketOptions_Exponential) ProtoMessage()    {}
func (*Distribution_BucketOptions_Exponential) Descriptor() ([]byte, []int) {
	return fileDescriptorDistribution, []int{0, 1, 1}
}

func (m *Distribution_BucketOptions_Exponential) GetNumFiniteBuckets() int32 {
	if m != nil {
		return m.NumFiniteBuckets
	}
	return 0
}

func (m *Distribution_BucketOptions_Exponential) GetGrowthFactor() float64 {
	if m != nil {
		return m.GrowthFactor
	}
	return 0
}

func (m *Distribution_BucketOptions_Exponential) GetScale() float64 {
	if m != nil {
		return m.Scale
	}
	return 0
}

// A set of buckets with arbitrary widths.
//
// Defines `size(bounds) + 1` (= N) buckets with these boundaries for
// bucket i:
//
//    Upper bound (0 <= i < N-1):     bounds[i]
//    Lower bound (1 <= i < N);       bounds[i - 1]
//
// There must be at least one element in `bounds`.  If `bounds` has only one
// element, there are no finite buckets, and that single element is the
// common boundary of the overflow and underflow buckets.
type Distribution_BucketOptions_Explicit struct {
	// The values must be monotonically increasing.
	Bounds []float64 `protobuf:"fixed64,1,rep,packed,name=bounds" json:"bounds,omitempty"`
}

func (m *Distribution_BucketOptions_Explicit) Reset()         { *m = Distribution_BucketOptions_Explicit{} }
func (m *Distribution_BucketOptions_Explicit) String() string { return proto.CompactTextString(m) }
func (*Distribution_BucketOptions_Explicit) ProtoMessage()    {}
func (*Distribution_BucketOptions_Explicit) Descriptor() ([]byte, []int) {
	return fileDescriptorDistribution, []int{0, 1, 2}
}

func (m *Distribution_BucketOptions_Explicit) GetBounds() []float64 {
	if m != nil {
		return m.Bounds
	}
	return nil
}

func init() {
	proto.RegisterType((*Distribution)(nil), "google.api.Distribution")
	proto.RegisterType((*Distribution_Range)(nil), "google.api.Distribution.Range")
	proto.RegisterType((*Distribution_BucketOptions)(nil), "google.api.Distribution.BucketOptions")
	proto.RegisterType((*Distribution_BucketOptions_Linear)(nil), "google.api.Distribution.BucketOptions.Linear")
	proto.RegisterType((*Distribution_BucketOptions_Exponential)(nil), "google.api.Distribution.BucketOptions.Exponential")
	proto.RegisterType((*Distribution_BucketOptions_Explicit)(nil), "google.api.Distribution.BucketOptions.Explicit")
}

func init() { proto.RegisterFile("google/api/distribution.proto", fileDescriptorDistribution) }

var fileDescriptorDistribution = []byte{
	// 516 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x53, 0x5d, 0x6f, 0xd3, 0x30,
	0x14, 0x6d, 0x96, 0x7e, 0xc0, 0xed, 0x07, 0xc5, 0x0c, 0x14, 0x22, 0x3e, 0xaa, 0x4d, 0xa0, 0x4a,
	0x40, 0x2a, 0x15, 0x24, 0xde, 0xcb, 0x98, 0xfa, 0x00, 0xda, 0x64, 0x24, 0x9e, 0x90, 0x22, 0x27,
	0x71, 0x3a, 0x8b, 0xc6, 0x0e, 0xb1, 0xbd, 0x95, 0x77, 0xfe, 0x14, 0xff, 0x0e, 0xc5, 0x76, 0xb6,
	0x0c, 0x09, 0xa9, 0xbc, 0xf9, 0xde, 0x73, 0xee, 0x39, 0xe7, 0x5e, 0x25, 0xf0, 0x74, 0x23, 0xc4,
	0x66, 0x4b, 0x17, 0xa4, 0x64, 0x8b, 0x8c, 0x49, 0x55, 0xb1, 0x44, 0x2b, 0x26, 0x78, 0x54, 0x56,
	0x42, 0x09, 0x04, 0x16, 0x8e, 0x48, 0xc9, 0xc2, 0x27, 0x2d, 0x2a, 0xe1, 0x5c, 0x28, 0x52, 0x13,
	0xa5, 0x65, 0x86, 0x8f, 0x1d, 0x6a, 0xaa, 0x44, 0xe7, 0x0b, 0xc2, 0x7f, 0x3a, 0xe8, 0xf9, 0xdf,
	0x90, 0x62, 0x05, 0x95, 0x8a, 0x14, 0xa5, 0x25, 0x1c, 0xfd, 0x1a, 0xc0, 0xe8, 0xa4, 0x65, 0x8e,
	0x0e, 0xa1, 0x97, 0x0a, 0xcd, 0x55, 0xe0, 0xcd, 0xbc, 0xb9, 0x8f, 0x6d, 0x81, 0x10, 0x74, 0x0b,
	0x4a, 0x78, 0x70, 0x30, 0xf3, 0xe6, 0x1e, 0x36, 0x6f, 0xf4, 0x1e, 0x02, 0xa9, 0x8b, 0x58, 0xe4,
	0xb1, 0xfc, 0xa1, 0x49, 0x45, 0xb3, 0x38, 0xa3, 0x97, 0xcc, 0x24, 0x0b, 0x7c, 0xc3, 0x7b, 0x28,
	0x75, 0x71, 0x96, 0x7f, 0xb1, 0xe8, 0x49, 0x03, 0xa2, 0x77, 0xd0, 0xab, 0x08, 0xdf, 0xd0, 0xa0,
	0x3b, 0xf3, 0xe6, 0xc3, 0xe5, 0xb3, 0xe8, 0x66, 0xd3, 0xa8, 0x9d, 0x25, 0xc2, 0x35, 0x0b, 0x5b,
	0x32, 0xfa, 0x0c, 0x93, 0x44, 0xa7, 0xdf, 0xa9, 0x8a, 0x45, 0x69, 0xb6, 0x0f, 0xfa, 0x66, 0xfc,
	0xe5, 0x3f, 0xc7, 0x57, 0x86, 0x7e, 0x66, 0xd9, 0x78, 0x9c, 0xb4, 0x4b, 0x74, 0x0c, 0xae, 0x11,
	0x9b, 0x0d, 0x65, 0x30, 0x98, 0xf9, 0x73, 0x1f, 0x8f, 0x6c, 0xf3, 0x83, 0xe9, 0x85, 0xaf, 0xa0,
	0x67, 0x32, 0xa0, 0x29, 0xf8, 0x05, 0xe3, 0xe6, 0x26, 0x1e, 0xae, 0x9f, 0xa6, 0x43, 0x76, 0xee,
	0x20, 0xf5, 0x33, 0xfc, 0xdd, 0x85, 0xf1, 0x2d, 0x4b, 0xf4, 0x15, 0x26, 0x5b, 0xc6, 0x29, 0xa9,
	0x62, 0xab, 0x2a, 0x8d, 0xc0, 0x70, 0xf9, 0x66, 0xbf, 0xc8, 0xd1, 0x27, 0x33, 0xbc, 0xee, 0xe0,
	0xb1, 0x95, 0xb1, 0xa8, 0x44, 0x14, 0x1e, 0xd0, 0x5d, 0x29, 0x38, 0xe5, 0x8a, 0x91, 0xed, 0xb5,
	0xf8, 0x81, 0x11, 0x5f, 0xee, 0x29, 0xfe, 0xf1, 0x46, 0x61, 0xdd, 0xc1, 0xa8, 0x25, 0xd8, 0xd8,
	0x7c, 0x83, 0x29, 0xdd, 0x95, 0x5b, 0x96, 0x32, 0x75, 0xed, 0xe1, 0x1b, 0x8f, 0xc5, 0xfe, 0x1e,
	0x66, 0x7c, 0xdd, 0xc1, 0xf7, 0x1a, 0x29, 0xa7, 0x1e, 0x66, 0xd0, 0xb7, 0xfb, 0xa1, 0xd7, 0x80,
	0xb8, 0x2e, 0xe2, 0x9c, 0x71, 0xa6, 0xe8, 0xad, 0x53, 0xf5, 0xf0, 0x94, 0xeb, 0xe2, 0xd4, 0x00,
	0x4d, 0xaa, 0x43, 0xe8, 0x5d, 0xb1, 0x4c, 0x5d, 0xb8, 0xd3, 0xdb, 0x02, 0x3d, 0x82, 0xbe, 0xc8,
	0x73, 0x49, 0x95, 0xfb, 0xf4, 0x5c, 0x15, 0x5e, 0xc2, 0xb0, 0xb5, 0xe8, 0x7f, 0x5a, 0x1d, 0xc3,
	0x78, 0x53, 0x89, 0x2b, 0x75, 0x11, 0xe7, 0x24, 0x55, 0xa2, 0x72, 0x96, 0x23, 0xdb, 0x3c, 0x35,
	0xbd, 0x3a, 0x8f, 0x4c, 0xc9, 0x96, 0x3a, 0x63, 0x5b, 0x84, 0x47, 0x70, 0xa7, 0x59, 0xbe, 0xce,
	0x96, 0x08, 0xcd, 0xb3, 0xda, 0xc8, 0xaf, 0xb3, 0xd9, 0x6a, 0x75, 0x17, 0x06, 0xee, 0x53, 0x5e,
	0xbd, 0x80, 0x49, 0x2a, 0x8a, 0xd6, 0x55, 0x57, 0xf7, 0xdb, 0x67, 0x3d, 0xaf, 0xff, 0xd5, 0x73,
	0x2f, 0xe9, 0x9b, 0x9f, 0xf6, 0xed, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x4d, 0x31, 0xa0, 0xa1,
	0x3b, 0x04, 0x00, 0x00,
}
