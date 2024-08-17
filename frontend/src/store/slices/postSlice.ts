import { createSlice } from '@reduxjs/toolkit';
import { Post } from '../../types/post';
import { fetchPost } from '../thunks/postThunk';
import { updatePost } from '../thunks/postUpdateThunk';

export interface PostState {
  data: Post;
  status: 'idle' | 'loading' | 'success' | 'failed';
  error: string | null;
}

const initialState: PostState = {
  data: {
    id: 0,
    title: "",
    content: "",
    category: "",
    status_id: 0
  },
  status: 'idle',
  error: null,
};

const postSlice = createSlice({
  name: 'post',
  initialState,
  reducers: {
    resetStatus: (state) => {
      state.status = 'idle';
      state.error = null;
    }
  },
  extraReducers: (builder) => {
    builder
      .addCase(fetchPost.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchPost.fulfilled, (state, action) => {
        state.status = 'success';
        state.data = action.payload;
        state.error = null;
      })
      .addCase(fetchPost.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload ?? 'fetch failed';
      })
      .addCase(updatePost.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(updatePost.fulfilled, (state, action) => {
        state.status = 'success';
        state.data = action.payload;
        state.error = null;
      })
      .addCase(updatePost.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload ?? 'fetch failed';
      })
  },
});

export const { resetStatus } = postSlice.actions;
export default postSlice.reducer;