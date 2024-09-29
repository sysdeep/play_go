import React from 'react';

interface ActionsFrameProps {
  id: string;
}
export default function ActionsFrame({ id }: ActionsFrameProps) {
  return (
    <div className='_card'>
      <h2>Actions</h2>
      <div>
        <ul>
          <li>
            <a href='/container/{id}/action/start'>Start TODO</a>
          </li>
          <li>
            <a href='/container/{id}/action/stop'>Stop TODO</a>
          </li>
          <li>
            <a href='/container/{id}/action/kill'>Kill TODO</a>
          </li>
          <li>
            <a href='/container/{id}/action/restart'>Restart TODO</a>
          </li>
          <li>
            <a href='/container/{id}/action/pause'>Pause TODO</a>
          </li>
          <li>
            <a href='/container/{id}/action/resume'>Resume TODO</a>
          </li>
          <li>
            <a href='/container/{id}/action/remove'>Remove TODO</a>
          </li>
        </ul>
      </div>
    </div>
  );
}
