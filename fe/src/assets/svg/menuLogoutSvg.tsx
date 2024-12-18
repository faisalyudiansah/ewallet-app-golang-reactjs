import React from 'react';
import { PropsColorMenuSvg } from '@/constants/types/props';

export const MenuLogoutSvg: React.FC<PropsColorMenuSvg> = ({ color }) => {
  return (
    <svg
      width="25"
      height="19"
      viewBox="0 0 25 19"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M16 0.5H2C0.89 0.5 0 1.39 0 2.5V6.5H2V2.5H16V16.5H2V12.5H0V16.5C0 17.0304 0.210714 17.5391 0.585786 17.9142C0.960859 18.2893 1.46957 18.5 2 18.5H16C16.5304 18.5 17.0391 18.2893 17.4142 17.9142C17.7893 17.5391 18 17.0304 18 16.5V2.5C18 1.39 17.1 0.5 16 0.5ZM7.08 13.08L8.5 14.5L13.5 9.5L8.5 4.5L7.08 5.91L9.67 8.5H0V10.5H9.67L7.08 13.08Z"
        fill={color}
      />
    </svg>
  );
};
