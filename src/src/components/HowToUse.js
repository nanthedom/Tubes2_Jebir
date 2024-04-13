import React from "react";
import Header from './Header'
import './styles.css';

function HowToUse() {
    return (
        <div>
            <div>
                <Header />
            </div>
            <h1 className="page-title">How To Use Jebir Wikirace (contoh)</h1>
            <div className="mid-container">
                <div className="step">
                    <div className="container-2">
                        <span className="number">
                            1
                        </span>
                    </div>
                    <div className="step-text">
                        Masukkan judul artikel Wikipedia awal pada kolom "Start".
                    </div>
                </div>
                <div className="step-1">
                    <div className="container-3">
                        <span className="number-1">
                            2
                        </span>
                    </div>
                    <div className="step-text-1">
                        Masukkan judul artikel Wikipedia tujuan pada kolom "End".
                    </div>
                </div>
                <div className="step-2">
                    <div className="container-1">
                        <span className="number-2">
                            3
                        </span>
                    </div>
                    <div className="step-text-2">
                        Masukkan judul artikel Wikipedia awal pada kolom "Start".
                    </div>
                </div>
                <div className="step-3">
                    <div className="container">
                        <span className="number-3">
                            4
                        </span>
                    </div>
                    <div className="step-text-3">
                        Hasil pencarian akan ditampilkan pada layar.
                    </div>
                </div>
            </div>
        </div>
    )
}

export default HowToUse;