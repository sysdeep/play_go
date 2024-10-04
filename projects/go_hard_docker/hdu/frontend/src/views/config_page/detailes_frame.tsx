import React from 'react';
import { ApiFullConfigModel } from '../../services/configs_service';

interface DetailsFrameProps {
  config: ApiFullConfigModel;
}

export default function DetailsFrame({ config }: DetailsFrameProps) {
  return (
    <div>
      <h2>Secret details</h2>
      <div>
        <table>
          <tbody>
            <tr>
              <td>ID</td>
              <td>{config.config.id}</td>
            </tr>
            <tr>
              <td>Name</td>
              <td>{config.config.name}</td>
            </tr>
            <tr>
              <td>Created</td>
              <td>{config.config.created}</td>
            </tr>
            <tr>
              <td>Updated</td>
              <td>{config.config.updated}</td>
            </tr>
          </tbody>
        </table>

        <pre>{config.config.data_text}</pre>

        {/* <div>
      <!-- <a href="/volumes/actions/remove/{ . }" class="button error"> -->
      <!--   <i class="fa fa-trash-o" aria-hidden="true"></i> -->
      <!--   Remove -->
      <!-- </a> -->
    </div> */}
      </div>
    </div>
  );
}
