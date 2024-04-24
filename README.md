# Tugas Besar II Strategi Algoritma (IF2211)
## Kelompok Jebir:
* Muhammad Naufal Aulia 			(13522074)
* Keanu Amadius Gonza Wrahatno		(13522082)
* Naufal Adnan					(13522116)


## Table of Contents
* [Website Jebir Wikirace](#jebir)
* [Screenshots](#screenshots)
* [Dependencies](#dependencies)
* [How to Use](#how-to-use)


## Jebir Wikirace: Implementasi Algoritma BFS dan IDS  <a href="jebir"></a>
> Jebir WikiRace adalah sebuah website solver atas permainan WikiRace yang memungkinkan pengguna untuk melakukan pencarian artikel Wikipedia dari artikel awal ke artikel target dengan rute penjelajahan artikel terdekat melalui hipertaut artikel yang ada di dalamnya.

Dalam rangka menemukan rute penjelajahan terpendek, digunakan pendekatan algoritma BFS dan IDS. 


## Screenshots <a href="screenshots"></a>
![Example screenshot](todo.gif)

## Dependencies <a href="dependencies"></a>
- Go 
- Node.js
- Docker desktop

## How to Use <a href="how-to-use"></a>
0. Siapkan requirement jika belum di-install:
    - Node.js (https://nodejs.org/en) 
    - Docker desktop (https://www.docker.com/products/docker-desktop/) 
    - Yarn

1. Download source code (.zip) pada link berikut:
    ```
    https://github.com/haziqam/tubes1-IF2211-game-engine/releases/tag/v1.1.0
    ```
2. Extract zip tersebut, lalu masuk ke folder hasil extractnya dan buka terminal
3. Masuk ke root directory dari project (sesuaikan dengan nama rilis terbaru)
    ```
    cd tubes1-IF2211-game-engine-1.1.0
    ```
4. Install dependencies menggunakan Yarn
    ```
    yarn
    ```
5. Setup default environment variable dengan menjalankan script berikut
Untuk Windows
    ```
    ./scripts/copy-env.bat
    ```
    Untuk Linux / (possibly) macOS
    ```
    chmod +x ./scripts/copy-env.sh
    ./scripts/copy-env.sh
    ```
6. Setup local database (buka aplikasi docker desktop terlebih dahulu, lalu jalankan command berikut di terminal)
    ```
    docker compose up -d database
    ```
    Lalu jalankan script berikut. Untuk Windows
    ```
    ./scripts/setup-db-prisma.bat
    ```
    Untuk Linux / (possibly) macOS
    ```
    chmod +x ./scripts/setup-db-prisma.sh
    ./scripts/setup-db-prisma.sh
    ```
7. Jika sudah, lakukan build kemudian run untuk menjalankan website
    ```
    npm run build
    ```
    ```
    npm run start
    ```
    Kunjungi frontend melalui `http://localhost:8082/`.
8. Untuk menjalankan bot, clone repository ini dengan
    ```
    git clone https://github.com/TazakiN/Tubes1_Nanang-Boneng.git
    ```
9. Masuk ke src directory dari project 
    ```
    cd Tubes1_Nanang-Boneng/src
    ```
10. Install dependencies menggunakan pip
    ```
    pip install -r requirjements.txt
    ```
11. Run bot di dalam direktori src dengan:
    - hanya 1 bot:
    ```
    python main.py --logic NanangBoneng --email=nanang_boneng_2@example.com --name=nanang --password=nanang_final_password --team etimo    
    ```
    - lebih dari 1 bot bersamaan (windows):
    ```
    ./run-bots.bat
    ```
    - lebih dari 1 bot bersamaan (linux/macOS):
    ```
    ./run-bots.sh
    ```
    sesuaikan script pada  `run-bots` tersebut

