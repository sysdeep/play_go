import { ApiContainerResponseModel } from '../../services/containers_service';
import React from 'react';

interface VolumesFrameProps {
  container: ApiContainerResponseModel;
}

export default function VolumesFrame({ container }: VolumesFrameProps) {
  const rows_view = container.mounts.map((volume, idx) => {
    return (
      <tr key={idx}>
        <td>
          <a href={'/volumes/' + volume.name}>{volume.name}</a>
        </td>
        <td>{volume.destination}</td>
      </tr>
    );
  });
  return (
    <div>
      <h2>Volumes</h2>
      <div>
        <table>
          <thead>
            <tr>
              <th>Host/volume</th>
              <th>Path in container</th>
            </tr>
          </thead>
          <tbody>{rows_view}</tbody>
        </table>
      </div>
    </div>
  );
}
