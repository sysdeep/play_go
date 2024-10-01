import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import TotalReport from './total_report';
import {
  ApiSecretListModel,
  SecretsService,
} from '../../services/secrets_service';
import IconSecrets from '../../components/icon_secrets';
import SecretsTable from './secrets_table';

export default function SecretsPage() {
  const secrets_service = useMemo(() => {
    return new SecretsService();
  }, []);

  const [secrets, setSecrets] = useState<ApiSecretListModel[]>([]);

  const refresh = () => {
    secrets_service
      .get_secrets()
      .then((secrets: ApiSecretListModel[]) => {
        setSecrets(secrets);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page secrets mounted');
    refresh();
  }, []);

  return (
    <div>
      <PageTitle>
        <IconSecrets />
        &nbsp; Secrets
      </PageTitle>

      {/* // TODO //{' '} */}
      {/* <div>
        //{' '}
        <a href='/volumes/actions/prune' class='button error'>
          // <i class='fa fa-trash-o' aria-hidden='true'></i>
          // Prune //{' '}
        </a>
        //{' '}
      </div> */}

      <SecretsTable secrets={secrets} />
      <TotalReport total={secrets.length} />
    </div>
  );
}
