import { useCallback, useEffect, useState } from 'react';
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
  const [offset, setOffset] = useState(0);

  const fetchPostsData = useCallback(
    (offset: number) => {
      dispatch(fetchPosts({ limit: 5, offset: offset, statusID: 1 }));
    },
    [dispatch]
  );

  useEffect(() => {
    dispatch(resetStatus());
    fetchPostsData(offset);
  }, [dispatch, offset]);

  if (status === 'loading') {
    return <Loader />;
  }

  if (error) {
    return <PageError error={error} />;
  }

  const handleNext = () => {
    setOffset(offset + 5);
  };

  const handlePrevious = () => {
    if (offset > 5) {
      setOffset(offset - 5);
    }
  };

  return (
    <main className={style.main}>
      {data.map((post) => {
        return (
          <div key={post.id}>
            <h3>{post.title}</h3>
            <h4>{post.category}</h4>
          </div>
        );
      })}
      <button onClick={handlePrevious}>Previous</button>
      <button onClick={handleNext}>Next</button>
    </main>
  );
};
