import React from 'react';
import style from './sidebar.module.css';
import { MenuItem } from 'src/constants/types/typeSidebar';
import { useLocation } from 'react-router-dom';

export const MenuSideBar: React.FC<{
  item: MenuItem;
  isSidebarOpen: boolean;
}> = ({ item, isSidebarOpen }) => {
  const location = useLocation();
  const isActive = (path: string) => location.pathname === path;

  return (
    <span className={style['list-menu-sidebar']}>
      <item.Icon
        color={item.path && isActive(item.path) ? '#4D47C3' : '#95999E'}
      />
      <span
        className={`${style['sidebar-name-list-menu']} ${
          item.path && isActive(item.path) ? style['sidebar-active'] : ''
        } ${isSidebarOpen ? style['sidebar-name-list-menu-show'] : ''}`}
      >
        {item.label}
      </span>
    </span>
  );
};
