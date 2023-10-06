package types

const OCIContentDescriptor = "application/vnd.oci.descriptor.v1+json"

const OCIImageIndex = "application/vnd.oci.image.index.v1+json"

const OCIManifestSchema1 = "application/vnd.oci.image.manifest.v1+json"

const OCIConfigJSON = "application/vnd.oci.image.config.v1+json"

const OCILayer = "application/vnd.oci.image.layer.v1.tar+gzip"

const OCILayerZStd = "application/vnd.oci.image.layer.v1.tar+zstd"

const OCIRestrictedLayer = "application/vnd.oci.image.layer.nondistributable.v1.tar+gzip"

const OCIUncompressedLayer = "application/vnd.oci.image.layer.v1.tar"

const OCIUncompressedRestrictedLayer = "application/vnd.oci.image.layer.nondistributable.v1.tar"

const DockerManifestSchema1 = "application/vnd.docker.distribution.manifest.v1+json"

const DockerManifestSchema1Signed = "application/vnd.docker.distribution.manifest.v1+prettyjws"

const DockerManifestSchema2 = "application/vnd.docker.distribution.manifest.v2+json"

const DockerManifestList = "application/vnd.docker.distribution.manifest.list.v2+json"

const DockerLayer = "application/vnd.docker.image.rootfs.diff.tar.gzip"

const DockerConfigJSON = "application/vnd.docker.container.image.v1+json"

const DockerPluginConfig = "application/vnd.docker.plugin.v1+json"

const DockerForeignLayer = "application/vnd.docker.image.rootfs.foreign.diff.tar.gzip"

const DockerUncompressedLayer = "application/vnd.docker.image.rootfs.diff.tar"

const OCIVendorPrefix = "vnd.oci"

const DockerVendorPrefix = "vnd.docker"

type MediaType string

func (m MediaType) IsDistributable() bool {
	panic("stub")
}

func (m MediaType) IsImage() bool {
	panic("stub")
}

func (m MediaType) IsIndex() bool {
	panic("stub")
}

func (m MediaType) IsConfig() bool {
	panic("stub")
}

type Embedme interface{}
