import { ApiInfoModel } from '../../services/info_service';
import React from 'react';

interface ClientFrameProps {
  info: ApiInfoModel;
}

export default function ClientFrame({ info }: ClientFrameProps) {
  return (
    <div>
      <h2>Client</h2>

      <table className='table-small'>
        <tbody>
          <tr>
            <td>DaemonHost</td>
            <td className='text-right'> {info.daemon_host}</td>
          </tr>
          <tr>
            <td>ClientVersion</td>
            <td className='text-right'>{info.client_version}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}
