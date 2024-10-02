interface ApiSystemModel {
  // ID: 'RIC3:VAEK:G7YV:R6I5:4NTJ:VP3L:SC3L:KII6:Q24U:O3ER:LUEL:N2XC';
  // Containers: 2;
  // ContainersRunning: 2;
  // ContainersPaused: 0;
  // ContainersStopped: 0;
  // Images: 30;
  // Driver: 'overlay2';
  // KernelVersion: '6.8.0-45-generic';
  // OperatingSystem: 'Ubuntu 22.04.5 LTS';
  // OSVersion: '22.04';
  // OSType: 'linux';
  ncpu: number;
  mem_total: number;
  name: string;
  server_version: string;
  default_runtime: string;
  // Swarm: {
  //   NodeID: 'tgyx3z0ro5qukwx6l5ukq1a76';
  //   NodeAddr: '192.168.200.52';
  // };
  // DefaultAddressPools: [
  //   {
  //     Base: '172.27.0.0/16';
  //     Size: 24;
  //   },
  // ];
}

export interface ApiInfoModel {
  daemon_host: string;
  client_version: string;
  system: ApiSystemModel;
}

export default class InfoService {
  constructor() {
    console.log('volumes_service created');
  }

  async get_info(): Promise<ApiInfoModel> {
    const response = await fetch('http://localhost:1313/api/info');

    const data = (await response.json()) as ApiInfoModel;

    return data;
  }

  // async remove_image(id: string): Promise<void> {
  //   await fetch('http://localhost:1313/api/images/' + id, {
  //     method: 'DELETE',
  //   });

  //   return;
  // }
}
