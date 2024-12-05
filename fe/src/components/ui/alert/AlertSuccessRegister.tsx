import React from 'react';
import style from './alertSuccessRegister.module.css';

export const AlertSuccessRegister: React.FC<{
  message: string;
}> = ({ message }) => {
  return (
    <div className={style['modal']}>
      <div className={style['modalContent']}>{message}</div>
    </div>
  );
};
