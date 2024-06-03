import React from 'react';

const MasterHeader: React.FC = () => {
    return (
        <div className="navbar bg-main-background">
            <div className="navbar-start">

                <div className="flex-1">
                    <a className="btn btn-ghost bg-main-card text-base">Instructions</a>
                </div>
            </div>

            <div className="navbar-center">
                <a className="text-xl">Title 123</a>
            </div>

            <div className="navbar-end">
                <ul className="menu menu-horizontal px-1">
                    <li><a>Feedback</a></li>
                    <li><a>Templates</a></li>
                    <li>
                        <details>
                            <summary>
                                Account
                            </summary>
                            <ul className="p-2 bg-main-background rounded-t-none">
                                <li><a>Settings</a></li>
                                <li><a>Log Out</a></li>
                            </ul>
                        </details>
                    </li>
                </ul>
            </div>
        </div>
    );
}

export default MasterHeader;