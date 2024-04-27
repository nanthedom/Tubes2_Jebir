import React from "react";
import Header from './Header'
import './About.css';
import nopalImage from './assets/nopal.png';
import gonzaImage from './assets/gonza.png';
import adnanImage from './assets/adnan.png';

function About() {
    return (
        <div>
            <div>
                <Header />
            </div>
            <div>
                <h1 className="page-title">About Jebir Wikirace</h1>
                <div className="content-container">
                    <p className="content">
                        Jebir Wikirace adalah website application yang memungkinkan pengguna untuk melakukan pencarian rute artikel Wikipedia terpendek dari artikel awal ke artikel tujuan menggunakan dua opsi pendekatan algoritma: Breadth-First Search (BFS) atau Iterative Deepening Search (IDS). Hasil pencarian pada website ini akan menampilkan: jumlah artikel yang diperiksa, jumlah artikel yang dilalui, rute penjelajahan artikel, dan waktu pencarian. Website ini dibangun menggunakan Golang sebagai backend dan React sebagai frontend.
                        <br/>
                        <br/>
                        Website ini merupakan realisasi atas Tugas Besar II Strategi Algoritma (IF2211). Dibuat oleh kelompok Jebir dengan anggota:
                    </p>
                </div>
                
                <div className="blank"></div>
                <div className="anggota-container">
                    <div className="anggota-wrapper">
                        <div className="nama-anggota">Muhammad Naufal Aulia</div>
                        <img className="gambar-anggota" src={nopalImage} alt="inigambar"/>
                        <div className="nim-anggota">13522074</div>
                    </div>
                    <div className="anggota-wrapper">
                        <div className="nama-anggota">Keanu Amadius Gonza Wrahatno</div>
                        <img className="gambar-anggota" src={gonzaImage} alt="inigambar"/>
                        <div className="nim-anggota">13522082</div>
                    </div>
                    <div className="anggota-wrapper">
                        <div className="nama-anggota">Naufal Adnan</div>
                        <img className="gambar-anggota" src={adnanImage} alt="inigambar"/>
                        <div className="nim-anggota">13522116</div>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default About;