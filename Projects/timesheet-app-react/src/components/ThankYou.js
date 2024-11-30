import React from 'react';
import { useNavigate } from 'react-router-dom';
import styles from './ThankYou.module.css';

function ThankYou() {
    const navigate = useNavigate();

    const handleBackToTimesheet = () => {
        navigate('/timesheet');
    };

    return (
        <div className={styles.container}>
            <h1>Thank You</h1>
            <p>Your timesheet has been successfully saved.</p>
            <button onClick={handleBackToTimesheet} className={styles.backButton}>Back to Timesheet</button>
        </div>
    );
}

export default ThankYou;