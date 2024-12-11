import React from 'react';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import '@testing-library/jest-dom';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import Login from './Login';

// Mock axios
jest.mock('axios');

// Mock react-router-dom
jest.mock('react-router-dom', () => ({
  ...jest.requireActual('react-router-dom'),
  useNavigate: jest.fn(),
}));

describe('Login Component', () => {
  const mockNavigate = jest.fn();

  beforeEach(() => {
    useNavigate.mockImplementation(() => mockNavigate);
  });

  test('renders login form', () => {
    render(<Login />);
    expect(screen.getByPlaceholderText('Username / Email')).toBeInTheDocument();
    expect(screen.getByPlaceholderText('Password')).toBeInTheDocument();
    expect(screen.getByText('Log In Now')).toBeInTheDocument();
  });

  test('handles successful login', async () => {
    axios.post.mockResolvedValueOnce({ data: { token: 'fakeToken', refreshToken: 'fakeRefreshToken' } });

    render(<Login />);

    fireEvent.change(screen.getByPlaceholderText('Username / Email'), { target: { value: 'testuser' } });
    fireEvent.change(screen.getByPlaceholderText('Password'), { target: { value: 'password123' } });
    fireEvent.click(screen.getByText('Log In Now'));

    await waitFor(() => {
      expect(axios.post).toHaveBeenCalledWith('http://localhost:8080/login', { username: 'testuser', password: 'password123' });
      expect(localStorage.setItem).toHaveBeenCalledWith('jwtToken', 'fakeToken');
      expect(localStorage.setItem).toHaveBeenCalledWith('refreshToken', 'fakeRefreshToken');
      expect(mockNavigate).toHaveBeenCalledWith('/timesheet');
    });
  });

  test('handles login failure', async () => {
    axios.post.mockRejectedValueOnce(new Error('Login failed'));
    const alertMock = jest.spyOn(window, 'alert').mockImplementation(() => {});

    render(<Login />);

    fireEvent.change(screen.getByPlaceholderText('Username / Email'), { target: { value: 'testuser' } });
    fireEvent.change(screen.getByPlaceholderText('Password'), { target: { value: 'wrongpassword' } });
    fireEvent.click(screen.getByText('Log In Now'));

    await waitFor(() => {
      expect(axios.post).toHaveBeenCalledWith('http://localhost:8080/login', { username: 'testuser', password: 'wrongpassword' });
      expect(alertMock).toHaveBeenCalledWith('Login failed. Please check your credentials.');
    });

    alertMock.mockRestore();
  });

  test('toggles password visibility', () => {
    render(<Login />);
    const passwordInput = screen.getByPlaceholderText('Password');
    const toggleButton = screen.getByRole('button', { name: /toggle password visibility/i });

    expect(passwordInput).toHaveAttribute('type', 'password');
    fireEvent.click(toggleButton);
    expect(passwordInput).toHaveAttribute('type', 'text');
    fireEvent.click(toggleButton);
    expect(passwordInput).toHaveAttribute('type', 'password');
  });

  test('shows forgot password form', () => {
    render(<Login />);
    fireEvent.click(screen.getByText('Forgot Password?'));
    expect(screen.getByPlaceholderText('Registered Email')).toBeInTheDocument();
    expect(screen.getByText('Send Reset Link')).toBeInTheDocument();
  });

  test('handles forgot password submission', async () => {
    axios.post.mockResolvedValueOnce({});
    const alertMock = jest.spyOn(window, 'alert').mockImplementation(() => {});

    render(<Login />);
    fireEvent.click(screen.getByText('Forgot Password?'));

    fireEvent.change(screen.getByPlaceholderText('Registered Email'), { target: { value: 'test@example.com' } });
    fireEvent.click(screen.getByText('Send Reset Link'));

    await waitFor(() => {
      expect(axios.post).toHaveBeenCalledWith('http://localhost:8080/forgot-password', { email: 'test@example.com' });
      expect(alertMock).toHaveBeenCalledWith('Password reset link sent to your email.');
    });

    alertMock.mockRestore();
  });
});