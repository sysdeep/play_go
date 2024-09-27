import ImageListModel from '../models/image_list_model';

interface ApiImageListModel {
  containers: number;
  created: string;
  id: string;
  tags: string[];
  size: number;
}

interface ApiImagesListModel {
  images: ApiImageListModel[];
  total: number;
}

export default class ImagesService {
  constructor() {
    console.log('images_service created');
  }

  async get_images(): Promise<ImageListModel[]> {
    let response = await fetch('http://localhost:1313/api/images');

    let data = (await response.json()) as ApiImagesListModel;
    console.log(data);

    let dataset = data.images.map((model) => {
      let dmodel: ImageListModel = {
        id: model.id,
        created: model.created,
        tags: model.tags,
        size: model.size,
      };
      return dmodel;
    });

    return dataset;
  }

  async remove_image(id: string): Promise<void> {
    await fetch('http://localhost:1313/api/images/' + id, {
      method: 'DELETE',
    });

    return;
  }
}
