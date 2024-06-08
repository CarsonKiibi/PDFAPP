import React, { useState, ChangeEvent } from 'react';
import './App.css';
import Sidebar from './components/Sidebar';
import MasterHeader from './components/MasterHeader';
import DocContainer from './components/DocContainer';
import EditHeader from './components/EditHeader';
import PdfHeader from './components/PdfHeader';
import TextAreaWithLineNumbers from './components/TextAreaWithLineNumbers';

interface TokenError {
  line: number;
  column: number;
  message: string;
}

function App() {
  const [text, setText] = useState<string>("My name is {Bob} and I like {Shiny Green Apples}");
  const [errors, setErrors] = useState<TokenError[]>([]);

  const handleTextChange = (e: ChangeEvent<HTMLTextAreaElement>) => {
    setText(e.target.value);
  };

  const handleSubmit = async (event: React.FormEvent) => {
    event.preventDefault();
    const response = await fetch('/api/parse', {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ text }),
    });
    const result = await response.json();
    if (result.errors) {
      setErrors(result.errors);
    } else {
      setErrors([]);
    }
  };

  return (
    <div className="flex flex-col h-screen">
      <MasterHeader />
      <div className="flex flex-1 flex-col lg:flex-row overflow-hidden">
        <div className="flex flex-col w-full mx-2 flex-1 overflow-hidden">
          <EditHeader />
          <TextAreaWithLineNumbers text={text} onChange={handleTextChange} errors={errors} />
        </div>

        <div className="flex flex-col w-full mx-3 flex-1 overflow-hidden">
          <PdfHeader />
          <DocContainer />
        </div>
      </div>
    </div>
  );
}

export default App;

