import PageTitle from '../../components/PageTitle';
import React, { useEffect, useMemo, useState } from 'react';
import ContainersTable from './ContainersTable';
import ContainerListModel from '../../models/container_list_model';
import ContainersService from '../../services/containers_service';
import TotalReport from './TotalReport';

export default function ContainersPage() {
  const containers_service = useMemo(() => {
    return new ContainersService();
  }, []);

  const [containers, setContainers] = useState<ContainerListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  // const [filter, setFilter] = useState<FilterModel>({ dates: [] });

  const refresh = () => {
    setLoading(true);
    containers_service
      .get_containers()
      .then((containers: ContainerListModel[]) => {
        setContainers(containers);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    console.log('containers page mounted');
    refresh();
  }, []);

  // const remove_image = (id: string) => {
  //   console.log('remove', id);

  //   containers_service
  //     .remove_image(id)
  //     .then(() => {
  //       refresh();
  //     })
  //     .catch((err) => {
  //       console.log(err);
  //     });
  // };

  // const on_date = (date: string) => {
  //   let result = [];
  //   if (filter.dates.includes(date)) {
  //     result = filter.dates.filter((d) => d !== date);
  //   } else {
  //     result = [...filter.dates, date];
  //   }
  //   setFilter({ ...filter, dates: result });
  // };

  return (
    <div>
      <PageTitle>Containers</PageTitle>
      <ContainersTable containers={containers} />
      <TotalReport total={containers.length} />
      {/* <FilterPanel filter={filter} on_date={on_date} />
      <div>
        <span>loading: {loading}</span>
        <button className='button' onClick={() => refresh()}>
          Refresh
        </button>
      </div>
      <ImagesTable
        images={images}
        filter={filter}
        on_remove={remove_image}
        on_date={on_date}
      />
       */}
    </div>
  );
}
