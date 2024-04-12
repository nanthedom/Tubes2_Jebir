import React, { useState } from 'react';
import './FormInput.css';

const FormInput = ({ onSubmit }) => {
  const [formData, setFormData] = useState({
    startArticle: '',
    endArticle: ''
  });

  const [submittedData, setSubmittedData] = useState(null);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
    setSubmittedData(formData); // Setelah submit, simpan data di state untuk ditampilkan
  };

  return (
    <div>
      <form onSubmit={handleSubmit} className="inline-form">
        <label>
          Start <br />
          <input
            type="text"
            name="startArticle"
            value={formData.startArticle}
            onChange={handleChange}
            className="inline-input"
          />
        </label>
        <label>
          End <br />
          <input
            type="text"
            name="endArticle"
            value={formData.endArticle}
            onChange={handleChange}
            className="inline-input"
          />
        </label>
        <button type="submit" className="inline-button">Find!</button>
      </form>

      {submittedData && (
        <div className="finding-route">
          <h2>Finding Route:</h2>
          <p>{submittedData.startArticle} To {submittedData.endArticle}</p>
        </div>
      )}
    </div>
  );
};

export default FormInput;
