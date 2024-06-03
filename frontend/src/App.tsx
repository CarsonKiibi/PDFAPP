import React from 'react';
import './App.css';
import Sidebar from './components/Sidebar';
import MasterHeader from './components/MasterHeader';
import DocContainer from './components/DocContainer';
import EditHeader from './components/EditHeader';
import PdfHeader from './components/PdfHeader';

function App() {
  return (
    <div className="flex flex-col h-screen">
      <MasterHeader />
      <div className="flex flex-1 flex-col lg:flex-row overflow-hidden">
        <div className="flex flex-col w-full mx-2 flex-1 overflow-hidden">
          <EditHeader />
          <DocContainer />
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
