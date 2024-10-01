import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import TotalReport from './total_report';
import {
  ApiConfigListModel,
  ConfigsServices,
} from '../../services/configs_service';
import IconConfigs from '../../components/icon_configs';
import ConfigsTable from './configs_table';

export default function ConfigsPage() {
  const configs_service = useMemo(() => {
    return new ConfigsServices();
  }, []);

  const [configs, setConfigs] = useState<ApiConfigListModel[]>([]);

  const refresh = () => {
    configs_service
      .get_configs()
      .then((configs: ApiConfigListModel[]) => {
        setConfigs(configs);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page configs mounted');
    refresh();
  }, []);

  return (
    <div>
      <PageTitle>
        <IconConfigs />
        &nbsp; Configs
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

      <ConfigsTable configs={configs} />
      <TotalReport total={configs.length} />
    </div>
  );
}
