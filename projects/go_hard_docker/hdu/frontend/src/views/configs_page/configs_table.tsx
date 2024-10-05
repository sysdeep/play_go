import { Link } from 'react-router-dom';
import { ApiConfigListModel } from '../../services/configs_service';
import React from 'react';
import { route, join_url } from '../../routes';

interface ConfigsTableProps {
  configs: ApiConfigListModel[];
}

export default function ConfigsTable({ configs }: ConfigsTableProps) {
  const configs_view = configs.map((config, idx) => {
    return (
      <tr key={idx}>
        <td>
          <Link to={join_url(route.config, config.id)}>{config.name}</Link>
        </td>
        <td> {config.created} </td>
        <td> {config.updated} </td>
        <td>
          {/* <a
            href={'/configs/actions/remove/' + config.name}
            className='button1 error'
          >
            <i className='fa fa-trash-o' aria-hidden='true'></i>
            Remove TODO
          </a> */}
        </td>
      </tr>
    );
  });
  return (
    <table className='table-small'>
      <thead>
        <tr>
          <th>Name</th>
          <th>Created</th>
          <th>Updated</th>
          <th>Options</th>
        </tr>
      </thead>
      <tbody>{configs_view}</tbody>
    </table>
  );
}