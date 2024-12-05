import React, { useState } from 'react';
import { Outlet } from 'react-router-dom';
import { Sidebar } from './components/layout/sidebar/Sidebar';

export const Layout: React.FC = () => {
  const [isSidebarOpen, setIsSidebarOpen] = useState(false);

  const toggleSidebar = () => {
    setIsSidebarOpen(!isSidebarOpen);
  };
  return (
    <>
      <div className={`container ${isSidebarOpen ? 'sidebar-open' : ''}`}>
        <Sidebar toggleSidebar={toggleSidebar} isSidebarOpen={isSidebarOpen} />
        <div className="content">
          <Outlet />
        </div>
      </div>
    </>
  );
};
