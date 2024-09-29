import React from 'react';

export default function NetworksFrame() {
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
          <tbody>
            range $endpoint, $ep_settings := .Network.Networks
            <tr>
              <td>
                <a href='/networks/ $ep_settings.NetworkID '> $endpoint </a>
              </td>
              <td> $ep_settings.IPAddress </td>
              <td> $ep_settings.Gateway </td>
              <td> $ep_settings.MacAddress </td>
              {/* <!-- TODO --> */}
              {/* <!-- <td> --> */}
              {/* <!--   <a href="TODO">Leave TODO</a> --> */}
              {/* <!-- </td> --> */}
            </tr>
            end
          </tbody>
        </table>
      </div>
    </div>
  );
}
