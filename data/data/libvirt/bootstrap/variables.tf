variable "addresses" {
  type        = list(string)
  default     = []
  description = "IP addresses to assign to the bootstrap node."
}

variable "base_volume_id" {
  type        = string
  description = "The ID of the base volume for the bootstrap node."
}

variable "cluster_id" {
  type        = string
  description = "The identifier for the cluster."
}

variable "cluster_domain" {
  type        = string
  description = "The domain for the cluster that all DNS records must belong"
}

variable "ignition" {
  type        = string
  description = "The content of the bootstrap ignition file."
}

variable "pool" {
  type        = string
  description = "The name of the storage pool."
}
