import React, { useState, useEffect, ChangeEvent } from 'react';

interface TokenError {
    line: number;
    column: number;
    message: string;
}

interface TextAreaWithLineNumbersProps {
    text: string;
    onChange: (e: ChangeEvent<HTMLTextAreaElement>) => void;
    errors: TokenError[];
}

const TextAreaWithLineNumbers: React.FC<TextAreaWithLineNumbersProps> = ({ text, onChange, errors }) => {
    const [lineNumbers, setLineNumbers] = useState<number[]>([]);

    useEffect(() => {
        const lines = text.split('\n').length;
        setLineNumbers(Array.from({ length: lines }, (_, i) => i + 1));
    }, [text]);

    const handleScroll = (e: React.UIEvent<HTMLTextAreaElement>) => {
        const element = e.target as HTMLTextAreaElement;
        const lineNumbersElement = document.getElementById('line-numbers');
        if (lineNumbersElement) {
            lineNumbersElement.scrollTop = element.scrollTop;
        }
    };

    return (
        <div className="relative flex">
            <div className="bg-gray-100 border-r border-gray-300 text-right py-2 pr-4 select-none text-black" id="line-numbers">
                {lineNumbers.map((number) => (
                    <div key={number} className="h-6">{number}</div>
                ))}
            </div>
            <div className="relative flex-grow">
                <textarea
                    value={text}
                    onChange={onChange}
                    onScroll={handleScroll}
                    className="w-full h-64 p-2 border-none outline-none resize-none font-mono leading-6 text-black"
                    rows={10}
                ></textarea>
                {errors.map((error, idx) => (
                    <div
                        key={idx}
                        className="absolute bg-red-500 h-6 w-1"
                        style={{
                            top: `${(error.line - 1) * 24}px`, // Adjust based on line height
                            left: `${error.column * 8}px`, // Adjust based on column width
                        }}
                        title={error.message}
                    ></div>
                ))}
            </div>
        </div>
    );
};

export default TextAreaWithLineNumbers;
