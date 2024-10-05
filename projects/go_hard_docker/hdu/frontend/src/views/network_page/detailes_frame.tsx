import React from 'react';
import { format_size } from '../../utils/humanize';
import { ApiFullNetworkModel } from '../../services/networks_service';

interface DetailsFrameProps {
  network: ApiFullNetworkModel;
}

export default function DetailsFrame({ network }: DetailsFrameProps) {
  return (
    <div>
      <h2>Network details</h2>
      <div>
        <table>
          <tbody>
            <tr>
              <td>ID</td>
              <td>{network.network.id}</td>
            </tr>
            <tr>
              <td>Name</td>
              <td>{network.network.name}</td>
            </tr>
            <tr>
              <td>Created</td>
              <td>{network.network.created}</td>
            </tr>
            <tr>
              <td>Driver</td>
              <td>{network.network.driver}</td>
            </tr>
            <tr>
              <td>Scope</td>
              <td>{network.network.scope}</td>
            </tr>
            <tr>
              <td>Internal</td>
              <td>{network.network.internal}</td>
            </tr>
            <tr>
              <td>Attachable</td>
              <td>{network.network.attachable}</td>
            </tr>
            <tr>
              <td>Ingress</td>
              <td>{network.network.ingress}</td>
            </tr>
          </tbody>
        </table>
      </div>
    </div>
  );
}