import React from 'react';

export default function VolumesFrame() {
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
          <tbody>
            range .Mounts
            <tr>
              <td>
                <a href='/volumes/.Name'>.Name</a>
              </td>
              <td>.Destination</td>
            </tr>
            end
          </tbody>
        </table>
      </div>
    </div>
  );
}
