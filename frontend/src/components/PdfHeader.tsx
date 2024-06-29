import React from 'react';

interface PdfHeaderProps {
    onCompile: () => void;
}

const PdfHeader: React.FC<PdfHeaderProps> = ({ onCompile }) => {
    return (
        <div className="navbar bg-base-100 rounded-t-xl">
            <button onClick={onCompile}>Compile</button>
        </div>
    );
};

export default PdfHeader;