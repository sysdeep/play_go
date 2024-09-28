import React from 'react';

import ContainersPage from './views/containers_page/ContainersPage';
import ContainerPage from './views/container_page/container_page';
import ImagesPage from './views/images_page/ImagesPage';
import ErrorPage from './views/error-page';
import Application from './views/Application';

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
    ],
  },
];
