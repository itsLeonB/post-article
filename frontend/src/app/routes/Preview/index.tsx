import { useCallback, useEffect } from 'react';
import { useAppDispatch } from '../../../hooks/useAppDispatch';
import { useAppSelector } from '../../../hooks/useAppSelector';
import { fetchPosts } from '../../../store/thunks/postListThunk';
import { resetStatus } from '../../../store/slices/postListSlice';
import Loader from '../../../components/Loader';
import PageError from '../../../components/PageError';

import style from './style.module.css';

export const Preview: React.FC = () => {
  const dispatch = useAppDispatch();
  const { data, status, error } = useAppSelector((state) => state.postList);

  const fetchPostsData = useCallback(() => {
    dispatch(fetchPosts());
  }, [dispatch]);

  useEffect(() => {
    dispatch(resetStatus());
    fetchPostsData();
  }, [dispatch, fetchPostsData]);

  if (status === 'loading') {
    return <Loader />;
  }

  if (error) {
    return <PageError error={error} />;
  }

  return (
    <main className={style.main}>
      {data
        .filter((item) => {
          return item.status_id === 1;
        })
        .map((post) => {
          return (
            <div key={post.id}>
              <h3>{post.title}</h3>
              <h4>{post.category}</h4>
            </div>
          );
        })}
    </main>
  );
};
