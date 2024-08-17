import { useCallback, useEffect } from 'react';
import { useAppDispatch } from '../../../hooks/useAppDispatch';
import { useAppSelector } from '../../../hooks/useAppSelector';
import { fetchPosts } from '../../../store/thunks/postListThunk';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import { fetchPostStatuses } from '../../../store/thunks/postStatusThunk';
import { Link } from 'react-router-dom';
import Loader from '../../../components/Loader';
import PageError from '../../../components/PageError';
import { resetStatus } from '../../../store/slices/postListSlice';
import { Post } from '../../../types/post';
import { updatePost } from '../../../store/thunks/postUpdateThunk';
import { toast, ToastContainer } from 'react-toastify';

import 'react-tabs/style/react-tabs.css';

const Listing: React.FC = () => {
  const dispatch = useAppDispatch();
  const { data, status, error } = useAppSelector((state) => state.postList);
  const postStatusState = useAppSelector((state) => state.postStatus);

  const fetchPostsData = useCallback(() => {
    dispatch(fetchPosts());
  }, [dispatch]);

  useEffect(() => {
    dispatch(resetStatus());
    fetchPostsData();
  }, [dispatch, fetchPostsData]);

  useEffect(() => {
    if (postStatusState.status === 'idle') {
      dispatch(fetchPostStatuses());
    }
  }, [postStatusState.status, dispatch]);

  if (status === 'loading') {
    return <Loader />;
  }

  if (error) {
    return <PageError error={error} />;
  }

  const trash = async (trashPost: Post) => {
    try {
      const trashingPost: Post = {
        id: trashPost.id,
        title: trashPost.title,
        content: trashPost.content,
        category: trashPost.category,
        status_id: 3
      };

      await dispatch(updatePost(trashingPost)).unwrap();
      fetchPostsData();
    } catch (error) {
      toast.error('Failed to trash post', { theme: 'dark' });
    }
  };

  return (
    <main>
      <ToastContainer />
      <Tabs>
        <TabList>
          {postStatusState.data.map((status) => {
            return <Tab key={status.id}>{status.name}</Tab>;
          })}
        </TabList>
        {postStatusState.data.map((status) => {
          return (
            <TabPanel key={status.id}>
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
                      <h4>Action</h4>
                    </td>
                  </tr>
                </thead>
                <tbody>
                  {data
                    .filter((item) => {
                      return item.status_id === status.id;
                    })
                    .map((post) => {
                      return (
                        <tr key={post.id}>
                          <td>{post.title}</td>
                          <td>{post.category}</td>
                          <td>
                            <button>
                              <Link to={`/${post.id}`}>Edit</Link>
                            </button>
                            {status.id !== 3 && (
                              <button onClick={() => trash(post)}>Trash</button>
                            )}
                          </td>
                        </tr>
                      );
                    })}
                </tbody>
              </table>
            </TabPanel>
          );
        })}
      </Tabs>
    </main>
  );
};

export default Listing;
