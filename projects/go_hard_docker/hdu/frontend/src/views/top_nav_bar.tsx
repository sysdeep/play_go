import React from 'react';
import { Link } from 'react-router-dom';
import IconContainers from '../components/icon_containers';
import IconImages from '../components/icon_images';
import IconVolumes from '../components/icon_volumes';
import IconNetworks from '../components/icon_networks';
import IconConfigs from '../components/icon_configs';
import IconSecrets from '../components/icon_secrets';
import IconHome from '../components/icon_home';
import { route } from '../routes';

export default function TopNavBar() {
  return (
    <nav className='nav'>
      <div className='nav-left'>
        <Link to={'/'}>
          <IconHome />
          &nbsp; Main
        </Link>
        <Link to={route.containers}>
          {/* <i className='fa fa-cubes mr-1' aria-hidden='true'></i> */}
          <IconContainers />
          &nbsp; Containers
        </Link>
        <Link to={route.images}>
          <IconImages />
          &nbsp; Images
        </Link>
        <Link to={route.volumes}>
          <IconVolumes />
          &nbsp; Volumes
        </Link>
        <Link to={route.networks}>
          <IconNetworks />
          &nbsp; Networks
        </Link>
        <Link to={route.configs}>
          <IconConfigs />
          &nbsp; Configs
        </Link>
        <Link to={route.secrets}>
          <IconSecrets />
          &nbsp; Secrets
        </Link>
      </div>
    </nav>
  );
}
