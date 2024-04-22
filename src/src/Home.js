import React, { useState } from 'react';
import Header from './components/Header';
import FormInput from './components/FormInput';
import Result from './components/Result';

function Home() {
  const [formData, setFormData] = useState(null);
  const [updateTrigger, setUpdateTrigger] = useState(false);

  const handleFormSubmit = (formData) => {
    setFormData(formData);
    setUpdateTrigger(prevState => !prevState);
  };

  return (
    <div>
      <div>
        <Header />
      </div>
      <div className="FormInput">
        <FormInput onFormSubmit={handleFormSubmit} />
      </div>
      <div>
        {formData && <Result formData={formData} updateTrigger={updateTrigger} />}
      </div>
    </div>
  );
}

export default Home;
