import React, { useEffect, useMemo } from 'react';
import { useParams } from 'react-router-dom';

import ActionsFrame from './actions_frame';
import PageTitle from '../../components/PageTitle';
import ContainersService from '../../services/containers_service';

export default function ContainerPage() {
  const { id } = useParams();
  
  const containers_service = useMemo(() => {
    return new ContainersService();
  }, []);

  const refresh = () => {
    // setLoading(true);
    containers_service
      .get_container(id)
      .then(data => {
        console.log(data)
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        // setLoading(false);
      });
  };

  useEffect(() => {
    console.log('container page mounted');
    refresh();
  }, []);

  return (
    <div>
      <PageTitle>Container: {id}</PageTitle>
      <ActionsFrame />
    </div>
  );
}
