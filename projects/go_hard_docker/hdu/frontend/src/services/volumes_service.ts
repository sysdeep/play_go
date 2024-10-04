export default class VolumesService {
  constructor() {
    console.log('volumes_service created');
  }

  async get_volumes(): Promise<ApiVolumeListModel[]> {
    const response = await fetch('http://localhost:1313/api/volumes');

    const data = (await response.json()) as ApiVolumesListModel;

    return data.volumes || [];
  }

  async get_volume(id: string): Promise<ApiFullVolumeModel> {
    const response = await fetch('http://localhost:1313/api/volumes/' + id);

    const data = (await response.json()) as ApiFullVolumeModel;

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

// volume model ---------------------------------------------------------------
export interface ApiFullVolumeModel {
  volume: ApiVolumeModel;
}
interface ApiVolumeModel {
  name: string;
  created: string;
  driver: string;
  mount_point: string;
}
