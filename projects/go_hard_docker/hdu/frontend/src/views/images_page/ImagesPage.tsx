import React, { useState, useEffect, useCallback, useMemo } from 'react';
import PageTitle from '../../components/PageTitle';
import ImagesTable from './ImagesTable';
import ImageListModel from '../../models/image_list_model';
import ImagesService from '../../services/images_service';
import FilterPanel from './FilterPanel';
import FilterModel from './FilterModel';
import TotalReport from './TotalReport';

export default function ImagesPage() {
  const images_service = useMemo(() => {
    return new ImagesService();
  }, []);

  const [images, setImages] = useState<ImageListModel[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [filter, setFilter] = useState<FilterModel>({ dates: [] });

  const refresh = () => {
    setLoading(true);
    images_service
      .get_images()
      .then((images: ImageListModel[]) => {
        setImages(images);
      })
      .catch((err) => {
        console.log(err);
      })
      .finally(() => {
        setLoading(false);
      });
  };

  useEffect(() => {
    console.log('page images mounted');
    refresh();
  }, []);

  const remove_image = (id: string) => {
    console.log('remove', id);

    images_service
      .remove_image(id)
      .then(() => {
        refresh();
      })
      .catch((err) => {
        console.log(err);
      });
  };

  const on_date = (date: string) => {
    let result = [];
    if (filter.dates.includes(date)) {
      result = filter.dates.filter((d) => d !== date);
    } else {
      result = [...filter.dates, date];
    }
    setFilter({ ...filter, dates: result });
  };

  return (
    <div>
      <PageTitle>Images</PageTitle>
      <FilterPanel filter={filter} on_date={on_date} />
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
      <TotalReport total={images.length} />
    </div>
  );
}
