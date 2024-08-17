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
        return { Component: () => <Edit /> };
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
