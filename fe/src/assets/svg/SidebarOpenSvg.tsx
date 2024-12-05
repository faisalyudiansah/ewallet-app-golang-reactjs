import React from 'react';

export const SidebarOpenSvg: React.FC<{
  customClass?: string;
}> = ({ customClass }) => {
  return (
    <svg
      className={customClass}
      width="27"
      height="27"
      viewBox="0 0 27 27"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M2.7 0H0V27H2.7V14.85H21.8295L14.4045 22.275L16.308 24.192L27 13.5L16.308 2.808L14.4045 4.725L21.8295 12.15H2.7V0Z"
        fill="#95999E"
      />
    </svg>
  );
};
