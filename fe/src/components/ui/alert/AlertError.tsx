import React from 'react';
import style from './alertError.module.css';

export const AlertError: React.FC<{
  message: string | null;
}> = ({ message }) => {
  return (
    <div className={style['modal']}>
      <div className={style['modalContent']}>{message}</div>
    </div>
  );
};
