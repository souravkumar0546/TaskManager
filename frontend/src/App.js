import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import './assets/styles/App.css'; 
import NavBar from './components/NavBar';
import Home from './pages/Home';
import Profile from './pages/Profile';
import Login from './pages/Login';
import SignUp from './pages/SignUp';
import AddNewTask from './pages/AddNewTask';
import ApiService from './services/apiService';

const App = () => {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [avatar, setAvatar] = useState('');
  const [avatarKey, setAvatarKey] = useState(0); // State to force re-render

  useEffect(() => {
    const checkAuthentication = async () => {
      try {
        const userProfile = await ApiService.getUserProfile();
        setAvatar(userProfile.user.avatar); // Set avatar URL from user profile
        setIsAuthenticated(true);
      } catch (error) {
        console.error('Failed to fetch user profile:', error);
        setIsAuthenticated(false);
      }
    };

    checkAuthentication();
  }, []);

  const handleLogin = async () => {
    setIsAuthenticated(true);
    try {
      const userProfile = await ApiService.getUserProfile();
      setAvatar(userProfile.user.avatar);
    } catch (error) {
      console.error('Failed to fetch user profile after login:', error);
      setIsAuthenticated(false);
    }
  };

  const handleLogout = async () => {
    try {
      await ApiService.logout();
      setIsAuthenticated(false);
      setAvatar('');
    } catch (error) {
      console.error('Failed to log out:', error);
    }
  };

  // Function to update avatar URL after profile picture update
  const updateAvatar = (newAvatar) => {
    setAvatar(newAvatar);
    setAvatarKey(prevKey => prevKey + 1); // Increment the key to force re-render
  };

  return (
    <Router>
      <NavBar isAuthenticated={isAuthenticated} onLogout={handleLogout} avatar={avatar} key={avatarKey} />
      <Routes>
        <Route path="/" element={<Navigate to={isAuthenticated ? "/home" : "/login"} />} />
        <Route path="/home" element={isAuthenticated ? <Home /> : <Navigate to="/login" />} />
        <Route path="/profile" element={isAuthenticated ? <Profile updateAvatar={updateAvatar} /> : <Navigate to="/login" />} />
        <Route path="/login" element={<Login onLogin={handleLogin} />} />
        <Route path="/signup" element={<SignUp/>} />
        <Route path="/logout" element={<Navigate to="/login" />} />
        <Route path="/AddNewTask" element={isAuthenticated ? <AddNewTask /> : <Navigate to="/login" />} />
      </Routes>
    </Router>
  );
};

export default App;
