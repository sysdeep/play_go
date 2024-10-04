import React from 'react';

import Application from './views/application';
import ConfigsPage from './views/configs_page/configs_page';
import ContainerPage from './views/container_page/container_page';
import ContainersPage from './views/containers_page/containers_page';
import ErrorPage from './views/error-page';
import HomePage from './views/home_page/home_page';
import ImagePage from './views/image_page/image_page';
import ImagesPage from './views/images_page/images_page';
import NetworksPage from './views/networks_page/networks_page';
import SecretsPage from './views/secrets_page/secrets_page';
import VolumesPage from './views/volumes_page/volumes_page';

export const route = {
  images: '/images',
  containers: '/containers',
  volumes: '/volumes',
  networks: '/networks',
  configs: '/configs',
  secrets: '/secrets',
};

export const routes = [
  // {
  //   path: '/demo',
  //   element: <div>Hello world!</div>,
  // },
  {
    path: '/',
    // element: <div>Hello world!</div>,
    // element: <Application />,
    // element: <Root />,
    element: <Application />,
    errorElement: <ErrorPage />,

    children: [
      {
        path: '/',
        element: <HomePage />,
      },
      {
        path: route.containers,
        element: <ContainersPage />,
      },
      {
        path: '/container/:id',
        element: <ContainerPage />,
      },
      {
        path: route.images,
        element: <ImagesPage />,
      },
      {
        path: '/image/:id',
        element: <ImagePage />,
      },
      {
        path: route.volumes,
        element: <VolumesPage />,
      },
      {
        path: route.networks,
        element: <NetworksPage />,
      },
      {
        path: route.configs,
        element: <ConfigsPage />,
      },
      {
        path: route.secrets,
        element: <SecretsPage />,
      },
    ],
  },
];
