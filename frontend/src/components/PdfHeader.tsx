import React from 'react';

const PdfHeader: React.FC = () => {
    return (
        <div className="navbar bg-base-100 rounded-t-xl">
            <a className="btn btn-ghost bg-main-accent text-main-background ml-2">Compile</a>
        </div>
    );
}

export default PdfHeader;