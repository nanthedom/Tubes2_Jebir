import React, { useState, useEffect } from 'react';
import './Result.css';

const Result = ({ result }) => {
    const [checkedArticle, setCheckedArticle] = useState(0);
    const [clickArticle, setClickArticle] = useState(0);
    const [excTime, setExcTime] = useState(0);

    useEffect(() => {
        console.log(result); // Tambahkan ini untuk memeriksa nilai result
        if (result) {
            setCheckedArticle(result.checkedArticle);
            setClickArticle(result.clickArticle);
            setExcTime(result.excTime);
        }
    }, [result]);
    

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
                    Waktu pencarian: {excTime}ms
                </div>
            </div>
        </div>
    );
}

export default Result;
