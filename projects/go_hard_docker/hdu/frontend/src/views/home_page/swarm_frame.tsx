import { ApiInfoModel } from '../../services/info_service';
import React from 'react';

interface SwarmFrameProps {
  info: ApiInfoModel;
}
export default function SwarmFrame({ info }: SwarmFrameProps) {
  return (
    <div>
      <h2>Sawrm</h2>

      <table className='table-small'>
        <tbody>
          <tr>
            <td>Node id</td>
            <td className='text-right'>{info.system.swarm.node_id}</td>
          </tr>
          <tr>
            <td>Node address</td>
            <td className='text-right'>{info.system.swarm.node_addr}</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}
