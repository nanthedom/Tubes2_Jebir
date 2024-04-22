import React, { useState, useEffect } from 'react';
import './FormInput.css';
import arrow from './assets/arrow.png';

const FormInput = ({ onFormSubmit }) => {
  const [formData, setFormData] = useState({
    startArticle: '',
    startUrl: '',
    endArticle: '',
    endUrl: '',
    algoritma: 'BFS'
  });

  const [submittedData, setSubmittedData] = useState(null);
  const [startSuggestions, setStartSuggestions] = useState([]);
  const [endSuggestions, setEndSuggestions] = useState([]);
  const [startSuggestionSelected, setStartSuggestionSelected] = useState(false);
  const [endSuggestionSelected, setEndSuggestionSelected] = useState(false);

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
    if (formData.startArticle.length > 0 && !startSuggestionSelected) {
      fetchSuggestions(formData.startArticle, setStartSuggestions);
    } else {
      setStartSuggestions([]);
    }
  }, [formData.startArticle, startSuggestionSelected]);

  useEffect(() => {
    if (formData.endArticle.length > 0 && !endSuggestionSelected) {
      fetchSuggestions(formData.endArticle, setEndSuggestions);
    } else {
      setEndSuggestions([]);
    }
  }, [formData.endArticle, endSuggestionSelected]);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });
    name === 'startArticle' ? setStartSuggestionSelected(false) : setEndSuggestionSelected(false);

    if (name === 'startArticle') {
      setFormData((prevData) => ({ ...prevData, startUrl: '' }));
    } else if (name === 'endArticle') {
      setFormData((prevData) => ({ ...prevData, endUrl: '' }));
    }
  };

  const handleSelectSuggestion = (value, field, urlField, setSelected) => {
    const encodedValue = value.replace(/ /g, "_");
    setFormData({ ...formData, [field]: value, [urlField]: `https://en.wikipedia.org/wiki/${encodedValue}` });
    field === 'startArticle' ? setStartSuggestions([]) : setEndSuggestions([]);
    setSelected(true);
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    onFormSubmit(formData);
    setSubmittedData(formData);
  };

  const handleSwapInputs = () => {
    setFormData({
      ...formData,
      startArticle: formData.endArticle,
      startUrl: formData.endUrl,
      endArticle: formData.startArticle,
      endUrl: formData.startUrl
    });
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
            {startSuggestions.length > 0 && !startSuggestionSelected && (
              <div className="suggestion-container">
                <ul className="suggestion-list">
                  {startSuggestions.map((suggestion, index) => (
                    <li
                      key={index}
                      className="suggestion"
                      onClick={() => handleSelectSuggestion(suggestion, 'startArticle', 'startUrl', setStartSuggestionSelected)}
                    >
                      {suggestion}
                    </li>
                  ))}
                </ul>
              </div>
            )}
          </div>
        </label>
        <button type="button" onClick={handleSwapInputs} className="inline-button">
          <img className="arrow-img" src={arrow} alt="" /></button>
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
            {endSuggestions.length > 0 && !endSuggestionSelected && (
              <div className="suggestion-container">
                <ul className="suggestion-list">
                  {endSuggestions.map((suggestion, index) => (
                    <li
                      key={index}
                      className="suggestion"
                      onClick={() => handleSelectSuggestion(suggestion, 'endArticle', 'endUrl', setEndSuggestionSelected)}
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
          Algoritma <br />
          <select
            name="algoritma"
            value={formData.algoritma}
            onChange={handleChange}
            className="inline-input"
          >
            <option className="suggestion" value="BFS">BFS</option>
            <option className="suggestion" value="IDS">IDS</option>
          </select>
        </label>
        <button type="submit" className="inline-button">Find!</button>
      </form>

      {submittedData && (
        <div className="finding-route">
          {formData.endArticle.length === 0 || formData.startArticle.length === 0 ? (
            <p>Start and End cannot be blank!</p>
          ) : !startSuggestionSelected || !endSuggestionSelected ? (
            <p>Please select a suggestion for both start and end articles.</p>
          ) : (
            <div>
              <h2>Finding Route:</h2>
              <p>
                {submittedData.startArticle} To {submittedData.endArticle}
              </p>
            </div>
          )}
        </div>
      )}
    </div>
  );
};

export default FormInput;
