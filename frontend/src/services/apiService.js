const API_URL = "https://taskmanager-9ij0.onrender.com";


const ApiService = {
  signUp: async (credentials) => {
    const response = await fetch(`${API_URL}/signup`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(credentials),
    });
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Sign Up failed');
    }
    return await response.json();
  },
  
  login: async (credentials) => {
    try {
      const response = await fetch(`${API_URL}/login`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        credentials: 'include', 
        body: JSON.stringify(credentials),
      });
      
      if (!response.ok) {
        const errorData = await response.json();
        throw new Error(errorData.message || 'Login failed');
      }

      return await response.json();
    } catch (error) {
      console.error('Login failed:', error);
      throw error; // Propagate the error back to the component
    }
  },
  
  logout: async () => {
    const response = await fetch(`${API_URL}/logout`, {
      method: 'POST',
      headers: {},
      credentials: 'include',
    });
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Logout failed');
    }
    return await response.json();
  },
  
  getUserProfile: async () => {
    const response = await fetch(`${API_URL}/api/user`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include', // Include cookies in the request
    });

    if (!response.ok) {
      if (response.status === 401) {
        throw new Error('Unauthorized. Please log in again.');
      }
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to fetch profile');
    }

    return await response.json();
  },

  updateProfilePicture: async (profilePicture) => {
    const formData = new FormData();
    formData.append('avatar', profilePicture);

    const response = await fetch(`${API_URL}/api/user/avatar`, {
      method: 'POST',
      headers: {
      },
      credentials: 'include',
      body: formData,
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to update profile picture');
    }

    return await response.json();
  },

  createTask: async (task) => {
    const response = await fetch(`${API_URL}/api/tasks`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include', // Include cookies in the request
      body: JSON.stringify(task),
    });

    if (!response.ok) {
      if (response.status === 401) {
        throw new Error('Unauthorized. Please log in again.');
      }
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to create task');
    }

    return await response.json();
  },

  getTasks: async () => {
    const response = await fetch(`${API_URL}/api/tasks`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include', // Include cookies in the request
    });

    if (!response.ok) {
      if (response.status === 401) {
        throw new Error('Unauthorized. Please log in again.');
      }
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to fetch tasks');
    }

    return await response.json();
  },

  UpdateTask: async (taskId, status) => {
    const response = await fetch(`${API_URL}/api/tasks/${taskId}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      credentials: 'include',
      body: JSON.stringify({ status }),
    });
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to update task status');
    }
    return await response.json();
  },

  deleteTask: async (taskId) => {
    const response = await fetch(`${API_URL}/api/tasks/${taskId}`, {
      method: 'DELETE',
      headers: {
      },
      credentials: 'include', // Include cookies in the request
    });
    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || 'Failed to delete task');
    }
    return await response.text(); // Return text as delete doesn't return json
  },

  // Add more endpoints as needed
};

export default ApiService;
