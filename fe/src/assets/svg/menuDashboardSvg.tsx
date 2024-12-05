import { PropsColorMenuSvg } from '@/constants/types/props';
import React from 'react';

export const MenuDashboardSvg: React.FC<PropsColorMenuSvg> = ({ color }) => {
  return (
    <svg
      width="24"
      height="25"
      viewBox="0 0 24 25"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M21 16.5V4.5H3V16.5H21ZM21 2.5C21.5304 2.5 22.0391 2.71071 22.4142 3.08579C22.7893 3.46086 23 3.96957 23 4.5V16.5C23 17.0304 22.7893 17.5391 22.4142 17.9142C22.0391 18.2893 21.5304 18.5 21 18.5H14V20.5H16V22.5H8V20.5H10V18.5H3C1.89 18.5 1 17.6 1 16.5V4.5C1 3.39 1.89 2.5 3 2.5H21ZM5 6.5H14V11.5H5V6.5ZM15 6.5H19V8.5H15V6.5ZM19 9.5V14.5H15V9.5H19ZM5 12.5H9V14.5H5V12.5ZM10 12.5H14V14.5H10V12.5Z"
        fill={color}
      />
    </svg>
  );
};
