export interface ApiConfigListModel {
  id: string;
  name: number;
  created: string;
  updated: string;
}

interface ApiConfigsListModel {
  configs: ApiConfigListModel[];
  total: number;
}

export class ConfigsServices {
  constructor() {
    console.log('configs service created');
  }

  async get_configs(): Promise<ApiConfigListModel[]> {
    const response = await fetch('http://localhost:1313/api/configs');

    const data = (await response.json()) as ApiConfigsListModel;

    return data.configs || [];
  }

  // async remove_image(id: string): Promise<void> {
  //   await fetch('http://localhost:1313/api/images/' + id, {
  //     method: 'DELETE',
  //   });

  //   return;
  // }
}
