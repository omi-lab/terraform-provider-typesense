terraform {
  required_providers {
    typesense = {
      source = "ronati/typesense"
    }
  }
}

provider "typesense" {
  api_key     = var.typesense_api_key
  api_address = var.typesense_api_address
}

resource "typesense_collection" "local_provider_test" {
  name = var.collection_name

  embed_api_keys = {
    embeddings_v2 = var.jina_api_key
  }

  fields {
    name  = "title"
    type  = "string"
    index = true
    facet = false
  }

  fields {
    name     = "category"
    type     = "string"
    index    = true
    facet    = true
    optional = true
  }

  fields {
    name     = "embeddings_v2"
    type     = "float[]"
    index    = true
    facet    = false
    optional = true
    sort     = false
    infix    = false
    store    = true
    num_dim  = 1024
    embed {
      from = ["name", "description"]
      model_config {
        model_name = "openai/jina-clip-v2"
        url        = "https://api.jina.ai/v1"
      }
    }
  }
}

output "collection_name" {
  description = "Created collection name"
  value       = typesense_collection.local_provider_test.name
}
