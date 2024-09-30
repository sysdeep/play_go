export interface ApiVolumeListModel {
  id: string;
  created: string;
  name: number;
  stack_name: number;
  mount_point: string;
  driver: string;
}

interface ApiVolumesListModel {
  volumes: ApiVolumeListModel[];
  total: number;
}

export default class VolumesService {
  constructor() {
    console.log('volumes_service created');
  }

  async get_volumes(): Promise<ApiVolumeListModel[]> {
    let response = await fetch('http://localhost:1313/api/volumes');

    let data = (await response.json()) as ApiVolumesListModel;

    return data.volumes;
  }

  // async remove_image(id: string): Promise<void> {
  //   await fetch('http://localhost:1313/api/images/' + id, {
  //     method: 'DELETE',
  //   });

  //   return;
  // }
}
