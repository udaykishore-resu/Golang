import React from 'react';
import { useNavigate } from 'react-router-dom';
import styles from './Welcome.module.css';

function Welcome() {
    const navigate = useNavigate();

    const handleLogin = () => {
        navigate('/login');
    };

    return (
        <div className={styles.container}>
            <h1>Welcome to Timesheet App</h1>
            <button onClick={handleLogin} className={styles.loginButton}>Login</button>
        </div>
    );
}

export default Welcome;