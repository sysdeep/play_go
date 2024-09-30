import React from 'react';
import { createRoot } from 'react-dom/client';

import { RouterProvider, createHashRouter } from 'react-router-dom';

import 'chota';
import './style.css';
import 'bootstrap-icons/font/bootstrap-icons.css';

import { routes } from './routes';

// Say something
console.log('[ERWT] : Renderer execution started');

// // Application to Render
// const app = <Application />;

// const router = createBrowserRouter([
const router = createHashRouter(routes);

// Render application in DOM
// createRoot(document.getElementById('app')).render(app);
createRoot(document.getElementById('app')).render(
  <RouterProvider router={router} />,
);
