import React, { FC } from "react";

interface SidebarProps {
  children: React.ReactNode;
}

const Sidebar: FC<SidebarProps> = ({ children }) => {
  return (
    <div className="flex bg-base-100">
      <div className="w-80 bg-base-200 p-4 calc-dvh-69">
        <ul className="menu text-base-content">
          <li>
            <a>Sidebar Item 1</a>
          </li>
          <li>
            <a>Sidebar Item 2</a>
          </li>
        </ul>
      </div>
      <div className="flex-1">{children}</div>
    </div>
  );
};


export default Sidebar;
