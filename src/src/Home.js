import React from 'react';
import './App.css';
import FormInput from './components/FormInput';

function App() {
  const handleSubmit = async (formData) => {
    try {
      const response = await fetch('http://localhost:8000/save', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify(formData)
      });
      
      if (!response.ok) {
        throw new Error('Failed to save data');
      }

      console.log('Data saved successfully');
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <div>
      <div className="header">
        <a href="#default" className="logo">Jebir Wikirace</a>
        <div className="header-right">
          <a className="active" href="#home">Home</a>
          <a href="/how-to-use">How-To Use</a>
          <a href="/about">About</a>
        </div>
      </div>
      <div className="FormInput">
          <FormInput onSubmit={handleSubmit} />
      </div>
    </div>
  );
}

export default App;