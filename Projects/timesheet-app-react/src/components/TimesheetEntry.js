// src/components/TimesheetEntry.js

import React, { useState, useEffect } from 'react';
import api from '../services/api';
import styles from './TimesheetEntry.module.css';

function TimesheetEntry() {
  const [formData, setFormData] = useState({
    projectId: '',
    subprojectId: '',
    jiraId: '',
    taskDescription: '',
    hoursSpent: 0,
    comments: ''
  });
  const [projects, setProjects] = useState([]);
  const [subprojects, setSubprojects] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchProjects();
  }, []);

  const fetchProjects = async () => {
    setIsLoading(true);
    try {
      console.log('Fetching projects...');
      const response = await api.get('/projects');
      console.log('Projects response:', response.data);
      setProjects(response.data);
    } catch (err) {
      console.error('Error fetching projects:', err);
      setError('Failed to fetch projects. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  const fetchSubprojects = async (projectId) => {
    if (!projectId) {
      setSubprojects([]);
      return;
    }
    setIsLoading(true);
    try {
      console.log(`Fetching subprojects for project ${projectId}...`);
      const response = await api.get(`/subprojects?project_id=${projectId}`);
      console.log('Subprojects response:', response.data);
      setSubprojects(response.data);
    } catch (err) {
      console.error('Error fetching subprojects:', err);
      setError('Failed to fetch subprojects. Please try again.');
    } finally {
      setIsLoading(false);
    }
  };

  const handleInputChange = (e) => {
    const { name, value } = e.target;
    setFormData(prevData => ({
      ...prevData,
      [name]: value
    }));

    if (name === 'projectId') {
      setFormData(prevData => ({ ...prevData, subprojectId: '' }));
      fetchSubprojects(value);
    }
  };

  const handleHoursChange = (increment) => {
    setFormData(prev => ({
      ...prev,
      hoursSpent: Math.max(0, parseFloat(prev.hoursSpent) + increment)
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      console.log('Submitting timesheet:', formData);
      await api.post('/timesheet', formData);
      alert('Timesheet entry submitted successfully!');
      // Reset form
      setFormData({
        projectId: '',
        subprojectId: '',
        jiraId: '',
        taskDescription: '',
        hoursSpent: 0,
        comments: ''
      });
    } catch (error) {
      console.error('Error submitting timesheet:', error);
      alert('Failed to submit timesheet entry. Please try again.');
    }
  };

  if (isLoading && projects.length === 0) return <div>Loading projects...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div className={styles.container}>
      <h2>Timesheet Entry</h2>
      <form onSubmit={handleSubmit} className={styles.form}>
        <div className={styles.formGroup}>
          <label htmlFor="projectId">Project ID:</label>
          <select
            id="projectId"
            name="projectId"
            value={formData.projectId}
            onChange={handleInputChange}
            required
          >
            <option value="">Select a project</option>
            {projects.map(project => (
              <option key={project.project_id} value={project.project_id}>
                {project.project_id} - {project.project_name}
              </option>
            ))}
          </select>
        </div>

        <div className={styles.formGroup}>
          <label htmlFor="subprojectId">Subproject ID:</label>
          <select
            id="subprojectId"
            name="subprojectId"
            value={formData.subprojectId}
            onChange={handleInputChange}
            required
            disabled={!formData.projectId}
          >
            <option value="">Select a subproject</option>
            {subprojects.map(subproject => (
              <option key={subproject.sub_project_id} value={subproject.sub_project_id}>
                {subproject.sub_project_id} - {subproject.sub_project_name}
              </option>
            ))}
          </select>
        </div>

        <div className={styles.formGroup}>
          <label htmlFor="jiraId">Jira ID:</label>
          <input
            type="text"
            id="jiraId"
            name="jiraId"
            value={formData.jiraId}
            onChange={handleInputChange}
            required
          />
        </div>

        <div className={styles.formGroup}>
          <label htmlFor="taskDescription">Task Description:</label>
          <textarea
            id="taskDescription"
            name="taskDescription"
            value={formData.taskDescription}
            onChange={handleInputChange}
            required
          />
        </div>

        <div className={styles.formGroup}>
          <label htmlFor="hoursSpent">Hours Spent:</label>
          <div className={styles.hoursInput}>
            <button type="button" onClick={() => handleHoursChange(-0.5)}>-</button>
            <input
              type="number"
              id="hoursSpent"
              name="hoursSpent"
              value={formData.hoursSpent}
              onChange={handleInputChange}
              step="0.5"
              min="0"
              required
            />
            <button type="button" onClick={() => handleHoursChange(0.5)}>+</button>
          </div>
        </div>

        <div className={styles.formGroup}>
          <label htmlFor="comments">Comments:</label>
          <textarea
            id="comments"
            name="comments"
            value={formData.comments}
            onChange={handleInputChange}
          />
        </div>

        <button type="submit" className={styles.submitButton}>Submit</button>
      </form>
    </div>
  );
}

export default TimesheetEntry;