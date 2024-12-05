import React from 'react';
import style from './dashboard.module.css';
import { DataUserProfile } from 'src/constants/response/resProfile';

export const DashboardHeadertitle: React.FC<{
  dataUserProfile: DataUserProfile | null;
}> = ({ dataUserProfile }) => {
  return (
    <div>
      <span className={style['header-text']}>
        Halo, {dataUserProfile?.name}!
      </span>
    </div>
  );
};
