import React from "react";


export default function DetailsFrame(){
    return (

<div>
  <h2>Container details</h2>
  <div>
    <table>
      <tbody>
        <tr>
          <td>Image</td>
          <td>
            <a href="/images/{{ .Container.Image }}">{{ .Config.Image }}</a>
            {{ .Container.Image }}
          </td>
        </tr>
        <tr>
          <td>Ports</td>
          <td>
            {{ range $pt, $ports := .Network.Ports}}
            <ul>
              {{ range $ports}}
              <li>{{ .HostIP }}:{{ .HostPort }} -> {{ $pt }}</li>
              {{ end }}
            </ul>
            {{ end }}
          </td>
        </tr>
        <tr>
          <td>CMD</td>
          <td>{{ .Config.Cmd }}</td>
        </tr>
        <tr>
          <td>Entrypoint</td>
          <td>{{ .Config.Entrypoint }}</td>
        </tr>
        <tr>
          <td>ENV</td>
          <td>
            <table>
              <tbody>
                {{range .Config.Env}}
                <tr>
                  <td>{{ . }}</td>
                </tr>
                {{end}}
              </tbody>
            </table>
          </td>
        </tr>
        <tr>
          <td>Restart policies</td>
          <td>TODO</td>
        </tr>
        <tr>
          <td></td>
          <td></td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
    )
}