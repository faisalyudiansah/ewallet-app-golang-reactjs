import React from 'react';

export const ExpenseSvg: React.FC<{ customClass?: string }> = ({
  customClass,
}) => {
  return (
    <svg
      className={customClass}
      width="81"
      height="56"
      viewBox="0 0 81 56"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M77.4993 31.429C78.3296 39.7897 78.3302 44.3816 77.4993 52.3771C69.5032 53.2081 64.9117 53.2075 56.5507 52.3771"
        stroke="black"
        strokeWidth="5.61877"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
      <path
        d="M77.4666 52.2661L43.9507 18.7503L31.3505 31.3505L3 3"
        stroke="black"
        strokeWidth="4.92707"
        strokeLinecap="round"
        strokeLinejoin="round"
      />
    </svg>
  );
};
