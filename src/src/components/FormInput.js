import React, { useState } from 'react';

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
      <form onSubmit={handleSubmit}>
        <label>
          Start Article:
          <input
            type="text"
            name="startArticle"
            value={formData.startArticle}
            onChange={handleChange}
          />
        </label>
        <br />
        <label>
          End Article:
          <input
            type="text"
            name="endArticle"
            value={formData.endArticle}
            onChange={handleChange}
          />
        </label>
        <br />
        <button type="submit">Start Game</button>
      </form>

      {submittedData && (
        <div>
          <h2>Submitted Data:</h2>
          <p>Start Article: {submittedData.startArticle}</p>
          <p>End Article: {submittedData.endArticle}</p>
        </div>
      )}
    </div>
  );
};

export default FormInput;
