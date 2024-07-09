import React, { useState, useEffect } from 'react';
import TaskList from '../components/TaskList';
import { useNavigate } from 'react-router-dom';
import ApiService from '../services/apiService';

const Home = () => {
  const [tasks, setTasks] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [filter, setFilter] = useState('All');
  const [error, setError] = useState('');
  const navigate = useNavigate();

  const fetchTasks = async () => {
    try {
      const fetchedTasks = await ApiService.getTasks();
      setTasks(fetchedTasks.tasks);
    } catch (error) {
      console.error('Failed to fetch tasks:', error);
      setError(error.message);
    }
  };

  useEffect(() => {
    fetchTasks();
  }, []);

  const filteredTasks = tasks.filter(task => {
    return (
      (filter === 'All' || task.status === filter) &&
      task.title.toLowerCase().includes(searchTerm.toLowerCase())
    );
  });

  const handleAddTaskClick = () => {
    navigate('/AddNewTask');
  };

  return (
    <div className="container mt-4">
      <div className="d-flex justify-content-between align-items-center mb-3">
        <h4>My Tasks</h4>
        <button className="btn btn-primary" onClick={handleAddTaskClick}>Add Task</button>
      </div>

      <input
        type="text"
        className="form-control mb-3"
        placeholder="Search Tasks"
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />

      <select
        className="form-select mb-3"
        style={{ width: '25%' }}
        value={filter}
        onChange={(e) => setFilter(e.target.value)}
      >
        <option value="All">All</option>
        <option value="To Do">To Do</option>
        <option value="In Progress">In Progress</option>
        <option value="Done">Done</option>
      </select>

      {error && <p className="text-danger">{error}</p>}

      <TaskList tasks={filteredTasks} fetchTasks={fetchTasks} />
    </div>
  );
};

export default Home;
