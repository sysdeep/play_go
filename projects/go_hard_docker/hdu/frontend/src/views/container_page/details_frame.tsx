import React from 'react';
import { ApiContainerResponseModel } from '../../services/containers_service';

interface DetailsFrameProps {
  container: ApiContainerResponseModel;
}

export default function DetailsFrame({ container }: DetailsFrameProps) {
  return (
    <div>
      <h2>Container details</h2>
      <div>
        <table>
          <tbody>
            <tr>
              <td>Image</td>
              <td>
                <a href='/images/{container.config.image}'>
                  {container.config.image}
                </a>
                {/* {container.config.image} */}
              </td>
            </tr>
            <tr>
              <td>Ports</td>
              <td>
                range $pt, $ports := .Network.Ports
                <ul>
                  {/* {container.network.ports} */}
                  range $ports
                  <li> .HostIP : .HostPort "-" $pt </li>
                  end
                </ul>
                end
              </td>
            </tr>
            <tr>
              <td>CMD</td>
              <td>{container.config.cmd}</td>
            </tr>
            <tr>
              <td>Entrypoint</td>
              <td>{container.config.entrypoint}</td>
            </tr>
            <tr>
              <td>ENV</td>
              <td>
                <EnvTable env={container.config.env} />
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
  );
}

interface EnvTableProps {
  env: string[];
}
function EnvTable({ env }: EnvTableProps) {
  const rows = env.map((row, idx) => {
    return (
      <tr key='idx'>
        <td>{row}</td>
      </tr>
    );
  });

  return (
    <table>
      <tbody>{rows}</tbody>
    </table>
  );
}
