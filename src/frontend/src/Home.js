import React, { useState, useEffect, useRef } from 'react';
import Header from './components/Header';
import FormInput from './components/FormInput';
import Result from './components/Result';

function Home() {
  const [formData, setFormData] = useState(null);
  const [updateTrigger, setUpdateTrigger] = useState(false);
  const resultRef = useRef(null); 

  const handleFormSubmit = (formData) => {
    setFormData(formData);
    setUpdateTrigger(prevTrigger => !prevTrigger); 
    resultRef.current.scrollIntoView({ behavior: 'smooth' });
  };

  useEffect(() => {
    if (formData) {
      resultRef.current.scrollIntoView({ behavior: 'smooth' });
    }
  }, [formData]);

  return (
    <div>
      <div>
        <Header />
      </div>
      <div className="FormInput">
        <FormInput onFormSubmit={handleFormSubmit} />
      </div>
      <div>
        <div ref={resultRef}> {}
          {formData && <Result formData={formData} updateTrigger={updateTrigger} />}
        </div>
      </div>
    </div>
  );
}

export default Home;
