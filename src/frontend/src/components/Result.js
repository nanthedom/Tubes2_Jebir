import React, { useState, useEffect } from 'react';
import './Result.css';
import GraphVisualization from './GraphVisualization';

const wikipediaFunFacts = [
    "Wikipedia was launched on January 15, 2001, by Jimmy Wales and Larry Sanger. It has since grown to become one of the largest and most popular reference websites on the internet.",
    "Wikipedia is available in over 300 languages, making it one of the most multilingual websites in the world. This allows people from diverse backgrounds to access and contribute to its content.",
    "The name 'Wikipedia' is a portmanteau of the words 'wiki' (a type of collaborative website) and 'encyclopedia.' It reflects the collaborative nature of the platform, where anyone can edit and contribute to articles.",
    "Naufal Adnan and Naufal Aulia turned out to be different people.",
    "Keanu Gonza: not an actor, not a football player."
];

const Result = ({ formData, updateTrigger }) => {
    const [countPath, setCountPath] = useState(0);
    const [checkedArticle, setCheckedArticle] = useState(0);
    const [clickArticle, setClickArticle] = useState(0);
    const [excTime, setExcTime] = useState(0);
    const [paths, setPaths] = useState([]);
    const [loading, setLoading] = useState(true);
    const [funFactIndex, setFunFactIndex] = useState(0);

    useEffect(() => {
        const fetchData = async () => {
            if (formData) {
                try {
                    setLoading(true);
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
                    setCountPath((data.paths).length);
                    setLoading(false);
                    console.log('Data fetched successfully:', data);
                } catch (error) {
                    console.error('Error fetching data:', error);
                    setLoading(false);
                }
            }
        };

        fetchData();
    }, [formData, updateTrigger]);

    useEffect(() => {
        // fun fact setiap 6.6 detik
        if (loading) {
            const timer = setInterval(() => {
                setFunFactIndex(prevIndex => (prevIndex + 1) % wikipediaFunFacts.length);
            }, 6600);
            return () => clearInterval(timer);
        }
    }, [loading]);

    return (
        <div>
            <h2 className="result-title">Result</h2>
            <div className="result-container">
                {loading ? (
                    <div className="loading-container">
                        <div className="loading-spinner"></div>
                        <div className="fun-fact">{wikipediaFunFacts[funFactIndex]}</div>
                    </div>
                ) : (
                    <>
                        <div className="article-checked">
                            Article checked: {checkedArticle} <br />
                        </div>
                        <div className="count-path">
                            Count path: {countPath} <br />
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
                    </>
                )}
            </div>
        </div>
    );
}

export default Result;
