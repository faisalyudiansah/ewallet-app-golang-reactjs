import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { MenuDashboardSvg } from 'src/assets/svg/menuDashboardSvg';
import { MenuTransactionSvg } from 'src/assets/svg/menuTransactionSvg';
import { MenuTransferSvg } from 'src/assets/svg/menuTransferSvg';
import { MenuTopUpSvg } from 'src/assets/svg/menuTopUpSvg';
import { MenuLogoutSvg } from 'src/assets/svg/menuLogoutSvg';
import style from './sidebar.module.css';
import { MenuItem } from 'src/constants/types/typeSidebar';
import { SidebarOpenSvg } from 'src/assets/svg/SidebarOpenSvg';
import { SidebarCloseSvg } from 'src/assets/svg/SidebarCloseSvg';
import { MenuSideBar } from './MenuSideBar';
import { ModalConfirmLogout } from 'src/components/ui/modal/ModalConfirmLogout';

export const Sidebar: React.FC<{
  toggleSidebar: () => void;
  isSidebarOpen: boolean;
}> = ({ toggleSidebar, isSidebarOpen }) => {
  const [modalConfirm, setModalConfirm] = useState(false);

  const menuItems: MenuItem[] = [
    { path: '/', label: 'Dashboard', Icon: MenuDashboardSvg },
    { path: '/transactions', label: 'Transactions', Icon: MenuTransactionSvg },
    { path: '/transfer', label: 'Transfer', Icon: MenuTransferSvg },
    { path: '/top-up', label: 'Top Up', Icon: MenuTopUpSvg },
    { path: null, label: 'Logout', Icon: MenuLogoutSvg },
  ];

  return (
    <>
      {modalConfirm && (
        <ModalConfirmLogout closeModal={() => setModalConfirm(false)} />
      )}
      <nav className={style['container-sidebar']}>
        <span
          className={`${
            isSidebarOpen
              ? style['title-sidebar-big-version']
              : style['title-sidebar']
          }`}
        >
          Sea Wallet
        </span>
        <div className={style['container-menu']}>
          <ul className={style['container-list-menu']}>
            {menuItems.map((item) => (
              <li key={item.path}>
                {!item.path ? (
                  <span
                    role="presentation"
                    onClick={() => setModalConfirm(!modalConfirm)}
                  >
                    <MenuSideBar item={item} isSidebarOpen={isSidebarOpen} />
                  </span>
                ) : (
                  <Link to={item.path}>
                    <MenuSideBar item={item} isSidebarOpen={isSidebarOpen} />
                  </Link>
                )}
              </li>
            ))}
          </ul>
        </div>
        <div
          className={style['sidebar-open-svg']}
          onClick={toggleSidebar}
          role="presentation"
        >
          {isSidebarOpen ? (
            <SidebarCloseSvg customClass={style['openCLose-svg']} />
          ) : (
            <SidebarOpenSvg customClass={style['openCLose-svg']} />
          )}
        </div>
      </nav>
    </>
  );
};
