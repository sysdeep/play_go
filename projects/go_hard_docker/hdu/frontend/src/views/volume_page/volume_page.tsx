import { useParams } from 'react-router-dom';
import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import DetailsFrame from './detailes_frame';
import VolumesService, {
  ApiFullVolumeModel,
} from '../../services/volumes_service';
import IconVolumes from '../../components/icon_volumes';

export default function VolumePage() {
  const { id } = useParams();

  const volume_service = useMemo(() => {
    return new VolumesService();
  }, []);

  const [volume, setVolume] = useState<ApiFullVolumeModel | null>(null);

  const refresh = () => {
    volume_service
      .get_volume(id)
      .then((volume) => {
        setVolume(volume);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page volume mounted');
    refresh();
  }, []);

  const body = () => {
    if (volume) {
      return (
        <div>
          <DetailsFrame volume={volume} />
        </div>
      );
    }

    return <p>no volume</p>;
  };

  return (
    <div>
      <PageTitle>
        <IconVolumes />
        &nbsp; Volume: {id}
      </PageTitle>

      {body()}
    </div>
  );
}
