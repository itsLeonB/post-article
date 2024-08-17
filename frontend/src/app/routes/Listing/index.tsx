import { useEffect } from 'react';
import { useAppDispatch } from '../../../hooks/useAppDispatch';
import { useAppSelector } from '../../../hooks/useAppSelector';
import { fetchPosts } from '../../../store/thunks/postListThunk';

const Listing: React.FC = () => {
  const dispatch = useAppDispatch();
  const { data, status, error } = useAppSelector((state) => state.postList);

  useEffect(() => {
    if (status === 'idle') {
      dispatch(fetchPosts());
    }
  }, [status, dispatch]);

  if (status === 'loading') {
    return <h1>Loading...</h1>;
  }

  if (error) {
    return <h1>{error}</h1>;
  }

  return (
    <main>
      <table>
        <thead>
          <tr>
            <td>
              <h4>Title</h4>
            </td>
            <td>
              <h4>Category</h4>
            </td>
            <td>
              <h4>Status</h4>
            </td>
          </tr>
        </thead>
        <tbody>
          {data.map((post) => {
            return (
              <tr key={post.title}>
                <td>{post.category}</td>
                <td>{post.category}</td>
                <td>{post.status}</td>
              </tr>
            );
          })}
        </tbody>
      </table>
    </main>
  );
};

export default Listing;
