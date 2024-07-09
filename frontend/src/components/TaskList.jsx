import React from 'react';
import ApiService from '../services/apiService';
import '../assets/styles/TaskList.css'; 

const TaskList = ({ tasks, fetchTasks }) => {
  const handleChangeStatus = async (event, taskId) => {
    try {
      const newStatus = event.target.value;
      await ApiService.UpdateTask(taskId, newStatus);
      fetchTasks(); 
      console.log(`Updated status for task ${taskId} to ${newStatus}`);
    } catch (error) {
      console.error(`Failed to update status for task ${taskId}`, error);
    }
  };

  const handleDeleteTask = async (taskId,title) => {
    try {
      await ApiService.deleteTask(taskId);
      fetchTasks();
      alert(`Deleted task- ${title}`);
    } catch (error) {
      console.error(`Failed to delete task ${taskId}`, error);
    }
  };

  const getTaskClass = (status) => {
    switch (status) {
      case 'To Do':
        return 'task-to-do';
      case 'In Progress':
        return 'task-in-progress';
      case 'Done':
        return 'task-done';
      default:
        return '';
    }
  };

  return (
    <ul className="list-group">
      {tasks.map(task => (
        <li key={task.id} className={`list-group-item d-flex flex-column flex-md-row justify-content-between align-items-center ${getTaskClass(task.status)}`}>
          <div className="flex-grow-1 mb-3 mb-md-0" style={{width: "75%"}}>
            <h5>{task.title}</h5>
            <p className="mb-0">{task.description}</p>
          </div>
          <div className="d-flex align-items-center mt-2 mt-md-0">
            <select
              className="form-select mb-2 mb-md-0 me-md-2"
              value={task.status}
              onChange={(e) => handleChangeStatus(e, task.id)}
              title={task.status} // Tooltip with current status
            >
              <option value="To Do">To Do</option>
              <option value="In Progress">In Progress</option>
              <option value="Done">Done</option>
            </select>
            <button className="btn btn-danger icon-light-black" aria-label="Delete Task" onClick={() => handleDeleteTask(task.id,task.title)}>
              <i className="bi bi-trash"></i> 
            </button>
          </div>
        </li>
      ))}
    </ul>
  );
};

export default TaskList;
