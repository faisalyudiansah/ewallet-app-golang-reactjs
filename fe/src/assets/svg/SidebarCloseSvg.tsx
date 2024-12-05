import React from 'react';

export const SidebarCloseSvg: React.FC<{
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
        d="M24.3 27H27V0H24.3V12.15H5.1705L12.5955 4.725L10.692 2.808L0 13.5L10.692 24.192L12.5955 22.275L5.1705 14.85H24.3V27Z"
        fill="#95999E"
      />
    </svg>
  );
};
