import ContainerListModel from '../models/container_list_model';

interface ApiImageListModel {
  containers: number;
  created: string;
  id: string;
  tags: string[];
  size: number;
}

interface ApiImagesListModel {
  images: ApiImageListModel[];
  total: number;
}

export default class ContainersService {
  constructor() {
    console.log('containers_service created');
  }

  async get_containers(): Promise<ContainerListModel[]> {
    let response = await fetch('http://localhost:1313/api/containers');

    let data = (await response.json()) as ApiImagesListModel;
    console.log(data);

    let dataset = data.images.map((model) => {
      let dmodel: ContainerListModel = {
        id: model.id,
        created: model.created,

        // TODO
        name: '',
        image: '',
        state: '',
        // tags: model.tags,
        // size: model.size,
      };
      return dmodel;
    });

    return dataset;
  }

  async remove_container(id: string): Promise<void> {
    await fetch('http://localhost:1313/api/containers/' + id, {
      method: 'DELETE',
    });

    return;
  }
}
