provider "libvirt" {
  uri = var.libvirt_uri
}

resource "libvirt_pool" "storage_pool" {
  name = var.cluster_id
  type = "dir"
  path = "/var/lib/libvirt/openshift-images/${var.cluster_id}"
}

module "volume" {
  source = "./volume"

  cluster_id = var.cluster_id
  image      = var.os_image
  pool       = libvirt_pool.storage_pool.name
}

module "bootstrap" {
  source = "./bootstrap"

  cluster_domain = var.cluster_domain
  addresses      = [var.libvirt_bootstrap_ip]
  base_volume_id = module.volume.coreos_base_volume_id
  cluster_id     = var.cluster_id
  ignition       = var.ignition_bootstrap
  pool           = libvirt_pool.storage_pool.name
}

resource "libvirt_volume" "master" {
  count          = var.master_count
  name           = "${var.cluster_id}-master-${count.index}"
  base_volume_id = module.volume.coreos_base_volume_id
  pool           = libvirt_pool.storage_pool.name
}

resource "libvirt_ignition" "master" {
  name    = "${var.cluster_id}-master.ign"
  content = var.ignition_master
  pool    = libvirt_pool.storage_pool.name
}


resource "libvirt_domain" "master" {
  count = var.master_count

  name = "${var.cluster_id}-master-${count.index}"

  memory = var.libvirt_master_memory
  vcpu   = var.libvirt_master_vcpu

  coreos_ignition = libvirt_ignition.master.id

  disk {
    volume_id = element(libvirt_volume.master.*.id, count.index)
  }

  console {
    type        = "pty"
    target_port = 0
  }

  cpu = {
    mode = "host-passthrough"
  }

  network_interface {
    bridge = "br-em1"
    mac = "52:54:00:4e:f5:0${count.index+1}"
  }
}

data "libvirt_network_dns_host_template" "bootstrap" {
  count    = var.bootstrap_dns ? 1 : 0
  ip       = var.libvirt_bootstrap_ip
  hostname = "api.${var.cluster_domain}"
}

data "libvirt_network_dns_host_template" "masters" {
  count    = var.master_count
  ip       = var.libvirt_master_ips[count.index]
  hostname = "api.${var.cluster_domain}"
}

data "libvirt_network_dns_host_template" "bootstrap_int" {
  count    = var.bootstrap_dns ? 1 : 0
  ip       = var.libvirt_bootstrap_ip
  hostname = "api-int.${var.cluster_domain}"
}

data "libvirt_network_dns_host_template" "masters_int" {
  count    = var.master_count
  ip       = var.libvirt_master_ips[count.index]
  hostname = "api-int.${var.cluster_domain}"
}

data "libvirt_network_dns_host_template" "etcds" {
  count    = var.master_count
  ip       = var.libvirt_master_ips[count.index]
  hostname = "etcd-${count.index}.${var.cluster_domain}"
}

data "libvirt_network_dns_srv_template" "etcd_cluster" {
  count    = var.master_count
  service  = "etcd-server-ssl"
  protocol = "tcp"
  domain   = var.cluster_domain
  port     = 2380
  weight   = 10
  target   = "etcd-${count.index}.${var.cluster_domain}"
}

