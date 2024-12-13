// src/redux/store.js
import { configureStore } from '@reduxjs/toolkit';
import authReducer from './authSlice';
import projectsReducer from './projectsSlice';
import timesheetReducer from './timesheetSlice';
import welcomeReducer from './welcomeSlice';
import thankYouReducer from './thankYouSlice';

const store = configureStore({
  reducer: {
    welcome: welcomeReducer,
    auth: authReducer,
    projects: projectsReducer,
    timesheet: timesheetReducer,
    thankYou: thankYouReducer,
  },
});

export default store;
