package codec

import (
	"github.com/golang/snappy"
	commonpb "go.temporal.io/api/common/v1"
	"go.temporal.io/sdk/converter"
)

func NewSnappyDataConverter() converter.DataConverter {
	return converter.NewCodecDataConverter(converter.GetDefaultDataConverter(), NewSnappyCodec())
}

func NewSnappyCodec() converter.PayloadCodec {
	return Codec{}
}

type Codec struct{}

func (Codec) Encode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	result := make([]*commonpb.Payload, len(payloads))

	for i, p := range payloads {
		original, err := p.Marshal()
		if err != nil {
			return payloads, err
		}
		compressed := snappy.Encode(nil, original)
		result[i] = &commonpb.Payload{
			Metadata: map[string][]byte{converter.MetadataEncoding: []byte("binary/snappy")},
			Data:     compressed,
		}
	}

	return result, nil
}

func (Codec) Decode(payloads []*commonpb.Payload) ([]*commonpb.Payload, error) {
	result := make([]*commonpb.Payload, len(payloads))

	for i, p := range payloads {
		if string(p.Metadata[converter.MetadataEncoding]) != "binary/snappy" {
			result[i] = p
			continue
		}
		original, err := snappy.Decode(nil, p.Data)
		if err != nil {
			return payloads, err
		}
		result[i] = &commonpb.Payload{}
		err = result[i].Unmarshal(original)
		if err != nil {
			return payloads, err
		}
	}

	return result, nil
}
