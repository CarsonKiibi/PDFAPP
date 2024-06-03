import React from 'react';

const EditHeader: React.FC = () => {
    return (
        <div className="navbar bg-base-100 rounded-t-xl">
            <ul className="space-x-5 ml-2">
                <li>Bold</li>
                <li>Italicize</li>
                <li>Underline</li>
            </ul>
        </div>
    );
}

export default EditHeader;