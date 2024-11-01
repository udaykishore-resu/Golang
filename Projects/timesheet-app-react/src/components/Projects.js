// src/components/Projects.js

import React, { useState, useEffect } from 'react';
import api from '../services/api';

const Projects = () => {
  const [projects, setProjects] = useState([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchProjects();
  }, []);

  const fetchProjects = async () => {
    setIsLoading(true);
    try {
      const response = await api.get('/projects');
      setProjects(response.data);
      setIsLoading(false);
    } catch (err) {
      setError(err.response?.data?.message || 'Failed to fetch projects. Please try again later.');
      setIsLoading(false);
      console.error('Error fetching projects:', err);
      
      // Handle unauthorized access
      if (err.response && err.response.status === 401) {
        // Redirect to login page or show login modal
        console.log('User is not authenticated. Redirecting to login...');
        // Implement your redirect logic here, e.g.:
        // history.push('/login');
      }
    }
  };

  if (isLoading) {
    return <div>Loading projects...</div>;
  }

  if (error) {
    return <div>{error}</div>;
  }

  return (
    <div>
      <h2>Projects</h2>
      {projects.length === 0 ? (
        <p>No projects available.</p>
      ) : (
        <ul>
          {projects.map((project) => (
            <li key={project.project_id}>
              {project.project_id} - {project.project_name}
            </li>
          ))}
        </ul>
      )}
    </div>
  );
};

export default Projects;