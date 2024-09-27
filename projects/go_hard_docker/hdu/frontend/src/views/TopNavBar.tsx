import React from 'react';
import { Outlet, Link } from 'react-router-dom';

export default function TopNavBar() {
  return (
    <nav className='nav'>
      <div className='nav-left'>
        <Link to={'/'}>Main</Link>
        <Link to='/containers'>
          <i className='fa fa-cubes mr-1' aria-hidden='true'></i>
          Containers
        </Link>
        <Link to='/images'>
          <i className='fa fa-hdd-o mr-1' aria-hidden='true'></i>
          Images
        </Link>
        <Link to='/volumes'>
          <i className='fa fa-folder mr-1' aria-hidden='true'></i>
          Volumes
        </Link>
        <Link to='/networks'>
          <i className='fa fa-sitemap mr-1' aria-hidden='true'></i>
          Networks
        </Link>
        <Link to='/configs'>
          <i className='fa fa-inbox mr-1' aria-hidden='true'></i>
          Configs
        </Link>
        <Link to='/secrets'>
          <i className='fa fa-lock mr-1' aria-hidden='true'></i>
          Secrets
        </Link>
      </div>
    </nav>
  );
}
