package registry_client

// type RegistryClient interface {
// 	GetCatalog(n int) (Catalog, error)
// 	GetRepository(image_name string) (Repository, error)
// 	GetManivestV2(image_name string, tag_name string) (ManifestV2, error)
// 	RemoveManifest(image_name string, digest string) error
// }

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
	ConfigDescriptor  Descriptor
	LayersDescriptors []Descriptor
	TotalSize         int64
	ContentDigest     string
}

type Descriptor struct {
	MediaType string
	Size      int64
	Digest    string
}

// @staticmethod
// def from_response(data: dict) -> 'Descriptor':
//     return Descriptor(media_type=data.get('mediaType'), size=data.get('size'), digest=data.get('digest'))
