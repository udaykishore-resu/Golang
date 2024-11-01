import React, { useState, useEffect } from 'react';
import axios from 'axios';
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

  useEffect(() => {
    console.log('Component mounted, calling fetchProjects');
    fetchProjects();
  }, []);

  useEffect(() => {
    console.log('Projects state updated:', projects);
  }, [projects]);

  const fetchProjects = async () => {
    setIsLoading(true);
    try {
      console.log('Attempting to fetch projects...');
      const response = await axios.get('http://localhost:8080/projects');
      console.log('API response received:', response);
      console.log('Raw API response data:', JSON.stringify(response.data, null, 2));
      
      if (Array.isArray(response.data) && response.data.length > 0) {
        console.log('Projects found:', response.data.length);
        const formattedProjects = response.data.map(project => ({
          id: project.project_id,
          name: `${project.project_id} - ${project.project_name}`
        }));
        console.log('Formatted projects:', formattedProjects);
        setProjects(formattedProjects);
      } else {
        console.error('Unexpected API response structure or empty array');
        setProjects([]);
      }
    } catch (error) {
      console.error('Error fetching projects:', error);
      if (error.response) {
        console.error('Error response:', error.response.data);
        console.error('Error status:', error.response.status);
      } else if (error.request) {
        console.error('No response received:', error.request);
      } else {
        console.error('Error message:', error.message);
      }
      setProjects([]);
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
      // For now, we're not handling subprojects
      // You may need to fetch subprojects based on the selected project
      setSubprojects([]);
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
      await axios.post('http://localhost:8080/timesheet', formData);
      alert('Timesheet entry submitted successfully!');
      // Reset form or redirect as needed
    } catch (error) {
      console.error('Error submitting timesheet:', error);
      alert('Failed to submit timesheet entry. Please try again.');
    }
  };

  return (
    <div className={styles.container}>
      <h2>Timesheet Entry</h2>
      {/* Temporary button for manual fetching */}
      <button onClick={fetchProjects}>Fetch Projects Manually</button>
      <form onSubmit={handleSubmit} className={styles.form}>
        <div className={styles.formGroup}>
          <label htmlFor="projectId">Project ID:</label>
          {isLoading ? (
            <p>Loading projects...</p>
          ) : projects.length > 0 ? (
            <select
              id="projectId"
              name="projectId"
              value={formData.projectId}
              onChange={handleInputChange}
              required
            >
              <option value="">Select a project</option>
              {projects.map(project => (
                <option key={project.id} value={project.id}>{project.name}</option>
              ))}
            </select>
          ) : (
            <p>No projects available. {projects.length === 0 ? '(Empty array)' : '(Unexpected data)'}</p>
          )}
        </div>

        <div className={styles.formGroup}>
          <label htmlFor="subprojectId">Subproject ID:</label>
          <select
            id="subprojectId"
            name="subprojectId"
            value={formData.subprojectId}
            onChange={handleInputChange}
            required
          >
            <option value="">Select a subproject</option>
            {subprojects.map(subproject => (
              <option key={subproject.id} value={subproject.id}>{subproject.name}</option>
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