import { Link } from 'react-router-dom';
import { ApiSecretListModel } from '../../services/secrets_service';
import React from 'react';

interface SecretsTableProps {
  secrets: ApiSecretListModel[];
}

export default function SecretsTable({ secrets }: SecretsTableProps) {
  const secrets_view = secrets.map((secret, idx) => {
    return (
      <tr key={idx}>
        <td>
          <Link to={'/secret/' + secret.id}>{secret.name}</Link>
        </td>
        <td> {secret.created} </td>
        <td> {secret.updated} </td>
        {/* <td>
          <a
            href={'/secrets/actions/remove/' + secret.name}
            className='button1 error'
          >
            <i className='fa fa-trash-o' aria-hidden='true'></i>
            Remove TODO
          </a>
        </td> */}
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
      <tbody>{secrets_view}</tbody>
    </table>
  );
}
