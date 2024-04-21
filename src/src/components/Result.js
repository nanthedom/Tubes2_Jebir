import React, { useState, useEffect } from 'react';
import './Result.css';

const Result = ({ formData }) => {
    const [checkedArticle, setCheckedArticle] = useState(0);
    const [clickArticle, setClickArticle] = useState(0);
    const [excTime, setExcTime] = useState(0);
    const [paths, setPaths] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            if (formData) {
                try {
                    // Make API request with full URLs
                    const response = await fetch('http://localhost:8000/', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify(formData),
                    });
                    const data = await response.json();
                    // Update state based on data received
                    setCheckedArticle(data.checkedArticle);
                    setClickArticle(data.clickArticle);
                    setExcTime(data.excTime);
                    setPaths(data.paths);
                    console.log('Data fetched successfully:', data);
                } catch (error) {
                    console.error('Error fetching data:', error);
                }
            }
        };

        fetchData(); // Call the async function immediately
    }, [formData]);

    return (
        <div>
            <h2 className="result-title">Result</h2>
            <div className="result-container">
                <div className="artikel-diperiksa">
                    Artikel diperiksa: {checkedArticle}
                </div>
                <div className="artikel-dilalui">
                    Artikel dilalui: {clickArticle}
                </div>
                <div className="waktu-pencarian">
                    Waktu pencarian: {excTime}
                </div>
                <div className="paths">
                    {paths.map((path, index) => (
                        <div key={index}>{path}</div>
                    ))}
                </div>
            </div>
        </div>
    );
}

export default Result;
