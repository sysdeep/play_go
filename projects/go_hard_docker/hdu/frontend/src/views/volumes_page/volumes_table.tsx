import { ApiVolumeListModel } from '@src/services/volumes_service';
import React from 'react';

interface VolumesTableProps {
  volumes: ApiVolumeListModel[];
}

export default function VolumesTable({ volumes }: VolumesTableProps) {
  const volumes_view = volumes.map((volume, idx) => {
    return (
      <tr key={idx}>
        <td>
          <a href={'/volumes/' + volume.name}>{volume.name} TODO</a>
        </td>
        <td> {volume.stack_name} </td>
        <td> {volume.driver} </td>
        {/* <!-- <td> .Mountpoint </td> --> */}
        <td> {volume.created} </td>
        <td>
          <a
            href={'/volumes/actions/remove/' + volume.name}
            className='button1 error'
          >
            <i className='fa fa-trash-o' aria-hidden='true'></i>
            Remove TODO
          </a>
        </td>
      </tr>
    );
  });
  return (
    <table className='table-small'>
      <thead>
        <tr>
          <th>Name</th>
          <th>Stack</th>
          <th>Driver</th>
          {/* <!-- <th>Mount point</th> --> */}
          <th>Created</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>{volumes_view}</tbody>
    </table>
  );
}
