import { createAsyncThunk } from '@reduxjs/toolkit';
import { APIResponse } from '../../types/api';
import { PostStatus } from '../../types/post';

export const fetchPostStatuses = createAsyncThunk<
  PostStatus[],
  void,
  {
    rejectValue: string;
  }
>(
  'fetchPostStatuses',
  async (_, { rejectWithValue }) => {
    const apiUrl = `${import.meta.env.VITE_BASE_URL}/article-statuses`;
    try {
      const response = await fetch(apiUrl);

      if (!response.ok) {
        throw new Error(`fetch error: ${response.status} - ${response.body}`);
      }

      const data: APIResponse<PostStatus[]> = await response.json();

      if (!data.success) {
        return rejectWithValue(data.error.message)
      }

      return data.data;
    } catch (error) {
      return rejectWithValue((error as Error).message);
    }
  }
);