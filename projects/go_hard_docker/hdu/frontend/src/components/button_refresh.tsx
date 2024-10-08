import React from 'react';
import IconRefresh from './icon_refresh';

interface ButtonRemoveProps {
  //   disabled: boolean;
  on_refresh(): void;
}
export default function ButtonRefresh({
  //   disabled = false,
  on_refresh,
}: ButtonRemoveProps) {
  const on_click = (e: any) => {
    e.preventDefault();
    on_refresh();
  };

  return (
    <button className='button secondary' onClick={on_click}>
      <IconRefresh />
      &nbsp; Refresh
    </button>
  );
}
