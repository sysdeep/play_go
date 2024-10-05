export class NetworksService {
  constructor() {
    console.log('networks_service created');
  }

  async get_networks(): Promise<ApiNetworkListModel[]> {
    const response = await fetch('http://localhost:1313/api/networks');

    const data = (await response.json()) as ApiNetworksListModel;

    if (data.networks.length > 0) {
      let net = await this.get_network(data.networks[0].id);
      console.log(net);
    }

    return data.networks || [];
  }

  async get_network(id: string): Promise<ApiFullNetworkModel> {
    const response = await fetch('http://localhost:1313/api/networks/' + id);

    const data = (await response.json()) as ApiFullNetworkModel;

    return data;
  }

  // async remove_image(id: string): Promise<void> {
  //   await fetch('http://localhost:1313/api/images/' + id, {
  //     method: 'DELETE',
  //   });

  //   return;
  // }
}

// list models ----------------------------------------------------------------
export interface ApiNetworkListModel {
  id: string;
  name: number;
  created: string;
  driver: string;
}

interface ApiNetworksListModel {
  networks: ApiNetworkListModel[];
  total: number;
}
// network model --------------------------------------------------------------
interface ApiNetworkModel {
  name: string;
  id: string;
  created: string;
  scope: string;
  driver: string;
  internal: boolean;
  attachable: boolean;
  ingress: boolean;
}

export interface ApiNetworkContainerModel {
  id: string;
  name: string;
  endpoint_id: string;
  mac_address: string;
  ip_v4_address: string;
  ip_v6_address: string;
}

export interface ApiFullNetworkModel {
  network: ApiNetworkModel;
  containers: ApiNetworkContainerModel[];
}
