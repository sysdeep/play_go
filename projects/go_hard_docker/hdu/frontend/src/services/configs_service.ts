export class ConfigsServices {
  constructor() {
    console.log('configs service created');
  }

  async get_configs(): Promise<ApiConfigListModel[]> {
    const response = await fetch('http://localhost:1313/api/configs');

    const data = (await response.json()) as ApiConfigsListModel;

    return data.configs || [];
  }

  async get_config(id: string): Promise<ApiFullConfigModel> {
    const response = await fetch('http://localhost:1313/api/configs/' + id);

    const data = (await response.json()) as ApiFullConfigModel;

    return data;
  }

  async remove_config(id: string): Promise<void> {
    await fetch('http://localhost:1313/api/configs/' + id, {
      method: 'DELETE',
    });

    return;
  }
}

// list models ----------------------------------------------------------------
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

// config model ---------------------------------------------------------------
interface ApiConfigModel {
  id: string;
  name: string;
  created: string;
  updated: string;
  data_text: string;
}

export interface ApiFullConfigModel {
  config: ApiConfigModel;
}
