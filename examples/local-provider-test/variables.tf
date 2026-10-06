variable "typesense_api_key" {
  description = "Typesense API key"
  type        = string
  sensitive   = true
}

variable "typesense_api_address" {
  description = "Typesense server address, e.g. http://localhost:8108"
  type        = string
}

variable "collection_name" {
  description = "Collection name to create"
  type        = string
  default     = "local-provider-test-collection"
}

variable "jina_api_key" {
  description = "Jina API key"
  type        = string
  sensitive   = true
}
