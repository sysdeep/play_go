package registry_client

/*
https://metanit.com/go/tutorial/9.6.php

*/
import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

const (
	// scheme version of manifest file
	// for details about scheme version goto https://docs.docker.com/registry/spec/manifest-v2-2/
	manifestSchemeV2 = "application/vnd.docker.distribution.manifest.v2+json"

	//  It uniquely identifies content by taking a collision-resistant hash of the bytes.
	contentDigestHeader = "docker-content-digest"
)

type HTTPRegistryClient struct {
	address string
	client  *http.Client
}

// create client
func NewHTTPRegistryClient(address string) *HTTPRegistryClient {

	// for https ignoring
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr, Timeout: 5 * time.Second}

	return &HTTPRegistryClient{address, &client}
}

// GetCatalog
func (c *HTTPRegistryClient) GetCatalog(n int) (Catalog, error) {

	// make endpoint address
	url := c.make_url(fmt.Sprintf("/v2/_catalog?n={%d}", n))

	// fetch
	body, err := c.make_get(url)
	if err != nil {
		return Catalog{}, err
	}

	// parse
	result := catalogResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Catalog{}, err
	}

	return Catalog{
		Repositories: result.Repositories,
	}, nil

}

// GetRepository
func (c *HTTPRegistryClient) GetRepository(image_name string) (Repository, error) {
	url := c.make_url(fmt.Sprintf("/v2/%s/tags/list", image_name))

	body, err := c.make_get(url)
	if err != nil {
		return Repository{}, err
	}

	result := repositoryResponse{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return Repository{}, err
	}

	return Repository{
		Name: result.Name,
		Tags: result.Tags,
	}, nil

}

// GetManivestV2
func (c *HTTPRegistryClient) GetManivestV2(image_name string, tag_name string) (ManifestV2, error) {
	url := c.make_url(fmt.Sprintf("/v2/%s/manifests/%s", image_name, tag_name))

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ManifestV2{}, err
	}
	req.Header.Add("Accept", manifestSchemeV2)
	res, err := c.client.Do(req)

	if err != nil {
		return ManifestV2{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return ManifestV2{}, err
	}

	schema := manifestV2ResponseSchema{}
	err = json.Unmarshal(body, &schema)
	if err != nil {
		return ManifestV2{}, err
	}

	content_digest := res.Header.Get(contentDigestHeader)

	var layers_descriptors []Descriptor

	for _, ld := range schema.LayersDescriptors {
		layers_descriptors = append(layers_descriptors, c.cd_from_response(ld))
	}

	return ManifestV2{
		SchemaVersion:     schema.SchemaVersion,
		MediaType:         schema.MediaType,
		ConfigDescriptor:  c.cd_from_response(schema.ConfigDescriptor),
		LayersDescriptors: layers_descriptors,
		TotalSize:         schema.CalculateCompressedImageSize(),
		ContentDigest:     content_digest,
	}, nil

}

// DeleteTag will delete the manifest identified by name and reference. Note that a manifest can only be deleted by digest.
// A digest can be fetched from manifest get response header 'docker-content-digest'
// после удаления необходимо выполнить чистку
// docker exec -it registry bin/registry garbage-collect  /etc/docker/registry/config.yml
func (c *HTTPRegistryClient) RemoveManifest(image_name string, digest string) error {
	// curl -v --silent -H "Accept: application/vnd.docker.distribution.manifest.v2+json" \
	// -X DELETE http://127.0.0.1:5000/v2/ubuntu/manifests/sha256:7cc0576c7c0ec2384de5cbf245f41567e922aab1b075f3e8ad565f508032df17

	fmt.Println("Client - remove: ", image_name, "/", digest)
	// TODO:
	// url, _ := url.JoinPath(c.address, "v2", reposytoryName, "/manifests/", digest)
	// c.logger.Debug("sending request: " + url)
	//
	// fmt.Println("-----------------------------------------")
	// fmt.Println(url)
	// fmt.Println("-----------------------------------------")
	//
	// req, err := http.NewRequest("DELETE", url, nil)
	// if err != nil {
	// 	return err
	// }
	//
	// req.Header.Add("Accept", manifestSchemeV2)
	// res, err := c.http_client.Do(req)
	// if err != nil {
	// 	return err
	// }
	// defer res.Body.Close()
	//
	// body, err := io.ReadAll(res.Body)
	// if err != nil {
	// 	return err
	// }
	//
	// fmt.Println("Delete body result ==================================================")
	// fmt.Println("Status: ", res.Status)
	// fmt.Println("StatusCode: ", res.StatusCode)
	// fmt.Println("Body:", string(body))
	// fmt.Println("=====================================================================")
	return nil

}

// private --------------------------------------------------------------------
func (c *HTTPRegistryClient) make_get(url string) ([]byte, error) {

	resp, err := c.client.Get(url)
	if err != nil {
		fmt.Println(err)
		return make([]byte, 0), err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, err
}

func (c *HTTPRegistryClient) make_url(part string) string {

	start := strings.TrimSuffix(c.address, "/")
	end := strings.TrimPrefix(part, "/")

	return start + "/" + end
}

func (c *HTTPRegistryClient) cd_from_response(data schema2Descriptor) Descriptor {
	return Descriptor{
		MediaType: data.MediaType,
		Size:      data.Size,
		Digest:    data.Digest,
	}
}

// NOTE: идея определять транспорт, но и без этого работает
// func makeHttpClient(address string) *http.Client{
// 	if strings.HasPrefix(address, "https:"){
//
// 	}
//
// }

// http models ----------------------------------------------------------------
type catalogResponse struct {
	Repositories []string `json:"repositories"`
}

type repositoryResponse struct {
	Name string   `json:"name"`
	Tags []string `json:"tags"`
}

/*
import requests
import urllib3

from registry_cli.registry_client.client_interface import ClientInterface
from registry_cli.registry_client.models import Catalog, Descriptor, ManifestV2, Repository
from registry_cli.registry_client.registry_client_params import RegistryClientParams


class RegistryClient(ClientInterface):
    _timeout = 5
    _manifest_scheme_v2_header = "application/vnd.docker.distribution.manifest.v2+json"
    _content_digest_header = "docker-content-digest"

    def __init__(self, params: RegistryClientParams):
        self._params = params
        urllib3.disable_warnings()

    def get_catalog(self, n: int) -> Catalog:
        url = self._make_url(f"/v2/_catalog?n={n}")
        resp = requests.get(url, verify=False, timeout=self._timeout)
        data = resp.json()

        return Catalog(repositories=[*data['repositories']])

    def get_repository(self, image_name: str) -> Repository:

        url = self._make_url(f"/v2/{image_name}/tags/list")
        resp = requests.get(url, verify=False, timeout=self._timeout)
        data = resp.json()

        tags = data['tags'] or []
        return Repository(name=data['name'], tags=tags)

    def get_manifest_v2(self, image_name: str, tag_name: str) -> ManifestV2:
        url = self._make_url(f"/v2/{image_name}/manifests/{tag_name}")
        headers = {
            'Accept': self._manifest_scheme_v2_header
        }
        resp = requests.get(url, verify=False, timeout=self._timeout, headers=headers)
        data = resp.json()

        result = ManifestV2(
            schema_version=data.get('schemaVersion'),
            media_type=data.get('mediaType'),
            config_descriptor=Descriptor.from_response(data.get('config')),
            layers_descriptors=[Descriptor.from_response(r) for r in data.get('layers')],

            total_size=0,
            content_digest=resp.headers.get(self._content_digest_header)
        )

        result.total_size = sum(d.size for d in result.layers_descriptors)

        return result

    def remove_manifest(self, image_name: str, digest: str):
        url = self._make_url(f"/v2/{image_name}/manifests/{digest}")
        headers = {
            'Accept': self._manifest_scheme_v2_header
        }
        resp = requests.delete(url, verify=False, timeout=self._timeout, headers=headers)
        resp.raise_for_status()

    def _make_url(self, part: str) -> str:
        start = self._params.registry_address
        if start.endswith('/'):
            start = start[:-1]

        end = part
        if end.startswith('/'):
            end = end[1:]

        return start + '/' + end



*/
