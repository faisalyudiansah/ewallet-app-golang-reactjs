import React from 'react';
import style from './alertSuccess.module.css';

export const AlertSuccess: React.FC<{
  message: string | null;
}> = ({ message }) => {
  return (
    <div className={style['modal']}>
      <div className={style['modalContent']}>{message}</div>
    </div>
  );
};
