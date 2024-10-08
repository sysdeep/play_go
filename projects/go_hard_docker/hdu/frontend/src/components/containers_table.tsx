import IconContainers from '@src/components/icon_containers';
import { ApiContainerListModel } from '@src/models/api_container_list_model';
import { join_url, route } from '@src/routes';
import React from 'react';
import { Link } from 'react-router-dom';

interface ContainersFrameProps {
  containers: ApiContainerListModel[];
}

export default function ContainersTable({ containers }: ContainersFrameProps) {
  const rows_view = containers.map((container, idx) => {
    return (
      <tr key={idx}>
        <td>
          <Link to={join_url(route.container, container.id)}>
            {container.name}
          </Link>
        </td>
        <td>{container.state}</td>
        <td>{container.image}</td>
        <td>{container.created}</td>
      </tr>
    );
  });

  // TODO: возможно стоит завести отдельный компонент...
  return (
    <table className='table table-small striped'>
      <thead>
        <tr>
          <th>Name</th>
          <th>State</th>
          <th>Image</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody>{rows_view}</tbody>
    </table>
  );
}
