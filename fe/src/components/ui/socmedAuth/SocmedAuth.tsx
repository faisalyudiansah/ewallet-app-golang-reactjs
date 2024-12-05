import React from 'react';
import { SocialMediaButton } from 'src/constants/types/typeButtonSvg';
import { FacebookSvg } from 'src/assets/svg/FacebookSvg';
import { AppleSvg } from 'src/assets/svg/AppleSvg';
import { GoogleSvg } from 'src/assets/svg/GoogleSvg';
import { PropsSocmedAuth } from '@/constants/types/props';

export const SocmedAuth: React.FC<PropsSocmedAuth> = ({ style }) => {
  const socialMediaButtons: SocialMediaButton[] = [
    { id: 1, SvgComponent: FacebookSvg },
    { id: 2, SvgComponent: AppleSvg },
    { id: 3, SvgComponent: GoogleSvg },
  ];

  const btnSocmedAuth = (socmedId: number) => {
    switch (socmedId) {
      case 1:
        window.location.href = '#';
        break;
      case 2:
        window.location.href = '#';
        break;
      case 3:
        window.location.href = '#';
        break;
      default:
        break;
    }
  };

  const handleKeyPress = (
    event: React.KeyboardEvent<HTMLDivElement>,
    socmedId: number,
  ) => {
    if (event.key === 'Enter' || event.key === ' ') {
      btnSocmedAuth(socmedId);
    }
  };
  return (
    <>
      <span
        className={`${style['text-secondary']} ${style['or-continue-with']}`}
      >
        or continue with
      </span>
      <div className={style['container-btn-alternative-login']}>
        {socialMediaButtons.map((button) => (
          <div
            key={button.id}
            onKeyDown={(e) => handleKeyPress(e, button.id)}
            className={style['svg-link']}
            role="button"
            tabIndex={0}
          >
            <button.SvgComponent />
          </div>
        ))}
      </div>
    </>
  );
};
