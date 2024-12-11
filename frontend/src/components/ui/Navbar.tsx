
import Link from "next/link";
import React from "react";

const Navbar: React.FC = () => {

  return (
    <>
      <div className="navbar bg-base-200">
        <div className="flex-1">
          <Link href="/" className="btn btn-ghost text-xl">FlowForge</Link>
        </div>
        <div className="flex-none">
          <ul className="menu menu-horizontal px-1">
            <li>
              <Link href="/flowchart">Visualizer</Link>
            </li>
          </ul>
        </div>
      </div>
    </>
  );
};

export default Navbar;
