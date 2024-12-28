// src/redux/thankYouSlice.js
import { createSlice } from '@reduxjs/toolkit';

const thankYouSlice = createSlice({
  name: 'thankYou',
  initialState: {
    message: 'Your timesheet has been successfully saved.'
  },
  reducers: {
    setThankYouMessage: (state, action) => {
      state.message = action.payload;
    }
  }
});

export const { setThankYouMessage } = thankYouSlice.actions;
export default thankYouSlice.reducer;
