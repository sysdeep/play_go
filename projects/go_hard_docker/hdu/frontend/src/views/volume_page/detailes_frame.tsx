import { ApiFullImageModel } from '../../services/images_service';
import React from 'react';
import { format_size } from '../../utils/humanize';
import { ApiFullVolumeModel } from '../../services/volumes_service';

interface DetailsFrameProps {
  volume: ApiFullVolumeModel;
}

export default function DetailsFrame({ volume }: DetailsFrameProps) {
  return (
    <div>
      <h2>Volume details</h2>
      <div>
        <table>
          <tbody>
            <tr>
              <td>ID</td>
              <td>{volume.volume.name}</td>
            </tr>
            <tr>
              <td>Created</td>
              <td>{volume.volume.created}</td>
            </tr>
            <tr>
              <td>Mount path</td>
              <td>{volume.volume.mount_point}</td>
            </tr>
            <tr>
              <td>Driver</td>
              <td>{volume.volume.driver}</td>
            </tr>
          </tbody>
        </table>

        {/* <div>
          <a
            href='/volumes/actions/remove/{ volume.volume.Name }'
            className='button error'
          >
            <i className='fa fa-trash-o' aria-hidden='true'></i>
            Remove
          </a>
        </div> */}
      </div>
    </div>
  );
}
