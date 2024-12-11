import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
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
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetchProjects();
  }, []);

  const fetchProjects = async () => {
    setIsLoading(true);
    try {
      const response = await api.get('/projects');
      setProjects(response.data);
    } catch (err) {
      setError('Failed to fetch projects. Please try again.');
      console.error('Error fetching projects:', err);
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
      const response = await api.get(`/subprojects?project_id=${projectId}`);
      
      // Check if response and response.data are valid
      if (response && Array.isArray(response.data)) {
        setSubprojects(response.data);
        
        // If no subprojects are found, set subprojectId to the projectId
        if (response.data.length === 0) {
          setFormData(prev => ({ ...prev, subprojectId: projectId }));
          alert('No subprojects found for this project. You can submit using the project ID.');
        }
      } else {
        console.error('Unexpected response format:', response);
        setError('No subprojects found for this project.');
      }
      
    } catch (err) {
      setError('Failed to fetch subprojects. Please try again.');
      console.error('Error fetching subprojects:', err);
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
    // Ensure HoursSpent does not exceed 10
    const newHoursSpent = Math.max(0, Math.min(10, parseFloat(formData.hoursSpent || 0) + increment));
    setFormData(prev => ({
      ...prev,
      hoursSpent: newHoursSpent
    }));
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setIsLoading(true);

    const timesheetData = {
        ProjectID: parseInt(formData.projectId),
        SubProjectID: parseInt(formData.subprojectId),
        JiraSnowID: formData.jiraId,
        TaskDescription: formData.taskDescription,
        HoursSpent: parseInt(formData.hoursSpent),
        Comments: formData.comments
    };

    try {
        const response = await api.post('/timesheet', timesheetData);
        console.log('Timesheet submitted successfully:', response.data);
        navigate('/thank-you');
    } catch (error) {
        console.error('Error submitting timesheet:', error.response?.data || error.message);
        alert('Failed to submit timesheet entry. Please try again.');
    } finally {
        setIsLoading(false);
    }
};

  if (isLoading && projects.length === 0) return <div className={styles.loading}>Loading projects...</div>;
  if (error) return <div className={styles.error}>{error}</div>;

  return (
    <div className={styles.container}>
      <div className={styles.formWrapper}>
        <h2>Timesheet Entry</h2>
        <form onSubmit={handleSubmit} className={styles.form}>
          <div className={styles.formGroup}>
            <label htmlFor="projectId">Project:</label>
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
            <label htmlFor="subprojectId">Subproject:</label>
            <select
              id="subprojectId"
              name="subprojectId"
              value={formData.subprojectId}
              onChange={handleInputChange}
              required
              disabled={!formData.projectId}
            >
              <option value="">Select a subproject</option>
              {subprojects.length > 0 ? (
                subprojects.map(subproject => (
                  <option key={subproject.sub_project_id} value={subproject.sub_project_id}>
                    {subproject.sub_project_id} - {subproject.sub_project_name}
                  </option>
                ))
              ) : (
                formData.projectId && (
                  <option value={formData.projectId}>
                    {formData.projectId} - No Subprojects Available
                  </option>
                )
              )}
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
                max="10" // Limit input to a maximum of 10 hours
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

          <button type="submit" className={styles.submitButton} disabled={isLoading}>
            {isLoading ? 'Submitting...' : 'Submit Timesheet'}
          </button>
        </form>
      </div>
    </div>
  );
}

export default TimesheetEntry;