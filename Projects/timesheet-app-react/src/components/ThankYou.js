// src/components/ThankYou.js
import React from 'react';
import { useNavigate } from 'react-router-dom';
import { useSelector } from 'react-redux';
import styles from './ThankYou.module.css';

function ThankYou() {
    const navigate = useNavigate();
    const message = useSelector(state => state.thankYou.message);

    const handleBackToTimesheet = () => {
        navigate('/timesheet');
    };

    return (
        <div className={styles.container}>
            <h1>Thank You</h1>
            <p>{message}</p>
            <button onClick={handleBackToTimesheet} className={styles.backButton}>Back to Timesheet</button>
        </div>
    );
}

export default ThankYou;
