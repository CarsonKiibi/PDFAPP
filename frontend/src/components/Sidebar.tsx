import React from 'react';

const Sidebar: React.FC = () => {
  return (
    <div className="navbar bg-base-100">
  <div className="flex-1">
    <a className="btn btn-ghost text-xl">Instructions</a>
  </div>
  <div className="flex-none">
    <ul className="menu menu-horizontal px-1">
      <li><a>Link</a></li>
      <li>
        <details>
          <summary>
            Parent
          </summary>
          <ul className="p-2 bg-base-100 rounded-t-none">
            <li><a>idk</a></li>
            <li><a>Account</a></li>
          </ul>
        </details>
      </li>
    </ul>
  </div>
</div>
  );
}

export default Sidebar;