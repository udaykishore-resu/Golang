import React, { useState, useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { loginUser, forgotPassword, clearError } from '../redux/authSlice';
import styles from './Login.module.css';

function Login() {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [email, setEmail] = useState('');
    const [showForgotPassword, setShowForgotPassword] = useState(false);
    const [showPassword, setShowPassword] = useState(false);

    const dispatch = useDispatch();
    const navigate = useNavigate();
    const { loading, error, user } = useSelector((state) => state.auth);

    useEffect(() => {
        if (user) {
            navigate('/timesheet');
        }
    }, [user, navigate]);

    useEffect(() => {
        if (error) {
            alert(error);
            dispatch(clearError());
        }
    }, [error, dispatch]);

    const handleSubmit = async (e) => {
        e.preventDefault();
        dispatch(loginUser({ username, password }));
    };

    const handleForgotPassword = async (e) => {
        e.preventDefault();
        dispatch(forgotPassword(email));
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
                            <button type="submit" className={styles.login__submit} disabled={loading}>
                                <span className={styles.button__text}>{loading ? 'Logging in...' : 'Log In Now'}</span>
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
                            <button type="submit" className={styles.login__submit} disabled={loading}>
                                <span className={styles.button__text}>{loading ? 'Sending...' : 'Send Reset Link'}</span>
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
