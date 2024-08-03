package registry_client

type RegistryClient interface {
	GetCatalog(n int) Catalog
	GetRepository(image_name string) Repository
	GetManivestV2(image_name string, tag_name string)
	RemoveManifest(image_name string, digest string)
}

type Catalog struct {
	Repositories []string
}

type Repository struct {
	Name string
	Tags []string
}

type ManifestV2 struct {
	SchemaVersion     int
	MediaType         string
	DonfigDescriptor  Descriptor
	LayersDescriptors []Descriptor
	TotalSize         int
	ContentDigest     string
}

type Descriptor struct {
	MediaType string
	Size      int
	Digest    string
}

// @staticmethod
// def from_response(data: dict) -> 'Descriptor':
//     return Descriptor(media_type=data.get('mediaType'), size=data.get('size'), digest=data.get('digest'))
