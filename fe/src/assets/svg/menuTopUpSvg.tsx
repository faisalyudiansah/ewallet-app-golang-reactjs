import React from 'react';
import { PropsColorMenuSvg } from '@/constants/types/props';

export const MenuTopUpSvg: React.FC<PropsColorMenuSvg> = ({ color }) => {
  return (
    <svg
      width="24"
      height="25"
      viewBox="0 0 24 25"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M18 15.3571V18.7857H14.4V21.0714H18V24.5H20.4V21.0714H24V18.7857H20.4V15.3571H18ZM12.312 21.9171L10.8 23.3571L9 21.6429L7.2 23.3571L5.4 21.6429L3.6 23.3571L1.8 21.6429L0 23.3571V0.5L1.8 2.21429L3.6 0.5L5.4 2.21429L7.2 0.5L9 2.21429L10.8 0.5L12.6 2.21429L14.4 0.5L16.2 2.21429L18 0.5L19.8 2.21429L21.6 0.5V13.4714C20.844 13.22 20.04 13.0714 19.2 13.0714V3.92857H2.4V19.9286H12C12 20.58 12.12 21.3229 12.312 21.9171ZM13.848 15.3571C13.2 16.02 12.732 16.7857 12.42 17.6429H3.6V15.3571H13.848ZM3.6 10.7857H18V13.0714H3.6V10.7857ZM3.6 6.21429H18V8.5H3.6V6.21429Z"
        fill={color}
      />
    </svg>
  );
};
