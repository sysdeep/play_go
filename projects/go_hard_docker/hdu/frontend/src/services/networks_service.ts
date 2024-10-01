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

    return data.networks || [];
  }

  // async remove_image(id: string): Promise<void> {
  //   await fetch('http://localhost:1313/api/images/' + id, {
  //     method: 'DELETE',
  //   });

  //   return;
  // }
}
