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

type HTTPRegistryClient struct {
	address string
	client  *http.Client
}

func NewHTTPRegistryClient(address string) *HTTPRegistryClient {

	// for https ignoring
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := http.Client{Transport: tr, Timeout: 5 * time.Second}

	return &HTTPRegistryClient{address, &client}
}

func (c *HTTPRegistryClient) GetCatalog(n int) Catalog {

	url := c.make_url(fmt.Sprintf("/v2/_catalog?n={%d}", n))
	fmt.Println(url)

	resp, err := c.client.Get(url)
	if err != nil {
		fmt.Println(err)
		return Catalog{}
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	result := catalogResponse{}
	json.Unmarshal(body, &result)

	return Catalog{
		Repositories: result.Repositories,
	}
}

func (c *HTTPRegistryClient) make_url(part string) string {

	start := strings.TrimSuffix(c.address, "/")
	end := strings.TrimPrefix(part, "/")

	return start + "/" + end

}

type catalogResponse struct {
	Repositories []string `json:"repositories"`
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
