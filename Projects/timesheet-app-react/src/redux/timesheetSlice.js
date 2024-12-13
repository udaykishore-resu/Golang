// src/redux/timesheetSlice.js
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import api from '../services/api';

export const fetchProjects = createAsyncThunk(
  'timesheet/fetchProjects',
  async (_, { rejectWithValue }) => {
    try {
      const response = await api.get('/projects');
      return response.data;
    } catch (error) {
      return rejectWithValue('Failed to fetch projects. Please try again.');
    }
  }
);

export const fetchSubprojects = createAsyncThunk(
  'timesheet/fetchSubprojects',
  async (projectId, { rejectWithValue }) => {
    try {
      const response = await api.get(`/subprojects?project_id=${projectId}`);
      return response.data;
    } catch (error) {
      return rejectWithValue('Failed to fetch subprojects. Please try again.');
    }
  }
);

export const submitTimesheet = createAsyncThunk(
  'timesheet/submitTimesheet',
  async (timesheetData, { rejectWithValue }) => {
    try {
      const response = await api.post('/timesheet', timesheetData);
      return response.data;
    } catch (error) {
      return rejectWithValue('Failed to submit timesheet entry. Please try again.');
    }
  }
);

const timesheetSlice = createSlice({
  name: 'timesheet',
  initialState: {
    formData: {
      projectId: '',
      subprojectId: '',
      jiraId: '',
      taskDescription: '',
      hoursSpent: 0,
      comments: ''
    },
    projects: [],
    subprojects: [],
    isLoading: false,
    error: null
  },
  reducers: {
    updateFormData: (state, action) => {
      state.formData = { ...state.formData, ...action.payload };
    },
    resetForm: (state) => {
      state.formData = {
        projectId: '',
        subprojectId: '',
        jiraId: '',
        taskDescription: '',
        hoursSpent: 0,
        comments: ''
      };
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchProjects.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(fetchProjects.fulfilled, (state, action) => {
        state.isLoading = false;
        state.projects = action.payload;
      })
      .addCase(fetchProjects.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload;
      })
      .addCase(fetchSubprojects.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(fetchSubprojects.fulfilled, (state, action) => {
        state.isLoading = false;
        state.subprojects = action.payload;
      })
      .addCase(fetchSubprojects.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload;
      })
      .addCase(submitTimesheet.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(submitTimesheet.fulfilled, (state) => {
        state.isLoading = false;
      })
      .addCase(submitTimesheet.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.payload;
      });
  }
});

export const { updateFormData, resetForm } = timesheetSlice.actions;
export default timesheetSlice.reducer;
