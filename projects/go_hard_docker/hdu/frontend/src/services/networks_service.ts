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

export class NetworksServices {
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

export interface ApiFullNetworkModel {
  network: any;
  containers: any;
}
