import { configureStore } from '@reduxjs/toolkit';
import postListReducer from './slices/postListSlice'

export const store = configureStore({
  reducer: {
    postList: postListReducer
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;