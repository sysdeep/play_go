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

interface ApiContainerModel {
  id: string;
  created: string;
  name: string;
  restart_count: number;
}
interface ApiContainerStateModel {
  status: string;
  started: string;
}
interface ApiContainerMountsModel {}
interface ApiContainerConfigModel {
  env: string[];
  cmd: string;
  image: string;
  entrypoint: string;
}
interface ApiNetworkSegment {
  gateway: string;
  ip_address: string;
  mac_address: string;
  network_id: string;
}
interface ApiPortSegment {
  host_ip: string;
  host_port: string;
}
type ApiNetworkMap = { [id: string]: ApiNetworkSegment };
type ApiPortMap = { [id: string]: ApiPortMap };
interface ApiContainerNetworkModel {
  networks: ApiNetworkMap;
  ports: ApiPortMap;
}

export interface ApiContainerResponseModel {
  container: ApiContainerModel;
  state: ApiContainerStateModel;
  mounts: ApiContainerMountsModel;
  config: ApiContainerConfigModel;
  network: ApiContainerNetworkModel;
}
