import React from 'react';

export default function SwarmFrame() {
  return (
    <div>
      <h2>Sawrm</h2>

      <table className='table-small'>
        <tbody>
          <tr>
            <td>Node id</td>
            <td className='text-right'> .SystemInfo.Swarm.NodeID TODO</td>
          </tr>
          <tr>
            <td>Node address</td>
            <td className='text-right'> .SystemInfo.Swarm.NodeAddr TODO</td>
          </tr>
          <tr>
            <td></td>
            <td className='text-right'></td>
          </tr>
          <tr>
            <td></td>
            <td className='text-right'></td>
          </tr>
        </tbody>
      </table>
    </div>
  );
}
