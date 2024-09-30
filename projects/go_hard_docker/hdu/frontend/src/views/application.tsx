import React from 'react';
import { Outlet } from 'react-router-dom';
import TopNavBar from './top_nav_bar';

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
