import React from "react";

export default function ActionsFrame(){
    return (
        <div className="_card">
  <h2>Actions</h2>
  <div>
    <ul>
      <li>
        <a href="/container/{{.Container.ID}}/action/start">Start TODO</a>
      </li>
      <li>
        <a href="/container/{{.Container.ID}}/action/stop">Stop TODO</a>
      </li>
      <li>
        <a href="/container/{{.Container.ID}}/action/kill">Kill TODO</a>
      </li>
      <li>
        <a href="/container/{{.Container.ID}}/action/restart">Restart TODO</a>
      </li>
      <li>
        <a href="/container/{{.Container.ID}}/action/pause">Pause TODO</a>
      </li>
      <li>
        <a href="/container/{{.Container.ID}}/action/resume">Resume TODO</a>
      </li>
      <li>
        <a href="/container/{{.Container.ID}}/action/remove">Remove TODO</a>
      </li>
    </ul>
  </div>
</div>

    )
}