import React from 'react';
import { PropsColorMenuSvg } from '@/constants/types/props';

export const MenuTransferSvg: React.FC<PropsColorMenuSvg> = ({ color }) => {
  return (
    <svg
      width="24"
      height="25"
      viewBox="0 0 24 25"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M6.66667 0.5C5.95942 0.5 5.28115 0.81607 4.78105 1.37868C4.28095 1.94129 4 2.70435 4 3.5V9.5H6.66667V3.5H17.3333V8H14L18.6667 13.25L23.3333 8H20V3.5C20 2.70435 19.719 1.94129 19.219 1.37868C18.7189 0.81607 18.0406 0.5 17.3333 0.5H6.66667ZM0 12.5V15.5H10.6667V12.5H0ZM0 17V20H10.6667V17H0ZM13.3333 17V20H24V17H13.3333ZM0 21.5V24.5H10.6667V21.5H0ZM13.3333 21.5V24.5H24V21.5H13.3333Z"
        fill={color}
      />
    </svg>
  );
};
