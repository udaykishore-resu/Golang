import React, { useState } from 'react';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import styles from './Login.module.css';

function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [showForgotPassword, setShowForgotPassword] = useState(false);
    const [showPassword, setShowPassword] = useState(false);
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

    const handleForgotPassword = async (e) => {
        e.preventDefault();
        try {
            await axios.post('http://localhost:8080/forgot-password', { email });
            alert('Password reset link sent to your email.');
            setShowForgotPassword(false);
        } catch (error) {
            console.error('Failed to send reset link:', error);
            alert('Failed to send reset link. Please try again.');
        }
    };
    
    return (
        <div className={styles.container}>
            <div className={styles.screen}>
                <div className={styles.screen__content}>
                    <h1 className={styles.app_title}>Timesheet App</h1>
                    {!showForgotPassword ? (
                        <form className={styles.login} onSubmit={handleSubmit}>
                            <div className={styles.login__field}>
                                <i className={`${styles.login__icon} fas fa-user`}></i>
                                <input 
                                    type="text" 
                                    className={styles.login__input} 
                                    placeholder="Username / Email"
                                    value={username}
                                    onChange={(e) => setUsername(e.target.value)}
                                    required
                                />
                            </div>
                            <div className={styles.login__field}>
                                <i className={`${styles.login__icon} fas fa-lock`}></i>
                                <input 
                                    type={showPassword ? "text" : "password"}
                                    className={styles.login__input} 
                                    placeholder="Password"
                                    value={password}
                                    onChange={(e) => setPassword(e.target.value)}
                                    required
                                />
                                <i 
                                    className={`${styles.eye__icon} fas ${showPassword ? 'fa-eye-slash' : 'fa-eye'}`}
                                    onClick={() => setShowPassword(!showPassword)}
                                ></i>
                            </div>
                            <button type="submit" className={styles.login__submit}>
                                <span className={styles.button__text}>Log In Now</span>
                                <i className={`${styles.button__icon} fas fa-chevron-right`}></i>
                            </button>
                            <div className={styles.forgot_password}>
                                <a href="#" onClick={() => setShowForgotPassword(true)}>Forgot Password?</a>
                            </div>
                        </form>
                    ) : (
                        <form className={styles.login} onSubmit={handleForgotPassword}>
                            <div className={styles.login__field}>
                                <i className={`${styles.login__icon} fas fa-envelope`}></i>
                                <input 
                                    type="email" 
                                    className={styles.login__input} 
                                    placeholder="Registered Email"
                                    value={email}
                                    onChange={(e) => setEmail(e.target.value)}
                                    required
                                />
                            </div>
                            <button type="submit" className={styles.login__submit}>
                                <span className={styles.button__text}>Send Reset Link</span>
                                <i className={`${styles.button__icon} fas fa-paper-plane`}></i>
                            </button>
                            <div className={styles.forgot_password}>
                                <a href="#" onClick={() => setShowForgotPassword(false)}>Back to Login</a>
                            </div>
                        </form>
                    )}
                </div>
                <div className={styles.screen__background}>
                    <span className={`${styles.screen__background__shape} ${styles.screen__background__shape4}`}></span>
                    <span className={`${styles.screen__background__shape} ${styles.screen__background__shape3}`}></span>		
                    <span className={`${styles.screen__background__shape} ${styles.screen__background__shape2}`}></span>
                    <span className={`${styles.screen__background__shape} ${styles.screen__background__shape1}`}></span>
                </div>		
            </div>
        </div>
    );
}

export default Login;