import React, { useState, useEffect } from 'react';
import { Document, Page, pdfjs } from 'react-pdf';
import 'react-pdf/dist/esm/Page/AnnotationLayer.css';
import 'react-pdf/dist/esm/Page/TextLayer.css';

// Set up the worker for PDF.js
pdfjs.GlobalWorkerOptions.workerSrc = `//cdnjs.cloudflare.com/ajax/libs/pdf.js/${pdfjs.version}/pdf.worker.min.js`;

interface PDFViewerProps {
  base64PDF?: string;
}

const PDFViewer: React.FC<PDFViewerProps> = ({ base64PDF }) => {
  const [numPages, setNumPages] = useState<number | null>(null);
  const [pdfData, setPdfData] = useState<string | null>(null);

  useEffect(() => {
    const defaultPDF = 'JVBERi0xLjMNCiXi48/TDQoNCjEgMCBvYmoNCjw8DQovVHlwZSAvQ2F0YWxvZw0KL091dGxpbmVzIDIgMCBSDQovUGFnZXMgMyAwIFINCj4+DQplbmRvYmoNCg0KMiAwIG9iag0KPDwNCi9UeXBlIC9PdXRsaW5lcw0KL0NvdW50IDANCj4+DQplbmRvYmoNCg0KMyAwIG9iag0KPDwNCi9UeXBlIC9QYWdlcw0KL0NvdW50IDENCi9LaWRzIFsgNCAwIFIgXQ0KPj4NCmVuZG9iag0KDQo0IDAgb2JqDQo8PA0KL1R5cGUgL1BhZ2UNCi9QYXJlbnQgMyAwIFINCi9SZXNvdXJjZXMgPDwNCi9Gb250IDw8DQovRjEgOSAwIFIgDQo+Pg0KL1Byb2NTZXQgOCAwIFINCj4+DQovTWVkaWFCb3ggWzAgMCA2MTIuMDAwMCA3OTIuMDAwMF0NCi9Db250ZW50cyA1IDAgUg0KPj4NCmVuZG9iag0KDQo1IDAgb2JqDQo8PCAvTGVuZ3RoIDY3NiA+Pg0Kc3RyZWFtDQoyIDAgMCAtMiAwIDc5MiBjbQ0KQlQNCi9GMSAyMCBUZg0KMSAwIDAgMSA3Mi4wMDAwIDcxMi4wMDAwIFRtDQooVGhpcyBpcyBhIHRlc3QgUERGIGZpbGUpIFRqDQpFVA0KQlQNCi9GMSAxNiBUZg0KMSAwIDAgMSA3Mi4wMDAwIDY4OC4wMDAwIFRtDQooUGxlYXNlIHJlcGxhY2UgdGhpcyB3aXRoIHlvdXIgb3duIGNvbnRlbnQuKSBUag0KRVQNCmVuZHN0cmVhbQ0KZW5kb2JqDQoNCjYgMCBvYmoNCjw8DQovVHlwZSAvRm9udA0KL1N1YnR5cGUgL1R5cGUxDQovTmFtZSAvRjENCi9CYXNlRm9udCAvSGVsdmV0aWNhDQovRW5jb2RpbmcgL1dpbkFuc2lFbmNvZGluZw0KPj4NCmVuZG9iag0KDQo3IDAgb2JqDQpbIC9QREYgL1RleHQgXQ0KZW5kb2JqDQoNCjggMCBvYmoNCjw8DQovQ3JlYXRvciAoUmF2ZSBcKGh0dHA6Ly93d3cubmV2cm9uYS5jb20vcmF2ZVwpKQ0KL1Byb2R1Y2VyIChOZXZyb25hIERlc2lnbnMpDQovQ3JlYXRpb25EYXRlIChEOjIwMDYwMzAxMDcyODI2KQ0KPj4NCmVuZG9iag0KDQp4cmVmDQowIDkNCjAwMDAwMDAwMDAgNjU1MzUgZg0KMDAwMDAwMDAxOSAwMDAwMCBuDQowMDAwMDAwMDkzIDAwMDAwIG4NCjAwMDAwMDAxNDcgMDAwMDAgbg0KMDAwMDAwMDIyMiAwMDAwMCBuDQowMDAwMDAwMzkwIDAwMDAwIG4NCjAwMDAwMDExMTYgMDAwMDAgbg0KMDAwMDAwMTI1NCAwMDAwMCBuDQowMDAwMDAxMjg0IDAwMDAwIG4NCnRyYWlsZXINCjw8DQovU2l6ZSA5DQovUm9vdCAxIDAgUg0KL0luZm8gOCAwIFINCj4+DQoNCnN0YXJ0eHJlZg0KMTQ1NA0KJSVFT0YNCg==';
    setPdfData(`data:application/pdf;base64,${base64PDF || defaultPDF}`);
  }, [base64PDF]);

  function onDocumentLoadSuccess({ numPages }: { numPages: number }) {
    setNumPages(numPages);
  }

  return (
    <div style={{ width: '100%', height: '100%', overflow: 'auto' }}>
      <Document
        file={pdfData}
        onLoadSuccess={onDocumentLoadSuccess}
        options={{
          cMapUrl: 'https://unpkg.com/pdfjs-dist@2.9.359/cmaps/',
          cMapPacked: true,
        }}
      >
        {Array.from(new Array(numPages), (el, index) => (
          <Page 
            key={`page_${index + 1}`}
            pageNumber={index + 1} 
            renderTextLayer={false}
            renderAnnotationLayer={false}
          />
        ))}
      </Document>
    </div>
  );
};

export default PDFViewer;
