import ImageListModel from '../models/image_list_model';

export default class ImagesService {
  constructor() {
    console.log('images_service created');
  }

  async get_images(): Promise<ImageListModel[]> {
    const response = await fetch('http://localhost:1313/api/images');

    const data = (await response.json()) as ApiImagesListModel;

    const images = data.images || [];
    const dataset = images.map((model) => {
      const dmodel: ImageListModel = {
        id: model.id,
        created: model.created,
        tags: model.tags,
        size: model.size,
      };
      return dmodel;
    });

    return dataset;
  }

  async get_image(id: string): Promise<ApiFullImageModel> {
    const response = (await fetch(
      'http://localhost:1313/api/images/' + id,
    ).then((data) => data.json())) as ApiFullImageModel;

    return response;
  }

  async remove_image(id: string): Promise<void> {
    await fetch('http://localhost:1313/api/images/' + id, {
      method: 'DELETE',
    });

    return;
  }
}

// images list ----------------------------------------------------------------
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

// image models ---------------------------------------------------------------
interface ApiImageModel {
  id: string;
  repo_tags: string[];
  parent: string;
  comment: string;
  created: string;
  size: number;
}

interface ApiImageHistoryModel {
  created: string;
  id: string;
  size: number;
  tags: string[];
}
interface ApiImageContainerModel {
  id: string;
  name: string;
}

export interface ApiFullImageModel {
  image: ApiImageModel;
  history: ApiImageHistoryModel[];
  containers: ApiImageContainerModel[];
}
