import React, { useEffect } from 'react';
import { pdfjs } from 'react-pdf';

const DisplayPDF: React.FC = () => {
    useEffect(() => {
        const url = 'http://localhost:8080/pdf';

        const fetchAndRenderPDF = async () => {
            try {
                const response = await fetch(url);
                const blob = await response.blob();
                const arrayBuffer = await blob.arrayBuffer();
                const typedarray = new Uint8Array(arrayBuffer);

                const pdf = await pdfjs.getDocument(typedarray).promise;

                for (let pageNum = 1; pageNum <= pdf.numPages; pageNum++) {
                    const page = await pdf.getPage(pageNum);
                    const scale = 5.0;
                    const viewport = page.getViewport({ scale });

                    const canvas = document.createElement('canvas');
                    canvas.className = 'pdfPage';
                    const context = canvas.getContext('2d');
                    canvas.height = viewport.height;
                    canvas.width = viewport.width;

                    const renderContext = {
                        canvasContext: context,
                        viewport: viewport,
                    };

                    if (context) {
                        await page.render({ canvasContext: context, viewport: viewport }).promise;
                    } else {
                        throw new Error('Canvas context is null.');
                    }
                    const pdfViewer = document.getElementById('pdfViewer');
                    if (pdfViewer) {
                        pdfViewer.appendChild(canvas);
                    } else {
                        throw new Error('PDF viewer container is not found.');
                    }
                }
            } catch (error) {
                console.error('Error fetching PDF:', error);
            }
        };

        fetchAndRenderPDF();
    }, []);

    return (
        <div id="pdfViewer" style={{ display: 'flex', flexDirection: 'column', backgroundColor: 'lightblue' }}></div>
    );
};

export default DisplayPDF;

