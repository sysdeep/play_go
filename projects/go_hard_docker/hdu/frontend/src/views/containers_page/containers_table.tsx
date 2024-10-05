import ContainerListModel from '../../models/container_list_model';
import React from 'react';
import { Link } from 'react-router-dom';
import { route, join_url } from '../../routes';

interface ContainersTableProps {
  containers: ContainerListModel[];
}

export default function ContainersTable({ containers }: ContainersTableProps) {
  const rows = containers.map((container, idx) => {
    return <ContainerRow container={container} key={idx} />;
  });

  return (
    <table className='table-small striped'>
      <thead>
        <tr>
          <th>Name</th>
          <th>State</th>
          <th>Image</th>
          <th>Created</th>
          {/* <th>IP Address</th>
      <th>Published Ports</th>
      <th>Actions</th> */}
        </tr>
      </thead>
      <tbody>{rows}</tbody>
    </table>
  );
}

interface ContainerRowProps {
  container: ContainerListModel;
}

function ContainerRow({ container }: ContainerRowProps) {
  return (
    <tr>
      <td>
        <Link to={join_url(route.container, container.id)}>
          {container.name}
        </Link>
      </td>
      <td>{container.state}</td>
      <td>{container.image}</td>
      <td>{container.created}</td>
      {/* <td></td>
      <td></td>
      <td>
        <a href="/containers/{{.ID}}/inspect">inspect</a>
        <a href="/containers/{{.ID}}/logs">logs</a>
      </td> */}
    </tr>
  );
}
