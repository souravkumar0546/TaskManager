import React, { useState, useEffect } from 'react';
import ApiService from '../services/apiService';

const Profile = ({ updateAvatar }) => {
  const [user, setUser] = useState({
    name: '',
    email: '',
    avatar: '', // New state for avatar URL
  });

  const fetchProfile = async () => {
    try {
      const userProfile = await ApiService.getUserProfile();
      console.log(userProfile);
      setUser({
        name: userProfile.user.name,
        email: userProfile.user.email,
        avatar: userProfile.user.avatar, // Set avatar URL from backend
      });
    } catch (error) {
      console.error('Failed to fetch user profile:', error);
    }
  };

  useEffect(() => {
    fetchProfile();
  }, []);

  const [newProfilePicture, setNewProfilePicture] = useState(null);
  const [inputkey,setInputKey]=useState(Date.now());

  const handleImageInput = (e) => {
    if (e.target.files.length > 0) {
      const file = e.target.files[0];
      setNewProfilePicture(file); // Set the new profile picture file
    }
  };

  const handleUpdate = async () => {
    if (newProfilePicture) {
      try {
        const updatedProfile = await ApiService.updateProfilePicture(newProfilePicture);
        console.log(updatedProfile);
        setUser({ ...user, avatar: updatedProfile.avatar }); // Update avatar URL in state
        updateAvatar(updatedProfile.avatar); // Update avatar URL in App state
        setNewProfilePicture(null); // Clear the new profile picture state
        setInputKey(Date.now());
      } catch (error) {
        console.error('Failed to update profile picture', error);
      }
    }
  };

  return (
    <div className="container text-center mt-4">
      <img
        alt={user.name}
        src={`${user.avatar}?timestamp=${new Date().getTime()}`}
        className="rounded-circle"
        style={{ width: '100px', height: '100px', margin: '0 auto' }}
      />
      <div className="row justify-content-center align-items-center mb-3">
        <h5 className="mt-3">{user.name}</h5>
        <p className="text-muted mb-3">{user.email}</p>
        <div className="col-auto">
          <p className="mb-0">New Profile Picture:</p>
        </div>
        <div className="col-auto">
          <input style={{width:"90%"}} key={inputkey} type="file" accept="image/*" onChange={handleImageInput} />
        </div>
      </div>
      <button className="btn btn-primary" onClick={handleUpdate}>
        Update
      </button>
    </div>
  );
};

export default Profile;
