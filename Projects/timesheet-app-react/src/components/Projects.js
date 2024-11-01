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
    try {
      const response = await api.get('/projects');
      setProjects(response.data);
      setIsLoading(false);
    } catch (err) {
      setError('Failed to fetch projects. Please try again.');
      setIsLoading(false);
    }
  };

  if (isLoading) return <div>Loading projects...</div>;
  if (error) return <div>{error}</div>;

  return (
    <div>
      <h2>Projects</h2>
      <ul>
        {projects.map(project => (
          <li key={project.project_id}>{project.project_id} - {project.project_name}</li>
        ))}
      </ul>
    </div>
  );
};

export default Projects;