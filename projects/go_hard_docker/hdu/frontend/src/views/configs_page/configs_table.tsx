import { ApiConfigListModel } from '../../services/configs_service';
import React from 'react';

interface ConfigsTableProps {
  configs: ApiConfigListModel[];
}

export default function ConfigsTable({ configs }: ConfigsTableProps) {
  const configs_view = configs.map((config, idx) => {
    return (
      <tr key={idx}>
        <td>
          <a href={'/configs/' + config.id}>{config.name} TODO</a>
        </td>
        <td> {config.created} </td>
        <td> {config.updated} </td>
        <td>
          <a
            href={'/configs/actions/remove/' + config.name}
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
