import { PropsColorMenuSvg } from '@/constants/types/props';
import React from 'react';

export const MenuTransactionSvg: React.FC<PropsColorMenuSvg> = ({ color }) => {
  return (
    <svg
      width="24"
      height="25"
      viewBox="0 0 24 25"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M5 6.5H23V18.5H5V6.5ZM14 9.5C14.7956 9.5 15.5587 9.81607 16.1213 10.3787C16.6839 10.9413 17 11.7044 17 12.5C17 13.2956 16.6839 14.0587 16.1213 14.6213C15.5587 15.1839 14.7956 15.5 14 15.5C13.2044 15.5 12.4413 15.1839 11.8787 14.6213C11.3161 14.0587 11 13.2956 11 12.5C11 11.7044 11.3161 10.9413 11.8787 10.3787C12.4413 9.81607 13.2044 9.5 14 9.5ZM9 8.5C9 9.03043 8.78929 9.53914 8.41421 9.91421C8.03914 10.2893 7.53043 10.5 7 10.5V14.5C7.53043 14.5 8.03914 14.7107 8.41421 15.0858C8.78929 15.4609 9 15.9696 9 16.5H19C19 15.9696 19.2107 15.4609 19.5858 15.0858C19.9609 14.7107 20.4696 14.5 21 14.5V10.5C20.4696 10.5 19.9609 10.2893 19.5858 9.91421C19.2107 9.53914 19 9.03043 19 8.5H9ZM1 10.5H3V20.5H19V22.5H1V10.5Z"
        fill={color}
      />
    </svg>
  );
};