import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import styles from './Login.module.css'; // Import the CSS module

function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        try {
            const response = await axios.post('http://localhost:8080/login', { username, password });
            if (response.data && response.data.token) {
                localStorage.setItem('jwtToken', response.data.token);
                localStorage.setItem('refreshToken', response.data.refreshToken);
                navigate('/timesheet');
            }
        } catch (error) {
            console.error('Login failed:', error);
            alert('Login failed. Please check your credentials.');
        }
    };
    
    return (
        <div className={styles.container}>
            <form className={styles.form} onSubmit={handleSubmit}>
                <h2>Timesheet App</h2>
                <div className={styles.inputGroup}>
                    <label>Username:</label>
                    <input
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required
                    />
                </div>
                <div className={styles.inputGroup}>
                    <label>Password:</label>
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <button type="submit">Login</button>
            </form>
        </div>
    );
}

export default Login;