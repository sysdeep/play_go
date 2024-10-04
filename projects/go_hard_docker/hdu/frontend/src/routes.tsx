import React from 'react';

import ContainersPage from './views/containers_page/containers_page';
import ContainerPage from './views/container_page/container_page';
import ImagesPage from './views/images_page/images_page';
import ErrorPage from './views/error-page';
import Application from './views/application';
import VolumesPage from './views/volumes_page/volumes_page';
import NetworksPage from './views/networks_page/networks_page';
import ConfigsPage from './views/configs_page/configs_page';
import SecretsPage from './views/secrets_page/secrets_page';
import HomePage from './views/home_page/home_page';

export const route = {
  images: '/images',
  containers: '/containers',
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
        path: '/containers',
        element: <ContainersPage />,
      },
      {
        path: '/container/:id',
        element: <ContainerPage />,
      },
      {
        path: '/images',
        element: <ImagesPage />,
      },
      {
        path: '/volumes',
        element: <VolumesPage />,
      },
      {
        path: '/networks',
        element: <NetworksPage />,
      },
      {
        path: '/configs',
        element: <ConfigsPage />,
      },
      {
        path: '/secrets',
        element: <SecretsPage />,
      },
    ],
  },
];
