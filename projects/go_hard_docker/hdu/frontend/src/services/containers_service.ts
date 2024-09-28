import ContainerListModel from '../models/container_list_model';

interface ApiContainerListModel {
  id: string;
  name: string;
  created: string;
  image: string;
  state: string;
}

interface ApiContainersListModel {
  containers: ApiContainerListModel[];
}

export default class ContainersService {
  constructor() {
    console.log('containers_service created');
  }

  async get_containers(): Promise<ContainerListModel[]> {
    let response = await fetch('http://localhost:1313/api/containers');

    let data = (await response.json()) as ApiContainersListModel;

    let dataset = data.containers.map((model) => {
      let dmodel: ContainerListModel = {
        id: model.id,
        created: model.created,
        name: model.name,
        image: model.image,
        state: model.state,
      };
      return dmodel;
    });

    return dataset;
  }

  // async remove_container(id: string): Promise<void> {
  //   await fetch('http://localhost:1313/api/containers/' + id, {
  //     method: 'DELETE',
  //   });

  //   return;
  // }

  async get_container(id: string): Promise<ApiContainerResponseModel> {
    let response = await fetch('http://localhost:1313/api/containers/' + id);
    let data = (await response.json()) as ApiContainerResponseModel;
    return data;
  }
}

interface ApiContainerModel {}
interface ApiContainerStateModel {}
interface ApiContainerMountsModel {}
interface ApiContainerConfigModel {}
interface ApiContainerNetworkModel {}

interface ApiContainerResponseModel {
  container: ApiContainerModel;
  state: ApiContainerStateModel;
  mounts: ApiContainerMountsModel;
  config: ApiContainerConfigModel;
  network: ApiContainerNetworkModel;
}
