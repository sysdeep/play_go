import { Link } from 'react-router-dom';
import { ApiNetworkListModel } from '../../services/networks_service';
import React from 'react';
import { route, join_url } from '../../routes';

interface NetworksTableProps {
  networks: ApiNetworkListModel[];
}

export default function NetworksTable({ networks }: NetworksTableProps) {
  const networks_view = networks.map((network, idx) => {
    return (
      <tr key={idx}>
        <td>
          <Link to={join_url(route.network, network.id)}>{network.name}</Link>
        </td>
        <td> {network.driver} </td>
        <td> {network.created} </td>
        <td>
          <a
            href={'/volumes/actions/remove/' + network.name}
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
    <table className='table-small striped'>
      <thead>
        <tr>
          <th>Name</th>
          <th>Driver</th>
          <th>Created</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>{networks_view}</tbody>
    </table>
  );
}
