import { createAsyncThunk } from '@reduxjs/toolkit';
import { APIResponse } from '../../types/api';
import { Post } from '../../types/post';

export const updatePost = createAsyncThunk<
  Post,
  Post,
  {
    rejectValue: string;
  }
>(
  'updatePost',
  async (updatingPost, { rejectWithValue }) => {
    const apiUrl = `${import.meta.env.VITE_BASE_URL}/article/${updatingPost.id}`;
    try {
      const response = await fetch(apiUrl, {
        method: 'PUT',
        body: JSON.stringify(updatingPost)
      });

      if (!response.ok) {
        throw new Error(`fetch error: ${response.status} - ${response.body}`);
      }

      const data: APIResponse<Post> = await response.json();

      if (!data.success) {
        return rejectWithValue(data.error.message)
      }

      return data.data;
    } catch (error) {
      return rejectWithValue((error as Error).message);
    }
  }
);