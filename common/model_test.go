package common_test

import (
	"testing"

	"github.com/QuantumNous/new-api/common"
	"github.com/QuantumNous/new-api/constant"
	"github.com/stretchr/testify/assert"
)

func TestWanEndpointsDistinguishImagesFromVideos(t *testing.T) {
	for _, name := range []string{
		"wan2.7-image-pro", "wan2.7-image", "wan2.6-image", "wan2.6-t2i",
		"wan2.5-t2i-preview", "wan2.2-t2i-flash", "wan2.2-t2i-plus",
		"wanx2.1-t2i-turbo", "wanx2.1-t2i-plus", "wanx2.0-t2i-turbo",
	} {
		t.Run(name, func(t *testing.T) {
			assert.Contains(t, common.GetEndpointTypesByChannelType(constant.ChannelTypeAli, name), constant.EndpointTypeImageGeneration)
		})
	}
	for _, name := range []string{
		"wanx2.1-t2v-plus", "wanx2.1-t2v-turbo", "wanx2.1-i2v-plus", "wanx2.1-i2v-turbo",
	} {
		t.Run(name, func(t *testing.T) {
			assert.NotContains(t, common.GetEndpointTypesByChannelType(constant.ChannelTypeAli, name), constant.EndpointTypeImageGeneration)
		})
	}
}

func TestIsImageGenerationModel(t *testing.T) {
	tests := []struct {
		name  string
		model string
		want  bool
	}{
		{name: "gpt-image-1", model: "gpt-image-1", want: true},
		{name: "gpt-image-1.5", model: "gpt-image-1.5", want: true},
		{name: "gpt-image versioned suffix", model: "gpt-image-2.5-sunburst", want: true},
		{name: "chatgpt-image-latest", model: "chatgpt-image-latest", want: true},
		{name: "dall-e", model: "dall-e-3", want: true},
		{name: "imagen prefix", model: "imagen-4.0-generate-001", want: true},
		{name: "flux", model: "flux.1-schnell", want: true},
		{name: "chat model", model: "gpt-4o", want: false},
		{name: "text model", model: "gpt-5.2", want: false},
		{name: "gemini model", model: "gemini-2.5-flash-image", want: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, common.IsImageGenerationModel(test.model))
		})
	}
}
