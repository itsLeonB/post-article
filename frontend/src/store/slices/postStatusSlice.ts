import { createSlice } from '@reduxjs/toolkit';
import { fetchPostStatuses } from '../thunks/postStatusThunk';
import { PostStatus } from '../../types/post';

export interface PostStatusState {
  data: PostStatus[];
  status: 'idle' | 'loading' | 'success' | 'failed';
  error: string | null;
}

const initialState: PostStatusState = {
  data: [],
  status: 'idle',
  error: null,
};

const postStatusSlice = createSlice({
  name: 'postStatus',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchPostStatuses.pending, (state) => {
        state.status = 'loading';
      })
      .addCase(fetchPostStatuses.fulfilled, (state, action) => {
        state.status = 'success';
        state.data = action.payload
        state.error = null;
      })
      .addCase(fetchPostStatuses.rejected, (state, action) => {
        state.status = 'failed';
        state.error = action.payload ?? 'fetch failed';
      });
  },
});

export default postStatusSlice.reducer;