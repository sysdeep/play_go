import React from 'react';
import { ApiContainerResponseModel } from '../../services/containers_service';

interface DetailsFrameProps {
  container: ApiContainerResponseModel;
}

export default function DetailsFrame({ container }: DetailsFrameProps) {
  const ports_view = Object.keys(container.network.ports).map(
    (port_name, idx) => {
      const values = container.network.ports[port_name];

      if (!values) {
        return <li key={idx * 100000}>{port_name} - not defined</li>;
      }

      return values.map((segment, idi) => {
        return (
          <li key={(idx + 1) * (idi + 2)}>
            {port_name} - {segment.host_ip}:{segment.host_port}
          </li>
        );
      });
    },
  );

  return (
    <div>
      <h2>Container details</h2>
      <div>
        <table>
          <tbody>
            <tr>
              <td>Image</td>
              <td>
                <a href={'/images/' + container.container.image}>
                  {container.config.image}
                </a>
                {/* {container.config.image} */}
              </td>
            </tr>
            <tr>
              <td>Ports</td>
              <td>
                <ul>{ports_view}</ul>
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
    return <li key={idx}>{row}</li>;
  });

  return <ul>{rows}</ul>;
}
