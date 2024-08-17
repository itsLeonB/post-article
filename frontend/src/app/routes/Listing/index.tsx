import { useEffect } from 'react';
import { useAppDispatch } from '../../../hooks/useAppDispatch';
import { useAppSelector } from '../../../hooks/useAppSelector';
import { fetchPosts } from '../../../store/thunks/postListThunk';
import { Tab, Tabs, TabList, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import { fetchPostStatuses } from '../../../store/thunks/postStatusThunk';

const Listing: React.FC = () => {
  const dispatch = useAppDispatch();
  const { data, status, error } = useAppSelector((state) => state.postList);
  const postStatusState = useAppSelector((state) => state.postStatus);

  useEffect(() => {
    if (status === 'idle') {
      dispatch(fetchPosts());
    }
  }, [status, dispatch]);

  useEffect(() => {
    if (postStatusState.status === 'idle') {
      dispatch(fetchPostStatuses());
    }
  }, [postStatusState.status, dispatch]);

  if (status === 'loading') {
    return <h1>Loading...</h1>;
  }

  if (error) {
    return <h1>{error}</h1>;
  }

  return (
    <main>
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
                        <tr key={post.title}>
                          <td>{post.title}</td>
                          <td>{post.category}</td>
                          <td>
                            <button>Edit</button>
                            <button>Trash</button>
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
