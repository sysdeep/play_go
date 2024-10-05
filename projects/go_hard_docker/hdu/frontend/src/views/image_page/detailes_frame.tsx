import { ApiFullImageModel } from '../../services/images_service';
import React from 'react';
import { format_size } from '../../utils/humanize';

interface DetailsFrameProps {
  image: ApiFullImageModel;
}

export default function DetailsFrame({ image }: DetailsFrameProps) {
  const tags_view = image.image.repo_tags.map((tag, idx) => {
    return <li key={idx}>{tag}</li>;
  });
  return (
    <div>
      <h2>Details</h2>
      <div>
        <table className='table-small'>
          <tbody>
            <tr>
              <td>ID</td>
              <td>{image.image.id}</td>
            </tr>
            <tr>
              <td>Tags</td>
              <td>
                <ul className='table-ui'>{tags_view}</ul>
              </td>
            </tr>
            <tr>
              <td>Parent</td>
              <td>{image.image.parent}</td>
            </tr>
            <tr>
              <td>Comment</td>
              <td>{image.image.comment}</td>
            </tr>
            <tr>
              <td>Size</td>
              <td>{format_size(image.image.size)}</td>
            </tr>
            <tr>
              <td>Created</td>
              <td>{image.image.created}</td>
            </tr>
          </tbody>
        </table>
        {/* TODO */}
        {/* <a href='/images/actions/remove/ .Image.ID ' className='button error'>
          <i className='fa fa-trash-o' aria-hidden='true'></i>
          Remove
        </a> */}
      </div>
    </div>
  );
}
