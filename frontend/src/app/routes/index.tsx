import { createBrowserRouter } from 'react-router-dom';
import Listing from './Listing';
import { ReactElement } from 'react';
import Navbar from '../../components/Navbar';

export const createRouter = () =>
  createBrowserRouter([
    {
      path: '/',
      element: (
        <Layout>
          <Listing />
        </Layout>
      )
    },
    {
      path: '/:id',
      lazy: async () => {
        const { Edit } = await import('./Edit');
        return {
          Component: () => (
            <Layout>
              <Edit />
            </Layout>
          )
        };
      }
    },
    {
      path: '/new',
      lazy: async () => {
        const { New } = await import('./New');
        return {
          Component: () => (
            <Layout>
              <New />
            </Layout>
          )
        };
      }
    },
    {
      path: '/preview',
      lazy: async () => {
        const { Preview } = await import('./Preview');
        return {
          Component: () => (
            <Layout>
              <Preview />
            </Layout>
          )
        };
      }
    }
  ]);

type LayoutProps = {
  children: ReactElement;
};

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <>
      <Navbar />
      {children}
    </>
  );
};
