import { createSlice } from '@reduxjs/toolkit';
import { Post } from '../../types/post';
import { fetchPosts } from '../thunks/postListThunk';

export interface PostListState {
  data: Post[];
  status: 'idle' | 'loading' | 'success' | 'failed';
  error: string | null;
}

const initialState: PostListState = {
  data: [],
  status: 'idle',
  error: null,
};

const postListSlice = createSlice({
  name: 'postList',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchPosts.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchPosts.fulfilled, (state, action) => {
        state.status = 'success';
        state.data = action.payload;
        state.error = null;
      })
      .addCase(fetchPosts.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload ?? 'fetch failed';
      });
  },
});

export default postListSlice.reducer;