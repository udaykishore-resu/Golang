// src/components/TimesheetEntry.js

import React, { useState, useEffect } from 'react';
import api from '../services/api';

const TimesheetEntry = () => {
  const [projects, setProjects] = useState([]);
  const [subprojects, setSubprojects] = useState([]);
  const [formData, setFormData] = useState({
    projectId: '',
    subprojectId: '',
    jiraId: '',
    taskDescription: '',
    hoursSpent: '',
    comments: ''
  });
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchProjects();
  }, []);

  const fetchProjects = async () => {
    try {
      const response = await api.get('/projects');
      setProjects(response.data);
      setIsLoading(false);
    } catch (err) {
      setError('Failed to fetch projects. Please try again.');
      setIsLoading(false);
    }
  };

  const fetchSubprojects = async (projectId) => {
    try {
      const response = await api.get(`/subprojects/${projectId}`);
      setSubprojects(response.data);
    } catch (err) {
      setError('Failed to fetch subprojects. Please try again.');
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData({ ...formData, [name]: value });

    if (name === 'projectId') {
      fetchSubprojects(value);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await api.post('/timesheet', formData);
      alert('Timesheet entry submitted successfully!');
      // Reset form or redirect as needed
    } catch (err) {
      setError('Failed to submit timesheet. Please try again.');
    }
  };

  if (isLoading) return <div>Loading...</div>;
  if (error) return <div>{error}</div>;

  return (
    <form onSubmit={handleSubmit}>
      <select
        name="projectId"
        value={formData.projectId}
        onChange={handleInputChange}
        required
      >
        <option value="">Select Project</option>
        {projects.map(project => (
          <option key={project.project_id} value={project.project_id}>
            {project.project_id} - {project.project_name}
          </option>
        ))}
      </select>

      <select
        name="subprojectId"
        value={formData.subprojectId}
        onChange={handleInputChange}
        required
      >
        <option value="">Select Subproject</option>
        {subprojects.map(subproject => (
          <option key={subproject.id} value={subproject.id}>
            {subproject.name}
          </option>
        ))}
      </select>

      <input
        type="text"
        name="jiraId"
        value={formData.jiraId}
        onChange={handleInputChange}
        placeholder="Jira ID"
        required
      />

      <textarea
        name="taskDescription"
        value={formData.taskDescription}
        onChange={handleInputChange}
        placeholder="Task Description"
        required
      />

      <input
        type="number"
        name="hoursSpent"
        value={formData.hoursSpent}
        onChange={handleInputChange}
        placeholder="Hours Spent"
        required
      />

      <textarea
        name="comments"
        value={formData.comments}
        onChange={handleInputChange}
        placeholder="Comments"
      />

      <button type="submit">Submit Timesheet</button>
    </form>
  );
};

export default TimesheetEntry;