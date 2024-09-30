import VolumesService from '../../services/volumes_service';
import IconVolumes from '../../components/icon_volumes';
import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import { ApiVolumeListModel } from '../../services/volumes_service';
import VolumesTable from './volumes_table';
import TotalReport from './total_report';

export default function VolumesPage() {
  const volumes_service = useMemo(() => {
    return new VolumesService();
  }, []);

  const [volumes, setVolumes] = useState<ApiVolumeListModel[]>([]);

  const refresh = () => {
    volumes_service
      .get_volumes()
      .then((volumes: ApiVolumeListModel[]) => {
        setVolumes(volumes);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page volumes mounted');
    refresh();
  }, []);

  return (
    <div>
      <PageTitle>
        <IconVolumes />
        &nbsp; Volumes
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

      <VolumesTable volumes={volumes} />
      <TotalReport total={volumes.length} />
    </div>
  );
}
