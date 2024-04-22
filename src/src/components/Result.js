import React, { useState, useEffect } from 'react';
import './Result.css';
import GraphVisualization from './GraphVisualization';

const Result = ({ formData, updateTrigger }) => {
    const [checkedArticle, setCheckedArticle] = useState(0);
    const [clickArticle, setClickArticle] = useState(0);
    const [excTime, setExcTime] = useState(0);
    const [paths, setPaths] = useState([]);

    useEffect(() => {
        const fetchData = async () => {
            if (formData) {
                try {
                    const response = await fetch('http://localhost:8000/', {
                        method: 'POST',
                        headers: {
                            'Content-Type': 'application/json',
                        },
                        body: JSON.stringify(formData),
                    });
                    const data = await response.json();
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

        fetchData();
    }, [formData, updateTrigger]);

    return (
        <div>
            <h2 className="result-title">Result</h2>
            <div className="result-container">
                <div className="article-checked">
                    Article checked: {checkedArticle} <br/>
                </div>
                <div className="article-clicked">
                    Article clicked: {clickArticle}
                </div>
                <div className="addition">
                    move node to adjust your view and
                    click node to visit page!
                </div>
                <div className="exc-time">
                    Execution time: {excTime}
                </div>
                <div className="graph">
                    <GraphVisualization paths={paths} />
                </div>
            </div>
        </div>
    );
}

export default Result;
