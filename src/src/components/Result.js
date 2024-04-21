// import React, { useState, useEffect } from 'react';
// import './Result.css';

// const Result = ({ result }) => {
//     const [path, setPath] = useState([]);
//     const [loading, setLoading] = useState(true);
    
//     useEffect(() => {
//         const fetchPath = async () => {
//         try {
//             const response = await fetch('http://localhost:8000/path');
//             const data = await response.json();
//             setPath(data.path);
//             setLoading(false);
//         } catch (error) {
//             console.error('Error fetching path:', error);
//         }
//         };
    
//         fetchPath();
//     }, []);
    
//     return (
//         <div>
//             <h2 className="result-title" >Result</h2>
//             {/* {loading ? (
//                 <p>Loading...</p>
//             ) : (
//                 <ul>
//                 {path.map((node, index) => (
//                     <li key={index}>{node}</li>
//                 ))}
//                 </ul>
//             )} */}
//             <div className="result-container">
//                 <div className="artikel-diperiksa">
//                     Artikel diperiksa: 0
//                 </div>
//                 <div className="artikel-dilalui">
//                     Artikel dilalui: 0
//                 </div>
//                 <div className="waktu-pencarian">
//                     Waktu pencarian: 0ms
//                 </div>
//             </div>
//         </div>
//     );
// }

// export default Result;
