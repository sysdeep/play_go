export class SecretsService {
  constructor() {
    console.log('secrets service created');
  }

  async get_secrets(): Promise<ApiSecretListModel[]> {
    const response = await fetch('http://localhost:1313/api/secrets');

    const data = (await response.json()) as ApiSecretsListModel;

    return data.secrets || [];
  }

  async get_secret(id: string): Promise<ApiFullSecretModel> {
    const response = await fetch('http://localhost:1313/api/secrets/' + id);

    const data = (await response.json()) as ApiFullSecretModel;

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
export interface ApiSecretListModel {
  id: string;
  name: number;
  created: string;
  updated: string;
}

interface ApiSecretsListModel {
  secrets: ApiSecretListModel[];
  total: number;
}

// Secret model ---------------------------------------------------------------
interface ApiSecretModel {
  id: string;
  name: string;
  created: string;
  updated: string;
}

export interface ApiFullSecretModel {
  secret: ApiSecretModel;
}
