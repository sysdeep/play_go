import ContainerListModel from '@src/models/container_list_model';
import React from 'react';

interface ContainersTableProps {
  containers: ContainerListModel[];
}

export default function ContainersTable({ containers }: ContainersTableProps) {
  const rows = containers.map((container, idx) => {
    return <ContainerRow container={container} key={idx} />;
  });

  return (
    <table>
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
        <a href='/containers/{container.id}'>{container.name}</a>
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
