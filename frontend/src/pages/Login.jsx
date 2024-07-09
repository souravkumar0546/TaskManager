import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import ApiService from '../services/apiService';

const Login = ({ onLogin }) => {
  const [email, setEmail] = useState('');
  const [emailError, setEmailError] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const checkAuthentication = async () => {
      try {
        await ApiService.getUserProfile();
        navigate('/home');
      } catch (error) {
      }
    };

    checkAuthentication();
  }, [navigate]);

  const validateEmail = () => {
    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!email) {
      setEmailError('Email is required.');
      return false;
    } else if (!emailPattern.test(email)) {
      setEmailError('Please enter a valid email address.');
      return false;
    } else {
      setEmailError('');
      return true;
    }
  };

  const handleBackendErrors = (error) => {
    setError(error.message || 'An unexpected error occurred. Please try again.');
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    const isEmailValid = validateEmail();
    if (!isEmailValid) {
      return;
    }

    setError('');

    try {
      await ApiService.login({ email, password });
      onLogin();
      navigate('/home');
    } catch (error) {
      handleBackendErrors(error);
    }
  };

  return (
    <div className="container">
      <h4>Login</h4>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="email" className="form-label">Email</label>
          <input
            type="email"
            className={`form-control ${emailError ? 'is-invalid' : ''}`}
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            onBlur={validateEmail}
            required
          />
          {emailError && <div className="invalid-feedback">{emailError}</div>}
        </div>
        <div className="mb-3">
          <label htmlFor="password" className="form-label">Password</label>
          <input
            type="password"
            className="form-control"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
        </div>
        {error && <div className="alert alert-danger">{error}</div>}
        <button type="submit" className="btn btn-primary">Login</button>
      </form>
    </div>
  );
};

export default Login;
