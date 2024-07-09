import React, { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import ApiService from '../services/apiService';

const NewTaskForm = () => {
  const navigate = useNavigate();
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [status, setStatus] = useState('To Do');
  const [error, setError] = useState('');

  const handleSubmit = async (event) => {
    event.preventDefault();
    try {
      const newTask = { title, description, status };
      await ApiService.createTask(newTask);
      alert(`Added new Task - ${title}`);
      navigate('/home');
    } catch (error) {
      console.error('Failed to create task:', error);
      setError(error.message); // Display error message
    }
  };

  const handleStatusChange = (event) => {
    setStatus(event.target.value);
  };

  return (
    <div className="container">
      <h4>Add New Task</h4>
      <form onSubmit={handleSubmit}>
        <div className="form-group mt-2">
          <label htmlFor="title">Title</label>
          <input
            type="text"
            className="form-control"
            id="title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
          />
        </div>
        <div className="form-group mt-2">
          <label htmlFor="description">Description</label>
          <textarea
            className="form-control"
            id="description"
            rows="3"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
          ></textarea>
        </div>
        <div className="form-group mt-2">
          <label htmlFor="status">Status</label>
          <select
            className="form-select"
            style={{width: "25%"}}
            id="status"
            value={status}
            onChange={handleStatusChange}
            required
          >
            <option value="To Do">To Do</option>
            <option value="In Progress">In Progress</option>
            <option value="Done">Done</option>
          </select>
        </div>
        {error && <p className="text-danger">{error}</p>}
        <br/>
        <div className="mt-2">
          <button type="submit" className="btn btn-primary">
            Add Task
          </button>
        </div>
      </form>
    </div>
  );
};

export default NewTaskForm;
