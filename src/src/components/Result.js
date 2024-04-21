import React, { useState, useEffect } from 'react';
import './Result.css';

const Result = ({ formData }) => {
    const [checkedArticle, setCheckedArticle] = useState(0);
    const [clickArticle, setClickArticle] = useState(0);
    const [excTime, setExcTime] = useState(0);

    useEffect(() => {
        const fetchData = async () => {
            if (formData) {
                try {
                    // untuk API request
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
                    console.log('Data fetched successfully:', data);
                } catch (error) {
                    console.error('Error fetching data:', error);
                }
            }
        };

        fetchData();
    }, [formData]);

    return (
        <div>
            <h2 className="result-title">Result</h2>
            <div className="result-container">
                <div className="artikel-diperiksa">
                    Article checked: {checkedArticle}
                </div>
                <div className="artikel-dilalui">
                    Article passed: {clickArticle}
                </div>
                <div className="waktu-pencarian">
                    Execution time: {excTime}
                </div>
            </div>
        </div>
    );
}

export default Result;
