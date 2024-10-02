import React from 'react';

export default function ServerFrame() {
  return (
    <div>
      <h2>Server</h2>

      <table className='table-small'>
        <tbody>
          <tr>
            <td>Hostname</td>
            <td className='text-right'> .SystemInfo.Name </td>
          </tr>
          <tr>
            <td>Server version</td>
            <td className='text-right'> .SystemInfo.ServerVersion </td>
          </tr>
          <tr>
            <td>
              {/* <!-- <i className="fa fa-cubes" aria-hidden="true"></i> --> */}
              <a href='/containers'>Containers</a>
            </td>
            <td className='text-right'>
              <span className='ml-2' title='total'>
                <i className='fa fa-cubes' aria-hidden='true'></i>
                .SystemInfo.Containers
              </span>
              <span className='ml-2' title='running'>
                <i className='fa fa-play text-success' aria-hidden='true'></i>
                .SystemInfo.ContainersRunning
              </span>
              <span className='ml-2' title='stopped'>
                <i className='fa fa-stop text-error' aria-hidden='true'></i>
                .SystemInfo.ContainersStopped
              </span>
              <span className='ml-2' title='paused'>
                <i className='fa fa-pause text-grey' aria-hidden='true'></i>
                .SystemInfo.ContainersPaused
              </span>
            </td>
          </tr>
          <tr>
            <td>
              <a href='/images'>Images</a>
            </td>
            <td className='text-right'> .SystemInfo.Images </td>
          </tr>
          <tr>
            <td>OperatingSystem</td>
            <td className='text-right'>
              .SystemInfo.OperatingSystem ( .SystemInfo.KernelVersion )
            </td>
          </tr>
          <tr>
            <td>Address pool</td>
            <td className='text-right'>
              range .SystemInfo.DefaultAddressPools [
              <span> .Base size: .Size </span>] end
            </td>
          </tr>
          <tr>
            <td>Default runtime</td>
            <td className='text-right'> .SystemInfo.DefaultRuntime</td>
          </tr>
          <tr>
            <td>HW</td>
            <td className='text-right'>
              <span className='ml-2'>
                <i className='fa fa-laptop' aria-hidden='true'></i>
                .SystemInfo.NCPU
              </span>
              <span className='ml-2'>
                <i className='fa fa-pie-chart' aria-hidden='true'></i>
                .SystemInfo.MemTotal
              </span>
            </td>
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
