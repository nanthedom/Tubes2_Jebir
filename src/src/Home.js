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
    <div className="App">
      <h1>Jebir Wiki-Race</h1>
      <FormInput onSubmit={handleSubmit} />
    </div>
  );
}

export default App;