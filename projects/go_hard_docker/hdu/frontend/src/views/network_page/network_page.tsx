import { useParams } from 'react-router-dom';
import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import DetailsFrame from './detailes_frame';
import {
  ApiFullNetworkModel,
  NetworksService,
} from '../../services/networks_service';
import IconNetworks from '../../components/icon_networks';
import ContainersFrame from './containers_frame';

export default function NetworkPage() {
  const { id } = useParams();

  const network_service = useMemo(() => {
    return new NetworksService();
  }, []);

  const [network, setNetwork] = useState<ApiFullNetworkModel | null>(null);

  const refresh = () => {
    network_service
      .get_network(id)
      .then((network) => {
        setNetwork(network);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page network mounted');
    refresh();
  }, []);

  const body = () => {
    if (network) {
      return (
        <div>
          <DetailsFrame network={network} />
          <ContainersFrame containers={network.containers} />
        </div>
      );
    }

    return <p>no network</p>;
  };

  return (
    <div>
      <PageTitle>
        <IconNetworks />
        &nbsp; Network {id}
      </PageTitle>

      {body()}
    </div>
  );
}
