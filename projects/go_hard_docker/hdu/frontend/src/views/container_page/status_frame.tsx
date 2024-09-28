import React from "react";


export default function StatusFrame(){
    return (

<div>
  <h2>Container status</h2>
  <div>
    <table>
      <tbody>
        <tr>
          <td>ID</td>
          <td>{{ .Container.ID}}</td>
        </tr>
        <tr>
          <td>Name</td>
          <td>{{.Container.Name}}</td>
        </tr>
        <tr>
          <td>Ip address</td>
          <td>TODO</td>
        </tr>
        <tr>
          <td>Status</td>
          <td>{{.State.Status}}</td>
        </tr>
        <tr>
          <td>Created</td>
          <td>{{ .Container.Created }}</td>
        </tr>
        <tr>
          <td>Start time</td>
          <td>{{.State.StartedAt}}</td>
        </tr>
        <tr>
          <td>RestartCount</td>
          <td>{{ .Container.RestartCount }}</td>
        </tr>
      </tbody>
    </table>
  </div>
  <div>
    <ul>
      <li>
        <a href="/container/{{ .Container.ID }}/logs">Logs TODO</a>
      </li>
      <li>
        <a href="/container/{{ .Container.ID }}/inspect">Inspect TODO</a>
      </li>
      <li>
        <a href="/container/{{ .Container.ID }}/stats">Stats TODO</a>
      </li>
      <li>
        <a href="/container/{{ .Container.ID }}/console">Console TODO</a>
      </li>
      <li>
        <a href="/container/{{ .Container.ID }}/attach">Attach TODO</a>
      </li>
    </ul>
  </div>
</div>
    )
}