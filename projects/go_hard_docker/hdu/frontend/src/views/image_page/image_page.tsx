import { useParams } from 'react-router-dom';
import IconImages from '../../components/icon_images';
import PageTitle from '../../components/page_title';
import React, { useEffect, useMemo, useState } from 'react';
import DetailsFrame from './detailes_frame';
import ContainersFrame from './containers_frame';
import HistoryFrame from './history_frame';
import ImagesService, {
  ApiFullImageModel,
} from '../../services/images_service';

export default function ImagePage() {
  const { id } = useParams();

  const images_service = useMemo(() => {
    return new ImagesService();
  }, []);

  const [image, setImage] = useState<ApiFullImageModel | null>(null);

  const refresh = () => {
    images_service
      .get_image(id)
      .then((image) => {
        console.log(image);
        setImage(image);
      })
      .catch((err) => {
        console.log(err);
      });
  };

  useEffect(() => {
    console.log('page image mounted');
    refresh();
  }, []);

  const body = () => {
    if (image) {
      return (
        <div>
          <DetailsFrame image={image} />
          <ContainersFrame image={image} />
          <HistoryFrame image={image} />
        </div>
      );
    }

    return <p>no image</p>;
  };

  const [, image_hash] = id.split(':');
  const page_title = image_hash.slice(0, 12);

  return (
    <div>
      <PageTitle>
        <IconImages />
        &nbsp; Image: {page_title}
      </PageTitle>

      {body()}
    </div>
  );
}
