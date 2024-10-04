import { useParams } from 'react-router-dom';
import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import DetailsFrame from './detailes_frame';
import {
  ApiFullSecretModel,
  SecretsService,
} from '../../services/secrets_service';
import IconSecrets from '../../components/icon_secrets';
import {
  ApiFullConfigModel,
  ConfigsServices,
} from '../../services/configs_service';

export default function ConfigPage() {
  const { id } = useParams();

  const config_service = useMemo(() => {
    return new ConfigsServices();
  }, []);

  const [config, setConfig] = useState<ApiFullConfigModel | null>(null);

  const refresh = () => {
    config_service
      .get_config(id)
      .then((config) => {
        setConfig(config);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page config mounted');
    refresh();
  }, []);

  const body = () => {
    if (config) {
      return (
        <div>
          <DetailsFrame config={config} />
        </div>
      );
    }

    return <p>no config</p>;
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
