import { useParams } from 'react-router-dom';
import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import DetailsFrame from './detailes_frame';
import {
  ApiFullSecretModel,
  SecretsService,
} from '../../services/secrets_service';
import IconSecrets from '../../components/icon_secrets';

export default function SecretPage() {
  const { id } = useParams();

  const secret_service = useMemo(() => {
    return new SecretsService();
  }, []);

  const [secret, setSecret] = useState<ApiFullSecretModel | null>(null);

  const refresh = () => {
    secret_service
      .get_secret(id)
      .then((secret) => {
        setSecret(secret);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page secret mounted');
    refresh();
  }, []);

  const body = () => {
    if (secret) {
      return (
        <div>
          <DetailsFrame secret={secret} />
        </div>
      );
    }

    return <p>no secret</p>;
  };

  return (
    <div>
      <PageTitle>
        <IconSecrets />
        &nbsp; Secret {id}
      </PageTitle>

      {body()}
    </div>
  );
}
