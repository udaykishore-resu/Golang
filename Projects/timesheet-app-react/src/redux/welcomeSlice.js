// src/redux/welcomeSlice.js
import { createSlice } from '@reduxjs/toolkit';

const welcomeSlice = createSlice({
  name: 'welcome',
  initialState: {
    message: 'Welcome to Timesheet App'
  },
  reducers: {
    setWelcomeMessage: (state, action) => {
      state.message = action.payload;
    }
  }
});

export const { setWelcomeMessage } = welcomeSlice.actions;
export default welcomeSlice.reducer;
