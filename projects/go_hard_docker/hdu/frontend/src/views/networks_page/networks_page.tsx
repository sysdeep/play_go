import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import NetworksTable from './networks_table';
import TotalReport from './total_report';
import {
  ApiNetworkListModel,
  NetworksService,
} from '../../services/networks_service';
import IconNetworks from '../../components/icon_networks';
import { useConfiguration } from '@src/store/configuration';

export default function NetworksPage() {
  const { configuration } = useConfiguration();

  const networks_service = useMemo(() => {
    return new NetworksService(configuration.base_url);
  }, []);

  const [networks, setNetworks] = useState<ApiNetworkListModel[]>([]);

  const refresh = () => {
    networks_service
      .get_networks()
      .then((networks: ApiNetworkListModel[]) => {
        setNetworks(networks);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page networks mounted');
    refresh();
  }, []);

  const on_remove = (id: string) => {
    networks_service
      .remove_network(id)
      .then(() => {
        refresh();
      })
      .catch((err) => {
        console.log(err);
      });
  };

  return (
    <div>
      <PageTitle>
        <IconNetworks />
        &nbsp; Networks
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

      <NetworksTable networks={networks} on_remove={on_remove} />
      <TotalReport total={networks.length} />
    </div>
  );
}
