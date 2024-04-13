import React, { useState, useEffect } from 'react';
import './FormInput.css';

const FormInput = ({ onSubmit }) => {
  const [formData, setFormData] = useState({
    startArticle: '',
    startUrl: '',
    endArticle: '',
    endUrl: ''
  });

  const [submittedData, setSubmittedData] = useState(null);
  const [startSuggestions, setStartSuggestions] = useState([]);
  const [endSuggestions, setEndSuggestions] = useState([]);

  const fetchSuggestions = async (inputValue, setter) => {
    try {
      const response = await fetch(
        `https://en.wikipedia.org/w/api.php?action=opensearch&limit=10&format=json&search=${inputValue}&origin=*`
      );
      const data = await response.json();
      const suggestions = data[1] || [];
      setter(suggestions);
    } catch (error) {
      console.error('Error fetching suggestions:', error);
    }
  };

  useEffect(() => {
    fetchSuggestions(formData.startArticle, setStartSuggestions);
  }, [formData.startArticle]);

  useEffect(() => {
    fetchSuggestions(formData.endArticle, setEndSuggestions);
  }, [formData.endArticle]);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSelectSuggestion = (value, field, urlField) => {
    const encodedValue = encodeURIComponent(value);
    setFormData({ ...formData, [field]: value, [urlField]: `https://en.wikipedia.org/wiki/${encodedValue}` });
    field === 'startArticle' ? setStartSuggestions([]) : setEndSuggestions([]);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onSubmit(formData);
    setSubmittedData(formData);
  };

  return (
    <div>
      <form onSubmit={handleSubmit} className="inline-form">
        <label>
          Start <br />
          <div className="autocomplete">
            <input
              type="text"
              name="startArticle"
              value={formData.startArticle}
              onChange={handleChange}
              className="inline-input"
            />
            {startSuggestions.length > 0 && (
              <div className="suggestion-container">
                <ul className="suggestion-list">
                  {startSuggestions.map((suggestion, index) => (
                    <li
                      key={index}
                      className="suggestion"
                      onClick={() => handleSelectSuggestion(suggestion, 'startArticle', 'startUrl')}
                    >
                      {suggestion}
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </div>
        </label>
        <label>
          End <br />
          <div className="autocomplete">
            <input
              type="text"
              name="endArticle"
              value={formData.endArticle}
              onChange={handleChange}
              className="inline-input"
            />
            {endSuggestions.length > 0 && (
              <div className="suggestion-container">
                <ul className="suggestion-list">
                  {endSuggestions.map((suggestion, index) => (
                    <li
                      key={index}
                      className="suggestion"
                      onClick={() => handleSelectSuggestion(suggestion, 'endArticle', 'endUrl')}
                    >
                      {suggestion}
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </div>
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
