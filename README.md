This repo illustrates the issue reported here: https://github.com/gogo/protobuf/issues/411

Specifically, message `Span` contains two fields `TraceID` and `SpanID` defined as `bytes`.
They are both backed by custom types `TraceID` and `SpanID`, where `TraceID` is a two-field struct
and `SpanID` is an alias to `uint64`.

The test `ids_test.go` does round-trip serialization using proto buffers and JSON.
If the method `func (s *SpanID) UnmarshalJSONPB` at the end of `ids.go` is commented
out or removed, the test fails, even though `SpanID` implements 
`func (t *T) UnmarshalJSON(data []byte) error {}` required for custom types per https://github.com/gogo/protobuf/blob/master/custom_types.md.
