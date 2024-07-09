import React, { useState } from 'react';
import ApiService from '../services/apiService';
import { useNavigate } from 'react-router-dom';


const SignUp = () => {
  const [name, setName] = useState('');
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const [confirmPassword, setConfirmPassword] = useState('');
  const [emailError, setEmailError] = useState('');
  const [passwordError, setPasswordError] = useState('');
  const [confirmPasswordError, setConfirmPasswordError] = useState('');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const validateEmail = () => {
    const emailPattern = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    if (!email) {
      setEmailError('Email is required.');
    } else if (!emailPattern.test(email)) {
      setEmailError('Please enter a valid email address.');
    } else {
      setEmailError('');
    }
  };

  const validatePassword = () => {
    const passwordPattern = /^(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])[A-Za-z0-9!@#$%^&*]{8,}$/;
    if (!password) {
      setPasswordError('Password is required.');
    } else if (!passwordPattern.test(password)) {
      setPasswordError('Password must contain at least eight characters, one uppercase letter, one number, and one special character.');
    } else {
      setPasswordError('');
    }
  };

  const validateConfirmPassword = () => {
    if (!confirmPassword) {
      setConfirmPasswordError('Confirm Password is required.');
    } else if (password !== confirmPassword) {
      setConfirmPasswordError('Passwords do not match.');
    } else {
      setConfirmPasswordError('');
    }
  };

  const handleBackendErrors = (error) => {
    setError(error.message || 'An unexpected error occurred. Please try again.');
  };

  const handleSubmit = async (event) => {
    event.preventDefault();

    validateEmail();
    validatePassword();
    validateConfirmPassword();

    if (emailError || passwordError || confirmPasswordError) {
      return;
    }

    try {
      const response = await ApiService.signUp({ email, name, password });
      // Handle successful signup (e.g., redirect to login page)
      console.log('Signup successful:', response);
      alert('Signup successful!');
      navigate('/login');

    } catch (error) {
      handleBackendErrors(error);
    }
  };

  return (
    <div className="container">
      <h4>Sign Up</h4>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">

        <label htmlFor="name" className="form-label">Name</label>
          <input
            type="name"
            className="form-control"
            id="name"
            value={name}
            onChange={(e) => setName(e.target.value)}
            required
          />

          <label htmlFor="email" className="form-label">Email</label>
          <input
            type="email"
            className="form-control"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            onBlur={validateEmail} // Validate email on blur
            required
          />
          {emailError && <div className="text-danger">{emailError}</div>}
        </div>
        <div className="mb-3">
          <label htmlFor="password" className="form-label">Password</label>
          <input
            type="password"
            className="form-control"
            id="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            onBlur={validatePassword} // Validate password on blur
            required
          />
          {passwordError && <div className="text-danger">{passwordError}</div>}
        </div>
        <div className="mb-3">
          <label htmlFor="confirmPassword" className="form-label">Confirm Password</label>
          <input
            type="password"
            className="form-control"
            id="confirmPassword"
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.target.value)}
            onBlur={validateConfirmPassword} // Validate confirm password on blur
            required
          />
          {confirmPasswordError && <div className="text-danger">{confirmPasswordError}</div>}
        </div>
        {error && <div className="alert alert-danger">{error}</div>}
        <button type="submit" className="btn btn-primary">Sign Up</button>
      </form>
    </div>
  );
};

export default SignUp;
