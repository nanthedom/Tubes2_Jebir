import React, { useState } from 'react';
import Header from './components/Header';
import FormInput from './components/FormInput';
import Result from './components/Result';

function App() {
  const [formData, setFormData] = useState(null);

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
      setFormData(formData); // Update formData state
    } catch (error) {
      console.error('Error:', error);
    }
  };

  return (
    <div>
      <div>
        <Header />
      </div>
      <div className="FormInput">
        <FormInput onSubmit={handleSubmit} />
      </div>
      <div> 
        <Result formData={formData} />
      </div>
    </div>
  );
}

export default App;