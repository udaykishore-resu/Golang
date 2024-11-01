import React, { useEffect, useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';

function Welcome() {
  const [message, setMessage] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    const fetchWelcomeMessage = async () => {
      try {
        const token = localStorage.getItem('token');
        if (!token) {
          navigate('/login');
          return;
        }
        const response = await axios.get('http://localhost:8080', {
          headers: { Authorization: `Bearer ${token}` }
        });
        setMessage(response.data.message);
      } catch (error) {
        console.error('Error fetching welcome message:', error);
        navigate('/login');
      }
    };

    fetchWelcomeMessage();
  }, [navigate]);

  return (
    <div>
      <h2>Welcome</h2>
      <p>{message}</p>
    </div>
  );
}

export default Welcome;