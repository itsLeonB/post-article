import { configureStore } from '@reduxjs/toolkit';
import postListReducer from './slices/postListSlice'
import postStatusReducer from './slices/postStatusSlice'

export const store = configureStore({
  reducer: {
    postList: postListReducer,
    postStatus: postStatusReducer
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;