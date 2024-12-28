// src/components/Projects.js
import React, { useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { fetchProjects } from '../redux/projectsSlice';

const Projects = () => {
  const dispatch = useDispatch();
  const { projects, isLoading, error } = useSelector((state) => state.projects);

  useEffect(() => {
    dispatch(fetchProjects());
  }, [dispatch]);

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
