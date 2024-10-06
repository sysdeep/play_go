import React from 'react';
import { createRoot } from 'react-dom/client';

import { RouterProvider, createHashRouter } from 'react-router-dom';

import 'chota';
import './style.css';
import 'bootstrap-icons/font/bootstrap-icons.css';

import { routes } from './routes';
import { useConfiguration } from './store/configuration';

// setup configuration
const { setConfiguration } = useConfiguration();
setConfiguration({
  // TODO: use global variable
  base_url: 'http://localhost:1313',
});

// setup router
// const router = createBrowserRouter([
const router = createHashRouter(routes);

// Render application in DOM
// createRoot(document.getElementById('app')).render(app);
createRoot(document.getElementById('app')).render(
  <RouterProvider router={router} />,
);
