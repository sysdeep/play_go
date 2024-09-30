import React from 'react';
import { ApiContainerResponseModel } from '../../services/containers_service';

interface NetworksFrameProps {
  container: ApiContainerResponseModel;
}

export default function NetworksFrame({ container }: NetworksFrameProps) {
  const networks_view = Object.keys(container.network.networks).map(
    (endpoint, idx) => {
      const net = container.network.networks[endpoint];
      return (
        <tr key={idx}>
          <td>
            <a href={'/networks/' + net.network_id}>{endpoint}</a>
          </td>
          <td>{net.ip_address}</td>
          <td>{net.gateway}</td>
          <td>{net.mac_address}</td>
          {/* <!-- TODO --> */}
          {/* <!-- <td> --> */}
          {/* <!--   <a href="TODO">Leave TODO</a> --> */}
          {/* <!-- </td> --> */}
        </tr>
      );
    },
  );
  return (
    <div>
      <h2>Networks</h2>
      {/* <!-- TODO --> */}
      {/* <!-- <div>Connect to: TODO</div> --> */}
      <div>
        <table>
          <thead>
            <tr>
              <th>Network</th>
              <th>IP Address</th>
              <th>Gateway</th>
              <th>MAC</th>
              {/* <!-- <th>Actions</th> --> */}
            </tr>
          </thead>
          <tbody>{networks_view}</tbody>
        </table>
      </div>
    </div>
  );
}
