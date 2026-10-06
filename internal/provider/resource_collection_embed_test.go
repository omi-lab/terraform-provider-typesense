package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func TestFlattenFieldEmbedOmitsCredentials(t *testing.T) {
	apiKey := "secret"
	url := "https://api.jina.ai/v1"
	embed := &fieldEmbedAPI{From: []string{"name"}}
	embed.ModelConfig.ModelName = "openai/jina-clip-v2"
	embed.ModelConfig.Url = &url
	embed.ModelConfig.ApiKey = &apiKey

	got := flattenFieldEmbed(embed)
	if got == nil || got.ModelConfig == nil {
		t.Fatal("expected embed model config")
	}
	if !got.ModelConfig.ApiKey.IsNull() {
		t.Fatalf("api_key = %q, want null so it is not stored inside fields", got.ModelConfig.ApiKey.ValueString())
	}
	if !got.ModelConfig.AccessToken.IsNull() || !got.ModelConfig.ClientSecret.IsNull() || !got.ModelConfig.RefreshToken.IsNull() {
		t.Fatal("credential fields must stay null")
	}
	if got.ModelConfig.ModelName.ValueString() != "openai/jina-clip-v2" {
		t.Fatalf("model_name = %q", got.ModelConfig.ModelName.ValueString())
	}
	if got.ModelConfig.Url.ValueString() != url {
		t.Fatalf("url = %q", got.ModelConfig.Url.ValueString())
	}
}

func TestFiledModelToApiFieldUsesEmbedAPIKeys(t *testing.T) {
	field := CollectionResourceFieldModel{
		Name: types.StringValue("embedding"),
		Type: types.StringValue("float[]"),
		Embed: &CollectionFieldEmbedModel{
			From: []types.String{types.StringValue("name")},
			ModelConfig: &CollectionFieldEmbedModelConfigModel{
				ModelName:    types.StringValue("openai/jina-clip-v2"),
				ApiKey:       types.StringNull(),
				AccessToken:  types.StringNull(),
				ClientSecret: types.StringNull(),
				RefreshToken: types.StringNull(),
			},
		},
	}

	got := filedModelToApiField(field, map[string]string{"embedding": "secret"})
	if got.Embed == nil || got.Embed.ModelConfig.ApiKey == nil || *got.Embed.ModelConfig.ApiKey != "secret" {
		t.Fatal("expected embed api key from embed_api_keys")
	}
}

func TestRotatedEmbedAPIKeyFields(t *testing.T) {
	keys := func(value string) types.Map {
		return types.MapValueMust(types.StringType, map[string]attr.Value{
			"embedding": types.StringValue(value),
		})
	}

	if len(rotatedEmbedAPIKeyFields(types.MapNull(types.StringType), keys("secret"))) != 0 {
		t.Fatal("recording a key that was not in state must not rotate the field")
	}
	if len(rotatedEmbedAPIKeyFields(keys("secret"), keys("secret"))) != 0 {
		t.Fatal("unchanged key must not rotate the field")
	}
	rotated := rotatedEmbedAPIKeyFields(keys("old"), keys("new"))
	if !rotated["embedding"] {
		t.Fatal("expected embedding key rotation")
	}
}
