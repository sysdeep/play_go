import React from 'react';
import { Outlet, Link } from 'react-router-dom';
import TopNavBar from './TopNavBar';

export default function Application() {
  return (
    <div className='container'>
      <TopNavBar />
      <div>
        <Outlet />
      </div>
    </div>
  );
}

// https://reactrouter.com/en/main/start/tutorial
