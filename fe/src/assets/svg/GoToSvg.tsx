import React from 'react';

export const GoToSvg: React.FC<{ customClass: string }> = ({
  customClass = '',
}) => {
  return (
    <svg
      className={customClass}
      width="22"
      height="17"
      viewBox="0 0 22 17"
      fill="none"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        d="M21.4976 9.22797L14.2996 16.4259C13.6569 17.0686 12.543 16.6187 12.543 15.6976V11.5844H6.71604C6.14621 11.5844 5.68777 11.126 5.68777 10.5562V6.44304C5.68777 5.87321 6.14621 5.41477 6.71604 5.41477H12.543V1.30165C12.543 0.384769 13.6526 -0.0693874 14.2996 0.573287L21.4976 7.77124C21.896 8.17398 21.896 8.82522 21.4976 9.22797ZM8.42984 16.2117V14.4979C8.42984 14.2151 8.19848 13.9838 7.9157 13.9838H4.31673C3.55837 13.9838 2.94569 13.3711 2.94569 12.6127V4.38649C2.94569 3.62813 3.55837 3.01545 4.31673 3.01545H7.9157C8.19848 3.01545 8.42984 2.78409 8.42984 2.50131V0.787512C8.42984 0.504735 8.19848 0.273372 7.9157 0.273372H4.31673C2.04595 0.273372 0.203613 2.11571 0.203613 4.38649V12.6127C0.203613 14.8835 2.04595 16.7258 4.31673 16.7258H7.9157C8.19848 16.7258 8.42984 16.4945 8.42984 16.2117Z"
        fill="black"
      />
    </svg>
  );
};